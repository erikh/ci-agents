package logsvc

import (
	transport "github.com/erikh/go-transport"
	"github.com/tinyci/ci-agents/ci-gen/grpc/handler"
	client "github.com/tinyci/ci-agents/clients/log"
	"github.com/tinyci/ci-agents/config"
	"github.com/tinyci/ci-agents/errors"
	logsvcpb "github.com/tinyci/ci-agents/gen/grpc/logsvc/pb"
	logsvcsvr "github.com/tinyci/ci-agents/gen/grpc/logsvc/server"
	"github.com/tinyci/ci-agents/gen/logsvc"
	"google.golang.org/grpc"
)

// MakeLogServer makes a logsvc.
func MakeLogServer() (*handler.H, chan struct{}, *LogJournal, error) {
	journal := &LogJournal{Journal: map[string][]*logsvc.PutPayload{}}

	logDispatch := DispatchTable{
		client.LevelDebug: func(wf Dispatcher, msg *logsvc.PutPayload) {
			journal.Append(client.LevelDebug, msg)
		},
		client.LevelError: func(wf Dispatcher, msg *logsvc.PutPayload) {
			journal.Append(client.LevelError, msg)
		},
		client.LevelInfo: func(wf Dispatcher, msg *logsvc.PutPayload) {
			journal.Append(client.LevelInfo, msg)
		},
	}

	h := &handler.H{
		Service: config.Service{
			Name: "logsvc",
		},
		UserConfig: config.UserConfig{
			Port: 6005,
			// FIXME this is really dumb and should be unnecessary
			Auth: config.AuthConfig{
				TokenCryptKey: "1431d583a48a00243cc3d3d596ed362d77c50be4848dbf0d2f52bab841f072f9",
			},
		},
	}

	t, err := transport.Listen(nil, "tcp", config.DefaultServices.Log.String())
	if err != nil {
		return nil, nil, nil, errors.New(err)
	}

	srv := grpc.NewServer()
	svc := logsvcsvr.New(logsvc.NewEndpoints(New(logDispatch)), nil)

	logsvcpb.RegisterLogsvcServer(srv, svc)

	doneChan, err := h.Boot(t, srv, make(chan struct{}))
	return h, doneChan, journal, errors.New(err)
}
