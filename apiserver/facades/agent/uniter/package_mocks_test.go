// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/uniter (interfaces: LXDProfileBackend,LXDProfileMachine,LXDProfileUnit,LXDProfileBackendV2,LXDProfileMachineV2,LXDProfileUnitV2,LXDProfileCharmV2,LXDProfileModelV2,NetworkService)
//
// Generated by this command:
//
//	mockgen -package uniter_test -destination package_mocks_test.go github.com/juju/juju/apiserver/facades/agent/uniter LXDProfileBackend,LXDProfileMachine,LXDProfileUnit,LXDProfileBackendV2,LXDProfileMachineV2,LXDProfileUnitV2,LXDProfileCharmV2,LXDProfileModelV2,NetworkService
//

// Package uniter_test is a generated GoMock package.
package uniter_test

import (
	context "context"
	reflect "reflect"

	uniter "github.com/juju/juju/apiserver/facades/agent/uniter"
	instance "github.com/juju/juju/core/instance"
	lxdprofile "github.com/juju/juju/core/lxdprofile"
	network "github.com/juju/juju/core/network"
	config "github.com/juju/juju/environs/config"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockLXDProfileBackend is a mock of LXDProfileBackend interface.
type MockLXDProfileBackend struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileBackendMockRecorder
}

// MockLXDProfileBackendMockRecorder is the mock recorder for MockLXDProfileBackend.
type MockLXDProfileBackendMockRecorder struct {
	mock *MockLXDProfileBackend
}

// NewMockLXDProfileBackend creates a new mock instance.
func NewMockLXDProfileBackend(ctrl *gomock.Controller) *MockLXDProfileBackend {
	mock := &MockLXDProfileBackend{ctrl: ctrl}
	mock.recorder = &MockLXDProfileBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileBackend) EXPECT() *MockLXDProfileBackendMockRecorder {
	return m.recorder
}

// Machine mocks base method.
func (m *MockLXDProfileBackend) Machine(arg0 string) (uniter.LXDProfileMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(uniter.LXDProfileMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockLXDProfileBackendMockRecorder) Machine(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockLXDProfileBackend)(nil).Machine), arg0)
}

// Unit mocks base method.
func (m *MockLXDProfileBackend) Unit(arg0 string) (uniter.LXDProfileUnit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(uniter.LXDProfileUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockLXDProfileBackendMockRecorder) Unit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockLXDProfileBackend)(nil).Unit), arg0)
}

// MockLXDProfileMachine is a mock of LXDProfileMachine interface.
type MockLXDProfileMachine struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileMachineMockRecorder
}

// MockLXDProfileMachineMockRecorder is the mock recorder for MockLXDProfileMachine.
type MockLXDProfileMachineMockRecorder struct {
	mock *MockLXDProfileMachine
}

// NewMockLXDProfileMachine creates a new mock instance.
func NewMockLXDProfileMachine(ctrl *gomock.Controller) *MockLXDProfileMachine {
	mock := &MockLXDProfileMachine{ctrl: ctrl}
	mock.recorder = &MockLXDProfileMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileMachine) EXPECT() *MockLXDProfileMachineMockRecorder {
	return m.recorder
}

// WatchLXDProfileUpgradeNotifications mocks base method.
func (m *MockLXDProfileMachine) WatchLXDProfileUpgradeNotifications(arg0 string) (state.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchLXDProfileUpgradeNotifications", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchLXDProfileUpgradeNotifications indicates an expected call of WatchLXDProfileUpgradeNotifications.
func (mr *MockLXDProfileMachineMockRecorder) WatchLXDProfileUpgradeNotifications(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchLXDProfileUpgradeNotifications", reflect.TypeOf((*MockLXDProfileMachine)(nil).WatchLXDProfileUpgradeNotifications), arg0)
}

// MockLXDProfileUnit is a mock of LXDProfileUnit interface.
type MockLXDProfileUnit struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileUnitMockRecorder
}

// MockLXDProfileUnitMockRecorder is the mock recorder for MockLXDProfileUnit.
type MockLXDProfileUnitMockRecorder struct {
	mock *MockLXDProfileUnit
}

// NewMockLXDProfileUnit creates a new mock instance.
func NewMockLXDProfileUnit(ctrl *gomock.Controller) *MockLXDProfileUnit {
	mock := &MockLXDProfileUnit{ctrl: ctrl}
	mock.recorder = &MockLXDProfileUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileUnit) EXPECT() *MockLXDProfileUnitMockRecorder {
	return m.recorder
}

// AssignedMachineId mocks base method.
func (m *MockLXDProfileUnit) AssignedMachineId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignedMachineId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignedMachineId indicates an expected call of AssignedMachineId.
func (mr *MockLXDProfileUnitMockRecorder) AssignedMachineId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedMachineId", reflect.TypeOf((*MockLXDProfileUnit)(nil).AssignedMachineId))
}

// Name mocks base method.
func (m *MockLXDProfileUnit) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockLXDProfileUnitMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockLXDProfileUnit)(nil).Name))
}

// Tag mocks base method.
func (m *MockLXDProfileUnit) Tag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockLXDProfileUnitMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockLXDProfileUnit)(nil).Tag))
}

// WatchLXDProfileUpgradeNotifications mocks base method.
func (m *MockLXDProfileUnit) WatchLXDProfileUpgradeNotifications() (state.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchLXDProfileUpgradeNotifications")
	ret0, _ := ret[0].(state.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchLXDProfileUpgradeNotifications indicates an expected call of WatchLXDProfileUpgradeNotifications.
func (mr *MockLXDProfileUnitMockRecorder) WatchLXDProfileUpgradeNotifications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchLXDProfileUpgradeNotifications", reflect.TypeOf((*MockLXDProfileUnit)(nil).WatchLXDProfileUpgradeNotifications))
}

// MockLXDProfileBackendV2 is a mock of LXDProfileBackendV2 interface.
type MockLXDProfileBackendV2 struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileBackendV2MockRecorder
}

// MockLXDProfileBackendV2MockRecorder is the mock recorder for MockLXDProfileBackendV2.
type MockLXDProfileBackendV2MockRecorder struct {
	mock *MockLXDProfileBackendV2
}

// NewMockLXDProfileBackendV2 creates a new mock instance.
func NewMockLXDProfileBackendV2(ctrl *gomock.Controller) *MockLXDProfileBackendV2 {
	mock := &MockLXDProfileBackendV2{ctrl: ctrl}
	mock.recorder = &MockLXDProfileBackendV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileBackendV2) EXPECT() *MockLXDProfileBackendV2MockRecorder {
	return m.recorder
}

// Charm mocks base method.
func (m *MockLXDProfileBackendV2) Charm(arg0 string) (uniter.LXDProfileCharmV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm", arg0)
	ret0, _ := ret[0].(uniter.LXDProfileCharmV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Charm indicates an expected call of Charm.
func (mr *MockLXDProfileBackendV2MockRecorder) Charm(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockLXDProfileBackendV2)(nil).Charm), arg0)
}

// Machine mocks base method.
func (m *MockLXDProfileBackendV2) Machine(arg0 string) (uniter.LXDProfileMachineV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(uniter.LXDProfileMachineV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockLXDProfileBackendV2MockRecorder) Machine(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockLXDProfileBackendV2)(nil).Machine), arg0)
}

// Model mocks base method.
func (m *MockLXDProfileBackendV2) Model() (uniter.LXDProfileModelV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(uniter.LXDProfileModelV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockLXDProfileBackendV2MockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockLXDProfileBackendV2)(nil).Model))
}

// Unit mocks base method.
func (m *MockLXDProfileBackendV2) Unit(arg0 string) (uniter.LXDProfileUnitV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(uniter.LXDProfileUnitV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockLXDProfileBackendV2MockRecorder) Unit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockLXDProfileBackendV2)(nil).Unit), arg0)
}

// MockLXDProfileMachineV2 is a mock of LXDProfileMachineV2 interface.
type MockLXDProfileMachineV2 struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileMachineV2MockRecorder
}

// MockLXDProfileMachineV2MockRecorder is the mock recorder for MockLXDProfileMachineV2.
type MockLXDProfileMachineV2MockRecorder struct {
	mock *MockLXDProfileMachineV2
}

// NewMockLXDProfileMachineV2 creates a new mock instance.
func NewMockLXDProfileMachineV2(ctrl *gomock.Controller) *MockLXDProfileMachineV2 {
	mock := &MockLXDProfileMachineV2{ctrl: ctrl}
	mock.recorder = &MockLXDProfileMachineV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileMachineV2) EXPECT() *MockLXDProfileMachineV2MockRecorder {
	return m.recorder
}

// CharmProfiles mocks base method.
func (m *MockLXDProfileMachineV2) CharmProfiles() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmProfiles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmProfiles indicates an expected call of CharmProfiles.
func (mr *MockLXDProfileMachineV2MockRecorder) CharmProfiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmProfiles", reflect.TypeOf((*MockLXDProfileMachineV2)(nil).CharmProfiles))
}

// ContainerType mocks base method.
func (m *MockLXDProfileMachineV2) ContainerType() instance.ContainerType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerType")
	ret0, _ := ret[0].(instance.ContainerType)
	return ret0
}

// ContainerType indicates an expected call of ContainerType.
func (mr *MockLXDProfileMachineV2MockRecorder) ContainerType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerType", reflect.TypeOf((*MockLXDProfileMachineV2)(nil).ContainerType))
}

// IsManual mocks base method.
func (m *MockLXDProfileMachineV2) IsManual() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsManual")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsManual indicates an expected call of IsManual.
func (mr *MockLXDProfileMachineV2MockRecorder) IsManual() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsManual", reflect.TypeOf((*MockLXDProfileMachineV2)(nil).IsManual))
}

// WatchInstanceData mocks base method.
func (m *MockLXDProfileMachineV2) WatchInstanceData() state.NotifyWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchInstanceData")
	ret0, _ := ret[0].(state.NotifyWatcher)
	return ret0
}

// WatchInstanceData indicates an expected call of WatchInstanceData.
func (mr *MockLXDProfileMachineV2MockRecorder) WatchInstanceData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchInstanceData", reflect.TypeOf((*MockLXDProfileMachineV2)(nil).WatchInstanceData))
}

// MockLXDProfileUnitV2 is a mock of LXDProfileUnitV2 interface.
type MockLXDProfileUnitV2 struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileUnitV2MockRecorder
}

// MockLXDProfileUnitV2MockRecorder is the mock recorder for MockLXDProfileUnitV2.
type MockLXDProfileUnitV2MockRecorder struct {
	mock *MockLXDProfileUnitV2
}

// NewMockLXDProfileUnitV2 creates a new mock instance.
func NewMockLXDProfileUnitV2(ctrl *gomock.Controller) *MockLXDProfileUnitV2 {
	mock := &MockLXDProfileUnitV2{ctrl: ctrl}
	mock.recorder = &MockLXDProfileUnitV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileUnitV2) EXPECT() *MockLXDProfileUnitV2MockRecorder {
	return m.recorder
}

// ApplicationName mocks base method.
func (m *MockLXDProfileUnitV2) ApplicationName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ApplicationName indicates an expected call of ApplicationName.
func (mr *MockLXDProfileUnitV2MockRecorder) ApplicationName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationName", reflect.TypeOf((*MockLXDProfileUnitV2)(nil).ApplicationName))
}

// AssignedMachineId mocks base method.
func (m *MockLXDProfileUnitV2) AssignedMachineId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignedMachineId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignedMachineId indicates an expected call of AssignedMachineId.
func (mr *MockLXDProfileUnitV2MockRecorder) AssignedMachineId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedMachineId", reflect.TypeOf((*MockLXDProfileUnitV2)(nil).AssignedMachineId))
}

// CharmURL mocks base method.
func (m *MockLXDProfileUnitV2) CharmURL() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmURL")
	ret0, _ := ret[0].(*string)
	return ret0
}

// CharmURL indicates an expected call of CharmURL.
func (mr *MockLXDProfileUnitV2MockRecorder) CharmURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmURL", reflect.TypeOf((*MockLXDProfileUnitV2)(nil).CharmURL))
}

// Name mocks base method.
func (m *MockLXDProfileUnitV2) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockLXDProfileUnitV2MockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockLXDProfileUnitV2)(nil).Name))
}

// Tag mocks base method.
func (m *MockLXDProfileUnitV2) Tag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockLXDProfileUnitV2MockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockLXDProfileUnitV2)(nil).Tag))
}

// MockLXDProfileCharmV2 is a mock of LXDProfileCharmV2 interface.
type MockLXDProfileCharmV2 struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileCharmV2MockRecorder
}

// MockLXDProfileCharmV2MockRecorder is the mock recorder for MockLXDProfileCharmV2.
type MockLXDProfileCharmV2MockRecorder struct {
	mock *MockLXDProfileCharmV2
}

// NewMockLXDProfileCharmV2 creates a new mock instance.
func NewMockLXDProfileCharmV2(ctrl *gomock.Controller) *MockLXDProfileCharmV2 {
	mock := &MockLXDProfileCharmV2{ctrl: ctrl}
	mock.recorder = &MockLXDProfileCharmV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileCharmV2) EXPECT() *MockLXDProfileCharmV2MockRecorder {
	return m.recorder
}

// LXDProfile mocks base method.
func (m *MockLXDProfileCharmV2) LXDProfile() lxdprofile.Profile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(lxdprofile.Profile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile.
func (mr *MockLXDProfileCharmV2MockRecorder) LXDProfile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockLXDProfileCharmV2)(nil).LXDProfile))
}

// MockLXDProfileModelV2 is a mock of LXDProfileModelV2 interface.
type MockLXDProfileModelV2 struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileModelV2MockRecorder
}

// MockLXDProfileModelV2MockRecorder is the mock recorder for MockLXDProfileModelV2.
type MockLXDProfileModelV2MockRecorder struct {
	mock *MockLXDProfileModelV2
}

// NewMockLXDProfileModelV2 creates a new mock instance.
func NewMockLXDProfileModelV2(ctrl *gomock.Controller) *MockLXDProfileModelV2 {
	mock := &MockLXDProfileModelV2{ctrl: ctrl}
	mock.recorder = &MockLXDProfileModelV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfileModelV2) EXPECT() *MockLXDProfileModelV2MockRecorder {
	return m.recorder
}

// ModelConfig mocks base method.
func (m *MockLXDProfileModelV2) ModelConfig(arg0 context.Context) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig", arg0)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockLXDProfileModelV2MockRecorder) ModelConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockLXDProfileModelV2)(nil).ModelConfig), arg0)
}

// Type mocks base method.
func (m *MockLXDProfileModelV2) Type() state.ModelType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(state.ModelType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockLXDProfileModelV2MockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockLXDProfileModelV2)(nil).Type))
}

// MockNetworkService is a mock of NetworkService interface.
type MockNetworkService struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceMockRecorder
}

// MockNetworkServiceMockRecorder is the mock recorder for MockNetworkService.
type MockNetworkServiceMockRecorder struct {
	mock *MockNetworkService
}

// NewMockNetworkService creates a new mock instance.
func NewMockNetworkService(ctrl *gomock.Controller) *MockNetworkService {
	mock := &MockNetworkService{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkService) EXPECT() *MockNetworkServiceMockRecorder {
	return m.recorder
}

// GetAllSubnets mocks base method.
func (m *MockNetworkService) GetAllSubnets(arg0 context.Context) (network.SubnetInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubnets", arg0)
	ret0, _ := ret[0].(network.SubnetInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubnets indicates an expected call of GetAllSubnets.
func (mr *MockNetworkServiceMockRecorder) GetAllSubnets(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubnets", reflect.TypeOf((*MockNetworkService)(nil).GetAllSubnets), arg0)
}

// SpaceByName mocks base method.
func (m *MockNetworkService) SpaceByName(arg0 context.Context, arg1 string) (*network.SpaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceByName", arg0, arg1)
	ret0, _ := ret[0].(*network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceByName indicates an expected call of SpaceByName.
func (mr *MockNetworkServiceMockRecorder) SpaceByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceByName", reflect.TypeOf((*MockNetworkService)(nil).SpaceByName), arg0, arg1)
}
