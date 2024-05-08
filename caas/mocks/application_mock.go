// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/caas (interfaces: Application)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/application_mock.go github.com/juju/juju/caas Application
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	caas "github.com/juju/juju/caas"
	watcher "github.com/juju/juju/core/watcher"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// ApplicationPodSpec mocks base method.
func (m *MockApplication) ApplicationPodSpec(arg0 caas.ApplicationConfig) (*v1.PodSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationPodSpec", arg0)
	ret0, _ := ret[0].(*v1.PodSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationPodSpec indicates an expected call of ApplicationPodSpec.
func (mr *MockApplicationMockRecorder) ApplicationPodSpec(arg0 any) *MockApplicationApplicationPodSpecCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationPodSpec", reflect.TypeOf((*MockApplication)(nil).ApplicationPodSpec), arg0)
	return &MockApplicationApplicationPodSpecCall{Call: call}
}

// MockApplicationApplicationPodSpecCall wrap *gomock.Call
type MockApplicationApplicationPodSpecCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationApplicationPodSpecCall) Return(arg0 *v1.PodSpec, arg1 error) *MockApplicationApplicationPodSpecCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationApplicationPodSpecCall) Do(f func(caas.ApplicationConfig) (*v1.PodSpec, error)) *MockApplicationApplicationPodSpecCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationApplicationPodSpecCall) DoAndReturn(f func(caas.ApplicationConfig) (*v1.PodSpec, error)) *MockApplicationApplicationPodSpecCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Delete mocks base method.
func (m *MockApplication) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockApplicationMockRecorder) Delete() *MockApplicationDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockApplication)(nil).Delete))
	return &MockApplicationDeleteCall{Call: call}
}

// MockApplicationDeleteCall wrap *gomock.Call
type MockApplicationDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationDeleteCall) Return(arg0 error) *MockApplicationDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationDeleteCall) Do(f func() error) *MockApplicationDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationDeleteCall) DoAndReturn(f func() error) *MockApplicationDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Ensure mocks base method.
func (m *MockApplication) Ensure(arg0 caas.ApplicationConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ensure", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ensure indicates an expected call of Ensure.
func (mr *MockApplicationMockRecorder) Ensure(arg0 any) *MockApplicationEnsureCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ensure", reflect.TypeOf((*MockApplication)(nil).Ensure), arg0)
	return &MockApplicationEnsureCall{Call: call}
}

// MockApplicationEnsureCall wrap *gomock.Call
type MockApplicationEnsureCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationEnsureCall) Return(arg0 error) *MockApplicationEnsureCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationEnsureCall) Do(f func(caas.ApplicationConfig) error) *MockApplicationEnsureCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationEnsureCall) DoAndReturn(f func(caas.ApplicationConfig) error) *MockApplicationEnsureCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Exists mocks base method.
func (m *MockApplication) Exists() (caas.DeploymentState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists")
	ret0, _ := ret[0].(caas.DeploymentState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockApplicationMockRecorder) Exists() *MockApplicationExistsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockApplication)(nil).Exists))
	return &MockApplicationExistsCall{Call: call}
}

// MockApplicationExistsCall wrap *gomock.Call
type MockApplicationExistsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationExistsCall) Return(arg0 caas.DeploymentState, arg1 error) *MockApplicationExistsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationExistsCall) Do(f func() (caas.DeploymentState, error)) *MockApplicationExistsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationExistsCall) DoAndReturn(f func() (caas.DeploymentState, error)) *MockApplicationExistsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Scale mocks base method.
func (m *MockApplication) Scale(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scale", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scale indicates an expected call of Scale.
func (mr *MockApplicationMockRecorder) Scale(arg0 any) *MockApplicationScaleCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scale", reflect.TypeOf((*MockApplication)(nil).Scale), arg0)
	return &MockApplicationScaleCall{Call: call}
}

// MockApplicationScaleCall wrap *gomock.Call
type MockApplicationScaleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationScaleCall) Return(arg0 error) *MockApplicationScaleCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationScaleCall) Do(f func(int) error) *MockApplicationScaleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationScaleCall) DoAndReturn(f func(int) error) *MockApplicationScaleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Service mocks base method.
func (m *MockApplication) Service() (*caas.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Service")
	ret0, _ := ret[0].(*caas.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Service indicates an expected call of Service.
func (mr *MockApplicationMockRecorder) Service() *MockApplicationServiceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Service", reflect.TypeOf((*MockApplication)(nil).Service))
	return &MockApplicationServiceCall{Call: call}
}

// MockApplicationServiceCall wrap *gomock.Call
type MockApplicationServiceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceCall) Return(arg0 *caas.Service, arg1 error) *MockApplicationServiceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceCall) Do(f func() (*caas.Service, error)) *MockApplicationServiceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceCall) DoAndReturn(f func() (*caas.Service, error)) *MockApplicationServiceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// State mocks base method.
func (m *MockApplication) State() (caas.ApplicationState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(caas.ApplicationState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// State indicates an expected call of State.
func (mr *MockApplicationMockRecorder) State() *MockApplicationStateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockApplication)(nil).State))
	return &MockApplicationStateCall{Call: call}
}

// MockApplicationStateCall wrap *gomock.Call
type MockApplicationStateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationStateCall) Return(arg0 caas.ApplicationState, arg1 error) *MockApplicationStateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationStateCall) Do(f func() (caas.ApplicationState, error)) *MockApplicationStateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationStateCall) DoAndReturn(f func() (caas.ApplicationState, error)) *MockApplicationStateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Trust mocks base method.
func (m *MockApplication) Trust(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trust", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Trust indicates an expected call of Trust.
func (mr *MockApplicationMockRecorder) Trust(arg0 any) *MockApplicationTrustCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trust", reflect.TypeOf((*MockApplication)(nil).Trust), arg0)
	return &MockApplicationTrustCall{Call: call}
}

// MockApplicationTrustCall wrap *gomock.Call
type MockApplicationTrustCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationTrustCall) Return(arg0 error) *MockApplicationTrustCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationTrustCall) Do(f func(bool) error) *MockApplicationTrustCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationTrustCall) DoAndReturn(f func(bool) error) *MockApplicationTrustCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Units mocks base method.
func (m *MockApplication) Units() ([]caas.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Units")
	ret0, _ := ret[0].([]caas.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units.
func (mr *MockApplicationMockRecorder) Units() *MockApplicationUnitsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockApplication)(nil).Units))
	return &MockApplicationUnitsCall{Call: call}
}

// MockApplicationUnitsCall wrap *gomock.Call
type MockApplicationUnitsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationUnitsCall) Return(arg0 []caas.Unit, arg1 error) *MockApplicationUnitsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationUnitsCall) Do(f func() ([]caas.Unit, error)) *MockApplicationUnitsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationUnitsCall) DoAndReturn(f func() ([]caas.Unit, error)) *MockApplicationUnitsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UnitsToRemove mocks base method.
func (m *MockApplication) UnitsToRemove(arg0 context.Context, arg1 int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitsToRemove", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitsToRemove indicates an expected call of UnitsToRemove.
func (mr *MockApplicationMockRecorder) UnitsToRemove(arg0, arg1 any) *MockApplicationUnitsToRemoveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitsToRemove", reflect.TypeOf((*MockApplication)(nil).UnitsToRemove), arg0, arg1)
	return &MockApplicationUnitsToRemoveCall{Call: call}
}

// MockApplicationUnitsToRemoveCall wrap *gomock.Call
type MockApplicationUnitsToRemoveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationUnitsToRemoveCall) Return(arg0 []string, arg1 error) *MockApplicationUnitsToRemoveCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationUnitsToRemoveCall) Do(f func(context.Context, int) ([]string, error)) *MockApplicationUnitsToRemoveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationUnitsToRemoveCall) DoAndReturn(f func(context.Context, int) ([]string, error)) *MockApplicationUnitsToRemoveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdatePorts mocks base method.
func (m *MockApplication) UpdatePorts(arg0 []caas.ServicePort, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePorts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePorts indicates an expected call of UpdatePorts.
func (mr *MockApplicationMockRecorder) UpdatePorts(arg0, arg1 any) *MockApplicationUpdatePortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePorts", reflect.TypeOf((*MockApplication)(nil).UpdatePorts), arg0, arg1)
	return &MockApplicationUpdatePortsCall{Call: call}
}

// MockApplicationUpdatePortsCall wrap *gomock.Call
type MockApplicationUpdatePortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationUpdatePortsCall) Return(arg0 error) *MockApplicationUpdatePortsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationUpdatePortsCall) Do(f func([]caas.ServicePort, bool) error) *MockApplicationUpdatePortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationUpdatePortsCall) DoAndReturn(f func([]caas.ServicePort, bool) error) *MockApplicationUpdatePortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateService mocks base method.
func (m *MockApplication) UpdateService(arg0 caas.ServiceParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockApplicationMockRecorder) UpdateService(arg0 any) *MockApplicationUpdateServiceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockApplication)(nil).UpdateService), arg0)
	return &MockApplicationUpdateServiceCall{Call: call}
}

// MockApplicationUpdateServiceCall wrap *gomock.Call
type MockApplicationUpdateServiceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationUpdateServiceCall) Return(arg0 error) *MockApplicationUpdateServiceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationUpdateServiceCall) Do(f func(caas.ServiceParam) error) *MockApplicationUpdateServiceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationUpdateServiceCall) DoAndReturn(f func(caas.ServiceParam) error) *MockApplicationUpdateServiceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Watch mocks base method.
func (m *MockApplication) Watch(arg0 context.Context) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockApplicationMockRecorder) Watch(arg0 any) *MockApplicationWatchCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockApplication)(nil).Watch), arg0)
	return &MockApplicationWatchCall{Call: call}
}

// MockApplicationWatchCall wrap *gomock.Call
type MockApplicationWatchCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationWatchCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockApplicationWatchCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationWatchCall) Do(f func(context.Context) (watcher.Watcher[struct{}], error)) *MockApplicationWatchCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationWatchCall) DoAndReturn(f func(context.Context) (watcher.Watcher[struct{}], error)) *MockApplicationWatchCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchReplicas mocks base method.
func (m *MockApplication) WatchReplicas() (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchReplicas")
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchReplicas indicates an expected call of WatchReplicas.
func (mr *MockApplicationMockRecorder) WatchReplicas() *MockApplicationWatchReplicasCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchReplicas", reflect.TypeOf((*MockApplication)(nil).WatchReplicas))
	return &MockApplicationWatchReplicasCall{Call: call}
}

// MockApplicationWatchReplicasCall wrap *gomock.Call
type MockApplicationWatchReplicasCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationWatchReplicasCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockApplicationWatchReplicasCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationWatchReplicasCall) Do(f func() (watcher.Watcher[struct{}], error)) *MockApplicationWatchReplicasCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationWatchReplicasCall) DoAndReturn(f func() (watcher.Watcher[struct{}], error)) *MockApplicationWatchReplicasCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
