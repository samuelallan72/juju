// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/changestream (interfaces: DBGetter,Logger,WatchableDBWorker,FileNotifyWatcher)
//
// Generated by this command:
//
//	mockgen -typed -package changestream -destination stream_mock_test.go github.com/juju/juju/internal/worker/changestream DBGetter,Logger,WatchableDBWorker,FileNotifyWatcher
//

// Package changestream is a generated GoMock package.
package changestream

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	sqlair "github.com/canonical/sqlair"
	changestream "github.com/juju/juju/core/changestream"
	database "github.com/juju/juju/core/database"
	gomock "go.uber.org/mock/gomock"
)

// MockDBGetter is a mock of DBGetter interface.
type MockDBGetter struct {
	ctrl     *gomock.Controller
	recorder *MockDBGetterMockRecorder
}

// MockDBGetterMockRecorder is the mock recorder for MockDBGetter.
type MockDBGetterMockRecorder struct {
	mock *MockDBGetter
}

// NewMockDBGetter creates a new mock instance.
func NewMockDBGetter(ctrl *gomock.Controller) *MockDBGetter {
	mock := &MockDBGetter{ctrl: ctrl}
	mock.recorder = &MockDBGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBGetter) EXPECT() *MockDBGetterMockRecorder {
	return m.recorder
}

// GetDB mocks base method.
func (m *MockDBGetter) GetDB(arg0 string) (database.TxnRunner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDB", arg0)
	ret0, _ := ret[0].(database.TxnRunner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDB indicates an expected call of GetDB.
func (mr *MockDBGetterMockRecorder) GetDB(arg0 any) *MockDBGetterGetDBCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDB", reflect.TypeOf((*MockDBGetter)(nil).GetDB), arg0)
	return &MockDBGetterGetDBCall{Call: call}
}

// MockDBGetterGetDBCall wrap *gomock.Call
type MockDBGetterGetDBCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDBGetterGetDBCall) Return(arg0 database.TxnRunner, arg1 error) *MockDBGetterGetDBCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDBGetterGetDBCall) Do(f func(string) (database.TxnRunner, error)) *MockDBGetterGetDBCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDBGetterGetDBCall) DoAndReturn(f func(string) (database.TxnRunner, error)) *MockDBGetterGetDBCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(arg0 any, arg1 ...any) *MockLoggerDebugfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
	return &MockLoggerDebugfCall{Call: call}
}

// MockLoggerDebugfCall wrap *gomock.Call
type MockLoggerDebugfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerDebugfCall) Return() *MockLoggerDebugfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerDebugfCall) Do(f func(string, ...any)) *MockLoggerDebugfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerDebugfCall) DoAndReturn(f func(string, ...any)) *MockLoggerDebugfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(arg0 any, arg1 ...any) *MockLoggerErrorfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
	return &MockLoggerErrorfCall{Call: call}
}

// MockLoggerErrorfCall wrap *gomock.Call
type MockLoggerErrorfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerErrorfCall) Return() *MockLoggerErrorfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerErrorfCall) Do(f func(string, ...any)) *MockLoggerErrorfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerErrorfCall) DoAndReturn(f func(string, ...any)) *MockLoggerErrorfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Infof mocks base method.
func (m *MockLogger) Infof(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(arg0 any, arg1 ...any) *MockLoggerInfofCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
	return &MockLoggerInfofCall{Call: call}
}

// MockLoggerInfofCall wrap *gomock.Call
type MockLoggerInfofCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerInfofCall) Return() *MockLoggerInfofCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerInfofCall) Do(f func(string, ...any)) *MockLoggerInfofCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerInfofCall) DoAndReturn(f func(string, ...any)) *MockLoggerInfofCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsTraceEnabled mocks base method.
func (m *MockLogger) IsTraceEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTraceEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTraceEnabled indicates an expected call of IsTraceEnabled.
func (mr *MockLoggerMockRecorder) IsTraceEnabled() *MockLoggerIsTraceEnabledCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTraceEnabled", reflect.TypeOf((*MockLogger)(nil).IsTraceEnabled))
	return &MockLoggerIsTraceEnabledCall{Call: call}
}

// MockLoggerIsTraceEnabledCall wrap *gomock.Call
type MockLoggerIsTraceEnabledCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerIsTraceEnabledCall) Return(arg0 bool) *MockLoggerIsTraceEnabledCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerIsTraceEnabledCall) Do(f func() bool) *MockLoggerIsTraceEnabledCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerIsTraceEnabledCall) DoAndReturn(f func() bool) *MockLoggerIsTraceEnabledCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Tracef mocks base method.
func (m *MockLogger) Tracef(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockLoggerMockRecorder) Tracef(arg0 any, arg1 ...any) *MockLoggerTracefCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockLogger)(nil).Tracef), varargs...)
	return &MockLoggerTracefCall{Call: call}
}

// MockLoggerTracefCall wrap *gomock.Call
type MockLoggerTracefCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerTracefCall) Return() *MockLoggerTracefCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerTracefCall) Do(f func(string, ...any)) *MockLoggerTracefCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerTracefCall) DoAndReturn(f func(string, ...any)) *MockLoggerTracefCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Warningf mocks base method.
func (m *MockLogger) Warningf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf.
func (mr *MockLoggerMockRecorder) Warningf(arg0 any, arg1 ...any) *MockLoggerWarningfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockLogger)(nil).Warningf), varargs...)
	return &MockLoggerWarningfCall{Call: call}
}

// MockLoggerWarningfCall wrap *gomock.Call
type MockLoggerWarningfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerWarningfCall) Return() *MockLoggerWarningfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerWarningfCall) Do(f func(string, ...any)) *MockLoggerWarningfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerWarningfCall) DoAndReturn(f func(string, ...any)) *MockLoggerWarningfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockWatchableDBWorker is a mock of WatchableDBWorker interface.
type MockWatchableDBWorker struct {
	ctrl     *gomock.Controller
	recorder *MockWatchableDBWorkerMockRecorder
}

// MockWatchableDBWorkerMockRecorder is the mock recorder for MockWatchableDBWorker.
type MockWatchableDBWorkerMockRecorder struct {
	mock *MockWatchableDBWorker
}

// NewMockWatchableDBWorker creates a new mock instance.
func NewMockWatchableDBWorker(ctrl *gomock.Controller) *MockWatchableDBWorker {
	mock := &MockWatchableDBWorker{ctrl: ctrl}
	mock.recorder = &MockWatchableDBWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatchableDBWorker) EXPECT() *MockWatchableDBWorkerMockRecorder {
	return m.recorder
}

// Kill mocks base method.
func (m *MockWatchableDBWorker) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockWatchableDBWorkerMockRecorder) Kill() *MockWatchableDBWorkerKillCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockWatchableDBWorker)(nil).Kill))
	return &MockWatchableDBWorkerKillCall{Call: call}
}

// MockWatchableDBWorkerKillCall wrap *gomock.Call
type MockWatchableDBWorkerKillCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatchableDBWorkerKillCall) Return() *MockWatchableDBWorkerKillCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatchableDBWorkerKillCall) Do(f func()) *MockWatchableDBWorkerKillCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatchableDBWorkerKillCall) DoAndReturn(f func()) *MockWatchableDBWorkerKillCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StdTxn mocks base method.
func (m *MockWatchableDBWorker) StdTxn(arg0 context.Context, arg1 func(context.Context, *sql.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StdTxn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StdTxn indicates an expected call of StdTxn.
func (mr *MockWatchableDBWorkerMockRecorder) StdTxn(arg0, arg1 any) *MockWatchableDBWorkerStdTxnCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StdTxn", reflect.TypeOf((*MockWatchableDBWorker)(nil).StdTxn), arg0, arg1)
	return &MockWatchableDBWorkerStdTxnCall{Call: call}
}

// MockWatchableDBWorkerStdTxnCall wrap *gomock.Call
type MockWatchableDBWorkerStdTxnCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatchableDBWorkerStdTxnCall) Return(arg0 error) *MockWatchableDBWorkerStdTxnCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatchableDBWorkerStdTxnCall) Do(f func(context.Context, func(context.Context, *sql.Tx) error) error) *MockWatchableDBWorkerStdTxnCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatchableDBWorkerStdTxnCall) DoAndReturn(f func(context.Context, func(context.Context, *sql.Tx) error) error) *MockWatchableDBWorkerStdTxnCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Subscribe mocks base method.
func (m *MockWatchableDBWorker) Subscribe(arg0 ...changestream.SubscriptionOption) (changestream.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(changestream.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockWatchableDBWorkerMockRecorder) Subscribe(arg0 ...any) *MockWatchableDBWorkerSubscribeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockWatchableDBWorker)(nil).Subscribe), arg0...)
	return &MockWatchableDBWorkerSubscribeCall{Call: call}
}

// MockWatchableDBWorkerSubscribeCall wrap *gomock.Call
type MockWatchableDBWorkerSubscribeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatchableDBWorkerSubscribeCall) Return(arg0 changestream.Subscription, arg1 error) *MockWatchableDBWorkerSubscribeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatchableDBWorkerSubscribeCall) Do(f func(...changestream.SubscriptionOption) (changestream.Subscription, error)) *MockWatchableDBWorkerSubscribeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatchableDBWorkerSubscribeCall) DoAndReturn(f func(...changestream.SubscriptionOption) (changestream.Subscription, error)) *MockWatchableDBWorkerSubscribeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Txn mocks base method.
func (m *MockWatchableDBWorker) Txn(arg0 context.Context, arg1 func(context.Context, *sqlair.TX) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Txn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Txn indicates an expected call of Txn.
func (mr *MockWatchableDBWorkerMockRecorder) Txn(arg0, arg1 any) *MockWatchableDBWorkerTxnCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Txn", reflect.TypeOf((*MockWatchableDBWorker)(nil).Txn), arg0, arg1)
	return &MockWatchableDBWorkerTxnCall{Call: call}
}

// MockWatchableDBWorkerTxnCall wrap *gomock.Call
type MockWatchableDBWorkerTxnCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatchableDBWorkerTxnCall) Return(arg0 error) *MockWatchableDBWorkerTxnCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatchableDBWorkerTxnCall) Do(f func(context.Context, func(context.Context, *sqlair.TX) error) error) *MockWatchableDBWorkerTxnCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatchableDBWorkerTxnCall) DoAndReturn(f func(context.Context, func(context.Context, *sqlair.TX) error) error) *MockWatchableDBWorkerTxnCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Wait mocks base method.
func (m *MockWatchableDBWorker) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockWatchableDBWorkerMockRecorder) Wait() *MockWatchableDBWorkerWaitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockWatchableDBWorker)(nil).Wait))
	return &MockWatchableDBWorkerWaitCall{Call: call}
}

// MockWatchableDBWorkerWaitCall wrap *gomock.Call
type MockWatchableDBWorkerWaitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatchableDBWorkerWaitCall) Return(arg0 error) *MockWatchableDBWorkerWaitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatchableDBWorkerWaitCall) Do(f func() error) *MockWatchableDBWorkerWaitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatchableDBWorkerWaitCall) DoAndReturn(f func() error) *MockWatchableDBWorkerWaitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockFileNotifyWatcher is a mock of FileNotifyWatcher interface.
type MockFileNotifyWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockFileNotifyWatcherMockRecorder
}

// MockFileNotifyWatcherMockRecorder is the mock recorder for MockFileNotifyWatcher.
type MockFileNotifyWatcherMockRecorder struct {
	mock *MockFileNotifyWatcher
}

// NewMockFileNotifyWatcher creates a new mock instance.
func NewMockFileNotifyWatcher(ctrl *gomock.Controller) *MockFileNotifyWatcher {
	mock := &MockFileNotifyWatcher{ctrl: ctrl}
	mock.recorder = &MockFileNotifyWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileNotifyWatcher) EXPECT() *MockFileNotifyWatcherMockRecorder {
	return m.recorder
}

// Changes mocks base method.
func (m *MockFileNotifyWatcher) Changes(arg0 string) (<-chan bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes", arg0)
	ret0, _ := ret[0].(<-chan bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Changes indicates an expected call of Changes.
func (mr *MockFileNotifyWatcherMockRecorder) Changes(arg0 any) *MockFileNotifyWatcherChangesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockFileNotifyWatcher)(nil).Changes), arg0)
	return &MockFileNotifyWatcherChangesCall{Call: call}
}

// MockFileNotifyWatcherChangesCall wrap *gomock.Call
type MockFileNotifyWatcherChangesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFileNotifyWatcherChangesCall) Return(arg0 <-chan bool, arg1 error) *MockFileNotifyWatcherChangesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFileNotifyWatcherChangesCall) Do(f func(string) (<-chan bool, error)) *MockFileNotifyWatcherChangesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFileNotifyWatcherChangesCall) DoAndReturn(f func(string) (<-chan bool, error)) *MockFileNotifyWatcherChangesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
