package logsvc

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/tinyci/ci-agents/gen/logsvc"
)

// LogJournal is a journal of log entries intended to facilitate mocking the logsvc.
type LogJournal struct {
	Journal map[string][]*logsvc.PutPayload
	mutex   sync.Mutex
}

// Tail echoes all journal entries to stdout
func (lj *LogJournal) Tail() {
	for {
		lj.mutex.Lock()
		for _, items := range lj.Journal {
			for _, item := range items {
				res := logrus.Fields{}
				for key, val := range item.Fields {
					res[key] = val
				}

				logrus.WithFields(res).Println(item.Message)
			}
		}
		// XXX manual version of Reset() to avoid deadlocking
		lj.Journal = map[string][]*logsvc.PutPayload{}
		lj.mutex.Unlock()
	}
}

// Reset resets the log journal, erasing all recorded messages.
func (lj *LogJournal) Reset() {
	lj.mutex.Lock()
	defer lj.mutex.Unlock()
	lj.Journal = map[string][]*logsvc.PutPayload{}
}

// Append appends a message.
func (lj *LogJournal) Append(level string, msg *logsvc.PutPayload) {
	lj.mutex.Lock()
	defer lj.mutex.Unlock()

	if _, ok := lj.Journal[level]; !ok {
		lj.Journal[level] = []*logsvc.PutPayload{}
	}

	lj.Journal[level] = append(lj.Journal[level], msg)
}
