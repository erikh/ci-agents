package test

import (
	"testing"
	"time"

	check "github.com/erikh/check"
	"github.com/tinyci/ci-agents/grpc/handler"
	"github.com/tinyci/ci-agents/model"
	"github.com/tinyci/ci-agents/testutil"
	"github.com/tinyci/ci-agents/testutil/testclients"
	"github.com/tinyci/ci-agents/testutil/testservers"
)

type datasvcSuite struct {
	model        *model.Model
	dataDoneChan chan struct{}
	dataServer   *handler.H
	logDoneChan  chan struct{}
	logServer    *handler.H
	client       *testclients.DataClient
}

var _ = check.Suite(&datasvcSuite{})

func TestDataSvc(t *testing.T) {
	check.TestingT(t)
}

func (ds *datasvcSuite) SetUpTest(c *check.C) {
	testutil.WipeDB(c)

	var err error
	ds.model, err = model.New(testutil.TestDBConfig)
	c.Assert(err, check.IsNil)

	ds.dataServer, ds.dataDoneChan, err = testservers.MakeDataServer()
	c.Assert(err, check.IsNil)

	ds.logServer, ds.logDoneChan, _, err = testservers.MakeLogServer()
	c.Assert(err, check.IsNil)

	ds.client, err = testclients.NewDataClient()
	c.Assert(err, check.IsNil)
}

func (ds *datasvcSuite) TearDownTest(c *check.C) {
	close(ds.logDoneChan)
	close(ds.dataDoneChan)
	time.Sleep(100 * time.Millisecond)
}
