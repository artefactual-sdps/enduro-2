// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-sdps/enduro/internal/package_ (interfaces: Service)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"
	time "time"

	package_ "github.com/artefactual-sdps/enduro/internal/api/gen/package_"
	package_0 "github.com/artefactual-sdps/enduro/internal/package_"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CompletePreservationAction mocks base method.
func (m *MockService) CompletePreservationAction(arg0 context.Context, arg1 uint, arg2 package_0.PreservationActionStatus, arg3 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompletePreservationAction", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompletePreservationAction indicates an expected call of CompletePreservationAction.
func (mr *MockServiceMockRecorder) CompletePreservationAction(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompletePreservationAction", reflect.TypeOf((*MockService)(nil).CompletePreservationAction), arg0, arg1, arg2, arg3)
}

// CompletePreservationTask mocks base method.
func (m *MockService) CompletePreservationTask(arg0 context.Context, arg1 uint, arg2 *string, arg3 package_0.PreservationTaskStatus, arg4 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompletePreservationTask", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompletePreservationTask indicates an expected call of CompletePreservationTask.
func (mr *MockServiceMockRecorder) CompletePreservationTask(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompletePreservationTask", reflect.TypeOf((*MockService)(nil).CompletePreservationTask), arg0, arg1, arg2, arg3, arg4)
}

// Create mocks base method.
func (m *MockService) Create(arg0 context.Context, arg1 *package_0.Package) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), arg0, arg1)
}

// CreatePreservationAction mocks base method.
func (m *MockService) CreatePreservationAction(arg0 context.Context, arg1 *package_0.PreservationAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePreservationAction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePreservationAction indicates an expected call of CreatePreservationAction.
func (mr *MockServiceMockRecorder) CreatePreservationAction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePreservationAction", reflect.TypeOf((*MockService)(nil).CreatePreservationAction), arg0, arg1)
}

// CreatePreservationTask mocks base method.
func (m *MockService) CreatePreservationTask(arg0 context.Context, arg1 *package_0.PreservationTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePreservationTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePreservationTask indicates an expected call of CreatePreservationTask.
func (mr *MockServiceMockRecorder) CreatePreservationTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePreservationTask", reflect.TypeOf((*MockService)(nil).CreatePreservationTask), arg0, arg1)
}

// Goa mocks base method.
func (m *MockService) Goa() package_.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Goa")
	ret0, _ := ret[0].(package_.Service)
	return ret0
}

// Goa indicates an expected call of Goa.
func (mr *MockServiceMockRecorder) Goa() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Goa", reflect.TypeOf((*MockService)(nil).Goa))
}

// SetLocation mocks base method.
func (m *MockService) SetLocation(arg0 context.Context, arg1 uint, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLocation", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLocation indicates an expected call of SetLocation.
func (mr *MockServiceMockRecorder) SetLocation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLocation", reflect.TypeOf((*MockService)(nil).SetLocation), arg0, arg1, arg2)
}

// SetStatus mocks base method.
func (m *MockService) SetStatus(arg0 context.Context, arg1 uint, arg2 package_0.Status) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockServiceMockRecorder) SetStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockService)(nil).SetStatus), arg0, arg1, arg2)
}

// SetStatusInProgress mocks base method.
func (m *MockService) SetStatusInProgress(arg0 context.Context, arg1 uint, arg2 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatusInProgress", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatusInProgress indicates an expected call of SetStatusInProgress.
func (mr *MockServiceMockRecorder) SetStatusInProgress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatusInProgress", reflect.TypeOf((*MockService)(nil).SetStatusInProgress), arg0, arg1, arg2)
}

// SetStatusPending mocks base method.
func (m *MockService) SetStatusPending(arg0 context.Context, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatusPending", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatusPending indicates an expected call of SetStatusPending.
func (mr *MockServiceMockRecorder) SetStatusPending(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatusPending", reflect.TypeOf((*MockService)(nil).SetStatusPending), arg0, arg1)
}

// UpdateWorkflowStatus mocks base method.
func (m *MockService) UpdateWorkflowStatus(arg0 context.Context, arg1 uint, arg2, arg3, arg4, arg5 string, arg6 package_0.Status, arg7 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkflowStatus", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWorkflowStatus indicates an expected call of UpdateWorkflowStatus.
func (mr *MockServiceMockRecorder) UpdateWorkflowStatus(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowStatus", reflect.TypeOf((*MockService)(nil).UpdateWorkflowStatus), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}
