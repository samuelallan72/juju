// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/changestream (interfaces: Subscription,EventSource)
//
// Generated by this command:
//
//	mockgen -typed -package domain -destination changestream_mock_test.go github.com/juju/juju/core/changestream Subscription,EventSource
//

// Package domain is a generated GoMock package.
package domain

import (
	reflect "reflect"

	changestream "github.com/juju/juju/core/changestream"
	gomock "go.uber.org/mock/gomock"
)

// MockSubscription is a mock of Subscription interface.
type MockSubscription struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionMockRecorder
}

// MockSubscriptionMockRecorder is the mock recorder for MockSubscription.
type MockSubscriptionMockRecorder struct {
	mock *MockSubscription
}

// NewMockSubscription creates a new mock instance.
func NewMockSubscription(ctrl *gomock.Controller) *MockSubscription {
	mock := &MockSubscription{ctrl: ctrl}
	mock.recorder = &MockSubscriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscription) EXPECT() *MockSubscriptionMockRecorder {
	return m.recorder
}

// Changes mocks base method.
func (m *MockSubscription) Changes() <-chan []changestream.ChangeEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes")
	ret0, _ := ret[0].(<-chan []changestream.ChangeEvent)
	return ret0
}

// Changes indicates an expected call of Changes.
func (mr *MockSubscriptionMockRecorder) Changes() *MockSubscriptionChangesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockSubscription)(nil).Changes))
	return &MockSubscriptionChangesCall{Call: call}
}

// MockSubscriptionChangesCall wrap *gomock.Call
type MockSubscriptionChangesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSubscriptionChangesCall) Return(arg0 <-chan []changestream.ChangeEvent) *MockSubscriptionChangesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSubscriptionChangesCall) Do(f func() <-chan []changestream.ChangeEvent) *MockSubscriptionChangesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSubscriptionChangesCall) DoAndReturn(f func() <-chan []changestream.ChangeEvent) *MockSubscriptionChangesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Done mocks base method.
func (m *MockSubscription) Done() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockSubscriptionMockRecorder) Done() *MockSubscriptionDoneCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockSubscription)(nil).Done))
	return &MockSubscriptionDoneCall{Call: call}
}

// MockSubscriptionDoneCall wrap *gomock.Call
type MockSubscriptionDoneCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSubscriptionDoneCall) Return(arg0 <-chan struct{}) *MockSubscriptionDoneCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSubscriptionDoneCall) Do(f func() <-chan struct{}) *MockSubscriptionDoneCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSubscriptionDoneCall) DoAndReturn(f func() <-chan struct{}) *MockSubscriptionDoneCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Unsubscribe mocks base method.
func (m *MockSubscription) Unsubscribe() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unsubscribe")
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockSubscriptionMockRecorder) Unsubscribe() *MockSubscriptionUnsubscribeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockSubscription)(nil).Unsubscribe))
	return &MockSubscriptionUnsubscribeCall{Call: call}
}

// MockSubscriptionUnsubscribeCall wrap *gomock.Call
type MockSubscriptionUnsubscribeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSubscriptionUnsubscribeCall) Return() *MockSubscriptionUnsubscribeCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSubscriptionUnsubscribeCall) Do(f func()) *MockSubscriptionUnsubscribeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSubscriptionUnsubscribeCall) DoAndReturn(f func()) *MockSubscriptionUnsubscribeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockEventSource is a mock of EventSource interface.
type MockEventSource struct {
	ctrl     *gomock.Controller
	recorder *MockEventSourceMockRecorder
}

// MockEventSourceMockRecorder is the mock recorder for MockEventSource.
type MockEventSourceMockRecorder struct {
	mock *MockEventSource
}

// NewMockEventSource creates a new mock instance.
func NewMockEventSource(ctrl *gomock.Controller) *MockEventSource {
	mock := &MockEventSource{ctrl: ctrl}
	mock.recorder = &MockEventSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventSource) EXPECT() *MockEventSourceMockRecorder {
	return m.recorder
}

// Subscribe mocks base method.
func (m *MockEventSource) Subscribe(arg0 ...changestream.SubscriptionOption) (changestream.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(changestream.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockEventSourceMockRecorder) Subscribe(arg0 ...any) *MockEventSourceSubscribeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockEventSource)(nil).Subscribe), arg0...)
	return &MockEventSourceSubscribeCall{Call: call}
}

// MockEventSourceSubscribeCall wrap *gomock.Call
type MockEventSourceSubscribeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEventSourceSubscribeCall) Return(arg0 changestream.Subscription, arg1 error) *MockEventSourceSubscribeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEventSourceSubscribeCall) Do(f func(...changestream.SubscriptionOption) (changestream.Subscription, error)) *MockEventSourceSubscribeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEventSourceSubscribeCall) DoAndReturn(f func(...changestream.SubscriptionOption) (changestream.Subscription, error)) *MockEventSourceSubscribeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
