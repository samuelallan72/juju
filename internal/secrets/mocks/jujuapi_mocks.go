// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/secrets (interfaces: JujuAPIClient)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/jujuapi_mocks.go github.com/juju/juju/internal/secrets JujuAPIClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	secrets0 "github.com/juju/juju/internal/secrets"
	provider "github.com/juju/juju/internal/secrets/provider"
	gomock "go.uber.org/mock/gomock"
)

// MockJujuAPIClient is a mock of JujuAPIClient interface.
type MockJujuAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockJujuAPIClientMockRecorder
}

// MockJujuAPIClientMockRecorder is the mock recorder for MockJujuAPIClient.
type MockJujuAPIClientMockRecorder struct {
	mock *MockJujuAPIClient
}

// NewMockJujuAPIClient creates a new mock instance.
func NewMockJujuAPIClient(ctrl *gomock.Controller) *MockJujuAPIClient {
	mock := &MockJujuAPIClient{ctrl: ctrl}
	mock.recorder = &MockJujuAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJujuAPIClient) EXPECT() *MockJujuAPIClientMockRecorder {
	return m.recorder
}

// GetBackendConfigForDrain mocks base method.
func (m *MockJujuAPIClient) GetBackendConfigForDrain(arg0 *string) (*provider.ModelBackendConfig, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackendConfigForDrain", arg0)
	ret0, _ := ret[0].(*provider.ModelBackendConfig)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBackendConfigForDrain indicates an expected call of GetBackendConfigForDrain.
func (mr *MockJujuAPIClientMockRecorder) GetBackendConfigForDrain(arg0 any) *MockJujuAPIClientGetBackendConfigForDrainCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendConfigForDrain", reflect.TypeOf((*MockJujuAPIClient)(nil).GetBackendConfigForDrain), arg0)
	return &MockJujuAPIClientGetBackendConfigForDrainCall{Call: call}
}

// MockJujuAPIClientGetBackendConfigForDrainCall wrap *gomock.Call
type MockJujuAPIClientGetBackendConfigForDrainCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJujuAPIClientGetBackendConfigForDrainCall) Return(arg0 *provider.ModelBackendConfig, arg1 string, arg2 error) *MockJujuAPIClientGetBackendConfigForDrainCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJujuAPIClientGetBackendConfigForDrainCall) Do(f func(*string) (*provider.ModelBackendConfig, string, error)) *MockJujuAPIClientGetBackendConfigForDrainCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJujuAPIClientGetBackendConfigForDrainCall) DoAndReturn(f func(*string) (*provider.ModelBackendConfig, string, error)) *MockJujuAPIClientGetBackendConfigForDrainCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetContentInfo mocks base method.
func (m *MockJujuAPIClient) GetContentInfo(arg0 *secrets.URI, arg1 string, arg2, arg3 bool) (*secrets0.ContentParams, *provider.ModelBackendConfig, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContentInfo", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*secrets0.ContentParams)
	ret1, _ := ret[1].(*provider.ModelBackendConfig)
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetContentInfo indicates an expected call of GetContentInfo.
func (mr *MockJujuAPIClientMockRecorder) GetContentInfo(arg0, arg1, arg2, arg3 any) *MockJujuAPIClientGetContentInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContentInfo", reflect.TypeOf((*MockJujuAPIClient)(nil).GetContentInfo), arg0, arg1, arg2, arg3)
	return &MockJujuAPIClientGetContentInfoCall{Call: call}
}

// MockJujuAPIClientGetContentInfoCall wrap *gomock.Call
type MockJujuAPIClientGetContentInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJujuAPIClientGetContentInfoCall) Return(arg0 *secrets0.ContentParams, arg1 *provider.ModelBackendConfig, arg2 bool, arg3 error) *MockJujuAPIClientGetContentInfoCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJujuAPIClientGetContentInfoCall) Do(f func(*secrets.URI, string, bool, bool) (*secrets0.ContentParams, *provider.ModelBackendConfig, bool, error)) *MockJujuAPIClientGetContentInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJujuAPIClientGetContentInfoCall) DoAndReturn(f func(*secrets.URI, string, bool, bool) (*secrets0.ContentParams, *provider.ModelBackendConfig, bool, error)) *MockJujuAPIClientGetContentInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetRevisionContentInfo mocks base method.
func (m *MockJujuAPIClient) GetRevisionContentInfo(arg0 *secrets.URI, arg1 int, arg2 bool) (*secrets0.ContentParams, *provider.ModelBackendConfig, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRevisionContentInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(*secrets0.ContentParams)
	ret1, _ := ret[1].(*provider.ModelBackendConfig)
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetRevisionContentInfo indicates an expected call of GetRevisionContentInfo.
func (mr *MockJujuAPIClientMockRecorder) GetRevisionContentInfo(arg0, arg1, arg2 any) *MockJujuAPIClientGetRevisionContentInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRevisionContentInfo", reflect.TypeOf((*MockJujuAPIClient)(nil).GetRevisionContentInfo), arg0, arg1, arg2)
	return &MockJujuAPIClientGetRevisionContentInfoCall{Call: call}
}

// MockJujuAPIClientGetRevisionContentInfoCall wrap *gomock.Call
type MockJujuAPIClientGetRevisionContentInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJujuAPIClientGetRevisionContentInfoCall) Return(arg0 *secrets0.ContentParams, arg1 *provider.ModelBackendConfig, arg2 bool, arg3 error) *MockJujuAPIClientGetRevisionContentInfoCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJujuAPIClientGetRevisionContentInfoCall) Do(f func(*secrets.URI, int, bool) (*secrets0.ContentParams, *provider.ModelBackendConfig, bool, error)) *MockJujuAPIClientGetRevisionContentInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJujuAPIClientGetRevisionContentInfoCall) DoAndReturn(f func(*secrets.URI, int, bool) (*secrets0.ContentParams, *provider.ModelBackendConfig, bool, error)) *MockJujuAPIClientGetRevisionContentInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSecretBackendConfig mocks base method.
func (m *MockJujuAPIClient) GetSecretBackendConfig(arg0 *string) (*provider.ModelBackendConfigInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretBackendConfig", arg0)
	ret0, _ := ret[0].(*provider.ModelBackendConfigInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretBackendConfig indicates an expected call of GetSecretBackendConfig.
func (mr *MockJujuAPIClientMockRecorder) GetSecretBackendConfig(arg0 any) *MockJujuAPIClientGetSecretBackendConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretBackendConfig", reflect.TypeOf((*MockJujuAPIClient)(nil).GetSecretBackendConfig), arg0)
	return &MockJujuAPIClientGetSecretBackendConfigCall{Call: call}
}

// MockJujuAPIClientGetSecretBackendConfigCall wrap *gomock.Call
type MockJujuAPIClientGetSecretBackendConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJujuAPIClientGetSecretBackendConfigCall) Return(arg0 *provider.ModelBackendConfigInfo, arg1 error) *MockJujuAPIClientGetSecretBackendConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJujuAPIClientGetSecretBackendConfigCall) Do(f func(*string) (*provider.ModelBackendConfigInfo, error)) *MockJujuAPIClientGetSecretBackendConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJujuAPIClientGetSecretBackendConfigCall) DoAndReturn(f func(*string) (*provider.ModelBackendConfigInfo, error)) *MockJujuAPIClientGetSecretBackendConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
