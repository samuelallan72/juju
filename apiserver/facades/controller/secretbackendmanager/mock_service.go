// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -typed -package secretbackendmanager -destination mock_service.go -source service.go BackendService
//

// Package secretbackendmanager is a generated GoMock package.
package secretbackendmanager

import (
	context "context"
	reflect "reflect"

	watcher "github.com/juju/juju/core/watcher"
	gomock "go.uber.org/mock/gomock"
)

// MockBackendService is a mock of BackendService interface.
type MockBackendService struct {
	ctrl     *gomock.Controller
	recorder *MockBackendServiceMockRecorder
}

// MockBackendServiceMockRecorder is the mock recorder for MockBackendService.
type MockBackendServiceMockRecorder struct {
	mock *MockBackendService
}

// NewMockBackendService creates a new mock instance.
func NewMockBackendService(ctrl *gomock.Controller) *MockBackendService {
	mock := &MockBackendService{ctrl: ctrl}
	mock.recorder = &MockBackendServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendService) EXPECT() *MockBackendServiceMockRecorder {
	return m.recorder
}

// RotateBackendToken mocks base method.
func (m *MockBackendService) RotateBackendToken(ctx context.Context, backendID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RotateBackendToken", ctx, backendID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RotateBackendToken indicates an expected call of RotateBackendToken.
func (mr *MockBackendServiceMockRecorder) RotateBackendToken(ctx, backendID any) *MockBackendServiceRotateBackendTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateBackendToken", reflect.TypeOf((*MockBackendService)(nil).RotateBackendToken), ctx, backendID)
	return &MockBackendServiceRotateBackendTokenCall{Call: call}
}

// MockBackendServiceRotateBackendTokenCall wrap *gomock.Call
type MockBackendServiceRotateBackendTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendServiceRotateBackendTokenCall) Return(arg0 error) *MockBackendServiceRotateBackendTokenCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendServiceRotateBackendTokenCall) Do(f func(context.Context, string) error) *MockBackendServiceRotateBackendTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendServiceRotateBackendTokenCall) DoAndReturn(f func(context.Context, string) error) *MockBackendServiceRotateBackendTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchSecretBackendRotationChanges mocks base method.
func (m *MockBackendService) WatchSecretBackendRotationChanges() (watcher.SecretBackendRotateWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchSecretBackendRotationChanges")
	ret0, _ := ret[0].(watcher.SecretBackendRotateWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchSecretBackendRotationChanges indicates an expected call of WatchSecretBackendRotationChanges.
func (mr *MockBackendServiceMockRecorder) WatchSecretBackendRotationChanges() *MockBackendServiceWatchSecretBackendRotationChangesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchSecretBackendRotationChanges", reflect.TypeOf((*MockBackendService)(nil).WatchSecretBackendRotationChanges))
	return &MockBackendServiceWatchSecretBackendRotationChangesCall{Call: call}
}

// MockBackendServiceWatchSecretBackendRotationChangesCall wrap *gomock.Call
type MockBackendServiceWatchSecretBackendRotationChangesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendServiceWatchSecretBackendRotationChangesCall) Return(arg0 watcher.SecretBackendRotateWatcher, arg1 error) *MockBackendServiceWatchSecretBackendRotationChangesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendServiceWatchSecretBackendRotationChangesCall) Do(f func() (watcher.SecretBackendRotateWatcher, error)) *MockBackendServiceWatchSecretBackendRotationChangesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendServiceWatchSecretBackendRotationChangesCall) DoAndReturn(f func() (watcher.SecretBackendRotateWatcher, error)) *MockBackendServiceWatchSecretBackendRotationChangesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
