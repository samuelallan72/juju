// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/upgradesteps (interfaces: UpgradeService)
//
// Generated by this command:
//
//	mockgen -typed -package upgradesteps -destination upgradeservice_mock_test.go github.com/juju/juju/internal/worker/upgradesteps UpgradeService
//

// Package upgradesteps is a generated GoMock package.
package upgradesteps

import (
	context "context"
	reflect "reflect"

	upgrade "github.com/juju/juju/core/upgrade"
	watcher "github.com/juju/juju/core/watcher"
	upgrade0 "github.com/juju/juju/domain/upgrade"
	gomock "go.uber.org/mock/gomock"
)

// MockUpgradeService is a mock of UpgradeService interface.
type MockUpgradeService struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeServiceMockRecorder
}

// MockUpgradeServiceMockRecorder is the mock recorder for MockUpgradeService.
type MockUpgradeServiceMockRecorder struct {
	mock *MockUpgradeService
}

// NewMockUpgradeService creates a new mock instance.
func NewMockUpgradeService(ctrl *gomock.Controller) *MockUpgradeService {
	mock := &MockUpgradeService{ctrl: ctrl}
	mock.recorder = &MockUpgradeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeService) EXPECT() *MockUpgradeServiceMockRecorder {
	return m.recorder
}

// ActiveUpgrade mocks base method.
func (m *MockUpgradeService) ActiveUpgrade(arg0 context.Context) (upgrade0.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveUpgrade", arg0)
	ret0, _ := ret[0].(upgrade0.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveUpgrade indicates an expected call of ActiveUpgrade.
func (mr *MockUpgradeServiceMockRecorder) ActiveUpgrade(arg0 any) *MockUpgradeServiceActiveUpgradeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveUpgrade", reflect.TypeOf((*MockUpgradeService)(nil).ActiveUpgrade), arg0)
	return &MockUpgradeServiceActiveUpgradeCall{Call: call}
}

// MockUpgradeServiceActiveUpgradeCall wrap *gomock.Call
type MockUpgradeServiceActiveUpgradeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceActiveUpgradeCall) Return(arg0 upgrade0.UUID, arg1 error) *MockUpgradeServiceActiveUpgradeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceActiveUpgradeCall) Do(f func(context.Context) (upgrade0.UUID, error)) *MockUpgradeServiceActiveUpgradeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceActiveUpgradeCall) DoAndReturn(f func(context.Context) (upgrade0.UUID, error)) *MockUpgradeServiceActiveUpgradeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetControllerDone mocks base method.
func (m *MockUpgradeService) SetControllerDone(arg0 context.Context, arg1 upgrade0.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetControllerDone", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetControllerDone indicates an expected call of SetControllerDone.
func (mr *MockUpgradeServiceMockRecorder) SetControllerDone(arg0, arg1, arg2 any) *MockUpgradeServiceSetControllerDoneCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetControllerDone", reflect.TypeOf((*MockUpgradeService)(nil).SetControllerDone), arg0, arg1, arg2)
	return &MockUpgradeServiceSetControllerDoneCall{Call: call}
}

// MockUpgradeServiceSetControllerDoneCall wrap *gomock.Call
type MockUpgradeServiceSetControllerDoneCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceSetControllerDoneCall) Return(arg0 error) *MockUpgradeServiceSetControllerDoneCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceSetControllerDoneCall) Do(f func(context.Context, upgrade0.UUID, string) error) *MockUpgradeServiceSetControllerDoneCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceSetControllerDoneCall) DoAndReturn(f func(context.Context, upgrade0.UUID, string) error) *MockUpgradeServiceSetControllerDoneCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetDBUpgradeFailed mocks base method.
func (m *MockUpgradeService) SetDBUpgradeFailed(arg0 context.Context, arg1 upgrade0.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDBUpgradeFailed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDBUpgradeFailed indicates an expected call of SetDBUpgradeFailed.
func (mr *MockUpgradeServiceMockRecorder) SetDBUpgradeFailed(arg0, arg1 any) *MockUpgradeServiceSetDBUpgradeFailedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDBUpgradeFailed", reflect.TypeOf((*MockUpgradeService)(nil).SetDBUpgradeFailed), arg0, arg1)
	return &MockUpgradeServiceSetDBUpgradeFailedCall{Call: call}
}

// MockUpgradeServiceSetDBUpgradeFailedCall wrap *gomock.Call
type MockUpgradeServiceSetDBUpgradeFailedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceSetDBUpgradeFailedCall) Return(arg0 error) *MockUpgradeServiceSetDBUpgradeFailedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceSetDBUpgradeFailedCall) Do(f func(context.Context, upgrade0.UUID) error) *MockUpgradeServiceSetDBUpgradeFailedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceSetDBUpgradeFailedCall) DoAndReturn(f func(context.Context, upgrade0.UUID) error) *MockUpgradeServiceSetDBUpgradeFailedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpgradeInfo mocks base method.
func (m *MockUpgradeService) UpgradeInfo(arg0 context.Context, arg1 upgrade0.UUID) (upgrade.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeInfo", arg0, arg1)
	ret0, _ := ret[0].(upgrade.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeInfo indicates an expected call of UpgradeInfo.
func (mr *MockUpgradeServiceMockRecorder) UpgradeInfo(arg0, arg1 any) *MockUpgradeServiceUpgradeInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeInfo", reflect.TypeOf((*MockUpgradeService)(nil).UpgradeInfo), arg0, arg1)
	return &MockUpgradeServiceUpgradeInfoCall{Call: call}
}

// MockUpgradeServiceUpgradeInfoCall wrap *gomock.Call
type MockUpgradeServiceUpgradeInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceUpgradeInfoCall) Return(arg0 upgrade.Info, arg1 error) *MockUpgradeServiceUpgradeInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceUpgradeInfoCall) Do(f func(context.Context, upgrade0.UUID) (upgrade.Info, error)) *MockUpgradeServiceUpgradeInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceUpgradeInfoCall) DoAndReturn(f func(context.Context, upgrade0.UUID) (upgrade.Info, error)) *MockUpgradeServiceUpgradeInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchForUpgradeState mocks base method.
func (m *MockUpgradeService) WatchForUpgradeState(arg0 context.Context, arg1 upgrade0.UUID, arg2 upgrade.State) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchForUpgradeState", arg0, arg1, arg2)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchForUpgradeState indicates an expected call of WatchForUpgradeState.
func (mr *MockUpgradeServiceMockRecorder) WatchForUpgradeState(arg0, arg1, arg2 any) *MockUpgradeServiceWatchForUpgradeStateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchForUpgradeState", reflect.TypeOf((*MockUpgradeService)(nil).WatchForUpgradeState), arg0, arg1, arg2)
	return &MockUpgradeServiceWatchForUpgradeStateCall{Call: call}
}

// MockUpgradeServiceWatchForUpgradeStateCall wrap *gomock.Call
type MockUpgradeServiceWatchForUpgradeStateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUpgradeServiceWatchForUpgradeStateCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockUpgradeServiceWatchForUpgradeStateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUpgradeServiceWatchForUpgradeStateCall) Do(f func(context.Context, upgrade0.UUID, upgrade.State) (watcher.Watcher[struct{}], error)) *MockUpgradeServiceWatchForUpgradeStateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUpgradeServiceWatchForUpgradeStateCall) DoAndReturn(f func(context.Context, upgrade0.UUID, upgrade.State) (watcher.Watcher[struct{}], error)) *MockUpgradeServiceWatchForUpgradeStateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
