// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package objectstores3caller

import (
	"testing"
	time "time"

	"go.uber.org/goleak"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	controller "github.com/juju/juju/controller"
	corehttp "github.com/juju/juju/core/http"
	"github.com/juju/juju/core/logger"
	watcher "github.com/juju/juju/core/watcher"
	"github.com/juju/juju/core/watcher/watchertest"
	loggertesting "github.com/juju/juju/internal/logger/testing"
	coretesting "github.com/juju/juju/internal/testing"
)

//go:generate go run go.uber.org/mock/mockgen -typed -package objectstores3caller -destination package_mock_test.go github.com/juju/juju/core/objectstore Client,Session
//go:generate go run go.uber.org/mock/mockgen -typed -package objectstores3caller -destination services_mocks_test.go github.com/juju/juju/internal/worker/objectstores3caller ControllerConfigService
//go:generate go run go.uber.org/mock/mockgen -typed -package objectstores3caller -destination http_mocks_test.go github.com/juju/juju/internal/s3client HTTPClient
//go:generate go run go.uber.org/mock/mockgen -typed -package objectstores3caller -destination clock_mocks_test.go github.com/juju/clock Clock
//go:generate go run go.uber.org/mock/mockgen -typed -package objectstores3caller -destination httpclient_mock_test.go github.com/juju/juju/core/http HTTPClientGetter

func TestPackage(t *testing.T) {
	defer goleak.VerifyNone(t)

	gc.TestingT(t)
}

type baseSuite struct {
	states chan string

	session                 *MockSession
	controllerConfigService *MockControllerConfigService
	clock                   *MockClock

	httpClientGetter *MockHTTPClientGetter
	httpClient       *MockHTTPClient

	logger logger.Logger
}

func (s *baseSuite) setupMocks(c *gc.C) *gomock.Controller {
	// Ensure we buffer the channel, this is because we might miss the
	// event if we're too quick at starting up.
	s.states = make(chan string, 1)

	ctrl := gomock.NewController(c)

	s.session = NewMockSession(ctrl)
	s.controllerConfigService = NewMockControllerConfigService(ctrl)
	s.clock = NewMockClock(ctrl)

	s.httpClientGetter = NewMockHTTPClientGetter(ctrl)
	s.httpClient = NewMockHTTPClient(ctrl)

	s.logger = loggertesting.WrapCheckLog(c)

	return ctrl
}

func (s *baseSuite) expectClock() {
	s.clock.EXPECT().Now().AnyTimes()
}

func (s *baseSuite) expectTimeAfter() {
	s.clock.EXPECT().After(gomock.Any()).DoAndReturn(func(_ time.Duration) <-chan time.Time {
		ch := make(chan time.Time)
		close(ch)
		return ch
	}).AnyTimes()
}

func (s *baseSuite) expectHTTPClient(c *gc.C) {
	s.httpClientGetter.EXPECT().GetHTTPClient(gomock.Any(), corehttp.S3Purpose).Return(s.httpClient, nil)
}

func (s *baseSuite) expectControllerConfig(c *gc.C, config controller.Config) {
	s.controllerConfigService.EXPECT().ControllerConfig(gomock.Any()).Return(config, nil)
}

func (s *baseSuite) expectControllerConfigWatch(c *gc.C) {
	s.controllerConfigService.EXPECT().WatchControllerConfig().DoAndReturn(func() (watcher.Watcher[[]string], error) {
		ch := make(chan []string)
		go func() {
			select {
			case ch <- []string{}:
			case <-time.After(coretesting.ShortWait * 10):
				c.Fatalf("timed out sending change")
			}
		}()
		return watchertest.NewMockStringsWatcher(ch), nil
	})
}

func (s *baseSuite) expectControllerConfigWatchWithChanges(c *gc.C, changes <-chan []string) {
	s.controllerConfigService.EXPECT().WatchControllerConfig().DoAndReturn(func() (watcher.Watcher[[]string], error) {
		return watchertest.NewMockStringsWatcher(changes), nil
	})
}

func (s *baseSuite) ensureStartup(c *gc.C) {
	select {
	case state := <-s.states:
		c.Assert(state, gc.Equals, stateStarted)
	case <-time.After(coretesting.ShortWait * 10):
		c.Fatalf("timed out waiting for startup")
	}
}

func (s *baseSuite) sendInitialChange(c *gc.C, changes chan<- []string) {
	done := make(chan struct{})
	go func() {
		defer close(done)

		select {
		case changes <- []string{}:
		case <-time.After(coretesting.ShortWait * 10):
			c.Fatalf("timed out sending change")
		}
	}()
	select {
	case <-done:
	case <-time.After(coretesting.ShortWait * 10):
		c.Fatalf("timed out waiting for change to be sent")
	}
}
