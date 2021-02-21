package logsvc

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/gen/logsvc"
	"google.golang.org/grpc/codes"
)

// Put submits a log message to the service, which in our current case, echoes it to stdout by way of sirupsen/logrus.
func (ls *LogServer) Put(ctx context.Context, lm *logsvc.PutPayload) error {
	dispatcher, ok := ls.DispatchTable[lm.Level]
	if !ok {
		return errors.Errorf("Invalid log level %q", lm.Level).ToGRPC(codes.FailedPrecondition)
	}

	fields := map[string]interface{}{}

	for key, value := range lm.Fields {
		fields[key] = value
	}

	dispatcher(logrus.WithFields(fields), lm)
	return nil
}
