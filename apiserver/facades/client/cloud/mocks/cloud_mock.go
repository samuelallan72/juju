// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/cloud (interfaces: CredentialService,CloudService,CloudAccessService)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/cloud_mock.go github.com/juju/juju/apiserver/facades/client/cloud CredentialService,CloudService,CloudAccessService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cloud "github.com/juju/juju/cloud"
	credential "github.com/juju/juju/core/credential"
	permission "github.com/juju/juju/core/permission"
	watcher "github.com/juju/juju/core/watcher"
	access "github.com/juju/juju/domain/access"
	service "github.com/juju/juju/domain/credential/service"
	gomock "go.uber.org/mock/gomock"
)

// MockCredentialService is a mock of CredentialService interface.
type MockCredentialService struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialServiceMockRecorder
}

// MockCredentialServiceMockRecorder is the mock recorder for MockCredentialService.
type MockCredentialServiceMockRecorder struct {
	mock *MockCredentialService
}

// NewMockCredentialService creates a new mock instance.
func NewMockCredentialService(ctrl *gomock.Controller) *MockCredentialService {
	mock := &MockCredentialService{ctrl: ctrl}
	mock.recorder = &MockCredentialServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredentialService) EXPECT() *MockCredentialServiceMockRecorder {
	return m.recorder
}

// AllCloudCredentialsForOwner mocks base method.
func (m *MockCredentialService) AllCloudCredentialsForOwner(arg0 context.Context, arg1 string) (map[credential.Key]cloud.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllCloudCredentialsForOwner", arg0, arg1)
	ret0, _ := ret[0].(map[credential.Key]cloud.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllCloudCredentialsForOwner indicates an expected call of AllCloudCredentialsForOwner.
func (mr *MockCredentialServiceMockRecorder) AllCloudCredentialsForOwner(arg0, arg1 any) *MockCredentialServiceAllCloudCredentialsForOwnerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllCloudCredentialsForOwner", reflect.TypeOf((*MockCredentialService)(nil).AllCloudCredentialsForOwner), arg0, arg1)
	return &MockCredentialServiceAllCloudCredentialsForOwnerCall{Call: call}
}

// MockCredentialServiceAllCloudCredentialsForOwnerCall wrap *gomock.Call
type MockCredentialServiceAllCloudCredentialsForOwnerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCredentialServiceAllCloudCredentialsForOwnerCall) Return(arg0 map[credential.Key]cloud.Credential, arg1 error) *MockCredentialServiceAllCloudCredentialsForOwnerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCredentialServiceAllCloudCredentialsForOwnerCall) Do(f func(context.Context, string) (map[credential.Key]cloud.Credential, error)) *MockCredentialServiceAllCloudCredentialsForOwnerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCredentialServiceAllCloudCredentialsForOwnerCall) DoAndReturn(f func(context.Context, string) (map[credential.Key]cloud.Credential, error)) *MockCredentialServiceAllCloudCredentialsForOwnerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CheckAndRevokeCredential mocks base method.
func (m *MockCredentialService) CheckAndRevokeCredential(arg0 context.Context, arg1 credential.Key, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAndRevokeCredential", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckAndRevokeCredential indicates an expected call of CheckAndRevokeCredential.
func (mr *MockCredentialServiceMockRecorder) CheckAndRevokeCredential(arg0, arg1, arg2 any) *MockCredentialServiceCheckAndRevokeCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAndRevokeCredential", reflect.TypeOf((*MockCredentialService)(nil).CheckAndRevokeCredential), arg0, arg1, arg2)
	return &MockCredentialServiceCheckAndRevokeCredentialCall{Call: call}
}

// MockCredentialServiceCheckAndRevokeCredentialCall wrap *gomock.Call
type MockCredentialServiceCheckAndRevokeCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCredentialServiceCheckAndRevokeCredentialCall) Return(arg0 error) *MockCredentialServiceCheckAndRevokeCredentialCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCredentialServiceCheckAndRevokeCredentialCall) Do(f func(context.Context, credential.Key, bool) error) *MockCredentialServiceCheckAndRevokeCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCredentialServiceCheckAndRevokeCredentialCall) DoAndReturn(f func(context.Context, credential.Key, bool) error) *MockCredentialServiceCheckAndRevokeCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CheckAndUpdateCredential mocks base method.
func (m *MockCredentialService) CheckAndUpdateCredential(arg0 context.Context, arg1 credential.Key, arg2 cloud.Credential, arg3 bool) ([]service.UpdateCredentialModelResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAndUpdateCredential", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]service.UpdateCredentialModelResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAndUpdateCredential indicates an expected call of CheckAndUpdateCredential.
func (mr *MockCredentialServiceMockRecorder) CheckAndUpdateCredential(arg0, arg1, arg2, arg3 any) *MockCredentialServiceCheckAndUpdateCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAndUpdateCredential", reflect.TypeOf((*MockCredentialService)(nil).CheckAndUpdateCredential), arg0, arg1, arg2, arg3)
	return &MockCredentialServiceCheckAndUpdateCredentialCall{Call: call}
}

// MockCredentialServiceCheckAndUpdateCredentialCall wrap *gomock.Call
type MockCredentialServiceCheckAndUpdateCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCredentialServiceCheckAndUpdateCredentialCall) Return(arg0 []service.UpdateCredentialModelResult, arg1 error) *MockCredentialServiceCheckAndUpdateCredentialCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCredentialServiceCheckAndUpdateCredentialCall) Do(f func(context.Context, credential.Key, cloud.Credential, bool) ([]service.UpdateCredentialModelResult, error)) *MockCredentialServiceCheckAndUpdateCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCredentialServiceCheckAndUpdateCredentialCall) DoAndReturn(f func(context.Context, credential.Key, cloud.Credential, bool) ([]service.UpdateCredentialModelResult, error)) *MockCredentialServiceCheckAndUpdateCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CloudCredential mocks base method.
func (m *MockCredentialService) CloudCredential(arg0 context.Context, arg1 credential.Key) (cloud.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudCredential", arg0, arg1)
	ret0, _ := ret[0].(cloud.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudCredential indicates an expected call of CloudCredential.
func (mr *MockCredentialServiceMockRecorder) CloudCredential(arg0, arg1 any) *MockCredentialServiceCloudCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudCredential", reflect.TypeOf((*MockCredentialService)(nil).CloudCredential), arg0, arg1)
	return &MockCredentialServiceCloudCredentialCall{Call: call}
}

// MockCredentialServiceCloudCredentialCall wrap *gomock.Call
type MockCredentialServiceCloudCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCredentialServiceCloudCredentialCall) Return(arg0 cloud.Credential, arg1 error) *MockCredentialServiceCloudCredentialCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCredentialServiceCloudCredentialCall) Do(f func(context.Context, credential.Key) (cloud.Credential, error)) *MockCredentialServiceCloudCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCredentialServiceCloudCredentialCall) DoAndReturn(f func(context.Context, credential.Key) (cloud.Credential, error)) *MockCredentialServiceCloudCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CloudCredentialsForOwner mocks base method.
func (m *MockCredentialService) CloudCredentialsForOwner(arg0 context.Context, arg1, arg2 string) (map[string]cloud.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudCredentialsForOwner", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]cloud.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudCredentialsForOwner indicates an expected call of CloudCredentialsForOwner.
func (mr *MockCredentialServiceMockRecorder) CloudCredentialsForOwner(arg0, arg1, arg2 any) *MockCredentialServiceCloudCredentialsForOwnerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudCredentialsForOwner", reflect.TypeOf((*MockCredentialService)(nil).CloudCredentialsForOwner), arg0, arg1, arg2)
	return &MockCredentialServiceCloudCredentialsForOwnerCall{Call: call}
}

// MockCredentialServiceCloudCredentialsForOwnerCall wrap *gomock.Call
type MockCredentialServiceCloudCredentialsForOwnerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCredentialServiceCloudCredentialsForOwnerCall) Return(arg0 map[string]cloud.Credential, arg1 error) *MockCredentialServiceCloudCredentialsForOwnerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCredentialServiceCloudCredentialsForOwnerCall) Do(f func(context.Context, string, string) (map[string]cloud.Credential, error)) *MockCredentialServiceCloudCredentialsForOwnerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCredentialServiceCloudCredentialsForOwnerCall) DoAndReturn(f func(context.Context, string, string) (map[string]cloud.Credential, error)) *MockCredentialServiceCloudCredentialsForOwnerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoveCloudCredential mocks base method.
func (m *MockCredentialService) RemoveCloudCredential(arg0 context.Context, arg1 credential.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCloudCredential", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCloudCredential indicates an expected call of RemoveCloudCredential.
func (mr *MockCredentialServiceMockRecorder) RemoveCloudCredential(arg0, arg1 any) *MockCredentialServiceRemoveCloudCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCloudCredential", reflect.TypeOf((*MockCredentialService)(nil).RemoveCloudCredential), arg0, arg1)
	return &MockCredentialServiceRemoveCloudCredentialCall{Call: call}
}

// MockCredentialServiceRemoveCloudCredentialCall wrap *gomock.Call
type MockCredentialServiceRemoveCloudCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCredentialServiceRemoveCloudCredentialCall) Return(arg0 error) *MockCredentialServiceRemoveCloudCredentialCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCredentialServiceRemoveCloudCredentialCall) Do(f func(context.Context, credential.Key) error) *MockCredentialServiceRemoveCloudCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCredentialServiceRemoveCloudCredentialCall) DoAndReturn(f func(context.Context, credential.Key) error) *MockCredentialServiceRemoveCloudCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateCloudCredential mocks base method.
func (m *MockCredentialService) UpdateCloudCredential(arg0 context.Context, arg1 credential.Key, arg2 cloud.Credential) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCloudCredential", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCloudCredential indicates an expected call of UpdateCloudCredential.
func (mr *MockCredentialServiceMockRecorder) UpdateCloudCredential(arg0, arg1, arg2 any) *MockCredentialServiceUpdateCloudCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCloudCredential", reflect.TypeOf((*MockCredentialService)(nil).UpdateCloudCredential), arg0, arg1, arg2)
	return &MockCredentialServiceUpdateCloudCredentialCall{Call: call}
}

// MockCredentialServiceUpdateCloudCredentialCall wrap *gomock.Call
type MockCredentialServiceUpdateCloudCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCredentialServiceUpdateCloudCredentialCall) Return(arg0 error) *MockCredentialServiceUpdateCloudCredentialCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCredentialServiceUpdateCloudCredentialCall) Do(f func(context.Context, credential.Key, cloud.Credential) error) *MockCredentialServiceUpdateCloudCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCredentialServiceUpdateCloudCredentialCall) DoAndReturn(f func(context.Context, credential.Key, cloud.Credential) error) *MockCredentialServiceUpdateCloudCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchCredential mocks base method.
func (m *MockCredentialService) WatchCredential(arg0 context.Context, arg1 credential.Key) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchCredential", arg0, arg1)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchCredential indicates an expected call of WatchCredential.
func (mr *MockCredentialServiceMockRecorder) WatchCredential(arg0, arg1 any) *MockCredentialServiceWatchCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchCredential", reflect.TypeOf((*MockCredentialService)(nil).WatchCredential), arg0, arg1)
	return &MockCredentialServiceWatchCredentialCall{Call: call}
}

// MockCredentialServiceWatchCredentialCall wrap *gomock.Call
type MockCredentialServiceWatchCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCredentialServiceWatchCredentialCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockCredentialServiceWatchCredentialCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCredentialServiceWatchCredentialCall) Do(f func(context.Context, credential.Key) (watcher.Watcher[struct{}], error)) *MockCredentialServiceWatchCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCredentialServiceWatchCredentialCall) DoAndReturn(f func(context.Context, credential.Key) (watcher.Watcher[struct{}], error)) *MockCredentialServiceWatchCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockCloudService is a mock of CloudService interface.
type MockCloudService struct {
	ctrl     *gomock.Controller
	recorder *MockCloudServiceMockRecorder
}

// MockCloudServiceMockRecorder is the mock recorder for MockCloudService.
type MockCloudServiceMockRecorder struct {
	mock *MockCloudService
}

// NewMockCloudService creates a new mock instance.
func NewMockCloudService(ctrl *gomock.Controller) *MockCloudService {
	mock := &MockCloudService{ctrl: ctrl}
	mock.recorder = &MockCloudServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudService) EXPECT() *MockCloudServiceMockRecorder {
	return m.recorder
}

// Cloud mocks base method.
func (m *MockCloudService) Cloud(arg0 context.Context, arg1 string) (*cloud.Cloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud", arg0, arg1)
	ret0, _ := ret[0].(*cloud.Cloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cloud indicates an expected call of Cloud.
func (mr *MockCloudServiceMockRecorder) Cloud(arg0, arg1 any) *MockCloudServiceCloudCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockCloudService)(nil).Cloud), arg0, arg1)
	return &MockCloudServiceCloudCall{Call: call}
}

// MockCloudServiceCloudCall wrap *gomock.Call
type MockCloudServiceCloudCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCloudServiceCloudCall) Return(arg0 *cloud.Cloud, arg1 error) *MockCloudServiceCloudCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCloudServiceCloudCall) Do(f func(context.Context, string) (*cloud.Cloud, error)) *MockCloudServiceCloudCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCloudServiceCloudCall) DoAndReturn(f func(context.Context, string) (*cloud.Cloud, error)) *MockCloudServiceCloudCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateCloud mocks base method.
func (m *MockCloudService) CreateCloud(arg0 context.Context, arg1 string, arg2 cloud.Cloud) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCloud", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCloud indicates an expected call of CreateCloud.
func (mr *MockCloudServiceMockRecorder) CreateCloud(arg0, arg1, arg2 any) *MockCloudServiceCreateCloudCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCloud", reflect.TypeOf((*MockCloudService)(nil).CreateCloud), arg0, arg1, arg2)
	return &MockCloudServiceCreateCloudCall{Call: call}
}

// MockCloudServiceCreateCloudCall wrap *gomock.Call
type MockCloudServiceCreateCloudCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCloudServiceCreateCloudCall) Return(arg0 error) *MockCloudServiceCreateCloudCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCloudServiceCreateCloudCall) Do(f func(context.Context, string, cloud.Cloud) error) *MockCloudServiceCreateCloudCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCloudServiceCreateCloudCall) DoAndReturn(f func(context.Context, string, cloud.Cloud) error) *MockCloudServiceCreateCloudCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteCloud mocks base method.
func (m *MockCloudService) DeleteCloud(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCloud", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCloud indicates an expected call of DeleteCloud.
func (mr *MockCloudServiceMockRecorder) DeleteCloud(arg0, arg1 any) *MockCloudServiceDeleteCloudCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCloud", reflect.TypeOf((*MockCloudService)(nil).DeleteCloud), arg0, arg1)
	return &MockCloudServiceDeleteCloudCall{Call: call}
}

// MockCloudServiceDeleteCloudCall wrap *gomock.Call
type MockCloudServiceDeleteCloudCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCloudServiceDeleteCloudCall) Return(arg0 error) *MockCloudServiceDeleteCloudCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCloudServiceDeleteCloudCall) Do(f func(context.Context, string) error) *MockCloudServiceDeleteCloudCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCloudServiceDeleteCloudCall) DoAndReturn(f func(context.Context, string) error) *MockCloudServiceDeleteCloudCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListAll mocks base method.
func (m *MockCloudService) ListAll(arg0 context.Context) ([]cloud.Cloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].([]cloud.Cloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockCloudServiceMockRecorder) ListAll(arg0 any) *MockCloudServiceListAllCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockCloudService)(nil).ListAll), arg0)
	return &MockCloudServiceListAllCall{Call: call}
}

// MockCloudServiceListAllCall wrap *gomock.Call
type MockCloudServiceListAllCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCloudServiceListAllCall) Return(arg0 []cloud.Cloud, arg1 error) *MockCloudServiceListAllCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCloudServiceListAllCall) Do(f func(context.Context) ([]cloud.Cloud, error)) *MockCloudServiceListAllCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCloudServiceListAllCall) DoAndReturn(f func(context.Context) ([]cloud.Cloud, error)) *MockCloudServiceListAllCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateCloud mocks base method.
func (m *MockCloudService) UpdateCloud(arg0 context.Context, arg1 cloud.Cloud) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCloud", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCloud indicates an expected call of UpdateCloud.
func (mr *MockCloudServiceMockRecorder) UpdateCloud(arg0, arg1 any) *MockCloudServiceUpdateCloudCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCloud", reflect.TypeOf((*MockCloudService)(nil).UpdateCloud), arg0, arg1)
	return &MockCloudServiceUpdateCloudCall{Call: call}
}

// MockCloudServiceUpdateCloudCall wrap *gomock.Call
type MockCloudServiceUpdateCloudCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCloudServiceUpdateCloudCall) Return(arg0 error) *MockCloudServiceUpdateCloudCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCloudServiceUpdateCloudCall) Do(f func(context.Context, cloud.Cloud) error) *MockCloudServiceUpdateCloudCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCloudServiceUpdateCloudCall) DoAndReturn(f func(context.Context, cloud.Cloud) error) *MockCloudServiceUpdateCloudCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockCloudAccessService is a mock of CloudAccessService interface.
type MockCloudAccessService struct {
	ctrl     *gomock.Controller
	recorder *MockCloudAccessServiceMockRecorder
}

// MockCloudAccessServiceMockRecorder is the mock recorder for MockCloudAccessService.
type MockCloudAccessServiceMockRecorder struct {
	mock *MockCloudAccessService
}

// NewMockCloudAccessService creates a new mock instance.
func NewMockCloudAccessService(ctrl *gomock.Controller) *MockCloudAccessService {
	mock := &MockCloudAccessService{ctrl: ctrl}
	mock.recorder = &MockCloudAccessServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudAccessService) EXPECT() *MockCloudAccessServiceMockRecorder {
	return m.recorder
}

// AllModelAccessForCloudCredential mocks base method.
func (m *MockCloudAccessService) AllModelAccessForCloudCredential(arg0 context.Context, arg1 credential.Key) ([]access.CredentialOwnerModelAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllModelAccessForCloudCredential", arg0, arg1)
	ret0, _ := ret[0].([]access.CredentialOwnerModelAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllModelAccessForCloudCredential indicates an expected call of AllModelAccessForCloudCredential.
func (mr *MockCloudAccessServiceMockRecorder) AllModelAccessForCloudCredential(arg0, arg1 any) *MockCloudAccessServiceAllModelAccessForCloudCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllModelAccessForCloudCredential", reflect.TypeOf((*MockCloudAccessService)(nil).AllModelAccessForCloudCredential), arg0, arg1)
	return &MockCloudAccessServiceAllModelAccessForCloudCredentialCall{Call: call}
}

// MockCloudAccessServiceAllModelAccessForCloudCredentialCall wrap *gomock.Call
type MockCloudAccessServiceAllModelAccessForCloudCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCloudAccessServiceAllModelAccessForCloudCredentialCall) Return(arg0 []access.CredentialOwnerModelAccess, arg1 error) *MockCloudAccessServiceAllModelAccessForCloudCredentialCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCloudAccessServiceAllModelAccessForCloudCredentialCall) Do(f func(context.Context, credential.Key) ([]access.CredentialOwnerModelAccess, error)) *MockCloudAccessServiceAllModelAccessForCloudCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCloudAccessServiceAllModelAccessForCloudCredentialCall) DoAndReturn(f func(context.Context, credential.Key) ([]access.CredentialOwnerModelAccess, error)) *MockCloudAccessServiceAllModelAccessForCloudCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreatePermission mocks base method.
func (m *MockCloudAccessService) CreatePermission(arg0 context.Context, arg1 permission.UserAccessSpec) (permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePermission", arg0, arg1)
	ret0, _ := ret[0].(permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePermission indicates an expected call of CreatePermission.
func (mr *MockCloudAccessServiceMockRecorder) CreatePermission(arg0, arg1 any) *MockCloudAccessServiceCreatePermissionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePermission", reflect.TypeOf((*MockCloudAccessService)(nil).CreatePermission), arg0, arg1)
	return &MockCloudAccessServiceCreatePermissionCall{Call: call}
}

// MockCloudAccessServiceCreatePermissionCall wrap *gomock.Call
type MockCloudAccessServiceCreatePermissionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCloudAccessServiceCreatePermissionCall) Return(arg0 permission.UserAccess, arg1 error) *MockCloudAccessServiceCreatePermissionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCloudAccessServiceCreatePermissionCall) Do(f func(context.Context, permission.UserAccessSpec) (permission.UserAccess, error)) *MockCloudAccessServiceCreatePermissionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCloudAccessServiceCreatePermissionCall) DoAndReturn(f func(context.Context, permission.UserAccessSpec) (permission.UserAccess, error)) *MockCloudAccessServiceCreatePermissionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadAllAccessForUserAndObjectType mocks base method.
func (m *MockCloudAccessService) ReadAllAccessForUserAndObjectType(arg0 context.Context, arg1 string, arg2 permission.ObjectType) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllAccessForUserAndObjectType", arg0, arg1, arg2)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllAccessForUserAndObjectType indicates an expected call of ReadAllAccessForUserAndObjectType.
func (mr *MockCloudAccessServiceMockRecorder) ReadAllAccessForUserAndObjectType(arg0, arg1, arg2 any) *MockCloudAccessServiceReadAllAccessForUserAndObjectTypeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllAccessForUserAndObjectType", reflect.TypeOf((*MockCloudAccessService)(nil).ReadAllAccessForUserAndObjectType), arg0, arg1, arg2)
	return &MockCloudAccessServiceReadAllAccessForUserAndObjectTypeCall{Call: call}
}

// MockCloudAccessServiceReadAllAccessForUserAndObjectTypeCall wrap *gomock.Call
type MockCloudAccessServiceReadAllAccessForUserAndObjectTypeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCloudAccessServiceReadAllAccessForUserAndObjectTypeCall) Return(arg0 []permission.UserAccess, arg1 error) *MockCloudAccessServiceReadAllAccessForUserAndObjectTypeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCloudAccessServiceReadAllAccessForUserAndObjectTypeCall) Do(f func(context.Context, string, permission.ObjectType) ([]permission.UserAccess, error)) *MockCloudAccessServiceReadAllAccessForUserAndObjectTypeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCloudAccessServiceReadAllAccessForUserAndObjectTypeCall) DoAndReturn(f func(context.Context, string, permission.ObjectType) ([]permission.UserAccess, error)) *MockCloudAccessServiceReadAllAccessForUserAndObjectTypeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadAllUserAccessForTarget mocks base method.
func (m *MockCloudAccessService) ReadAllUserAccessForTarget(arg0 context.Context, arg1 permission.ID) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllUserAccessForTarget", arg0, arg1)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllUserAccessForTarget indicates an expected call of ReadAllUserAccessForTarget.
func (mr *MockCloudAccessServiceMockRecorder) ReadAllUserAccessForTarget(arg0, arg1 any) *MockCloudAccessServiceReadAllUserAccessForTargetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllUserAccessForTarget", reflect.TypeOf((*MockCloudAccessService)(nil).ReadAllUserAccessForTarget), arg0, arg1)
	return &MockCloudAccessServiceReadAllUserAccessForTargetCall{Call: call}
}

// MockCloudAccessServiceReadAllUserAccessForTargetCall wrap *gomock.Call
type MockCloudAccessServiceReadAllUserAccessForTargetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCloudAccessServiceReadAllUserAccessForTargetCall) Return(arg0 []permission.UserAccess, arg1 error) *MockCloudAccessServiceReadAllUserAccessForTargetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCloudAccessServiceReadAllUserAccessForTargetCall) Do(f func(context.Context, permission.ID) ([]permission.UserAccess, error)) *MockCloudAccessServiceReadAllUserAccessForTargetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCloudAccessServiceReadAllUserAccessForTargetCall) DoAndReturn(f func(context.Context, permission.ID) ([]permission.UserAccess, error)) *MockCloudAccessServiceReadAllUserAccessForTargetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadUserAccessLevelForTarget mocks base method.
func (m *MockCloudAccessService) ReadUserAccessLevelForTarget(arg0 context.Context, arg1 string, arg2 permission.ID) (permission.Access, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserAccessLevelForTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(permission.Access)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserAccessLevelForTarget indicates an expected call of ReadUserAccessLevelForTarget.
func (mr *MockCloudAccessServiceMockRecorder) ReadUserAccessLevelForTarget(arg0, arg1, arg2 any) *MockCloudAccessServiceReadUserAccessLevelForTargetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserAccessLevelForTarget", reflect.TypeOf((*MockCloudAccessService)(nil).ReadUserAccessLevelForTarget), arg0, arg1, arg2)
	return &MockCloudAccessServiceReadUserAccessLevelForTargetCall{Call: call}
}

// MockCloudAccessServiceReadUserAccessLevelForTargetCall wrap *gomock.Call
type MockCloudAccessServiceReadUserAccessLevelForTargetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCloudAccessServiceReadUserAccessLevelForTargetCall) Return(arg0 permission.Access, arg1 error) *MockCloudAccessServiceReadUserAccessLevelForTargetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCloudAccessServiceReadUserAccessLevelForTargetCall) Do(f func(context.Context, string, permission.ID) (permission.Access, error)) *MockCloudAccessServiceReadUserAccessLevelForTargetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCloudAccessServiceReadUserAccessLevelForTargetCall) DoAndReturn(f func(context.Context, string, permission.ID) (permission.Access, error)) *MockCloudAccessServiceReadUserAccessLevelForTargetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdatePermission mocks base method.
func (m *MockCloudAccessService) UpdatePermission(arg0 context.Context, arg1 access.UpdatePermissionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePermission", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePermission indicates an expected call of UpdatePermission.
func (mr *MockCloudAccessServiceMockRecorder) UpdatePermission(arg0, arg1 any) *MockCloudAccessServiceUpdatePermissionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePermission", reflect.TypeOf((*MockCloudAccessService)(nil).UpdatePermission), arg0, arg1)
	return &MockCloudAccessServiceUpdatePermissionCall{Call: call}
}

// MockCloudAccessServiceUpdatePermissionCall wrap *gomock.Call
type MockCloudAccessServiceUpdatePermissionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCloudAccessServiceUpdatePermissionCall) Return(arg0 error) *MockCloudAccessServiceUpdatePermissionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCloudAccessServiceUpdatePermissionCall) Do(f func(context.Context, access.UpdatePermissionArgs) error) *MockCloudAccessServiceUpdatePermissionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCloudAccessServiceUpdatePermissionCall) DoAndReturn(f func(context.Context, access.UpdatePermissionArgs) error) *MockCloudAccessServiceUpdatePermissionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
