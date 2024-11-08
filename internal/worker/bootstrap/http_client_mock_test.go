// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/http (interfaces: HTTPClientGetter)
//
// Generated by this command:
//
//	mockgen -typed -package bootstrap -destination http_client_mock_test.go github.com/juju/juju/core/http HTTPClientGetter
//

// Package bootstrap is a generated GoMock package.
package bootstrap

import (
	context "context"
	reflect "reflect"

	http "github.com/juju/juju/core/http"
	gomock "go.uber.org/mock/gomock"
)

// MockHTTPClientGetter is a mock of HTTPClientGetter interface.
type MockHTTPClientGetter struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientGetterMockRecorder
}

// MockHTTPClientGetterMockRecorder is the mock recorder for MockHTTPClientGetter.
type MockHTTPClientGetterMockRecorder struct {
	mock *MockHTTPClientGetter
}

// NewMockHTTPClientGetter creates a new mock instance.
func NewMockHTTPClientGetter(ctrl *gomock.Controller) *MockHTTPClientGetter {
	mock := &MockHTTPClientGetter{ctrl: ctrl}
	mock.recorder = &MockHTTPClientGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClientGetter) EXPECT() *MockHTTPClientGetterMockRecorder {
	return m.recorder
}

// GetHTTPClient mocks base method.
func (m *MockHTTPClientGetter) GetHTTPClient(arg0 context.Context, arg1 http.Purpose) (http.HTTPClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPClient", arg0, arg1)
	ret0, _ := ret[0].(http.HTTPClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHTTPClient indicates an expected call of GetHTTPClient.
func (mr *MockHTTPClientGetterMockRecorder) GetHTTPClient(arg0, arg1 any) *MockHTTPClientGetterGetHTTPClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPClient", reflect.TypeOf((*MockHTTPClientGetter)(nil).GetHTTPClient), arg0, arg1)
	return &MockHTTPClientGetterGetHTTPClientCall{Call: call}
}

// MockHTTPClientGetterGetHTTPClientCall wrap *gomock.Call
type MockHTTPClientGetterGetHTTPClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockHTTPClientGetterGetHTTPClientCall) Return(arg0 http.HTTPClient, arg1 error) *MockHTTPClientGetterGetHTTPClientCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockHTTPClientGetterGetHTTPClientCall) Do(f func(context.Context, http.Purpose) (http.HTTPClient, error)) *MockHTTPClientGetterGetHTTPClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockHTTPClientGetterGetHTTPClientCall) DoAndReturn(f func(context.Context, http.Purpose) (http.HTTPClient, error)) *MockHTTPClientGetterGetHTTPClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
