package logsvc

import (
	"context"
	"net/http"
	"net/url"

	check "github.com/erikh/check"
	"github.com/sirupsen/logrus"
	"github.com/tinyci/ci-agents/ci-gen/grpc/types"
	"github.com/tinyci/ci-agents/clients/log"
)

func (ls *logsvcSuite) Test01Journal(c *check.C) {
	log.New().Info(context.Background(), "test")
	j := ls.journal.Journal[log.LevelInfo]
	c.Assert(len(j), check.Equals, 1)
	c.Assert(j[0].Message, check.Equals, "test")
}

func (ls *logsvcSuite) TestLevels(c *check.C) {
	l := log.New()
	ctx := context.Background()

	l.Info(ctx, "test")
	l.Infof(ctx, "test %d", 2)

	j := ls.journal.Journal[log.LevelInfo]
	c.Assert(len(j), check.Equals, 2)
	c.Assert(j[0].Message, check.Equals, "test")
	c.Assert(j[1].Message, check.Equals, "test 2")

	l.Error(ctx, "test")
	l.Errorf(ctx, "test %d", 2)

	j = ls.journal.Journal[log.LevelError]
	c.Assert(len(j), check.Equals, 2)
	c.Assert(j[0].Message, check.Equals, "test")
	c.Assert(j[1].Message, check.Equals, "test 2")

	l.Debug(ctx, "test")
	l.Debugf(ctx, "test %d", 2)

	j = ls.journal.Journal[log.LevelDebug]
	c.Assert(len(j), check.Equals, 2)
	c.Assert(j[0].Message, check.Equals, "test")
	c.Assert(j[1].Message, check.Equals, "test 2")
}

func (ls *logsvcSuite) TestFiltering(c *check.C) {
	l := log.New()
	ctx := context.Background()

	table := map[logrus.Level]struct {
		info  bool
		debug bool
		error bool
	}{
		logrus.WarnLevel: {
			info:  false,
			debug: false,
			error: true,
		},
		logrus.InfoLevel: {
			info:  true,
			error: true,
			debug: false,
		},
		logrus.PanicLevel: {
			info:  false,
			error: false,
			debug: false,
		},
		logrus.DebugLevel: {
			info:  true,
			error: true,
			debug: true,
		},
		logrus.ErrorLevel: {
			info:  false,
			error: true,
			debug: false,
		},
	}

	for level, params := range table {
		ls.service.changeLevel(level)

		l.Info(ctx, "test")
		l.Infof(ctx, "test %d", 2)
		j := ls.journal.Journal[log.LevelInfo]

		if params.info {
			c.Assert(len(j), check.Equals, 2)
		} else {
			c.Assert(len(j), check.Equals, 0)
		}

		l.Debug(ctx, "test")
		l.Debugf(ctx, "test %d", 2)

		j = ls.journal.Journal[log.LevelDebug]
		if params.debug {
			c.Assert(len(j), check.Equals, 2)
		} else {
			c.Assert(len(j), check.Equals, 0)
		}

		l.Error(ctx, "test")
		l.Errorf(ctx, "test %d", 2)

		j = ls.journal.Journal[log.LevelError]
		if params.error {
			c.Assert(len(j), check.Equals, 2)
		} else {
			c.Assert(len(j), check.Equals, 0)
		}

		ls.journal.Reset()
	}
}

func (ls *logsvcSuite) TestFields(c *check.C) {
	l := log.New()
	ctx := context.Background()

	wf := l.WithFields(log.FieldMap{"test": "one", "test2": "two"})

	wf.Info(ctx, "test")
	wf.Infof(ctx, "test %d", 2)
	wf.Debug(ctx, "test")
	wf.Debugf(ctx, "test %d", 2)
	wf.Error(ctx, "test")
	wf.Errorf(ctx, "test %d", 2)

	for _, messages := range ls.journal.Journal {
		c.Assert(len(messages), check.Equals, 2)
		c.Assert(messages[0].Message, check.Equals, "test")
		c.Assert(messages[1].Message, check.Equals, "test 2")

		for i := 0; i < 2; i++ {
			val, ok := messages[i].Fields.Fields["test"]
			c.Assert(ok, check.Equals, true)
			c.Assert(val.GetStringValue(), check.Equals, "one")

			val, ok = messages[i].Fields.Fields["test2"]
			c.Assert(ok, check.Equals, true)
			c.Assert(val.GetStringValue(), check.Equals, "two")
		}
	}
}

func (ls *logsvcSuite) TestService(c *check.C) {
	l := log.New()
	ctx := context.Background()

	ws := l.WithService("test")

	ws.Info(ctx, "test")
	ws.Infof(ctx, "test %d", 2)
	ws.Debug(ctx, "test")
	ws.Debugf(ctx, "test %d", 2)
	ws.Error(ctx, "test")
	ws.Errorf(ctx, "test %d", 2)

	for _, messages := range ls.journal.Journal {
		c.Assert(len(messages), check.Equals, 2)
		c.Assert(messages[0].Message, check.Equals, "test")
		c.Assert(messages[1].Message, check.Equals, "test 2")

		for i := 0; i < 2; i++ {
			val, ok := messages[i].Fields.Fields["service"]
			c.Assert(ok, check.Equals, true)
			c.Assert(val.GetStringValue(), check.Equals, "test")
		}
	}
}

func (ls *logsvcSuite) TestRequest(c *check.C) {
	l := log.New()

	u, err := url.Parse("http://dogfood.tinyci.org")
	c.Assert(err, check.IsNil)

	req := &http.Request{
		RemoteAddr: "127.0.0.1:1234",
		Method:     "GET",
		URL:        u,
	}

	ctx := context.Background()

	wr := l.WithRequest(req)

	wr.Info(ctx, "test")
	wr.Infof(ctx, "test %d", 2)
	wr.Debug(ctx, "test")
	wr.Debugf(ctx, "test %d", 2)
	wr.Error(ctx, "test")
	wr.Errorf(ctx, "test %d", 2)

	for _, messages := range ls.journal.Journal {
		c.Assert(len(messages), check.Equals, 2)
		c.Assert(messages[0].Message, check.Equals, "test")
		c.Assert(messages[1].Message, check.Equals, "test 2")

		for i := 0; i < 2; i++ {
			val, ok := messages[i].Fields.Fields["remote_addr"]
			c.Assert(ok, check.Equals, true)
			c.Assert(val.GetStringValue(), check.Equals, "127.0.0.1")

			val, ok = messages[i].Fields.Fields["request_method"]
			c.Assert(ok, check.Equals, true)
			c.Assert(val.GetStringValue(), check.Equals, "GET")

			val, ok = messages[i].Fields.Fields["request_url"]
			c.Assert(ok, check.Equals, true)
			c.Assert(val.GetStringValue(), check.Equals, "http://dogfood.tinyci.org")
		}
	}

	u, err = url.Parse("https://pensando.io")
	c.Assert(err, check.IsNil)

	req = &http.Request{
		Header:     http.Header{"X-Forwarded-For": []string{"1.2.3.4,5.6.7.8"}},
		RemoteAddr: "127.0.0.1:1234",
		Method:     "POST",
		URL:        u,
	}

	ls.journal.Reset()

	wr = l.WithRequest(req)

	wr.Info(ctx, "test")
	wr.Infof(ctx, "test %d", 2)
	wr.Debug(ctx, "test")
	wr.Debugf(ctx, "test %d", 2)
	wr.Error(ctx, "test")
	wr.Errorf(ctx, "test %d", 2)

	for _, messages := range ls.journal.Journal {
		c.Assert(len(messages), check.Equals, 2)
		c.Assert(messages[0].Message, check.Equals, "test")
		c.Assert(messages[1].Message, check.Equals, "test 2")

		for i := 0; i < 2; i++ {
			val, ok := messages[i].Fields.Fields["remote_addr"]
			c.Assert(ok, check.Equals, true)
			c.Assert(val.GetStringValue(), check.Equals, "1.2.3.4")

			val, ok = messages[i].Fields.Fields["request_method"]
			c.Assert(ok, check.Equals, true)
			c.Assert(val.GetStringValue(), check.Equals, "POST")

			val, ok = messages[i].Fields.Fields["request_url"]
			c.Assert(ok, check.Equals, true)
			c.Assert(val.GetStringValue(), check.Equals, "https://pensando.io")
		}
	}
}

func (ls *logsvcSuite) TestUser(c *check.C) {
	l := log.New()
	ctx := context.Background()

	user := &types.User{Id: 1, Username: "erikh"}
	wu := l.WithUser(user)

	wu.Info(ctx, "test")
	wu.Infof(ctx, "test %d", 2)
	wu.Debug(ctx, "test")
	wu.Debugf(ctx, "test %d", 2)
	wu.Error(ctx, "test")
	wu.Errorf(ctx, "test %d", 2)

	for _, messages := range ls.journal.Journal {
		c.Assert(len(messages), check.Equals, 2)
		c.Assert(messages[0].Message, check.Equals, "test")
		c.Assert(messages[1].Message, check.Equals, "test 2")

		for i := 0; i < 2; i++ {
			val, ok := messages[i].Fields.Fields["user_id"]
			c.Assert(ok, check.Equals, true)
			c.Assert(val.GetStringValue(), check.Equals, "1")

			val, ok = messages[i].Fields.Fields["username"]
			c.Assert(ok, check.Equals, true)
			c.Assert(val.GetStringValue(), check.Equals, "erikh")
		}
	}
}
