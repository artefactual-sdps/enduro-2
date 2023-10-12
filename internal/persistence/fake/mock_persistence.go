// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-sdps/enduro/internal/persistence (interfaces: Service)
//
// Generated by this command:
//
//	mockgen -typed -destination=./internal/persistence/fake/mock_persistence.go -package=fake github.com/artefactual-sdps/enduro/internal/persistence Service
//
// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	package_ "github.com/artefactual-sdps/enduro/internal/package_"
	persistence "github.com/artefactual-sdps/enduro/internal/persistence"
	gomock "go.uber.org/mock/gomock"
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

// CreatePackage mocks base method.
func (m *MockService) CreatePackage(arg0 context.Context, arg1 *package_.Package) (*package_.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePackage", arg0, arg1)
	ret0, _ := ret[0].(*package_.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePackage indicates an expected call of CreatePackage.
func (mr *MockServiceMockRecorder) CreatePackage(arg0, arg1 any) *ServiceCreatePackageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePackage", reflect.TypeOf((*MockService)(nil).CreatePackage), arg0, arg1)
	return &ServiceCreatePackageCall{Call: call}
}

// ServiceCreatePackageCall wrap *gomock.Call
type ServiceCreatePackageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ServiceCreatePackageCall) Return(arg0 *package_.Package, arg1 error) *ServiceCreatePackageCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ServiceCreatePackageCall) Do(f func(context.Context, *package_.Package) (*package_.Package, error)) *ServiceCreatePackageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ServiceCreatePackageCall) DoAndReturn(f func(context.Context, *package_.Package) (*package_.Package, error)) *ServiceCreatePackageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdatePackage mocks base method.
func (m *MockService) UpdatePackage(arg0 context.Context, arg1 uint, arg2 persistence.PackageUpdater) (*package_.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackage", arg0, arg1, arg2)
	ret0, _ := ret[0].(*package_.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePackage indicates an expected call of UpdatePackage.
func (mr *MockServiceMockRecorder) UpdatePackage(arg0, arg1, arg2 any) *ServiceUpdatePackageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackage", reflect.TypeOf((*MockService)(nil).UpdatePackage), arg0, arg1, arg2)
	return &ServiceUpdatePackageCall{Call: call}
}

// ServiceUpdatePackageCall wrap *gomock.Call
type ServiceUpdatePackageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ServiceUpdatePackageCall) Return(arg0 *package_.Package, arg1 error) *ServiceUpdatePackageCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ServiceUpdatePackageCall) Do(f func(context.Context, uint, persistence.PackageUpdater) (*package_.Package, error)) *ServiceUpdatePackageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ServiceUpdatePackageCall) DoAndReturn(f func(context.Context, uint, persistence.PackageUpdater) (*package_.Package, error)) *ServiceUpdatePackageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
