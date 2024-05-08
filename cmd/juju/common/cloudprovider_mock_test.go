// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/common (interfaces: TestCloudProvider)
//
// Generated by this command:
//
//	mockgen -typed -package common -destination cloudprovider_mock_test.go github.com/juju/juju/cmd/juju/common TestCloudProvider
//

// Package common is a generated GoMock package.
package common

import (
	context "context"
	reflect "reflect"

	jsonschema "github.com/juju/jsonschema"
	cloud "github.com/juju/juju/cloud"
	environs "github.com/juju/juju/environs"
	config "github.com/juju/juju/environs/config"
	envcontext "github.com/juju/juju/environs/envcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockTestCloudProvider is a mock of TestCloudProvider interface.
type MockTestCloudProvider struct {
	ctrl     *gomock.Controller
	recorder *MockTestCloudProviderMockRecorder
}

// MockTestCloudProviderMockRecorder is the mock recorder for MockTestCloudProvider.
type MockTestCloudProviderMockRecorder struct {
	mock *MockTestCloudProvider
}

// NewMockTestCloudProvider creates a new mock instance.
func NewMockTestCloudProvider(ctrl *gomock.Controller) *MockTestCloudProvider {
	mock := &MockTestCloudProvider{ctrl: ctrl}
	mock.recorder = &MockTestCloudProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestCloudProvider) EXPECT() *MockTestCloudProviderMockRecorder {
	return m.recorder
}

// CloudSchema mocks base method.
func (m *MockTestCloudProvider) CloudSchema() *jsonschema.Schema {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudSchema")
	ret0, _ := ret[0].(*jsonschema.Schema)
	return ret0
}

// CloudSchema indicates an expected call of CloudSchema.
func (mr *MockTestCloudProviderMockRecorder) CloudSchema() *MockTestCloudProviderCloudSchemaCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudSchema", reflect.TypeOf((*MockTestCloudProvider)(nil).CloudSchema))
	return &MockTestCloudProviderCloudSchemaCall{Call: call}
}

// MockTestCloudProviderCloudSchemaCall wrap *gomock.Call
type MockTestCloudProviderCloudSchemaCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTestCloudProviderCloudSchemaCall) Return(arg0 *jsonschema.Schema) *MockTestCloudProviderCloudSchemaCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTestCloudProviderCloudSchemaCall) Do(f func() *jsonschema.Schema) *MockTestCloudProviderCloudSchemaCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTestCloudProviderCloudSchemaCall) DoAndReturn(f func() *jsonschema.Schema) *MockTestCloudProviderCloudSchemaCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CredentialSchemas mocks base method.
func (m *MockTestCloudProvider) CredentialSchemas() map[cloud.AuthType]cloud.CredentialSchema {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialSchemas")
	ret0, _ := ret[0].(map[cloud.AuthType]cloud.CredentialSchema)
	return ret0
}

// CredentialSchemas indicates an expected call of CredentialSchemas.
func (mr *MockTestCloudProviderMockRecorder) CredentialSchemas() *MockTestCloudProviderCredentialSchemasCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialSchemas", reflect.TypeOf((*MockTestCloudProvider)(nil).CredentialSchemas))
	return &MockTestCloudProviderCredentialSchemasCall{Call: call}
}

// MockTestCloudProviderCredentialSchemasCall wrap *gomock.Call
type MockTestCloudProviderCredentialSchemasCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTestCloudProviderCredentialSchemasCall) Return(arg0 map[cloud.AuthType]cloud.CredentialSchema) *MockTestCloudProviderCredentialSchemasCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTestCloudProviderCredentialSchemasCall) Do(f func() map[cloud.AuthType]cloud.CredentialSchema) *MockTestCloudProviderCredentialSchemasCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTestCloudProviderCredentialSchemasCall) DoAndReturn(f func() map[cloud.AuthType]cloud.CredentialSchema) *MockTestCloudProviderCredentialSchemasCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DetectCredentials mocks base method.
func (m *MockTestCloudProvider) DetectCredentials(arg0 string) (*cloud.CloudCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectCredentials", arg0)
	ret0, _ := ret[0].(*cloud.CloudCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectCredentials indicates an expected call of DetectCredentials.
func (mr *MockTestCloudProviderMockRecorder) DetectCredentials(arg0 any) *MockTestCloudProviderDetectCredentialsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectCredentials", reflect.TypeOf((*MockTestCloudProvider)(nil).DetectCredentials), arg0)
	return &MockTestCloudProviderDetectCredentialsCall{Call: call}
}

// MockTestCloudProviderDetectCredentialsCall wrap *gomock.Call
type MockTestCloudProviderDetectCredentialsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTestCloudProviderDetectCredentialsCall) Return(arg0 *cloud.CloudCredential, arg1 error) *MockTestCloudProviderDetectCredentialsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTestCloudProviderDetectCredentialsCall) Do(f func(string) (*cloud.CloudCredential, error)) *MockTestCloudProviderDetectCredentialsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTestCloudProviderDetectCredentialsCall) DoAndReturn(f func(string) (*cloud.CloudCredential, error)) *MockTestCloudProviderDetectCredentialsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FinalizeCredential mocks base method.
func (m *MockTestCloudProvider) FinalizeCredential(arg0 environs.FinalizeCredentialContext, arg1 environs.FinalizeCredentialParams) (*cloud.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeCredential", arg0, arg1)
	ret0, _ := ret[0].(*cloud.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FinalizeCredential indicates an expected call of FinalizeCredential.
func (mr *MockTestCloudProviderMockRecorder) FinalizeCredential(arg0, arg1 any) *MockTestCloudProviderFinalizeCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeCredential", reflect.TypeOf((*MockTestCloudProvider)(nil).FinalizeCredential), arg0, arg1)
	return &MockTestCloudProviderFinalizeCredentialCall{Call: call}
}

// MockTestCloudProviderFinalizeCredentialCall wrap *gomock.Call
type MockTestCloudProviderFinalizeCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTestCloudProviderFinalizeCredentialCall) Return(arg0 *cloud.Credential, arg1 error) *MockTestCloudProviderFinalizeCredentialCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTestCloudProviderFinalizeCredentialCall) Do(f func(environs.FinalizeCredentialContext, environs.FinalizeCredentialParams) (*cloud.Credential, error)) *MockTestCloudProviderFinalizeCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTestCloudProviderFinalizeCredentialCall) DoAndReturn(f func(environs.FinalizeCredentialContext, environs.FinalizeCredentialParams) (*cloud.Credential, error)) *MockTestCloudProviderFinalizeCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Ping mocks base method.
func (m *MockTestCloudProvider) Ping(arg0 envcontext.ProviderCallContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockTestCloudProviderMockRecorder) Ping(arg0, arg1 any) *MockTestCloudProviderPingCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockTestCloudProvider)(nil).Ping), arg0, arg1)
	return &MockTestCloudProviderPingCall{Call: call}
}

// MockTestCloudProviderPingCall wrap *gomock.Call
type MockTestCloudProviderPingCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTestCloudProviderPingCall) Return(arg0 error) *MockTestCloudProviderPingCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTestCloudProviderPingCall) Do(f func(envcontext.ProviderCallContext, string) error) *MockTestCloudProviderPingCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTestCloudProviderPingCall) DoAndReturn(f func(envcontext.ProviderCallContext, string) error) *MockTestCloudProviderPingCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PrepareConfig mocks base method.
func (m *MockTestCloudProvider) PrepareConfig(arg0 context.Context, arg1 environs.PrepareConfigParams) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareConfig", arg0, arg1)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareConfig indicates an expected call of PrepareConfig.
func (mr *MockTestCloudProviderMockRecorder) PrepareConfig(arg0, arg1 any) *MockTestCloudProviderPrepareConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareConfig", reflect.TypeOf((*MockTestCloudProvider)(nil).PrepareConfig), arg0, arg1)
	return &MockTestCloudProviderPrepareConfigCall{Call: call}
}

// MockTestCloudProviderPrepareConfigCall wrap *gomock.Call
type MockTestCloudProviderPrepareConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTestCloudProviderPrepareConfigCall) Return(arg0 *config.Config, arg1 error) *MockTestCloudProviderPrepareConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTestCloudProviderPrepareConfigCall) Do(f func(context.Context, environs.PrepareConfigParams) (*config.Config, error)) *MockTestCloudProviderPrepareConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTestCloudProviderPrepareConfigCall) DoAndReturn(f func(context.Context, environs.PrepareConfigParams) (*config.Config, error)) *MockTestCloudProviderPrepareConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RegisterCredentials mocks base method.
func (m *MockTestCloudProvider) RegisterCredentials(arg0 cloud.Cloud) (map[string]*cloud.CloudCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterCredentials", arg0)
	ret0, _ := ret[0].(map[string]*cloud.CloudCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterCredentials indicates an expected call of RegisterCredentials.
func (mr *MockTestCloudProviderMockRecorder) RegisterCredentials(arg0 any) *MockTestCloudProviderRegisterCredentialsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCredentials", reflect.TypeOf((*MockTestCloudProvider)(nil).RegisterCredentials), arg0)
	return &MockTestCloudProviderRegisterCredentialsCall{Call: call}
}

// MockTestCloudProviderRegisterCredentialsCall wrap *gomock.Call
type MockTestCloudProviderRegisterCredentialsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTestCloudProviderRegisterCredentialsCall) Return(arg0 map[string]*cloud.CloudCredential, arg1 error) *MockTestCloudProviderRegisterCredentialsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTestCloudProviderRegisterCredentialsCall) Do(f func(cloud.Cloud) (map[string]*cloud.CloudCredential, error)) *MockTestCloudProviderRegisterCredentialsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTestCloudProviderRegisterCredentialsCall) DoAndReturn(f func(cloud.Cloud) (map[string]*cloud.CloudCredential, error)) *MockTestCloudProviderRegisterCredentialsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Validate mocks base method.
func (m *MockTestCloudProvider) Validate(arg0 context.Context, arg1, arg2 *config.Config) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1, arg2)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockTestCloudProviderMockRecorder) Validate(arg0, arg1, arg2 any) *MockTestCloudProviderValidateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTestCloudProvider)(nil).Validate), arg0, arg1, arg2)
	return &MockTestCloudProviderValidateCall{Call: call}
}

// MockTestCloudProviderValidateCall wrap *gomock.Call
type MockTestCloudProviderValidateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTestCloudProviderValidateCall) Return(arg0 *config.Config, arg1 error) *MockTestCloudProviderValidateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTestCloudProviderValidateCall) Do(f func(context.Context, *config.Config, *config.Config) (*config.Config, error)) *MockTestCloudProviderValidateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTestCloudProviderValidateCall) DoAndReturn(f func(context.Context, *config.Config, *config.Config) (*config.Config, error)) *MockTestCloudProviderValidateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Version mocks base method.
func (m *MockTestCloudProvider) Version() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(int)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockTestCloudProviderMockRecorder) Version() *MockTestCloudProviderVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockTestCloudProvider)(nil).Version))
	return &MockTestCloudProviderVersionCall{Call: call}
}

// MockTestCloudProviderVersionCall wrap *gomock.Call
type MockTestCloudProviderVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTestCloudProviderVersionCall) Return(arg0 int) *MockTestCloudProviderVersionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTestCloudProviderVersionCall) Do(f func() int) *MockTestCloudProviderVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTestCloudProviderVersionCall) DoAndReturn(f func() int) *MockTestCloudProviderVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
