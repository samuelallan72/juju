// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/http (interfaces: RoundTripper)
//
// Generated by this command:
//
//	mockgen -typed -package http -destination http_mock_test.go github.com/juju/juju/internal/http RoundTripper
//

// Package http is a generated GoMock package.
package http

import (
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRoundTripper is a mock of RoundTripper interface.
type MockRoundTripper struct {
	ctrl     *gomock.Controller
	recorder *MockRoundTripperMockRecorder
}

// MockRoundTripperMockRecorder is the mock recorder for MockRoundTripper.
type MockRoundTripperMockRecorder struct {
	mock *MockRoundTripper
}

// NewMockRoundTripper creates a new mock instance.
func NewMockRoundTripper(ctrl *gomock.Controller) *MockRoundTripper {
	mock := &MockRoundTripper{ctrl: ctrl}
	mock.recorder = &MockRoundTripperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoundTripper) EXPECT() *MockRoundTripperMockRecorder {
	return m.recorder
}

// RoundTrip mocks base method.
func (m *MockRoundTripper) RoundTrip(arg0 *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoundTrip", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoundTrip indicates an expected call of RoundTrip.
func (mr *MockRoundTripperMockRecorder) RoundTrip(arg0 any) *MockRoundTripperRoundTripCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoundTrip", reflect.TypeOf((*MockRoundTripper)(nil).RoundTrip), arg0)
	return &MockRoundTripperRoundTripCall{Call: call}
}

// MockRoundTripperRoundTripCall wrap *gomock.Call
type MockRoundTripperRoundTripCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRoundTripperRoundTripCall) Return(arg0 *http.Response, arg1 error) *MockRoundTripperRoundTripCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRoundTripperRoundTripCall) Do(f func(*http.Request) (*http.Response, error)) *MockRoundTripperRoundTripCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRoundTripperRoundTripCall) DoAndReturn(f func(*http.Request) (*http.Response, error)) *MockRoundTripperRoundTripCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
