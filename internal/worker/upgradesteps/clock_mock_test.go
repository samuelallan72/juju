// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/clock (interfaces: Clock)
//
// Generated by this command:
//
//	mockgen -typed -package upgradesteps -destination clock_mock_test.go github.com/juju/clock Clock
//

// Package upgradesteps is a generated GoMock package.
package upgradesteps

import (
	reflect "reflect"
	time "time"

	clock "github.com/juju/clock"
	gomock "go.uber.org/mock/gomock"
)

// MockClock is a mock of Clock interface.
type MockClock struct {
	ctrl     *gomock.Controller
	recorder *MockClockMockRecorder
}

// MockClockMockRecorder is the mock recorder for MockClock.
type MockClockMockRecorder struct {
	mock *MockClock
}

// NewMockClock creates a new mock instance.
func NewMockClock(ctrl *gomock.Controller) *MockClock {
	mock := &MockClock{ctrl: ctrl}
	mock.recorder = &MockClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClock) EXPECT() *MockClockMockRecorder {
	return m.recorder
}

// After mocks base method.
func (m *MockClock) After(arg0 time.Duration) <-chan time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "After", arg0)
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// After indicates an expected call of After.
func (mr *MockClockMockRecorder) After(arg0 any) *MockClockAfterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "After", reflect.TypeOf((*MockClock)(nil).After), arg0)
	return &MockClockAfterCall{Call: call}
}

// MockClockAfterCall wrap *gomock.Call
type MockClockAfterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClockAfterCall) Return(arg0 <-chan time.Time) *MockClockAfterCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClockAfterCall) Do(f func(time.Duration) <-chan time.Time) *MockClockAfterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClockAfterCall) DoAndReturn(f func(time.Duration) <-chan time.Time) *MockClockAfterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AfterFunc mocks base method.
func (m *MockClock) AfterFunc(arg0 time.Duration, arg1 func()) clock.Timer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterFunc", arg0, arg1)
	ret0, _ := ret[0].(clock.Timer)
	return ret0
}

// AfterFunc indicates an expected call of AfterFunc.
func (mr *MockClockMockRecorder) AfterFunc(arg0, arg1 any) *MockClockAfterFuncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterFunc", reflect.TypeOf((*MockClock)(nil).AfterFunc), arg0, arg1)
	return &MockClockAfterFuncCall{Call: call}
}

// MockClockAfterFuncCall wrap *gomock.Call
type MockClockAfterFuncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClockAfterFuncCall) Return(arg0 clock.Timer) *MockClockAfterFuncCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClockAfterFuncCall) Do(f func(time.Duration, func()) clock.Timer) *MockClockAfterFuncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClockAfterFuncCall) DoAndReturn(f func(time.Duration, func()) clock.Timer) *MockClockAfterFuncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NewTimer mocks base method.
func (m *MockClock) NewTimer(arg0 time.Duration) clock.Timer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTimer", arg0)
	ret0, _ := ret[0].(clock.Timer)
	return ret0
}

// NewTimer indicates an expected call of NewTimer.
func (mr *MockClockMockRecorder) NewTimer(arg0 any) *MockClockNewTimerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTimer", reflect.TypeOf((*MockClock)(nil).NewTimer), arg0)
	return &MockClockNewTimerCall{Call: call}
}

// MockClockNewTimerCall wrap *gomock.Call
type MockClockNewTimerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClockNewTimerCall) Return(arg0 clock.Timer) *MockClockNewTimerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClockNewTimerCall) Do(f func(time.Duration) clock.Timer) *MockClockNewTimerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClockNewTimerCall) DoAndReturn(f func(time.Duration) clock.Timer) *MockClockNewTimerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Now mocks base method.
func (m *MockClock) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now.
func (mr *MockClockMockRecorder) Now() *MockClockNowCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClock)(nil).Now))
	return &MockClockNowCall{Call: call}
}

// MockClockNowCall wrap *gomock.Call
type MockClockNowCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClockNowCall) Return(arg0 time.Time) *MockClockNowCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClockNowCall) Do(f func() time.Time) *MockClockNowCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClockNowCall) DoAndReturn(f func() time.Time) *MockClockNowCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
