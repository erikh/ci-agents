package assetsvc

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"time"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/tinyci/ci-agents/ci-gen/grpc/handler"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/gen/assetsvc"
	"google.golang.org/grpc/codes"
)

const (
	defaultLogsRoot   = "/var/tinyci/logs"
	logsRootConfigKey = "logs_root_path"
)

// AssetServer is the handler anchor for the GRPC network system.
type AssetServer struct {
	H *handler.H
}

func (as *AssetServer) getLogsRoot() string {
	p, ok := as.H.ServiceConfig[logsRootConfigKey].(string)
	if !ok {
		p = defaultLogsRoot
	}

	return p
}

func makeGRPCError(err error) error {
	if err != nil {
		var retErr *errors.Error
		switch err := err.(type) {
		case *errors.Error:
			if err == nil {
				return nil
			}
			retErr = err
		default:
			retErr = errors.New(err)
		}
		return retErr.ToGRPC(codes.FailedPrecondition)
	}

	return nil
}

// PutLog writes the log to disk
func (as *AssetServer) PutLog(ctx context.Context, ap assetsvc.PutLogServerStream) error {
	defer ap.Close()
	return makeGRPCError(as.submit(ap, as.getLogsRoot()))
}

// GetLog spills the log back to connecting websocket.
func (as *AssetServer) GetLog(ctx context.Context, id *assetsvc.GetLogPayload, ag assetsvc.GetLogServerStream) error {
	defer ag.Close()
	return makeGRPCError(as.attach(id.ID, ag, as.getLogsRoot()))
}

func (as *AssetServer) submit(ap assetsvc.PutLogServerStream, p string) (retErr *errors.Error) {
	if err := os.MkdirAll(p, 0700); err != nil {
		return errors.New(err)
	}

	ls, err := ap.Recv()
	if err != nil {
		return errors.New(err)
	}

	file := path.Join(p, fmt.Sprintf("%d", ls.ID))

	if _, err := os.Stat(file); err == nil {
		return errors.New("log already exists")
	}

	if _, err := os.Stat(file + ".writing"); err == nil {
		return errors.New("log writing is currently in progress for this ID")
	}

	writing, err := os.Create(file + ".writing")
	if err != nil {
		return errors.New(err)
	}
	defer func() {
		writing.Close()
		os.Remove(writing.Name())
	}()

	f, err := os.Create(file)
	if err != nil {
		return errors.New(err)
	}
	defer f.Close()

	for {
		if _, err := f.Write(ls.Chunk); err != nil {
			return errors.New(err)
		}
		if ls, err = ap.Recv(); err != nil {
			if err == io.EOF {
				return nil
			}
			return errors.New(err)
		}
	}
}

func write(ag assetsvc.GetLogServerStream, buf []byte) error {
	return ag.Send(&assetsvc.GetLogResult{Chunk: buf})
}

func (as *AssetServer) attach(id int64, ag assetsvc.GetLogServerStream, p string) (retErr error) {
	defer func() {
		if retErr != nil {
		} else {
			retErr = write(ag, []byte(color.New(color.FgGreen).Sprintln("---- LOG COMPLETE ----")))
		}
	}()

	file := path.Join(p, fmt.Sprintf("%d", id))

	log, err := os.Open(file) // #nosec
	if err != nil {
		return err
	}
	defer log.Close()

	buf := make([]byte, 256)
	var (
		last bool
	)

	writeBuf := []byte{}

	for {
		// these two conditionals kind of wrap the main body of the loop to allow
		// the secret sauce; thanks `perldoc -q tail`!
		n, readErr := log.Read(buf)
		if readErr == io.EOF {
			if _, err := os.Stat(file + ".writing"); err != nil {
				return nil // file is done writing
			}

			if _, err := log.Seek(0, io.SeekEnd); err != nil {
				last = true
			} else {
				time.Sleep(500 * time.Millisecond)
			}
		} else if readErr != nil {
			last = true
		}

		writeBuf = append(writeBuf, buf[:n]...)
		// websockets running in textencoding (Which xterm.js requires) require UTF8
		// clean strings to be passed on each write, otherwise it will break the
		// connection.
		// XXX this buffering can probably be abused somehow.
		if utf8.ValidString(string(writeBuf)) {
			if err := write(ag, writeBuf); err != nil {
				if err != io.EOF { // spam-free log experience
					return err
				}

				return nil
			}

			if last {
				if readErr != io.EOF {
					return readErr
				}
			}
			writeBuf = []byte{}
		}
	}
}
