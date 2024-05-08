// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/api/base (interfaces: APICallCloser)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/apibase_mock.go github.com/juju/juju/api/base APICallCloser
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	url "net/url"
	reflect "reflect"

	base "github.com/juju/juju/api/base"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
	httprequest "gopkg.in/httprequest.v1"
)

// MockAPICallCloser is a mock of APICallCloser interface.
type MockAPICallCloser struct {
	ctrl     *gomock.Controller
	recorder *MockAPICallCloserMockRecorder
}

// MockAPICallCloserMockRecorder is the mock recorder for MockAPICallCloser.
type MockAPICallCloserMockRecorder struct {
	mock *MockAPICallCloser
}

// NewMockAPICallCloser creates a new mock instance.
func NewMockAPICallCloser(ctrl *gomock.Controller) *MockAPICallCloser {
	mock := &MockAPICallCloser{ctrl: ctrl}
	mock.recorder = &MockAPICallCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPICallCloser) EXPECT() *MockAPICallCloserMockRecorder {
	return m.recorder
}

// APICall mocks base method.
func (m *MockAPICallCloser) APICall(arg0 context.Context, arg1 string, arg2 int, arg3, arg4 string, arg5, arg6 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APICall", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// APICall indicates an expected call of APICall.
func (mr *MockAPICallCloserMockRecorder) APICall(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *MockAPICallCloserAPICallCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APICall", reflect.TypeOf((*MockAPICallCloser)(nil).APICall), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	return &MockAPICallCloserAPICallCall{Call: call}
}

// MockAPICallCloserAPICallCall wrap *gomock.Call
type MockAPICallCloserAPICallCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallCloserAPICallCall) Return(arg0 error) *MockAPICallCloserAPICallCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallCloserAPICallCall) Do(f func(context.Context, string, int, string, string, any, any) error) *MockAPICallCloserAPICallCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallCloserAPICallCall) DoAndReturn(f func(context.Context, string, int, string, string, any, any) error) *MockAPICallCloserAPICallCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// BakeryClient mocks base method.
func (m *MockAPICallCloser) BakeryClient() base.MacaroonDischarger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BakeryClient")
	ret0, _ := ret[0].(base.MacaroonDischarger)
	return ret0
}

// BakeryClient indicates an expected call of BakeryClient.
func (mr *MockAPICallCloserMockRecorder) BakeryClient() *MockAPICallCloserBakeryClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BakeryClient", reflect.TypeOf((*MockAPICallCloser)(nil).BakeryClient))
	return &MockAPICallCloserBakeryClientCall{Call: call}
}

// MockAPICallCloserBakeryClientCall wrap *gomock.Call
type MockAPICallCloserBakeryClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallCloserBakeryClientCall) Return(arg0 base.MacaroonDischarger) *MockAPICallCloserBakeryClientCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallCloserBakeryClientCall) Do(f func() base.MacaroonDischarger) *MockAPICallCloserBakeryClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallCloserBakeryClientCall) DoAndReturn(f func() base.MacaroonDischarger) *MockAPICallCloserBakeryClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// BestFacadeVersion mocks base method.
func (m *MockAPICallCloser) BestFacadeVersion(arg0 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestFacadeVersion", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// BestFacadeVersion indicates an expected call of BestFacadeVersion.
func (mr *MockAPICallCloserMockRecorder) BestFacadeVersion(arg0 any) *MockAPICallCloserBestFacadeVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestFacadeVersion", reflect.TypeOf((*MockAPICallCloser)(nil).BestFacadeVersion), arg0)
	return &MockAPICallCloserBestFacadeVersionCall{Call: call}
}

// MockAPICallCloserBestFacadeVersionCall wrap *gomock.Call
type MockAPICallCloserBestFacadeVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallCloserBestFacadeVersionCall) Return(arg0 int) *MockAPICallCloserBestFacadeVersionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallCloserBestFacadeVersionCall) Do(f func(string) int) *MockAPICallCloserBestFacadeVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallCloserBestFacadeVersionCall) DoAndReturn(f func(string) int) *MockAPICallCloserBestFacadeVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockAPICallCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAPICallCloserMockRecorder) Close() *MockAPICallCloserCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAPICallCloser)(nil).Close))
	return &MockAPICallCloserCloseCall{Call: call}
}

// MockAPICallCloserCloseCall wrap *gomock.Call
type MockAPICallCloserCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallCloserCloseCall) Return(arg0 error) *MockAPICallCloserCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallCloserCloseCall) Do(f func() error) *MockAPICallCloserCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallCloserCloseCall) DoAndReturn(f func() error) *MockAPICallCloserCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConnectControllerStream mocks base method.
func (m *MockAPICallCloser) ConnectControllerStream(arg0 context.Context, arg1 string, arg2 url.Values, arg3 http.Header) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectControllerStream", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectControllerStream indicates an expected call of ConnectControllerStream.
func (mr *MockAPICallCloserMockRecorder) ConnectControllerStream(arg0, arg1, arg2, arg3 any) *MockAPICallCloserConnectControllerStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectControllerStream", reflect.TypeOf((*MockAPICallCloser)(nil).ConnectControllerStream), arg0, arg1, arg2, arg3)
	return &MockAPICallCloserConnectControllerStreamCall{Call: call}
}

// MockAPICallCloserConnectControllerStreamCall wrap *gomock.Call
type MockAPICallCloserConnectControllerStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallCloserConnectControllerStreamCall) Return(arg0 base.Stream, arg1 error) *MockAPICallCloserConnectControllerStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallCloserConnectControllerStreamCall) Do(f func(context.Context, string, url.Values, http.Header) (base.Stream, error)) *MockAPICallCloserConnectControllerStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallCloserConnectControllerStreamCall) DoAndReturn(f func(context.Context, string, url.Values, http.Header) (base.Stream, error)) *MockAPICallCloserConnectControllerStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConnectStream mocks base method.
func (m *MockAPICallCloser) ConnectStream(arg0 context.Context, arg1 string, arg2 url.Values) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectStream indicates an expected call of ConnectStream.
func (mr *MockAPICallCloserMockRecorder) ConnectStream(arg0, arg1, arg2 any) *MockAPICallCloserConnectStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectStream", reflect.TypeOf((*MockAPICallCloser)(nil).ConnectStream), arg0, arg1, arg2)
	return &MockAPICallCloserConnectStreamCall{Call: call}
}

// MockAPICallCloserConnectStreamCall wrap *gomock.Call
type MockAPICallCloserConnectStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallCloserConnectStreamCall) Return(arg0 base.Stream, arg1 error) *MockAPICallCloserConnectStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallCloserConnectStreamCall) Do(f func(context.Context, string, url.Values) (base.Stream, error)) *MockAPICallCloserConnectStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallCloserConnectStreamCall) DoAndReturn(f func(context.Context, string, url.Values) (base.Stream, error)) *MockAPICallCloserConnectStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HTTPClient mocks base method.
func (m *MockAPICallCloser) HTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockAPICallCloserMockRecorder) HTTPClient() *MockAPICallCloserHTTPClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockAPICallCloser)(nil).HTTPClient))
	return &MockAPICallCloserHTTPClientCall{Call: call}
}

// MockAPICallCloserHTTPClientCall wrap *gomock.Call
type MockAPICallCloserHTTPClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallCloserHTTPClientCall) Return(arg0 *httprequest.Client, arg1 error) *MockAPICallCloserHTTPClientCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallCloserHTTPClientCall) Do(f func() (*httprequest.Client, error)) *MockAPICallCloserHTTPClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallCloserHTTPClientCall) DoAndReturn(f func() (*httprequest.Client, error)) *MockAPICallCloserHTTPClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelTag mocks base method.
func (m *MockAPICallCloser) ModelTag() (names.ModelTag, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockAPICallCloserMockRecorder) ModelTag() *MockAPICallCloserModelTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockAPICallCloser)(nil).ModelTag))
	return &MockAPICallCloserModelTagCall{Call: call}
}

// MockAPICallCloserModelTagCall wrap *gomock.Call
type MockAPICallCloserModelTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallCloserModelTagCall) Return(arg0 names.ModelTag, arg1 bool) *MockAPICallCloserModelTagCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallCloserModelTagCall) Do(f func() (names.ModelTag, bool)) *MockAPICallCloserModelTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallCloserModelTagCall) DoAndReturn(f func() (names.ModelTag, bool)) *MockAPICallCloserModelTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RootHTTPClient mocks base method.
func (m *MockAPICallCloser) RootHTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RootHTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RootHTTPClient indicates an expected call of RootHTTPClient.
func (mr *MockAPICallCloserMockRecorder) RootHTTPClient() *MockAPICallCloserRootHTTPClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHTTPClient", reflect.TypeOf((*MockAPICallCloser)(nil).RootHTTPClient))
	return &MockAPICallCloserRootHTTPClientCall{Call: call}
}

// MockAPICallCloserRootHTTPClientCall wrap *gomock.Call
type MockAPICallCloserRootHTTPClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallCloserRootHTTPClientCall) Return(arg0 *httprequest.Client, arg1 error) *MockAPICallCloserRootHTTPClientCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallCloserRootHTTPClientCall) Do(f func() (*httprequest.Client, error)) *MockAPICallCloserRootHTTPClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallCloserRootHTTPClientCall) DoAndReturn(f func() (*httprequest.Client, error)) *MockAPICallCloserRootHTTPClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
