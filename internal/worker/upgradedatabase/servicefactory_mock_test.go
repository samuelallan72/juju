// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/servicefactory (interfaces: ControllerServiceFactory)
//
// Generated by this command:
//
//	mockgen -typed -package upgradedatabase -destination servicefactory_mock_test.go github.com/juju/juju/internal/servicefactory ControllerServiceFactory
//

// Package upgradedatabase is a generated GoMock package.
package upgradedatabase

import (
	reflect "reflect"

	service "github.com/juju/juju/domain/access/service"
	service0 "github.com/juju/juju/domain/autocert/service"
	service1 "github.com/juju/juju/domain/cloud/service"
	service2 "github.com/juju/juju/domain/controllerconfig/service"
	service3 "github.com/juju/juju/domain/controllernode/service"
	service4 "github.com/juju/juju/domain/credential/service"
	service5 "github.com/juju/juju/domain/externalcontroller/service"
	service6 "github.com/juju/juju/domain/flag/service"
	service7 "github.com/juju/juju/domain/model/service"
	service8 "github.com/juju/juju/domain/modeldefaults/service"
	service9 "github.com/juju/juju/domain/objectstore/service"
	service10 "github.com/juju/juju/domain/secretbackend/service"
	service11 "github.com/juju/juju/domain/upgrade/service"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerServiceFactory is a mock of ControllerServiceFactory interface.
type MockControllerServiceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockControllerServiceFactoryMockRecorder
}

// MockControllerServiceFactoryMockRecorder is the mock recorder for MockControllerServiceFactory.
type MockControllerServiceFactoryMockRecorder struct {
	mock *MockControllerServiceFactory
}

// NewMockControllerServiceFactory creates a new mock instance.
func NewMockControllerServiceFactory(ctrl *gomock.Controller) *MockControllerServiceFactory {
	mock := &MockControllerServiceFactory{ctrl: ctrl}
	mock.recorder = &MockControllerServiceFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerServiceFactory) EXPECT() *MockControllerServiceFactoryMockRecorder {
	return m.recorder
}

// Access mocks base method.
func (m *MockControllerServiceFactory) Access() *service.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Access")
	ret0, _ := ret[0].(*service.Service)
	return ret0
}

// Access indicates an expected call of Access.
func (mr *MockControllerServiceFactoryMockRecorder) Access() *MockControllerServiceFactoryAccessCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Access", reflect.TypeOf((*MockControllerServiceFactory)(nil).Access))
	return &MockControllerServiceFactoryAccessCall{Call: call}
}

// MockControllerServiceFactoryAccessCall wrap *gomock.Call
type MockControllerServiceFactoryAccessCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryAccessCall) Return(arg0 *service.Service) *MockControllerServiceFactoryAccessCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryAccessCall) Do(f func() *service.Service) *MockControllerServiceFactoryAccessCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryAccessCall) DoAndReturn(f func() *service.Service) *MockControllerServiceFactoryAccessCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AgentObjectStore mocks base method.
func (m *MockControllerServiceFactory) AgentObjectStore() *service9.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentObjectStore")
	ret0, _ := ret[0].(*service9.WatchableService)
	return ret0
}

// AgentObjectStore indicates an expected call of AgentObjectStore.
func (mr *MockControllerServiceFactoryMockRecorder) AgentObjectStore() *MockControllerServiceFactoryAgentObjectStoreCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentObjectStore", reflect.TypeOf((*MockControllerServiceFactory)(nil).AgentObjectStore))
	return &MockControllerServiceFactoryAgentObjectStoreCall{Call: call}
}

// MockControllerServiceFactoryAgentObjectStoreCall wrap *gomock.Call
type MockControllerServiceFactoryAgentObjectStoreCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryAgentObjectStoreCall) Return(arg0 *service9.WatchableService) *MockControllerServiceFactoryAgentObjectStoreCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryAgentObjectStoreCall) Do(f func() *service9.WatchableService) *MockControllerServiceFactoryAgentObjectStoreCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryAgentObjectStoreCall) DoAndReturn(f func() *service9.WatchableService) *MockControllerServiceFactoryAgentObjectStoreCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AutocertCache mocks base method.
func (m *MockControllerServiceFactory) AutocertCache() *service0.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutocertCache")
	ret0, _ := ret[0].(*service0.Service)
	return ret0
}

// AutocertCache indicates an expected call of AutocertCache.
func (mr *MockControllerServiceFactoryMockRecorder) AutocertCache() *MockControllerServiceFactoryAutocertCacheCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutocertCache", reflect.TypeOf((*MockControllerServiceFactory)(nil).AutocertCache))
	return &MockControllerServiceFactoryAutocertCacheCall{Call: call}
}

// MockControllerServiceFactoryAutocertCacheCall wrap *gomock.Call
type MockControllerServiceFactoryAutocertCacheCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryAutocertCacheCall) Return(arg0 *service0.Service) *MockControllerServiceFactoryAutocertCacheCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryAutocertCacheCall) Do(f func() *service0.Service) *MockControllerServiceFactoryAutocertCacheCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryAutocertCacheCall) DoAndReturn(f func() *service0.Service) *MockControllerServiceFactoryAutocertCacheCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Cloud mocks base method.
func (m *MockControllerServiceFactory) Cloud() *service1.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud")
	ret0, _ := ret[0].(*service1.WatchableService)
	return ret0
}

// Cloud indicates an expected call of Cloud.
func (mr *MockControllerServiceFactoryMockRecorder) Cloud() *MockControllerServiceFactoryCloudCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockControllerServiceFactory)(nil).Cloud))
	return &MockControllerServiceFactoryCloudCall{Call: call}
}

// MockControllerServiceFactoryCloudCall wrap *gomock.Call
type MockControllerServiceFactoryCloudCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryCloudCall) Return(arg0 *service1.WatchableService) *MockControllerServiceFactoryCloudCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryCloudCall) Do(f func() *service1.WatchableService) *MockControllerServiceFactoryCloudCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryCloudCall) DoAndReturn(f func() *service1.WatchableService) *MockControllerServiceFactoryCloudCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerConfig mocks base method.
func (m *MockControllerServiceFactory) ControllerConfig() *service2.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(*service2.WatchableService)
	return ret0
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerServiceFactoryMockRecorder) ControllerConfig() *MockControllerServiceFactoryControllerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerServiceFactory)(nil).ControllerConfig))
	return &MockControllerServiceFactoryControllerConfigCall{Call: call}
}

// MockControllerServiceFactoryControllerConfigCall wrap *gomock.Call
type MockControllerServiceFactoryControllerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryControllerConfigCall) Return(arg0 *service2.WatchableService) *MockControllerServiceFactoryControllerConfigCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryControllerConfigCall) Do(f func() *service2.WatchableService) *MockControllerServiceFactoryControllerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryControllerConfigCall) DoAndReturn(f func() *service2.WatchableService) *MockControllerServiceFactoryControllerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerNode mocks base method.
func (m *MockControllerServiceFactory) ControllerNode() *service3.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerNode")
	ret0, _ := ret[0].(*service3.Service)
	return ret0
}

// ControllerNode indicates an expected call of ControllerNode.
func (mr *MockControllerServiceFactoryMockRecorder) ControllerNode() *MockControllerServiceFactoryControllerNodeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerNode", reflect.TypeOf((*MockControllerServiceFactory)(nil).ControllerNode))
	return &MockControllerServiceFactoryControllerNodeCall{Call: call}
}

// MockControllerServiceFactoryControllerNodeCall wrap *gomock.Call
type MockControllerServiceFactoryControllerNodeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryControllerNodeCall) Return(arg0 *service3.Service) *MockControllerServiceFactoryControllerNodeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryControllerNodeCall) Do(f func() *service3.Service) *MockControllerServiceFactoryControllerNodeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryControllerNodeCall) DoAndReturn(f func() *service3.Service) *MockControllerServiceFactoryControllerNodeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Credential mocks base method.
func (m *MockControllerServiceFactory) Credential() *service4.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credential")
	ret0, _ := ret[0].(*service4.WatchableService)
	return ret0
}

// Credential indicates an expected call of Credential.
func (mr *MockControllerServiceFactoryMockRecorder) Credential() *MockControllerServiceFactoryCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credential", reflect.TypeOf((*MockControllerServiceFactory)(nil).Credential))
	return &MockControllerServiceFactoryCredentialCall{Call: call}
}

// MockControllerServiceFactoryCredentialCall wrap *gomock.Call
type MockControllerServiceFactoryCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryCredentialCall) Return(arg0 *service4.WatchableService) *MockControllerServiceFactoryCredentialCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryCredentialCall) Do(f func() *service4.WatchableService) *MockControllerServiceFactoryCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryCredentialCall) DoAndReturn(f func() *service4.WatchableService) *MockControllerServiceFactoryCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ExternalController mocks base method.
func (m *MockControllerServiceFactory) ExternalController() *service5.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalController")
	ret0, _ := ret[0].(*service5.WatchableService)
	return ret0
}

// ExternalController indicates an expected call of ExternalController.
func (mr *MockControllerServiceFactoryMockRecorder) ExternalController() *MockControllerServiceFactoryExternalControllerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalController", reflect.TypeOf((*MockControllerServiceFactory)(nil).ExternalController))
	return &MockControllerServiceFactoryExternalControllerCall{Call: call}
}

// MockControllerServiceFactoryExternalControllerCall wrap *gomock.Call
type MockControllerServiceFactoryExternalControllerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryExternalControllerCall) Return(arg0 *service5.WatchableService) *MockControllerServiceFactoryExternalControllerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryExternalControllerCall) Do(f func() *service5.WatchableService) *MockControllerServiceFactoryExternalControllerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryExternalControllerCall) DoAndReturn(f func() *service5.WatchableService) *MockControllerServiceFactoryExternalControllerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Flag mocks base method.
func (m *MockControllerServiceFactory) Flag() *service6.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flag")
	ret0, _ := ret[0].(*service6.Service)
	return ret0
}

// Flag indicates an expected call of Flag.
func (mr *MockControllerServiceFactoryMockRecorder) Flag() *MockControllerServiceFactoryFlagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flag", reflect.TypeOf((*MockControllerServiceFactory)(nil).Flag))
	return &MockControllerServiceFactoryFlagCall{Call: call}
}

// MockControllerServiceFactoryFlagCall wrap *gomock.Call
type MockControllerServiceFactoryFlagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryFlagCall) Return(arg0 *service6.Service) *MockControllerServiceFactoryFlagCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryFlagCall) Do(f func() *service6.Service) *MockControllerServiceFactoryFlagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryFlagCall) DoAndReturn(f func() *service6.Service) *MockControllerServiceFactoryFlagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Model mocks base method.
func (m *MockControllerServiceFactory) Model() *service7.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(*service7.Service)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockControllerServiceFactoryMockRecorder) Model() *MockControllerServiceFactoryModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockControllerServiceFactory)(nil).Model))
	return &MockControllerServiceFactoryModelCall{Call: call}
}

// MockControllerServiceFactoryModelCall wrap *gomock.Call
type MockControllerServiceFactoryModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryModelCall) Return(arg0 *service7.Service) *MockControllerServiceFactoryModelCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryModelCall) Do(f func() *service7.Service) *MockControllerServiceFactoryModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryModelCall) DoAndReturn(f func() *service7.Service) *MockControllerServiceFactoryModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelDefaults mocks base method.
func (m *MockControllerServiceFactory) ModelDefaults() *service8.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelDefaults")
	ret0, _ := ret[0].(*service8.Service)
	return ret0
}

// ModelDefaults indicates an expected call of ModelDefaults.
func (mr *MockControllerServiceFactoryMockRecorder) ModelDefaults() *MockControllerServiceFactoryModelDefaultsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelDefaults", reflect.TypeOf((*MockControllerServiceFactory)(nil).ModelDefaults))
	return &MockControllerServiceFactoryModelDefaultsCall{Call: call}
}

// MockControllerServiceFactoryModelDefaultsCall wrap *gomock.Call
type MockControllerServiceFactoryModelDefaultsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryModelDefaultsCall) Return(arg0 *service8.Service) *MockControllerServiceFactoryModelDefaultsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryModelDefaultsCall) Do(f func() *service8.Service) *MockControllerServiceFactoryModelDefaultsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryModelDefaultsCall) DoAndReturn(f func() *service8.Service) *MockControllerServiceFactoryModelDefaultsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SecretBackend mocks base method.
func (m *MockControllerServiceFactory) SecretBackend() *service10.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretBackend")
	ret0, _ := ret[0].(*service10.WatchableService)
	return ret0
}

// SecretBackend indicates an expected call of SecretBackend.
func (mr *MockControllerServiceFactoryMockRecorder) SecretBackend() *MockControllerServiceFactorySecretBackendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretBackend", reflect.TypeOf((*MockControllerServiceFactory)(nil).SecretBackend))
	return &MockControllerServiceFactorySecretBackendCall{Call: call}
}

// MockControllerServiceFactorySecretBackendCall wrap *gomock.Call
type MockControllerServiceFactorySecretBackendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactorySecretBackendCall) Return(arg0 *service10.WatchableService) *MockControllerServiceFactorySecretBackendCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactorySecretBackendCall) Do(f func() *service10.WatchableService) *MockControllerServiceFactorySecretBackendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactorySecretBackendCall) DoAndReturn(f func() *service10.WatchableService) *MockControllerServiceFactorySecretBackendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Upgrade mocks base method.
func (m *MockControllerServiceFactory) Upgrade() *service11.WatchableService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade")
	ret0, _ := ret[0].(*service11.WatchableService)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockControllerServiceFactoryMockRecorder) Upgrade() *MockControllerServiceFactoryUpgradeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockControllerServiceFactory)(nil).Upgrade))
	return &MockControllerServiceFactoryUpgradeCall{Call: call}
}

// MockControllerServiceFactoryUpgradeCall wrap *gomock.Call
type MockControllerServiceFactoryUpgradeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerServiceFactoryUpgradeCall) Return(arg0 *service11.WatchableService) *MockControllerServiceFactoryUpgradeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerServiceFactoryUpgradeCall) Do(f func() *service11.WatchableService) *MockControllerServiceFactoryUpgradeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerServiceFactoryUpgradeCall) DoAndReturn(f func() *service11.WatchableService) *MockControllerServiceFactoryUpgradeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
