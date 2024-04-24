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
	eventsource "github.com/juju/juju/core/watcher/eventsource"
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

// CreateCharmApplicationSecret mocks base method.
func (m *MockState) CreateCharmApplicationSecret(arg0 context.Context, arg1 int, arg2 *secrets.URI, arg3 string, arg4 secret.UpsertSecretParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCharmApplicationSecret", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCharmApplicationSecret indicates an expected call of CreateCharmApplicationSecret.
func (mr *MockStateMockRecorder) CreateCharmApplicationSecret(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCharmApplicationSecret", reflect.TypeOf((*MockState)(nil).CreateCharmApplicationSecret), arg0, arg1, arg2, arg3, arg4)
}

// CreateCharmUnitSecret mocks base method.
func (m *MockState) CreateCharmUnitSecret(arg0 context.Context, arg1 int, arg2 *secrets.URI, arg3 string, arg4 secret.UpsertSecretParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCharmUnitSecret", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCharmUnitSecret indicates an expected call of CreateCharmUnitSecret.
func (mr *MockStateMockRecorder) CreateCharmUnitSecret(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCharmUnitSecret", reflect.TypeOf((*MockState)(nil).CreateCharmUnitSecret), arg0, arg1, arg2, arg3, arg4)
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

// DeleteSecret mocks base method.
func (m *MockState) DeleteSecret(arg0 context.Context, arg1 *secrets.URI, arg2 []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockStateMockRecorder) DeleteSecret(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockState)(nil).DeleteSecret), arg0, arg1, arg2)
}

// GetConsumedSecretURIs mocks base method.
func (m *MockState) GetConsumedSecretURIs(arg0 context.Context, arg1 string, arg2 ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConsumedSecretURIs", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsumedSecretURIs indicates an expected call of GetConsumedSecretURIs.
func (mr *MockStateMockRecorder) GetConsumedSecretURIs(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsumedSecretURIs", reflect.TypeOf((*MockState)(nil).GetConsumedSecretURIs), varargs...)
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

// GetRevisionIDsForObsolete mocks base method.
func (m *MockState) GetRevisionIDsForObsolete(arg0 context.Context, arg1 secret.ApplicationOwners, arg2 secret.UnitOwners, arg3 ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRevisionIDsForObsolete", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRevisionIDsForObsolete indicates an expected call of GetRevisionIDsForObsolete.
func (mr *MockStateMockRecorder) GetRevisionIDsForObsolete(arg0, arg1, arg2 any, arg3 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRevisionIDsForObsolete", reflect.TypeOf((*MockState)(nil).GetRevisionIDsForObsolete), varargs...)
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

// GetSecretAccess mocks base method.
func (m *MockState) GetSecretAccess(arg0 context.Context, arg1 *secrets.URI, arg2 secret.AccessParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretAccess indicates an expected call of GetSecretAccess.
func (mr *MockStateMockRecorder) GetSecretAccess(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretAccess", reflect.TypeOf((*MockState)(nil).GetSecretAccess), arg0, arg1, arg2)
}

// GetSecretAccessScope mocks base method.
func (m *MockState) GetSecretAccessScope(arg0 context.Context, arg1 *secrets.URI, arg2 secret.AccessParams) (*secret.AccessScope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretAccessScope", arg0, arg1, arg2)
	ret0, _ := ret[0].(*secret.AccessScope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretAccessScope indicates an expected call of GetSecretAccessScope.
func (mr *MockStateMockRecorder) GetSecretAccessScope(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretAccessScope", reflect.TypeOf((*MockState)(nil).GetSecretAccessScope), arg0, arg1, arg2)
}

// GetSecretConsumer mocks base method.
func (m *MockState) GetSecretConsumer(arg0 context.Context, arg1 *secrets.URI, arg2 string) (*secrets.SecretConsumerMetadata, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretConsumer", arg0, arg1, arg2)
	ret0, _ := ret[0].(*secrets.SecretConsumerMetadata)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSecretConsumer indicates an expected call of GetSecretConsumer.
func (mr *MockStateMockRecorder) GetSecretConsumer(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretConsumer", reflect.TypeOf((*MockState)(nil).GetSecretConsumer), arg0, arg1, arg2)
}

// GetSecretGrants mocks base method.
func (m *MockState) GetSecretGrants(arg0 context.Context, arg1 *secrets.URI, arg2 secrets.SecretRole) ([]secret.GrantParams, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretGrants", arg0, arg1, arg2)
	ret0, _ := ret[0].([]secret.GrantParams)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretGrants indicates an expected call of GetSecretGrants.
func (mr *MockStateMockRecorder) GetSecretGrants(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretGrants", reflect.TypeOf((*MockState)(nil).GetSecretGrants), arg0, arg1, arg2)
}

// GetSecretRemoteConsumer mocks base method.
func (m *MockState) GetSecretRemoteConsumer(arg0 context.Context, arg1 *secrets.URI, arg2 string) (*secrets.SecretConsumerMetadata, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretRemoteConsumer", arg0, arg1, arg2)
	ret0, _ := ret[0].(*secrets.SecretConsumerMetadata)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSecretRemoteConsumer indicates an expected call of GetSecretRemoteConsumer.
func (mr *MockStateMockRecorder) GetSecretRemoteConsumer(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretRemoteConsumer", reflect.TypeOf((*MockState)(nil).GetSecretRemoteConsumer), arg0, arg1, arg2)
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

// GetURIByConsumerLabel mocks base method.
func (m *MockState) GetURIByConsumerLabel(arg0 context.Context, arg1, arg2 string) (*secrets.URI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURIByConsumerLabel", arg0, arg1, arg2)
	ret0, _ := ret[0].(*secrets.URI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURIByConsumerLabel indicates an expected call of GetURIByConsumerLabel.
func (mr *MockStateMockRecorder) GetURIByConsumerLabel(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURIByConsumerLabel", reflect.TypeOf((*MockState)(nil).GetURIByConsumerLabel), arg0, arg1, arg2)
}

// GetUserSecretURIByLabel mocks base method.
func (m *MockState) GetUserSecretURIByLabel(arg0 context.Context, arg1 string) (*secrets.URI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSecretURIByLabel", arg0, arg1)
	ret0, _ := ret[0].(*secrets.URI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSecretURIByLabel indicates an expected call of GetUserSecretURIByLabel.
func (mr *MockStateMockRecorder) GetUserSecretURIByLabel(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSecretURIByLabel", reflect.TypeOf((*MockState)(nil).GetUserSecretURIByLabel), arg0, arg1)
}

// GrantAccess mocks base method.
func (m *MockState) GrantAccess(arg0 context.Context, arg1 *secrets.URI, arg2 secret.GrantParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantAccess indicates an expected call of GrantAccess.
func (mr *MockStateMockRecorder) GrantAccess(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantAccess", reflect.TypeOf((*MockState)(nil).GrantAccess), arg0, arg1, arg2)
}

// InitialWatchStatementForConsumedSecrets mocks base method.
func (m *MockState) InitialWatchStatementForConsumedSecrets(arg0 string) (string, eventsource.NamespaceQuery) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitialWatchStatementForConsumedSecrets", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(eventsource.NamespaceQuery)
	return ret0, ret1
}

// InitialWatchStatementForConsumedSecrets indicates an expected call of InitialWatchStatementForConsumedSecrets.
func (mr *MockStateMockRecorder) InitialWatchStatementForConsumedSecrets(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitialWatchStatementForConsumedSecrets", reflect.TypeOf((*MockState)(nil).InitialWatchStatementForConsumedSecrets), arg0)
}

// InitialWatchStatementForObsoleteRevision mocks base method.
func (m *MockState) InitialWatchStatementForObsoleteRevision(arg0 secret.ApplicationOwners, arg1 secret.UnitOwners) (string, eventsource.NamespaceQuery) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitialWatchStatementForObsoleteRevision", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(eventsource.NamespaceQuery)
	return ret0, ret1
}

// InitialWatchStatementForObsoleteRevision indicates an expected call of InitialWatchStatementForObsoleteRevision.
func (mr *MockStateMockRecorder) InitialWatchStatementForObsoleteRevision(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitialWatchStatementForObsoleteRevision", reflect.TypeOf((*MockState)(nil).InitialWatchStatementForObsoleteRevision), arg0, arg1)
}

// ListCharmSecrets mocks base method.
func (m *MockState) ListCharmSecrets(arg0 context.Context, arg1 secret.ApplicationOwners, arg2 secret.UnitOwners) ([]*secrets.SecretMetadata, [][]*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCharmSecrets", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*secrets.SecretMetadata)
	ret1, _ := ret[1].([][]*secrets.SecretRevisionMetadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListCharmSecrets indicates an expected call of ListCharmSecrets.
func (mr *MockStateMockRecorder) ListCharmSecrets(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCharmSecrets", reflect.TypeOf((*MockState)(nil).ListCharmSecrets), arg0, arg1, arg2)
}

// ListCharmSecretsToDrain mocks base method.
func (m *MockState) ListCharmSecretsToDrain(arg0 context.Context, arg1 secret.ApplicationOwners, arg2 secret.UnitOwners) ([]*secrets.SecretMetadataForDrain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCharmSecretsToDrain", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*secrets.SecretMetadataForDrain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCharmSecretsToDrain indicates an expected call of ListCharmSecretsToDrain.
func (mr *MockStateMockRecorder) ListCharmSecretsToDrain(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCharmSecretsToDrain", reflect.TypeOf((*MockState)(nil).ListCharmSecretsToDrain), arg0, arg1, arg2)
}

// ListExternalSecretRevisions mocks base method.
func (m *MockState) ListExternalSecretRevisions(arg0 context.Context, arg1 *secrets.URI, arg2 ...int) ([]secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListExternalSecretRevisions", varargs...)
	ret0, _ := ret[0].([]secrets.ValueRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExternalSecretRevisions indicates an expected call of ListExternalSecretRevisions.
func (mr *MockStateMockRecorder) ListExternalSecretRevisions(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExternalSecretRevisions", reflect.TypeOf((*MockState)(nil).ListExternalSecretRevisions), varargs...)
}

// ListGrantedSecretsForBackend mocks base method.
func (m *MockState) ListGrantedSecretsForBackend(arg0 context.Context, arg1 string, arg2 []secret.AccessParams, arg3 secrets.SecretRole) ([]*secrets.SecretRevisionRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGrantedSecretsForBackend", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*secrets.SecretRevisionRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGrantedSecretsForBackend indicates an expected call of ListGrantedSecretsForBackend.
func (mr *MockStateMockRecorder) ListGrantedSecretsForBackend(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGrantedSecretsForBackend", reflect.TypeOf((*MockState)(nil).ListGrantedSecretsForBackend), arg0, arg1, arg2, arg3)
}

// ListSecrets mocks base method.
func (m *MockState) ListSecrets(arg0 context.Context, arg1 *secrets.URI, arg2 *int, arg3 secret.Labels) ([]*secrets.SecretMetadata, [][]*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*secrets.SecretMetadata)
	ret1, _ := ret[1].([][]*secrets.SecretRevisionMetadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockStateMockRecorder) ListSecrets(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockState)(nil).ListSecrets), arg0, arg1, arg2, arg3)
}

// ListUserSecretsToDrain mocks base method.
func (m *MockState) ListUserSecretsToDrain(arg0 context.Context) ([]*secrets.SecretMetadataForDrain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserSecretsToDrain", arg0)
	ret0, _ := ret[0].([]*secrets.SecretMetadataForDrain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserSecretsToDrain indicates an expected call of ListUserSecretsToDrain.
func (mr *MockStateMockRecorder) ListUserSecretsToDrain(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserSecretsToDrain", reflect.TypeOf((*MockState)(nil).ListUserSecretsToDrain), arg0)
}

// RevokeAccess mocks base method.
func (m *MockState) RevokeAccess(arg0 context.Context, arg1 *secrets.URI, arg2 secret.AccessParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeAccess indicates an expected call of RevokeAccess.
func (mr *MockStateMockRecorder) RevokeAccess(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAccess", reflect.TypeOf((*MockState)(nil).RevokeAccess), arg0, arg1, arg2)
}

// SaveSecretConsumer mocks base method.
func (m *MockState) SaveSecretConsumer(arg0 context.Context, arg1 *secrets.URI, arg2 string, arg3 *secrets.SecretConsumerMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSecretConsumer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSecretConsumer indicates an expected call of SaveSecretConsumer.
func (mr *MockStateMockRecorder) SaveSecretConsumer(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSecretConsumer", reflect.TypeOf((*MockState)(nil).SaveSecretConsumer), arg0, arg1, arg2, arg3)
}

// SaveSecretRemoteConsumer mocks base method.
func (m *MockState) SaveSecretRemoteConsumer(arg0 context.Context, arg1 *secrets.URI, arg2 string, arg3 *secrets.SecretConsumerMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSecretRemoteConsumer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSecretRemoteConsumer indicates an expected call of SaveSecretRemoteConsumer.
func (mr *MockStateMockRecorder) SaveSecretRemoteConsumer(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSecretRemoteConsumer", reflect.TypeOf((*MockState)(nil).SaveSecretRemoteConsumer), arg0, arg1, arg2, arg3)
}

// UpdateRemoteSecretRevision mocks base method.
func (m *MockState) UpdateRemoteSecretRevision(arg0 context.Context, arg1 *secrets.URI, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRemoteSecretRevision", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRemoteSecretRevision indicates an expected call of UpdateRemoteSecretRevision.
func (mr *MockStateMockRecorder) UpdateRemoteSecretRevision(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRemoteSecretRevision", reflect.TypeOf((*MockState)(nil).UpdateRemoteSecretRevision), arg0, arg1, arg2)
}

// UpdateSecret mocks base method.
func (m *MockState) UpdateSecret(arg0 context.Context, arg1 *secrets.URI, arg2 secret.UpsertSecretParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecret indicates an expected call of UpdateSecret.
func (mr *MockStateMockRecorder) UpdateSecret(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockState)(nil).UpdateSecret), arg0, arg1, arg2)
}
