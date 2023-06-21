// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package dbaccessor

import (
	"context"
	sql "database/sql"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v3"
	"github.com/juju/worker/v3/dependency"
	"github.com/juju/worker/v3/workertest"
	gc "gopkg.in/check.v1"

	coredatabase "github.com/juju/juju/core/database"
	"github.com/juju/juju/database/app"
	"github.com/juju/juju/database/dqlite"
	"github.com/juju/juju/pubsub/apiserver"
	"github.com/juju/juju/testing"
)

type workerSuite struct {
	dbBaseSuite

	nodeManager *MockNodeManager
}

var _ = gc.Suite(&workerSuite{})

func (s *workerSuite) TestStartupNotExistingNodeThenCluster(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()
	s.expectTrackedDBKill()

	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(c.MkDir(), nil)
	mgrExp.IsExistingNode().Return(false, nil).Times(4)
	mgrExp.WithAddressOption("10.6.6.6").Return(nil)
	mgrExp.WithClusterOption([]string{"10.6.6.7"}).Return(nil)
	mgrExp.WithLogFuncOption().Return(nil)
	mgrExp.WithTLSOption().Return(nil, nil)
	mgrExp.WithTracingOption().Return(nil)
	mgrExp.IsBootstrappedNode(gomock.Any()).Return(false, nil)

	s.client.EXPECT().Cluster(gomock.Any()).Return(nil, nil)

	sync := s.expectNodeStartupAndShutdown(true)

	// When we are starting up as a new node,
	// we request details immediately.
	s.hub.EXPECT().Subscribe(apiserver.DetailsTopic, gomock.Any()).Return(func() {}, nil)
	s.hub.EXPECT().Publish(apiserver.DetailsRequestTopic, gomock.Any()).Return(func() {}, nil)

	w := s.newWorker(c)
	defer workertest.CleanKill(c, w)

	// Without a bind address for ourselves we keep waiting.
	select {
	case w.(*dbWorker).apiServerChanges <- apiserver.Details{
		Servers: map[string]apiserver.APIServer{
			"0": {ID: "0"},
			"1": {ID: "1", InternalAddress: "10.6.6.7:1234"},
		},
	}:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for cluster change to be processed")
	}

	// Without other cluster members we keep waiting.
	select {
	case w.(*dbWorker).apiServerChanges <- apiserver.Details{
		Servers: map[string]apiserver.APIServer{
			"0": {ID: "0", InternalAddress: "10.6.6.6:1234"},
		},
	}:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for cluster change to be processed")
	}

	// At this point, the Dqlite node is not started.
	// The worker is waiting for legitimate server detail messages.
	select {
	case <-sync:
		c.Fatal("Dqlite node should not be started yet.")
	case <-time.After(testing.ShortWait):
	}

	// Push a message onto the API details channel,
	// enabling node startup as a cluster member.
	select {
	case w.(*dbWorker).apiServerChanges <- apiserver.Details{
		Servers: map[string]apiserver.APIServer{
			"0": {ID: "0", InternalAddress: "10.6.6.6:1234"},
			"1": {ID: "1", InternalAddress: "10.6.6.7:1234"},
		},
	}:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for cluster change to be processed")
	}

	select {
	case <-sync:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for Dqlite node start")
	}

	s.client.EXPECT().Leader(gomock.Any()).Return(&dqlite.NodeInfo{
		ID:      1,
		Address: "10.10.1.1",
	}, nil)
	report := w.(interface{ Report() map[string]any }).Report()
	c.Assert(report, MapHasKeys, []string{
		"leader",
		"leader-id",
		"leader-role",
	})

	workertest.CleanKill(c, w)
}

func (s *workerSuite) TestWorkerStartupExistingNode(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()
	s.expectTrackedDBKill()

	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(c.MkDir(), nil)

	// If this is an existing node, we do not invoke the address or cluster
	// options, but if the node is not as bootstrapped, we do assume it is
	// part of a cluster, and uses the TLS option.
	// IsBootstrapped node is called twice - once to check the startup
	// conditions and then again upon worker shutdown.
	mgrExp.IsExistingNode().Return(true, nil)
	mgrExp.IsBootstrappedNode(gomock.Any()).Return(false, nil).Times(2)
	mgrExp.WithLogFuncOption().Return(nil)
	mgrExp.WithTLSOption().Return(nil, nil)
	mgrExp.WithTracingOption().Return(nil)

	s.client.EXPECT().Cluster(gomock.Any()).Return(nil, nil)

	sync := s.expectNodeStartupAndShutdown(true)

	s.hub.EXPECT().Subscribe(apiserver.DetailsTopic, gomock.Any()).Return(func() {}, nil)

	w := s.newWorker(c)
	defer workertest.CleanKill(c, w)

	select {
	case <-sync:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for Dqlite node start")
	}

	workertest.CleanKill(c, w)
}

func (s *workerSuite) TestWorkerStartupAsBootstrapNodeSingleServerNoRebind(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()
	s.expectTrackedDBKill()

	dataDir := c.MkDir()
	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(dataDir, nil).MinTimes(1)

	// If this is an existing node, we do not
	// invoke the address or cluster options.
	mgrExp.IsExistingNode().Return(true, nil).Times(3)
	mgrExp.IsBootstrappedNode(gomock.Any()).Return(true, nil).Times(4)
	mgrExp.WithLogFuncOption().Return(nil)
	mgrExp.WithTracingOption().Return(nil)

	s.client.EXPECT().Cluster(gomock.Any()).Return(nil, nil)

	sync := s.expectNodeStartupAndShutdown(false)

	s.hub.EXPECT().Subscribe(apiserver.DetailsTopic, gomock.Any()).Return(func() {}, nil)

	w := s.newWorker(c)
	defer workertest.CleanKill(c, w)

	select {
	case <-sync:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for Dqlite node start")
	}

	// At this point we have started successfully.
	// Push a message onto the API details channel.
	// A single server does not cause a binding change.
	select {
	case w.(*dbWorker).apiServerChanges <- apiserver.Details{
		Servers: map[string]apiserver.APIServer{
			"0": {ID: "0", InternalAddress: "10.6.6.6:1234"},
		},
	}:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for cluster change to be processed")
	}

	// Multiple servers still do not cause a binding change
	// if there is no internal address to bind to.
	select {
	case w.(*dbWorker).apiServerChanges <- apiserver.Details{
		Servers: map[string]apiserver.APIServer{
			"0": {ID: "0"},
			"1": {ID: "1", InternalAddress: "10.6.6.7:1234"},
			"2": {ID: "2", InternalAddress: "10.6.6.8:1234"},
		},
	}:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for cluster change to be processed")
	}

	workertest.CleanKill(c, w)
}

func (s *workerSuite) TestWorkerStartupAsBootstrapNodeThenReconfigure(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()
	s.expectTrackedDBKill()

	dataDir := c.MkDir()
	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(dataDir, nil).MinTimes(1)

	// If this is an existing node, we do not
	// invoke the address or cluster options.
	mgrExp.IsExistingNode().Return(true, nil).Times(2)
	gomock.InOrder(
		mgrExp.IsBootstrappedNode(gomock.Any()).Return(true, nil).Times(2),
		// This is the check at shutdown.
		mgrExp.IsBootstrappedNode(gomock.Any()).Return(false, nil))
	mgrExp.WithLogFuncOption().Return(nil)
	mgrExp.WithTracingOption().Return(nil)

	// These are the expectations around reconfiguring
	// the cluster and local node.
	mgrExp.ClusterServers(gomock.Any()).Return([]dqlite.NodeInfo{
		{
			ID:      3297041220608546238,
			Address: "127.0.0.1:17666",
			Role:    0,
		},
	}, nil)
	mgrExp.SetClusterServers(gomock.Any(), []dqlite.NodeInfo{
		{
			ID:      3297041220608546238,
			Address: "10.6.6.6:17666",
			Role:    0,
		},
	}).Return(nil)
	mgrExp.SetNodeInfo(dqlite.NodeInfo{
		ID:      3297041220608546238,
		Address: "10.6.6.6:17666",
		Role:    0,
	}).Return(nil)

	s.client.EXPECT().Cluster(gomock.Any()).Return(nil, nil)

	// Although the shut-down check for IsBootstrappedNode returns false,
	// this call to shut-down is actually run before reconfiguring the node.
	// When the loop exits, the node is already set to nil.
	sync := s.expectNodeStartupAndShutdown(false)

	s.hub.EXPECT().Subscribe(apiserver.DetailsTopic, gomock.Any()).Return(func() {}, nil)

	w := s.newWorker(c)
	defer workertest.DirtyKill(c, w)

	select {
	case <-sync:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for Dqlite node start")
	}

	// At this point we have started successfully.
	// Push a message onto the API details channel to simulate a move into HA.
	select {
	case w.(*dbWorker).apiServerChanges <- apiserver.Details{
		Servers: map[string]apiserver.APIServer{
			"0": {ID: "0", InternalAddress: "10.6.6.6:1234"},
			"1": {ID: "1", InternalAddress: "10.6.6.7:1234"},
			"2": {ID: "2", InternalAddress: "10.6.6.8:1234"},
		},
	}:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for cluster change to be processed")
	}

	err := workertest.CheckKilled(c, w)
	c.Assert(errors.Is(err, dependency.ErrBounce), jc.IsTrue)
}

func (s *workerSuite) TestEnsureNamespaceForController(c *gc.C) {
	defer s.setupMocks(c).Finish()

	w := &dbWorker{
		dbApp: s.dbApp,
	}

	err := w.ensureNamespace(coredatabase.ControllerNS)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *workerSuite) TestEnsureNamespaceForModelNotFound(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()

	dataDir := c.MkDir()
	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(dataDir, nil).MinTimes(1)

	// If this is an existing node, we do not
	// invoke the address or cluster options.
	mgrExp.IsExistingNode().Return(true, nil).Times(3)
	mgrExp.IsBootstrappedNode(gomock.Any()).Return(true, nil).Times(4)
	mgrExp.WithLogFuncOption().Return(nil)
	mgrExp.WithTracingOption().Return(nil)

	s.client.EXPECT().Cluster(gomock.Any()).Return(nil, nil)

	sync := s.expectNodeStartupAndShutdown(false)

	s.hub.EXPECT().Subscribe(apiserver.DetailsTopic, gomock.Any()).Return(func() {}, nil)

	trackedWorkerDB := newWorkerTrackedDB(s.TxnRunner())

	w := s.newWorkerWithDB(c, trackedWorkerDB)
	defer workertest.CleanKill(c, w)

	dbw := w.(*dbWorker)
	s.ensureStartup(c, dbw, sync)

	err := dbw.ensureNamespace("foo")
	c.Assert(err, jc.Satisfies, errors.IsNotFound)

	workertest.CleanKill(c, w)
}

func (s *workerSuite) TestEnsureNamespaceForModel(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()

	dataDir := c.MkDir()
	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(dataDir, nil).MinTimes(1)

	// If this is an existing node, we do not
	// invoke the address or cluster options.
	mgrExp.IsExistingNode().Return(true, nil).Times(3)
	mgrExp.IsBootstrappedNode(gomock.Any()).Return(true, nil).Times(4)
	mgrExp.WithLogFuncOption().Return(nil)
	mgrExp.WithTracingOption().Return(nil)

	s.client.EXPECT().Cluster(gomock.Any()).Return(nil, nil)

	sync := s.expectNodeStartupAndShutdown(false)

	s.hub.EXPECT().Subscribe(apiserver.DetailsTopic, gomock.Any()).Return(func() {}, nil)

	trackedWorkerDB := newWorkerTrackedDB(s.TxnRunner())

	w := s.newWorkerWithDB(c, trackedWorkerDB)
	defer workertest.CleanKill(c, w)

	ctx, cancel := context.WithTimeout(context.Background(), testing.LongWait)
	defer cancel()

	err := s.TxnRunner().StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		stmt := "INSERT INTO model_list (uuid) VALUES (?);"
		result, err := tx.ExecContext(ctx, stmt, "foo")
		c.Assert(err, jc.ErrorIsNil)

		num, err := result.RowsAffected()
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(num, gc.Equals, int64(1))

		return nil
	})
	c.Assert(err, jc.ErrorIsNil)

	dbw := w.(*dbWorker)
	s.ensureStartup(c, dbw, sync)

	err = dbw.ensureNamespace("foo")
	c.Assert(err, jc.ErrorIsNil)

	workertest.CleanKill(c, w)
}

func (s *workerSuite) TestEnsureNamespaceForModelWithCache(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()

	dataDir := c.MkDir()
	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(dataDir, nil).MinTimes(1)

	// If this is an existing node, we do not
	// invoke the address or cluster options.
	mgrExp.IsExistingNode().Return(true, nil).Times(3)
	mgrExp.IsBootstrappedNode(gomock.Any()).Return(true, nil).Times(4)
	mgrExp.WithLogFuncOption().Return(nil)
	mgrExp.WithTracingOption().Return(nil)

	s.client.EXPECT().Cluster(gomock.Any()).Return(nil, nil)

	sync := s.expectNodeStartupAndShutdown(false)

	s.hub.EXPECT().Subscribe(apiserver.DetailsTopic, gomock.Any()).Return(func() {}, nil)

	trackedWorkerDB := newWorkerTrackedDB(s.TxnRunner())

	w := s.newWorkerWithDB(c, trackedWorkerDB)
	defer workertest.CleanKill(c, w)

	ctx, cancel := context.WithTimeout(context.Background(), testing.LongWait)
	defer cancel()

	var attempt int
	err := s.TxnRunner().StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		attempt++

		stmt := "INSERT INTO model_list (uuid) VALUES (?);"
		result, err := tx.ExecContext(ctx, stmt, "foo")
		c.Assert(err, jc.ErrorIsNil)

		num, err := result.RowsAffected()
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(num, gc.Equals, int64(1))

		return nil
	})
	c.Assert(err, jc.ErrorIsNil)

	dbw := w.(*dbWorker)
	s.ensureStartup(c, dbw, sync)

	err = dbw.ensureNamespace("foo")
	c.Assert(err, jc.ErrorIsNil)

	// The second query will be cached.
	err = dbw.ensureNamespace("foo")
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(attempt, gc.Equals, 1)

	workertest.CleanKill(c, w)
}

func (s *workerSuite) TestCloseDatabaseForController(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()

	dataDir := c.MkDir()
	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(dataDir, nil).MinTimes(1)

	// If this is an existing node, we do not
	// invoke the address or cluster options.
	mgrExp.IsExistingNode().Return(true, nil).Times(3)
	mgrExp.IsBootstrappedNode(gomock.Any()).Return(true, nil).Times(4)
	mgrExp.WithLogFuncOption().Return(nil)
	mgrExp.WithTracingOption().Return(nil)

	s.client.EXPECT().Cluster(gomock.Any()).Return(nil, nil)

	sync := s.expectNodeStartupAndShutdown(false)

	s.hub.EXPECT().Subscribe(apiserver.DetailsTopic, gomock.Any()).Return(func() {}, nil)

	trackedWorkerDB := newWorkerTrackedDB(s.TxnRunner())

	w := s.newWorkerWithDB(c, trackedWorkerDB)
	defer workertest.CleanKill(c, w)

	ctx, cancel := context.WithTimeout(context.Background(), testing.LongWait)
	defer cancel()

	err := s.TxnRunner().StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		stmt := "INSERT INTO model_list (uuid) VALUES (?);"
		result, err := tx.ExecContext(ctx, stmt, "foo")
		c.Assert(err, jc.ErrorIsNil)

		num, err := result.RowsAffected()
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(num, gc.Equals, int64(1))

		return nil
	})
	c.Assert(err, jc.ErrorIsNil)

	dbw := w.(*dbWorker)
	s.ensureStartup(c, dbw, sync)

	err = dbw.closeDatabase(coredatabase.ControllerNS)
	c.Assert(err, gc.ErrorMatches, "cannot close controller database")

	workertest.CleanKill(c, w)
}

func (s *workerSuite) TestCloseDatabaseForModel(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()

	dataDir := c.MkDir()
	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(dataDir, nil).MinTimes(1)

	// If this is an existing node, we do not
	// invoke the address or cluster options.
	mgrExp.IsExistingNode().Return(true, nil).Times(3)
	mgrExp.IsBootstrappedNode(gomock.Any()).Return(true, nil).Times(4)
	mgrExp.WithLogFuncOption().Return(nil)
	mgrExp.WithTracingOption().Return(nil)

	s.client.EXPECT().Cluster(gomock.Any()).Return(nil, nil)

	sync := s.expectNodeStartupAndShutdown(false)

	s.hub.EXPECT().Subscribe(apiserver.DetailsTopic, gomock.Any()).Return(func() {}, nil)

	trackedWorkerDB := newWorkerTrackedDB(s.TxnRunner())

	w := s.newWorkerWithDB(c, trackedWorkerDB)
	defer workertest.CleanKill(c, w)

	ctx, cancel := context.WithTimeout(context.Background(), testing.LongWait)
	defer cancel()

	err := s.TxnRunner().StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		stmt := "INSERT INTO model_list (uuid) VALUES (?);"
		result, err := tx.ExecContext(ctx, stmt, "foo")
		c.Assert(err, jc.ErrorIsNil)

		num, err := result.RowsAffected()
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(num, gc.Equals, int64(1))

		return nil
	})
	c.Assert(err, jc.ErrorIsNil)

	dbw := w.(*dbWorker)
	s.ensureStartup(c, dbw, sync)

	_, err = dbw.GetDB("foo")
	c.Assert(err, jc.ErrorIsNil)

	err = dbw.closeDatabase("foo")
	c.Assert(err, jc.ErrorIsNil)

	workertest.CleanKill(c, w)
}

func (s *workerSuite) TestCloseDatabaseForUnknownModel(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()

	dataDir := c.MkDir()
	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(dataDir, nil).MinTimes(1)

	// If this is an existing node, we do not
	// invoke the address or cluster options.
	mgrExp.IsExistingNode().Return(true, nil).Times(3)
	mgrExp.IsBootstrappedNode(gomock.Any()).Return(true, nil).Times(4)
	mgrExp.WithLogFuncOption().Return(nil)
	mgrExp.WithTracingOption().Return(nil)

	s.client.EXPECT().Cluster(gomock.Any()).Return(nil, nil)

	sync := s.expectNodeStartupAndShutdown(false)

	s.hub.EXPECT().Subscribe(apiserver.DetailsTopic, gomock.Any()).Return(func() {}, nil)

	trackedWorkerDB := newWorkerTrackedDB(s.TxnRunner())

	w := s.newWorkerWithDB(c, trackedWorkerDB)
	defer workertest.CleanKill(c, w)

	dbw := w.(*dbWorker)
	s.ensureStartup(c, dbw, sync)

	err := dbw.closeDatabase("foo")
	c.Assert(err, gc.ErrorMatches, `stopping worker: worker "foo" not found`)

	workertest.CleanKill(c, w)
}

func (s *workerSuite) ensureStartup(c *gc.C, w *dbWorker, sync <-chan struct{}) {
	select {
	case <-sync:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for Dqlite node start")
	}

	// At this point we have started successfully.
	// Push a message onto the API details channel.
	// A single server does not cause a binding change.
	select {
	case w.apiServerChanges <- apiserver.Details{
		Servers: map[string]apiserver.APIServer{
			"0": {ID: "0", InternalAddress: "10.6.6.6:1234"},
		},
	}:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for cluster change to be processed")
	}

	// Multiple servers still do not cause a binding change
	// if there is no internal address to bind to.
	select {
	case w.apiServerChanges <- apiserver.Details{
		Servers: map[string]apiserver.APIServer{
			"0": {ID: "0"},
			"1": {ID: "1", InternalAddress: "10.6.6.7:1234"},
			"2": {ID: "2", InternalAddress: "10.6.6.8:1234"},
		},
	}:
	case <-time.After(testing.LongWait):
		c.Fatal("timed out waiting for cluster change to be processed")
	}
}

func (s *workerSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := s.baseSuite.setupMocks(c)
	s.nodeManager = NewMockNodeManager(ctrl)
	return ctrl
}

func (s *workerSuite) expectNodeStartupAndShutdown(handover bool) chan struct{} {
	sync := make(chan struct{})

	appExp := s.dbApp.EXPECT()
	appExp.Ready(gomock.Any()).Return(nil)
	appExp.Client(gomock.Any()).Return(s.client, nil).MinTimes(1)
	appExp.ID().DoAndReturn(func() uint64 {
		close(sync)
		return uint64(666)
	})

	if handover {
		appExp.Handover(gomock.Any()).Return(nil)
	}

	appExp.Close().Return(nil)

	return sync
}

func (s *workerSuite) newWorker(c *gc.C) worker.Worker {
	return s.newWorkerWithDB(c, s.trackedDB)
}

func (s *workerSuite) newWorkerWithDB(c *gc.C, db TrackedDB) worker.Worker {
	cfg := WorkerConfig{
		NodeManager:  s.nodeManager,
		Clock:        s.clock,
		Hub:          s.hub,
		ControllerID: "0",
		Logger:       s.logger,
		NewApp: func(string, ...app.Option) (DBApp, error) {
			return s.dbApp, nil
		},
		NewDBWorker: func(context.Context, DBApp, string, ...TrackedDBWorkerOption) (TrackedDB, error) {
			return db, nil
		},
		MetricsCollector: &Collector{},
	}

	w, err := NewWorker(cfg)
	c.Assert(err, jc.ErrorIsNil)
	return w
}