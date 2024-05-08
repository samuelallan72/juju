// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/charm/v13 (interfaces: Charm)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/charm_mock.go github.com/juju/charm/v13 Charm
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	charm "github.com/juju/charm/v13"
	gomock "go.uber.org/mock/gomock"
)

// MockCharm is a mock of Charm interface.
type MockCharm struct {
	ctrl     *gomock.Controller
	recorder *MockCharmMockRecorder
}

// MockCharmMockRecorder is the mock recorder for MockCharm.
type MockCharmMockRecorder struct {
	mock *MockCharm
}

// NewMockCharm creates a new mock instance.
func NewMockCharm(ctrl *gomock.Controller) *MockCharm {
	mock := &MockCharm{ctrl: ctrl}
	mock.recorder = &MockCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharm) EXPECT() *MockCharmMockRecorder {
	return m.recorder
}

// Actions mocks base method.
func (m *MockCharm) Actions() *charm.Actions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions")
	ret0, _ := ret[0].(*charm.Actions)
	return ret0
}

// Actions indicates an expected call of Actions.
func (mr *MockCharmMockRecorder) Actions() *MockCharmActionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockCharm)(nil).Actions))
	return &MockCharmActionsCall{Call: call}
}

// MockCharmActionsCall wrap *gomock.Call
type MockCharmActionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmActionsCall) Return(arg0 *charm.Actions) *MockCharmActionsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmActionsCall) Do(f func() *charm.Actions) *MockCharmActionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmActionsCall) DoAndReturn(f func() *charm.Actions) *MockCharmActionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Config mocks base method.
func (m *MockCharm) Config() *charm.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*charm.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockCharmMockRecorder) Config() *MockCharmConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockCharm)(nil).Config))
	return &MockCharmConfigCall{Call: call}
}

// MockCharmConfigCall wrap *gomock.Call
type MockCharmConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmConfigCall) Return(arg0 *charm.Config) *MockCharmConfigCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmConfigCall) Do(f func() *charm.Config) *MockCharmConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmConfigCall) DoAndReturn(f func() *charm.Config) *MockCharmConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Manifest mocks base method.
func (m *MockCharm) Manifest() *charm.Manifest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manifest")
	ret0, _ := ret[0].(*charm.Manifest)
	return ret0
}

// Manifest indicates an expected call of Manifest.
func (mr *MockCharmMockRecorder) Manifest() *MockCharmManifestCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manifest", reflect.TypeOf((*MockCharm)(nil).Manifest))
	return &MockCharmManifestCall{Call: call}
}

// MockCharmManifestCall wrap *gomock.Call
type MockCharmManifestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmManifestCall) Return(arg0 *charm.Manifest) *MockCharmManifestCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmManifestCall) Do(f func() *charm.Manifest) *MockCharmManifestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmManifestCall) DoAndReturn(f func() *charm.Manifest) *MockCharmManifestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Meta mocks base method.
func (m *MockCharm) Meta() *charm.Meta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(*charm.Meta)
	return ret0
}

// Meta indicates an expected call of Meta.
func (mr *MockCharmMockRecorder) Meta() *MockCharmMetaCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockCharm)(nil).Meta))
	return &MockCharmMetaCall{Call: call}
}

// MockCharmMetaCall wrap *gomock.Call
type MockCharmMetaCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmMetaCall) Return(arg0 *charm.Meta) *MockCharmMetaCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmMetaCall) Do(f func() *charm.Meta) *MockCharmMetaCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmMetaCall) DoAndReturn(f func() *charm.Meta) *MockCharmMetaCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Revision mocks base method.
func (m *MockCharm) Revision() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revision")
	ret0, _ := ret[0].(int)
	return ret0
}

// Revision indicates an expected call of Revision.
func (mr *MockCharmMockRecorder) Revision() *MockCharmRevisionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revision", reflect.TypeOf((*MockCharm)(nil).Revision))
	return &MockCharmRevisionCall{Call: call}
}

// MockCharmRevisionCall wrap *gomock.Call
type MockCharmRevisionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmRevisionCall) Return(arg0 int) *MockCharmRevisionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmRevisionCall) Do(f func() int) *MockCharmRevisionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmRevisionCall) DoAndReturn(f func() int) *MockCharmRevisionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
