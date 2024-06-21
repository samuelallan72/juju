// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/provider/oci (interfaces: FirewallClient)
//
// Generated by this command:
//
//	mockgen -typed -package testing -destination testing/mocks_firewall.go -write_package_comment=false github.com/juju/juju/internal/provider/oci FirewallClient
package testing

import (
	context "context"
	reflect "reflect"

	core "github.com/oracle/oci-go-sdk/v65/core"
	gomock "go.uber.org/mock/gomock"
)

// MockFirewallClient is a mock of FirewallClient interface.
type MockFirewallClient struct {
	ctrl     *gomock.Controller
	recorder *MockFirewallClientMockRecorder
}

// MockFirewallClientMockRecorder is the mock recorder for MockFirewallClient.
type MockFirewallClientMockRecorder struct {
	mock *MockFirewallClient
}

// NewMockFirewallClient creates a new mock instance.
func NewMockFirewallClient(ctrl *gomock.Controller) *MockFirewallClient {
	mock := &MockFirewallClient{ctrl: ctrl}
	mock.recorder = &MockFirewallClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirewallClient) EXPECT() *MockFirewallClientMockRecorder {
	return m.recorder
}

// CreateSecurityList mocks base method.
func (m *MockFirewallClient) CreateSecurityList(arg0 context.Context, arg1 core.CreateSecurityListRequest) (core.CreateSecurityListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecurityList", arg0, arg1)
	ret0, _ := ret[0].(core.CreateSecurityListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecurityList indicates an expected call of CreateSecurityList.
func (mr *MockFirewallClientMockRecorder) CreateSecurityList(arg0, arg1 any) *MockFirewallClientCreateSecurityListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecurityList", reflect.TypeOf((*MockFirewallClient)(nil).CreateSecurityList), arg0, arg1)
	return &MockFirewallClientCreateSecurityListCall{Call: call}
}

// MockFirewallClientCreateSecurityListCall wrap *gomock.Call
type MockFirewallClientCreateSecurityListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFirewallClientCreateSecurityListCall) Return(arg0 core.CreateSecurityListResponse, arg1 error) *MockFirewallClientCreateSecurityListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFirewallClientCreateSecurityListCall) Do(f func(context.Context, core.CreateSecurityListRequest) (core.CreateSecurityListResponse, error)) *MockFirewallClientCreateSecurityListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFirewallClientCreateSecurityListCall) DoAndReturn(f func(context.Context, core.CreateSecurityListRequest) (core.CreateSecurityListResponse, error)) *MockFirewallClientCreateSecurityListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteSecurityList mocks base method.
func (m *MockFirewallClient) DeleteSecurityList(arg0 context.Context, arg1 core.DeleteSecurityListRequest) (core.DeleteSecurityListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecurityList", arg0, arg1)
	ret0, _ := ret[0].(core.DeleteSecurityListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSecurityList indicates an expected call of DeleteSecurityList.
func (mr *MockFirewallClientMockRecorder) DeleteSecurityList(arg0, arg1 any) *MockFirewallClientDeleteSecurityListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecurityList", reflect.TypeOf((*MockFirewallClient)(nil).DeleteSecurityList), arg0, arg1)
	return &MockFirewallClientDeleteSecurityListCall{Call: call}
}

// MockFirewallClientDeleteSecurityListCall wrap *gomock.Call
type MockFirewallClientDeleteSecurityListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFirewallClientDeleteSecurityListCall) Return(arg0 core.DeleteSecurityListResponse, arg1 error) *MockFirewallClientDeleteSecurityListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFirewallClientDeleteSecurityListCall) Do(f func(context.Context, core.DeleteSecurityListRequest) (core.DeleteSecurityListResponse, error)) *MockFirewallClientDeleteSecurityListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFirewallClientDeleteSecurityListCall) DoAndReturn(f func(context.Context, core.DeleteSecurityListRequest) (core.DeleteSecurityListResponse, error)) *MockFirewallClientDeleteSecurityListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSecurityList mocks base method.
func (m *MockFirewallClient) GetSecurityList(arg0 context.Context, arg1 core.GetSecurityListRequest) (core.GetSecurityListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecurityList", arg0, arg1)
	ret0, _ := ret[0].(core.GetSecurityListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecurityList indicates an expected call of GetSecurityList.
func (mr *MockFirewallClientMockRecorder) GetSecurityList(arg0, arg1 any) *MockFirewallClientGetSecurityListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecurityList", reflect.TypeOf((*MockFirewallClient)(nil).GetSecurityList), arg0, arg1)
	return &MockFirewallClientGetSecurityListCall{Call: call}
}

// MockFirewallClientGetSecurityListCall wrap *gomock.Call
type MockFirewallClientGetSecurityListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFirewallClientGetSecurityListCall) Return(arg0 core.GetSecurityListResponse, arg1 error) *MockFirewallClientGetSecurityListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFirewallClientGetSecurityListCall) Do(f func(context.Context, core.GetSecurityListRequest) (core.GetSecurityListResponse, error)) *MockFirewallClientGetSecurityListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFirewallClientGetSecurityListCall) DoAndReturn(f func(context.Context, core.GetSecurityListRequest) (core.GetSecurityListResponse, error)) *MockFirewallClientGetSecurityListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListSecurityLists mocks base method.
func (m *MockFirewallClient) ListSecurityLists(arg0 context.Context, arg1, arg2 *string) ([]core.SecurityList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecurityLists", arg0, arg1, arg2)
	ret0, _ := ret[0].([]core.SecurityList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecurityLists indicates an expected call of ListSecurityLists.
func (mr *MockFirewallClientMockRecorder) ListSecurityLists(arg0, arg1, arg2 any) *MockFirewallClientListSecurityListsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecurityLists", reflect.TypeOf((*MockFirewallClient)(nil).ListSecurityLists), arg0, arg1, arg2)
	return &MockFirewallClientListSecurityListsCall{Call: call}
}

// MockFirewallClientListSecurityListsCall wrap *gomock.Call
type MockFirewallClientListSecurityListsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFirewallClientListSecurityListsCall) Return(arg0 []core.SecurityList, arg1 error) *MockFirewallClientListSecurityListsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFirewallClientListSecurityListsCall) Do(f func(context.Context, *string, *string) ([]core.SecurityList, error)) *MockFirewallClientListSecurityListsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFirewallClientListSecurityListsCall) DoAndReturn(f func(context.Context, *string, *string) ([]core.SecurityList, error)) *MockFirewallClientListSecurityListsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}