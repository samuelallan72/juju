// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: ModelMigration,NotifyWatcher)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/state.go github.com/juju/juju/state ModelMigration,NotifyWatcher
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	migration "github.com/juju/juju/core/migration"
	permission "github.com/juju/juju/core/permission"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockModelMigration is a mock of ModelMigration interface.
type MockModelMigration struct {
	ctrl     *gomock.Controller
	recorder *MockModelMigrationMockRecorder
}

// MockModelMigrationMockRecorder is the mock recorder for MockModelMigration.
type MockModelMigrationMockRecorder struct {
	mock *MockModelMigration
}

// NewMockModelMigration creates a new mock instance.
func NewMockModelMigration(ctrl *gomock.Controller) *MockModelMigration {
	mock := &MockModelMigration{ctrl: ctrl}
	mock.recorder = &MockModelMigrationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelMigration) EXPECT() *MockModelMigrationMockRecorder {
	return m.recorder
}

// Attempt mocks base method.
func (m *MockModelMigration) Attempt() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attempt")
	ret0, _ := ret[0].(int)
	return ret0
}

// Attempt indicates an expected call of Attempt.
func (mr *MockModelMigrationMockRecorder) Attempt() *MockModelMigrationAttemptCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attempt", reflect.TypeOf((*MockModelMigration)(nil).Attempt))
	return &MockModelMigrationAttemptCall{Call: call}
}

// MockModelMigrationAttemptCall wrap *gomock.Call
type MockModelMigrationAttemptCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationAttemptCall) Return(arg0 int) *MockModelMigrationAttemptCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationAttemptCall) Do(f func() int) *MockModelMigrationAttemptCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationAttemptCall) DoAndReturn(f func() int) *MockModelMigrationAttemptCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EndTime mocks base method.
func (m *MockModelMigration) EndTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// EndTime indicates an expected call of EndTime.
func (mr *MockModelMigrationMockRecorder) EndTime() *MockModelMigrationEndTimeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndTime", reflect.TypeOf((*MockModelMigration)(nil).EndTime))
	return &MockModelMigrationEndTimeCall{Call: call}
}

// MockModelMigrationEndTimeCall wrap *gomock.Call
type MockModelMigrationEndTimeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationEndTimeCall) Return(arg0 time.Time) *MockModelMigrationEndTimeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationEndTimeCall) Do(f func() time.Time) *MockModelMigrationEndTimeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationEndTimeCall) DoAndReturn(f func() time.Time) *MockModelMigrationEndTimeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Id mocks base method.
func (m *MockModelMigration) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockModelMigrationMockRecorder) Id() *MockModelMigrationIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockModelMigration)(nil).Id))
	return &MockModelMigrationIdCall{Call: call}
}

// MockModelMigrationIdCall wrap *gomock.Call
type MockModelMigrationIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationIdCall) Return(arg0 string) *MockModelMigrationIdCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationIdCall) Do(f func() string) *MockModelMigrationIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationIdCall) DoAndReturn(f func() string) *MockModelMigrationIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InitiatedBy mocks base method.
func (m *MockModelMigration) InitiatedBy() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitiatedBy")
	ret0, _ := ret[0].(string)
	return ret0
}

// InitiatedBy indicates an expected call of InitiatedBy.
func (mr *MockModelMigrationMockRecorder) InitiatedBy() *MockModelMigrationInitiatedByCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitiatedBy", reflect.TypeOf((*MockModelMigration)(nil).InitiatedBy))
	return &MockModelMigrationInitiatedByCall{Call: call}
}

// MockModelMigrationInitiatedByCall wrap *gomock.Call
type MockModelMigrationInitiatedByCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationInitiatedByCall) Return(arg0 string) *MockModelMigrationInitiatedByCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationInitiatedByCall) Do(f func() string) *MockModelMigrationInitiatedByCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationInitiatedByCall) DoAndReturn(f func() string) *MockModelMigrationInitiatedByCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MinionReports mocks base method.
func (m *MockModelMigration) MinionReports() (*state.MinionReports, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinionReports")
	ret0, _ := ret[0].(*state.MinionReports)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MinionReports indicates an expected call of MinionReports.
func (mr *MockModelMigrationMockRecorder) MinionReports() *MockModelMigrationMinionReportsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinionReports", reflect.TypeOf((*MockModelMigration)(nil).MinionReports))
	return &MockModelMigrationMinionReportsCall{Call: call}
}

// MockModelMigrationMinionReportsCall wrap *gomock.Call
type MockModelMigrationMinionReportsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationMinionReportsCall) Return(arg0 *state.MinionReports, arg1 error) *MockModelMigrationMinionReportsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationMinionReportsCall) Do(f func() (*state.MinionReports, error)) *MockModelMigrationMinionReportsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationMinionReportsCall) DoAndReturn(f func() (*state.MinionReports, error)) *MockModelMigrationMinionReportsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelUUID mocks base method.
func (m *MockModelMigration) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockModelMigrationMockRecorder) ModelUUID() *MockModelMigrationModelUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockModelMigration)(nil).ModelUUID))
	return &MockModelMigrationModelUUIDCall{Call: call}
}

// MockModelMigrationModelUUIDCall wrap *gomock.Call
type MockModelMigrationModelUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationModelUUIDCall) Return(arg0 string) *MockModelMigrationModelUUIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationModelUUIDCall) Do(f func() string) *MockModelMigrationModelUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationModelUUIDCall) DoAndReturn(f func() string) *MockModelMigrationModelUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelUserAccess mocks base method.
func (m *MockModelMigration) ModelUserAccess(arg0 names.Tag) permission.Access {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUserAccess", arg0)
	ret0, _ := ret[0].(permission.Access)
	return ret0
}

// ModelUserAccess indicates an expected call of ModelUserAccess.
func (mr *MockModelMigrationMockRecorder) ModelUserAccess(arg0 any) *MockModelMigrationModelUserAccessCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUserAccess", reflect.TypeOf((*MockModelMigration)(nil).ModelUserAccess), arg0)
	return &MockModelMigrationModelUserAccessCall{Call: call}
}

// MockModelMigrationModelUserAccessCall wrap *gomock.Call
type MockModelMigrationModelUserAccessCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationModelUserAccessCall) Return(arg0 permission.Access) *MockModelMigrationModelUserAccessCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationModelUserAccessCall) Do(f func(names.Tag) permission.Access) *MockModelMigrationModelUserAccessCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationModelUserAccessCall) DoAndReturn(f func(names.Tag) permission.Access) *MockModelMigrationModelUserAccessCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Phase mocks base method.
func (m *MockModelMigration) Phase() (migration.Phase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Phase")
	ret0, _ := ret[0].(migration.Phase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Phase indicates an expected call of Phase.
func (mr *MockModelMigrationMockRecorder) Phase() *MockModelMigrationPhaseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Phase", reflect.TypeOf((*MockModelMigration)(nil).Phase))
	return &MockModelMigrationPhaseCall{Call: call}
}

// MockModelMigrationPhaseCall wrap *gomock.Call
type MockModelMigrationPhaseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationPhaseCall) Return(arg0 migration.Phase, arg1 error) *MockModelMigrationPhaseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationPhaseCall) Do(f func() (migration.Phase, error)) *MockModelMigrationPhaseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationPhaseCall) DoAndReturn(f func() (migration.Phase, error)) *MockModelMigrationPhaseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PhaseChangedTime mocks base method.
func (m *MockModelMigration) PhaseChangedTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhaseChangedTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// PhaseChangedTime indicates an expected call of PhaseChangedTime.
func (mr *MockModelMigrationMockRecorder) PhaseChangedTime() *MockModelMigrationPhaseChangedTimeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhaseChangedTime", reflect.TypeOf((*MockModelMigration)(nil).PhaseChangedTime))
	return &MockModelMigrationPhaseChangedTimeCall{Call: call}
}

// MockModelMigrationPhaseChangedTimeCall wrap *gomock.Call
type MockModelMigrationPhaseChangedTimeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationPhaseChangedTimeCall) Return(arg0 time.Time) *MockModelMigrationPhaseChangedTimeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationPhaseChangedTimeCall) Do(f func() time.Time) *MockModelMigrationPhaseChangedTimeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationPhaseChangedTimeCall) DoAndReturn(f func() time.Time) *MockModelMigrationPhaseChangedTimeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Refresh mocks base method.
func (m *MockModelMigration) Refresh() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh")
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh.
func (mr *MockModelMigrationMockRecorder) Refresh() *MockModelMigrationRefreshCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockModelMigration)(nil).Refresh))
	return &MockModelMigrationRefreshCall{Call: call}
}

// MockModelMigrationRefreshCall wrap *gomock.Call
type MockModelMigrationRefreshCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationRefreshCall) Return(arg0 error) *MockModelMigrationRefreshCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationRefreshCall) Do(f func() error) *MockModelMigrationRefreshCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationRefreshCall) DoAndReturn(f func() error) *MockModelMigrationRefreshCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetPhase mocks base method.
func (m *MockModelMigration) SetPhase(arg0 migration.Phase) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPhase", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPhase indicates an expected call of SetPhase.
func (mr *MockModelMigrationMockRecorder) SetPhase(arg0 any) *MockModelMigrationSetPhaseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPhase", reflect.TypeOf((*MockModelMigration)(nil).SetPhase), arg0)
	return &MockModelMigrationSetPhaseCall{Call: call}
}

// MockModelMigrationSetPhaseCall wrap *gomock.Call
type MockModelMigrationSetPhaseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationSetPhaseCall) Return(arg0 error) *MockModelMigrationSetPhaseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationSetPhaseCall) Do(f func(migration.Phase) error) *MockModelMigrationSetPhaseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationSetPhaseCall) DoAndReturn(f func(migration.Phase) error) *MockModelMigrationSetPhaseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetStatusMessage mocks base method.
func (m *MockModelMigration) SetStatusMessage(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatusMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatusMessage indicates an expected call of SetStatusMessage.
func (mr *MockModelMigrationMockRecorder) SetStatusMessage(arg0 any) *MockModelMigrationSetStatusMessageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatusMessage", reflect.TypeOf((*MockModelMigration)(nil).SetStatusMessage), arg0)
	return &MockModelMigrationSetStatusMessageCall{Call: call}
}

// MockModelMigrationSetStatusMessageCall wrap *gomock.Call
type MockModelMigrationSetStatusMessageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationSetStatusMessageCall) Return(arg0 error) *MockModelMigrationSetStatusMessageCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationSetStatusMessageCall) Do(f func(string) error) *MockModelMigrationSetStatusMessageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationSetStatusMessageCall) DoAndReturn(f func(string) error) *MockModelMigrationSetStatusMessageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StartTime mocks base method.
func (m *MockModelMigration) StartTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// StartTime indicates an expected call of StartTime.
func (mr *MockModelMigrationMockRecorder) StartTime() *MockModelMigrationStartTimeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTime", reflect.TypeOf((*MockModelMigration)(nil).StartTime))
	return &MockModelMigrationStartTimeCall{Call: call}
}

// MockModelMigrationStartTimeCall wrap *gomock.Call
type MockModelMigrationStartTimeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationStartTimeCall) Return(arg0 time.Time) *MockModelMigrationStartTimeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationStartTimeCall) Do(f func() time.Time) *MockModelMigrationStartTimeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationStartTimeCall) DoAndReturn(f func() time.Time) *MockModelMigrationStartTimeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StatusMessage mocks base method.
func (m *MockModelMigration) StatusMessage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusMessage")
	ret0, _ := ret[0].(string)
	return ret0
}

// StatusMessage indicates an expected call of StatusMessage.
func (mr *MockModelMigrationMockRecorder) StatusMessage() *MockModelMigrationStatusMessageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusMessage", reflect.TypeOf((*MockModelMigration)(nil).StatusMessage))
	return &MockModelMigrationStatusMessageCall{Call: call}
}

// MockModelMigrationStatusMessageCall wrap *gomock.Call
type MockModelMigrationStatusMessageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationStatusMessageCall) Return(arg0 string) *MockModelMigrationStatusMessageCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationStatusMessageCall) Do(f func() string) *MockModelMigrationStatusMessageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationStatusMessageCall) DoAndReturn(f func() string) *MockModelMigrationStatusMessageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SubmitMinionReport mocks base method.
func (m *MockModelMigration) SubmitMinionReport(arg0 names.Tag, arg1 migration.Phase, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitMinionReport", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitMinionReport indicates an expected call of SubmitMinionReport.
func (mr *MockModelMigrationMockRecorder) SubmitMinionReport(arg0, arg1, arg2 any) *MockModelMigrationSubmitMinionReportCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitMinionReport", reflect.TypeOf((*MockModelMigration)(nil).SubmitMinionReport), arg0, arg1, arg2)
	return &MockModelMigrationSubmitMinionReportCall{Call: call}
}

// MockModelMigrationSubmitMinionReportCall wrap *gomock.Call
type MockModelMigrationSubmitMinionReportCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationSubmitMinionReportCall) Return(arg0 error) *MockModelMigrationSubmitMinionReportCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationSubmitMinionReportCall) Do(f func(names.Tag, migration.Phase, bool) error) *MockModelMigrationSubmitMinionReportCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationSubmitMinionReportCall) DoAndReturn(f func(names.Tag, migration.Phase, bool) error) *MockModelMigrationSubmitMinionReportCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SuccessTime mocks base method.
func (m *MockModelMigration) SuccessTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuccessTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// SuccessTime indicates an expected call of SuccessTime.
func (mr *MockModelMigrationMockRecorder) SuccessTime() *MockModelMigrationSuccessTimeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuccessTime", reflect.TypeOf((*MockModelMigration)(nil).SuccessTime))
	return &MockModelMigrationSuccessTimeCall{Call: call}
}

// MockModelMigrationSuccessTimeCall wrap *gomock.Call
type MockModelMigrationSuccessTimeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationSuccessTimeCall) Return(arg0 time.Time) *MockModelMigrationSuccessTimeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationSuccessTimeCall) Do(f func() time.Time) *MockModelMigrationSuccessTimeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationSuccessTimeCall) DoAndReturn(f func() time.Time) *MockModelMigrationSuccessTimeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// TargetInfo mocks base method.
func (m *MockModelMigration) TargetInfo() (*migration.TargetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TargetInfo")
	ret0, _ := ret[0].(*migration.TargetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TargetInfo indicates an expected call of TargetInfo.
func (mr *MockModelMigrationMockRecorder) TargetInfo() *MockModelMigrationTargetInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TargetInfo", reflect.TypeOf((*MockModelMigration)(nil).TargetInfo))
	return &MockModelMigrationTargetInfoCall{Call: call}
}

// MockModelMigrationTargetInfoCall wrap *gomock.Call
type MockModelMigrationTargetInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationTargetInfoCall) Return(arg0 *migration.TargetInfo, arg1 error) *MockModelMigrationTargetInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationTargetInfoCall) Do(f func() (*migration.TargetInfo, error)) *MockModelMigrationTargetInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationTargetInfoCall) DoAndReturn(f func() (*migration.TargetInfo, error)) *MockModelMigrationTargetInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchMinionReports mocks base method.
func (m *MockModelMigration) WatchMinionReports() (state.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchMinionReports")
	ret0, _ := ret[0].(state.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchMinionReports indicates an expected call of WatchMinionReports.
func (mr *MockModelMigrationMockRecorder) WatchMinionReports() *MockModelMigrationWatchMinionReportsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchMinionReports", reflect.TypeOf((*MockModelMigration)(nil).WatchMinionReports))
	return &MockModelMigrationWatchMinionReportsCall{Call: call}
}

// MockModelMigrationWatchMinionReportsCall wrap *gomock.Call
type MockModelMigrationWatchMinionReportsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelMigrationWatchMinionReportsCall) Return(arg0 state.NotifyWatcher, arg1 error) *MockModelMigrationWatchMinionReportsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelMigrationWatchMinionReportsCall) Do(f func() (state.NotifyWatcher, error)) *MockModelMigrationWatchMinionReportsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelMigrationWatchMinionReportsCall) DoAndReturn(f func() (state.NotifyWatcher, error)) *MockModelMigrationWatchMinionReportsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockNotifyWatcher is a mock of NotifyWatcher interface.
type MockNotifyWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockNotifyWatcherMockRecorder
}

// MockNotifyWatcherMockRecorder is the mock recorder for MockNotifyWatcher.
type MockNotifyWatcherMockRecorder struct {
	mock *MockNotifyWatcher
}

// NewMockNotifyWatcher creates a new mock instance.
func NewMockNotifyWatcher(ctrl *gomock.Controller) *MockNotifyWatcher {
	mock := &MockNotifyWatcher{ctrl: ctrl}
	mock.recorder = &MockNotifyWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotifyWatcher) EXPECT() *MockNotifyWatcherMockRecorder {
	return m.recorder
}

// Changes mocks base method.
func (m *MockNotifyWatcher) Changes() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Changes indicates an expected call of Changes.
func (mr *MockNotifyWatcherMockRecorder) Changes() *MockNotifyWatcherChangesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockNotifyWatcher)(nil).Changes))
	return &MockNotifyWatcherChangesCall{Call: call}
}

// MockNotifyWatcherChangesCall wrap *gomock.Call
type MockNotifyWatcherChangesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNotifyWatcherChangesCall) Return(arg0 <-chan struct{}) *MockNotifyWatcherChangesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNotifyWatcherChangesCall) Do(f func() <-chan struct{}) *MockNotifyWatcherChangesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNotifyWatcherChangesCall) DoAndReturn(f func() <-chan struct{}) *MockNotifyWatcherChangesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Err mocks base method.
func (m *MockNotifyWatcher) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockNotifyWatcherMockRecorder) Err() *MockNotifyWatcherErrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockNotifyWatcher)(nil).Err))
	return &MockNotifyWatcherErrCall{Call: call}
}

// MockNotifyWatcherErrCall wrap *gomock.Call
type MockNotifyWatcherErrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNotifyWatcherErrCall) Return(arg0 error) *MockNotifyWatcherErrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNotifyWatcherErrCall) Do(f func() error) *MockNotifyWatcherErrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNotifyWatcherErrCall) DoAndReturn(f func() error) *MockNotifyWatcherErrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Kill mocks base method.
func (m *MockNotifyWatcher) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockNotifyWatcherMockRecorder) Kill() *MockNotifyWatcherKillCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockNotifyWatcher)(nil).Kill))
	return &MockNotifyWatcherKillCall{Call: call}
}

// MockNotifyWatcherKillCall wrap *gomock.Call
type MockNotifyWatcherKillCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNotifyWatcherKillCall) Return() *MockNotifyWatcherKillCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNotifyWatcherKillCall) Do(f func()) *MockNotifyWatcherKillCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNotifyWatcherKillCall) DoAndReturn(f func()) *MockNotifyWatcherKillCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Stop mocks base method.
func (m *MockNotifyWatcher) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockNotifyWatcherMockRecorder) Stop() *MockNotifyWatcherStopCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockNotifyWatcher)(nil).Stop))
	return &MockNotifyWatcherStopCall{Call: call}
}

// MockNotifyWatcherStopCall wrap *gomock.Call
type MockNotifyWatcherStopCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNotifyWatcherStopCall) Return(arg0 error) *MockNotifyWatcherStopCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNotifyWatcherStopCall) Do(f func() error) *MockNotifyWatcherStopCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNotifyWatcherStopCall) DoAndReturn(f func() error) *MockNotifyWatcherStopCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Wait mocks base method.
func (m *MockNotifyWatcher) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockNotifyWatcherMockRecorder) Wait() *MockNotifyWatcherWaitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockNotifyWatcher)(nil).Wait))
	return &MockNotifyWatcherWaitCall{Call: call}
}

// MockNotifyWatcherWaitCall wrap *gomock.Call
type MockNotifyWatcherWaitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNotifyWatcherWaitCall) Return(arg0 error) *MockNotifyWatcherWaitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNotifyWatcherWaitCall) Do(f func() error) *MockNotifyWatcherWaitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNotifyWatcherWaitCall) DoAndReturn(f func() error) *MockNotifyWatcherWaitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
