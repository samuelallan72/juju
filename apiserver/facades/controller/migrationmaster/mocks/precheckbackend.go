// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/migration (interfaces: PrecheckBackend)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	migration "github.com/juju/juju/migration"
	resource "github.com/juju/juju/resource"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
	version "github.com/juju/version"
	reflect "reflect"
)

// MockPrecheckBackend is a mock of PrecheckBackend interface
type MockPrecheckBackend struct {
	ctrl     *gomock.Controller
	recorder *MockPrecheckBackendMockRecorder
}

// MockPrecheckBackendMockRecorder is the mock recorder for MockPrecheckBackend
type MockPrecheckBackendMockRecorder struct {
	mock *MockPrecheckBackend
}

// NewMockPrecheckBackend creates a new mock instance
func NewMockPrecheckBackend(ctrl *gomock.Controller) *MockPrecheckBackend {
	mock := &MockPrecheckBackend{ctrl: ctrl}
	mock.recorder = &MockPrecheckBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrecheckBackend) EXPECT() *MockPrecheckBackendMockRecorder {
	return m.recorder
}

// AgentVersion mocks base method
func (m *MockPrecheckBackend) AgentVersion() (version.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentVersion")
	ret0, _ := ret[0].(version.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentVersion indicates an expected call of AgentVersion
func (mr *MockPrecheckBackendMockRecorder) AgentVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentVersion", reflect.TypeOf((*MockPrecheckBackend)(nil).AgentVersion))
}

// AllApplications mocks base method
func (m *MockPrecheckBackend) AllApplications() ([]migration.PrecheckApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllApplications")
	ret0, _ := ret[0].([]migration.PrecheckApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllApplications indicates an expected call of AllApplications
func (mr *MockPrecheckBackendMockRecorder) AllApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllApplications", reflect.TypeOf((*MockPrecheckBackend)(nil).AllApplications))
}

// AllMachines mocks base method
func (m *MockPrecheckBackend) AllMachines() ([]migration.PrecheckMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMachines")
	ret0, _ := ret[0].([]migration.PrecheckMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMachines indicates an expected call of AllMachines
func (mr *MockPrecheckBackendMockRecorder) AllMachines() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMachines", reflect.TypeOf((*MockPrecheckBackend)(nil).AllMachines))
}

// AllModelUUIDs mocks base method
func (m *MockPrecheckBackend) AllModelUUIDs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllModelUUIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllModelUUIDs indicates an expected call of AllModelUUIDs
func (mr *MockPrecheckBackendMockRecorder) AllModelUUIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllModelUUIDs", reflect.TypeOf((*MockPrecheckBackend)(nil).AllModelUUIDs))
}

// AllRelations mocks base method
func (m *MockPrecheckBackend) AllRelations() ([]migration.PrecheckRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRelations")
	ret0, _ := ret[0].([]migration.PrecheckRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRelations indicates an expected call of AllRelations
func (mr *MockPrecheckBackendMockRecorder) AllRelations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRelations", reflect.TypeOf((*MockPrecheckBackend)(nil).AllRelations))
}

// CloudCredential mocks base method
func (m *MockPrecheckBackend) CloudCredential(arg0 names.CloudCredentialTag) (state.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudCredential", arg0)
	ret0, _ := ret[0].(state.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudCredential indicates an expected call of CloudCredential
func (mr *MockPrecheckBackendMockRecorder) CloudCredential(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudCredential", reflect.TypeOf((*MockPrecheckBackend)(nil).CloudCredential), arg0)
}

// ControllerBackend mocks base method
func (m *MockPrecheckBackend) ControllerBackend() (migration.PrecheckBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerBackend")
	ret0, _ := ret[0].(migration.PrecheckBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerBackend indicates an expected call of ControllerBackend
func (mr *MockPrecheckBackendMockRecorder) ControllerBackend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerBackend", reflect.TypeOf((*MockPrecheckBackend)(nil).ControllerBackend))
}

// IsMigrationActive mocks base method
func (m *MockPrecheckBackend) IsMigrationActive(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMigrationActive", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMigrationActive indicates an expected call of IsMigrationActive
func (mr *MockPrecheckBackendMockRecorder) IsMigrationActive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMigrationActive", reflect.TypeOf((*MockPrecheckBackend)(nil).IsMigrationActive), arg0)
}

// IsUpgrading mocks base method
func (m *MockPrecheckBackend) IsUpgrading() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUpgrading")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUpgrading indicates an expected call of IsUpgrading
func (mr *MockPrecheckBackendMockRecorder) IsUpgrading() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUpgrading", reflect.TypeOf((*MockPrecheckBackend)(nil).IsUpgrading))
}

// ListPendingResources mocks base method
func (m *MockPrecheckBackend) ListPendingResources(arg0 string) ([]resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPendingResources", arg0)
	ret0, _ := ret[0].([]resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPendingResources indicates an expected call of ListPendingResources
func (mr *MockPrecheckBackendMockRecorder) ListPendingResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPendingResources", reflect.TypeOf((*MockPrecheckBackend)(nil).ListPendingResources), arg0)
}

// Model mocks base method
func (m *MockPrecheckBackend) Model() (migration.PrecheckModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(migration.PrecheckModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model
func (mr *MockPrecheckBackendMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockPrecheckBackend)(nil).Model))
}

// NeedsCleanup mocks base method
func (m *MockPrecheckBackend) NeedsCleanup() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedsCleanup")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NeedsCleanup indicates an expected call of NeedsCleanup
func (mr *MockPrecheckBackendMockRecorder) NeedsCleanup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedsCleanup", reflect.TypeOf((*MockPrecheckBackend)(nil).NeedsCleanup))
}
