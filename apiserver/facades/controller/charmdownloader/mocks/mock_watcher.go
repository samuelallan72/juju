// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: StringsWatcher)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/mock_watcher.go github.com/juju/juju/state StringsWatcher
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStringsWatcher is a mock of StringsWatcher interface.
type MockStringsWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockStringsWatcherMockRecorder
}

// MockStringsWatcherMockRecorder is the mock recorder for MockStringsWatcher.
type MockStringsWatcherMockRecorder struct {
	mock *MockStringsWatcher
}

// NewMockStringsWatcher creates a new mock instance.
func NewMockStringsWatcher(ctrl *gomock.Controller) *MockStringsWatcher {
	mock := &MockStringsWatcher{ctrl: ctrl}
	mock.recorder = &MockStringsWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStringsWatcher) EXPECT() *MockStringsWatcherMockRecorder {
	return m.recorder
}

// Changes mocks base method.
func (m *MockStringsWatcher) Changes() <-chan []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes")
	ret0, _ := ret[0].(<-chan []string)
	return ret0
}

// Changes indicates an expected call of Changes.
func (mr *MockStringsWatcherMockRecorder) Changes() *MockStringsWatcherChangesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockStringsWatcher)(nil).Changes))
	return &MockStringsWatcherChangesCall{Call: call}
}

// MockStringsWatcherChangesCall wrap *gomock.Call
type MockStringsWatcherChangesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStringsWatcherChangesCall) Return(arg0 <-chan []string) *MockStringsWatcherChangesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStringsWatcherChangesCall) Do(f func() <-chan []string) *MockStringsWatcherChangesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStringsWatcherChangesCall) DoAndReturn(f func() <-chan []string) *MockStringsWatcherChangesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Err mocks base method.
func (m *MockStringsWatcher) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockStringsWatcherMockRecorder) Err() *MockStringsWatcherErrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockStringsWatcher)(nil).Err))
	return &MockStringsWatcherErrCall{Call: call}
}

// MockStringsWatcherErrCall wrap *gomock.Call
type MockStringsWatcherErrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStringsWatcherErrCall) Return(arg0 error) *MockStringsWatcherErrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStringsWatcherErrCall) Do(f func() error) *MockStringsWatcherErrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStringsWatcherErrCall) DoAndReturn(f func() error) *MockStringsWatcherErrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Kill mocks base method.
func (m *MockStringsWatcher) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockStringsWatcherMockRecorder) Kill() *MockStringsWatcherKillCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockStringsWatcher)(nil).Kill))
	return &MockStringsWatcherKillCall{Call: call}
}

// MockStringsWatcherKillCall wrap *gomock.Call
type MockStringsWatcherKillCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStringsWatcherKillCall) Return() *MockStringsWatcherKillCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStringsWatcherKillCall) Do(f func()) *MockStringsWatcherKillCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStringsWatcherKillCall) DoAndReturn(f func()) *MockStringsWatcherKillCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Stop mocks base method.
func (m *MockStringsWatcher) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockStringsWatcherMockRecorder) Stop() *MockStringsWatcherStopCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockStringsWatcher)(nil).Stop))
	return &MockStringsWatcherStopCall{Call: call}
}

// MockStringsWatcherStopCall wrap *gomock.Call
type MockStringsWatcherStopCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStringsWatcherStopCall) Return(arg0 error) *MockStringsWatcherStopCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStringsWatcherStopCall) Do(f func() error) *MockStringsWatcherStopCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStringsWatcherStopCall) DoAndReturn(f func() error) *MockStringsWatcherStopCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Wait mocks base method.
func (m *MockStringsWatcher) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockStringsWatcherMockRecorder) Wait() *MockStringsWatcherWaitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockStringsWatcher)(nil).Wait))
	return &MockStringsWatcherWaitCall{Call: call}
}

// MockStringsWatcherWaitCall wrap *gomock.Call
type MockStringsWatcherWaitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStringsWatcherWaitCall) Return(arg0 error) *MockStringsWatcherWaitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStringsWatcherWaitCall) Do(f func() error) *MockStringsWatcherWaitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStringsWatcherWaitCall) DoAndReturn(f func() error) *MockStringsWatcherWaitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
