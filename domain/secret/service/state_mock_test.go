// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/secret/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -package service -destination state_mock_test.go github.com/juju/juju/domain/secret/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	secret "github.com/juju/juju/domain/secret"
	gomock "go.uber.org/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// CreateUserSecret mocks base method.
func (m *MockState) CreateUserSecret(arg0 context.Context, arg1 int, arg2 *secrets.URI, arg3 secret.UpsertSecretParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserSecret", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserSecret indicates an expected call of CreateUserSecret.
func (mr *MockStateMockRecorder) CreateUserSecret(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserSecret", reflect.TypeOf((*MockState)(nil).CreateUserSecret), arg0, arg1, arg2, arg3)
}

// GetModelUUID mocks base method.
func (m *MockState) GetModelUUID(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelUUID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelUUID indicates an expected call of GetModelUUID.
func (mr *MockStateMockRecorder) GetModelUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelUUID", reflect.TypeOf((*MockState)(nil).GetModelUUID), arg0)
}

// GetSecret mocks base method.
func (m *MockState) GetSecret(arg0 context.Context, arg1 *secrets.URI) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockStateMockRecorder) GetSecret(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockState)(nil).GetSecret), arg0, arg1)
}

// GetSecretRevision mocks base method.
func (m *MockState) GetSecretRevision(arg0 context.Context, arg1 *secrets.URI, arg2 int) (*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretRevision", arg0, arg1, arg2)
	ret0, _ := ret[0].(*secrets.SecretRevisionMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretRevision indicates an expected call of GetSecretRevision.
func (mr *MockStateMockRecorder) GetSecretRevision(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretRevision", reflect.TypeOf((*MockState)(nil).GetSecretRevision), arg0, arg1, arg2)
}

// GetSecretValue mocks base method.
func (m *MockState) GetSecretValue(arg0 context.Context, arg1 *secrets.URI, arg2 int) (secrets.SecretData, *secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretValue", arg0, arg1, arg2)
	ret0, _ := ret[0].(secrets.SecretData)
	ret1, _ := ret[1].(*secrets.ValueRef)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSecretValue indicates an expected call of GetSecretValue.
func (mr *MockStateMockRecorder) GetSecretValue(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockState)(nil).GetSecretValue), arg0, arg1, arg2)
}
