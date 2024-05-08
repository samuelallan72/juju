// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/common/secrets (interfaces: SecretBackendGetter,SecretService,SecretBackendService)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/commonsecrets_mock.go github.com/juju/juju/apiserver/common/secrets SecretBackendGetter,SecretService,SecretBackendService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/juju/juju/core/model"
	secrets "github.com/juju/juju/core/secrets"
	watcher "github.com/juju/juju/core/watcher"
	service "github.com/juju/juju/domain/secret/service"
	service0 "github.com/juju/juju/domain/secretbackend/service"
	gomock "go.uber.org/mock/gomock"
)

// MockSecretBackendGetter is a mock of SecretBackendGetter interface.
type MockSecretBackendGetter struct {
	ctrl     *gomock.Controller
	recorder *MockSecretBackendGetterMockRecorder
}

// MockSecretBackendGetterMockRecorder is the mock recorder for MockSecretBackendGetter.
type MockSecretBackendGetterMockRecorder struct {
	mock *MockSecretBackendGetter
}

// NewMockSecretBackendGetter creates a new mock instance.
func NewMockSecretBackendGetter(ctrl *gomock.Controller) *MockSecretBackendGetter {
	mock := &MockSecretBackendGetter{ctrl: ctrl}
	mock.recorder = &MockSecretBackendGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretBackendGetter) EXPECT() *MockSecretBackendGetterMockRecorder {
	return m.recorder
}

// GetSecretBackendID mocks base method.
func (m *MockSecretBackendGetter) GetSecretBackendID(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretBackendID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretBackendID indicates an expected call of GetSecretBackendID.
func (mr *MockSecretBackendGetterMockRecorder) GetSecretBackendID(arg0 any) *MockSecretBackendGetterGetSecretBackendIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretBackendID", reflect.TypeOf((*MockSecretBackendGetter)(nil).GetSecretBackendID), arg0)
	return &MockSecretBackendGetterGetSecretBackendIDCall{Call: call}
}

// MockSecretBackendGetterGetSecretBackendIDCall wrap *gomock.Call
type MockSecretBackendGetterGetSecretBackendIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretBackendGetterGetSecretBackendIDCall) Return(arg0 string, arg1 error) *MockSecretBackendGetterGetSecretBackendIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretBackendGetterGetSecretBackendIDCall) Do(f func(context.Context) (string, error)) *MockSecretBackendGetterGetSecretBackendIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretBackendGetterGetSecretBackendIDCall) DoAndReturn(f func(context.Context) (string, error)) *MockSecretBackendGetterGetSecretBackendIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockSecretService is a mock of SecretService interface.
type MockSecretService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretServiceMockRecorder
}

// MockSecretServiceMockRecorder is the mock recorder for MockSecretService.
type MockSecretServiceMockRecorder struct {
	mock *MockSecretService
}

// NewMockSecretService creates a new mock instance.
func NewMockSecretService(ctrl *gomock.Controller) *MockSecretService {
	mock := &MockSecretService{ctrl: ctrl}
	mock.recorder = &MockSecretServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretService) EXPECT() *MockSecretServiceMockRecorder {
	return m.recorder
}

// ChangeSecretBackend mocks base method.
func (m *MockSecretService) ChangeSecretBackend(arg0 context.Context, arg1 *secrets.URI, arg2 int, arg3 service.ChangeSecretBackendParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeSecretBackend", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeSecretBackend indicates an expected call of ChangeSecretBackend.
func (mr *MockSecretServiceMockRecorder) ChangeSecretBackend(arg0, arg1, arg2, arg3 any) *MockSecretServiceChangeSecretBackendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeSecretBackend", reflect.TypeOf((*MockSecretService)(nil).ChangeSecretBackend), arg0, arg1, arg2, arg3)
	return &MockSecretServiceChangeSecretBackendCall{Call: call}
}

// MockSecretServiceChangeSecretBackendCall wrap *gomock.Call
type MockSecretServiceChangeSecretBackendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceChangeSecretBackendCall) Return(arg0 error) *MockSecretServiceChangeSecretBackendCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceChangeSecretBackendCall) Do(f func(context.Context, *secrets.URI, int, service.ChangeSecretBackendParams) error) *MockSecretServiceChangeSecretBackendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceChangeSecretBackendCall) DoAndReturn(f func(context.Context, *secrets.URI, int, service.ChangeSecretBackendParams) error) *MockSecretServiceChangeSecretBackendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSecretBackendID mocks base method.
func (m *MockSecretService) GetSecretBackendID(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretBackendID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretBackendID indicates an expected call of GetSecretBackendID.
func (mr *MockSecretServiceMockRecorder) GetSecretBackendID(arg0 any) *MockSecretServiceGetSecretBackendIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretBackendID", reflect.TypeOf((*MockSecretService)(nil).GetSecretBackendID), arg0)
	return &MockSecretServiceGetSecretBackendIDCall{Call: call}
}

// MockSecretServiceGetSecretBackendIDCall wrap *gomock.Call
type MockSecretServiceGetSecretBackendIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceGetSecretBackendIDCall) Return(arg0 string, arg1 error) *MockSecretServiceGetSecretBackendIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceGetSecretBackendIDCall) Do(f func(context.Context) (string, error)) *MockSecretServiceGetSecretBackendIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceGetSecretBackendIDCall) DoAndReturn(f func(context.Context) (string, error)) *MockSecretServiceGetSecretBackendIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListCharmSecretsToDrain mocks base method.
func (m *MockSecretService) ListCharmSecretsToDrain(arg0 context.Context, arg1 ...service.CharmSecretOwner) ([]*secrets.SecretMetadataForDrain, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCharmSecretsToDrain", varargs...)
	ret0, _ := ret[0].([]*secrets.SecretMetadataForDrain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCharmSecretsToDrain indicates an expected call of ListCharmSecretsToDrain.
func (mr *MockSecretServiceMockRecorder) ListCharmSecretsToDrain(arg0 any, arg1 ...any) *MockSecretServiceListCharmSecretsToDrainCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCharmSecretsToDrain", reflect.TypeOf((*MockSecretService)(nil).ListCharmSecretsToDrain), varargs...)
	return &MockSecretServiceListCharmSecretsToDrainCall{Call: call}
}

// MockSecretServiceListCharmSecretsToDrainCall wrap *gomock.Call
type MockSecretServiceListCharmSecretsToDrainCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceListCharmSecretsToDrainCall) Return(arg0 []*secrets.SecretMetadataForDrain, arg1 error) *MockSecretServiceListCharmSecretsToDrainCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceListCharmSecretsToDrainCall) Do(f func(context.Context, ...service.CharmSecretOwner) ([]*secrets.SecretMetadataForDrain, error)) *MockSecretServiceListCharmSecretsToDrainCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceListCharmSecretsToDrainCall) DoAndReturn(f func(context.Context, ...service.CharmSecretOwner) ([]*secrets.SecretMetadataForDrain, error)) *MockSecretServiceListCharmSecretsToDrainCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListUserSecretsToDrain mocks base method.
func (m *MockSecretService) ListUserSecretsToDrain(arg0 context.Context) ([]*secrets.SecretMetadataForDrain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserSecretsToDrain", arg0)
	ret0, _ := ret[0].([]*secrets.SecretMetadataForDrain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserSecretsToDrain indicates an expected call of ListUserSecretsToDrain.
func (mr *MockSecretServiceMockRecorder) ListUserSecretsToDrain(arg0 any) *MockSecretServiceListUserSecretsToDrainCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserSecretsToDrain", reflect.TypeOf((*MockSecretService)(nil).ListUserSecretsToDrain), arg0)
	return &MockSecretServiceListUserSecretsToDrainCall{Call: call}
}

// MockSecretServiceListUserSecretsToDrainCall wrap *gomock.Call
type MockSecretServiceListUserSecretsToDrainCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceListUserSecretsToDrainCall) Return(arg0 []*secrets.SecretMetadataForDrain, arg1 error) *MockSecretServiceListUserSecretsToDrainCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceListUserSecretsToDrainCall) Do(f func(context.Context) ([]*secrets.SecretMetadataForDrain, error)) *MockSecretServiceListUserSecretsToDrainCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceListUserSecretsToDrainCall) DoAndReturn(f func(context.Context) ([]*secrets.SecretMetadataForDrain, error)) *MockSecretServiceListUserSecretsToDrainCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchSecretBackendChanged mocks base method.
func (m *MockSecretService) WatchSecretBackendChanged(arg0 context.Context) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchSecretBackendChanged", arg0)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchSecretBackendChanged indicates an expected call of WatchSecretBackendChanged.
func (mr *MockSecretServiceMockRecorder) WatchSecretBackendChanged(arg0 any) *MockSecretServiceWatchSecretBackendChangedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchSecretBackendChanged", reflect.TypeOf((*MockSecretService)(nil).WatchSecretBackendChanged), arg0)
	return &MockSecretServiceWatchSecretBackendChangedCall{Call: call}
}

// MockSecretServiceWatchSecretBackendChangedCall wrap *gomock.Call
type MockSecretServiceWatchSecretBackendChangedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceWatchSecretBackendChangedCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockSecretServiceWatchSecretBackendChangedCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceWatchSecretBackendChangedCall) Do(f func(context.Context) (watcher.Watcher[struct{}], error)) *MockSecretServiceWatchSecretBackendChangedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceWatchSecretBackendChangedCall) DoAndReturn(f func(context.Context) (watcher.Watcher[struct{}], error)) *MockSecretServiceWatchSecretBackendChangedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockSecretBackendService is a mock of SecretBackendService interface.
type MockSecretBackendService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretBackendServiceMockRecorder
}

// MockSecretBackendServiceMockRecorder is the mock recorder for MockSecretBackendService.
type MockSecretBackendServiceMockRecorder struct {
	mock *MockSecretBackendService
}

// NewMockSecretBackendService creates a new mock instance.
func NewMockSecretBackendService(ctrl *gomock.Controller) *MockSecretBackendService {
	mock := &MockSecretBackendService{ctrl: ctrl}
	mock.recorder = &MockSecretBackendServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretBackendService) EXPECT() *MockSecretBackendServiceMockRecorder {
	return m.recorder
}

// GetRevisionsToDrain mocks base method.
func (m *MockSecretBackendService) GetRevisionsToDrain(arg0 context.Context, arg1 model.UUID, arg2 []secrets.SecretExternalRevision) ([]service0.RevisionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRevisionsToDrain", arg0, arg1, arg2)
	ret0, _ := ret[0].([]service0.RevisionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRevisionsToDrain indicates an expected call of GetRevisionsToDrain.
func (mr *MockSecretBackendServiceMockRecorder) GetRevisionsToDrain(arg0, arg1, arg2 any) *MockSecretBackendServiceGetRevisionsToDrainCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRevisionsToDrain", reflect.TypeOf((*MockSecretBackendService)(nil).GetRevisionsToDrain), arg0, arg1, arg2)
	return &MockSecretBackendServiceGetRevisionsToDrainCall{Call: call}
}

// MockSecretBackendServiceGetRevisionsToDrainCall wrap *gomock.Call
type MockSecretBackendServiceGetRevisionsToDrainCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretBackendServiceGetRevisionsToDrainCall) Return(arg0 []service0.RevisionInfo, arg1 error) *MockSecretBackendServiceGetRevisionsToDrainCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretBackendServiceGetRevisionsToDrainCall) Do(f func(context.Context, model.UUID, []secrets.SecretExternalRevision) ([]service0.RevisionInfo, error)) *MockSecretBackendServiceGetRevisionsToDrainCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretBackendServiceGetRevisionsToDrainCall) DoAndReturn(f func(context.Context, model.UUID, []secrets.SecretExternalRevision) ([]service0.RevisionInfo, error)) *MockSecretBackendServiceGetRevisionsToDrainCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
