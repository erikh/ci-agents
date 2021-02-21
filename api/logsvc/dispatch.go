package logsvc

import (
	"fmt"
	"time"

	client "github.com/tinyci/ci-agents/clients/log"
	"github.com/tinyci/ci-agents/gen/logsvc"
)

// Dispatcher dispatches logs based on loglevel.
type Dispatcher interface {
	Debug(...interface{})
	Error(...interface{})
	Info(...interface{})
}

// DispatchTable is a level -> execution function map
type DispatchTable map[string]func(wf Dispatcher, msg *logsvc.PutPayload)

var logLevelDispatch = DispatchTable{
	client.LevelDebug: func(wf Dispatcher, msg *logsvc.PutPayload) {
		wf.Debug(formatMsg(msg))
	},
	client.LevelError: func(wf Dispatcher, msg *logsvc.PutPayload) {
		wf.Error(formatMsg(msg))
	},
	client.LevelInfo: func(wf Dispatcher, msg *logsvc.PutPayload) {
		wf.Info(formatMsg(msg))
	},
}

func formatMsg(msg *logsvc.PutPayload) string {
	fmt.Printf("msg: %+v\n", msg)
	return fmt.Sprintf("[%v][%s] %s", time.Unix(msg.At, 0), msg.Service, msg.Message)
}
