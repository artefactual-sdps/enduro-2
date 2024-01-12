// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-sdps/enduro/internal/storage/persistence (interfaces: Storage)
//
// Generated by this command:
//
//	mockgen -typed -destination=./internal/storage/persistence/fake/mock_persistence.go -package=fake github.com/artefactual-sdps/enduro/internal/storage/persistence Storage
//
// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	storage "github.com/artefactual-sdps/enduro/internal/api/gen/storage"
	types "github.com/artefactual-sdps/enduro/internal/storage/types"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// CreateLocation mocks base method.
func (m *MockStorage) CreateLocation(arg0 context.Context, arg1 *storage.Location, arg2 *types.LocationConfig) (*storage.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLocation", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLocation indicates an expected call of CreateLocation.
func (mr *MockStorageMockRecorder) CreateLocation(arg0, arg1, arg2 any) *StorageCreateLocationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLocation", reflect.TypeOf((*MockStorage)(nil).CreateLocation), arg0, arg1, arg2)
	return &StorageCreateLocationCall{Call: call}
}

// StorageCreateLocationCall wrap *gomock.Call
type StorageCreateLocationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StorageCreateLocationCall) Return(arg0 *storage.Location, arg1 error) *StorageCreateLocationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StorageCreateLocationCall) Do(f func(context.Context, *storage.Location, *types.LocationConfig) (*storage.Location, error)) *StorageCreateLocationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StorageCreateLocationCall) DoAndReturn(f func(context.Context, *storage.Location, *types.LocationConfig) (*storage.Location, error)) *StorageCreateLocationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreatePackage mocks base method.
func (m *MockStorage) CreatePackage(arg0 context.Context, arg1 *storage.Package) (*storage.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePackage", arg0, arg1)
	ret0, _ := ret[0].(*storage.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePackage indicates an expected call of CreatePackage.
func (mr *MockStorageMockRecorder) CreatePackage(arg0, arg1 any) *StorageCreatePackageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePackage", reflect.TypeOf((*MockStorage)(nil).CreatePackage), arg0, arg1)
	return &StorageCreatePackageCall{Call: call}
}

// StorageCreatePackageCall wrap *gomock.Call
type StorageCreatePackageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StorageCreatePackageCall) Return(arg0 *storage.Package, arg1 error) *StorageCreatePackageCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StorageCreatePackageCall) Do(f func(context.Context, *storage.Package) (*storage.Package, error)) *StorageCreatePackageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StorageCreatePackageCall) DoAndReturn(f func(context.Context, *storage.Package) (*storage.Package, error)) *StorageCreatePackageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListLocations mocks base method.
func (m *MockStorage) ListLocations(arg0 context.Context) (storage.LocationCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocations", arg0)
	ret0, _ := ret[0].(storage.LocationCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLocations indicates an expected call of ListLocations.
func (mr *MockStorageMockRecorder) ListLocations(arg0 any) *StorageListLocationsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocations", reflect.TypeOf((*MockStorage)(nil).ListLocations), arg0)
	return &StorageListLocationsCall{Call: call}
}

// StorageListLocationsCall wrap *gomock.Call
type StorageListLocationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StorageListLocationsCall) Return(arg0 storage.LocationCollection, arg1 error) *StorageListLocationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StorageListLocationsCall) Do(f func(context.Context) (storage.LocationCollection, error)) *StorageListLocationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StorageListLocationsCall) DoAndReturn(f func(context.Context) (storage.LocationCollection, error)) *StorageListLocationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListPackages mocks base method.
func (m *MockStorage) ListPackages(arg0 context.Context) (storage.PackageCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPackages", arg0)
	ret0, _ := ret[0].(storage.PackageCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPackages indicates an expected call of ListPackages.
func (mr *MockStorageMockRecorder) ListPackages(arg0 any) *StorageListPackagesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPackages", reflect.TypeOf((*MockStorage)(nil).ListPackages), arg0)
	return &StorageListPackagesCall{Call: call}
}

// StorageListPackagesCall wrap *gomock.Call
type StorageListPackagesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StorageListPackagesCall) Return(arg0 storage.PackageCollection, arg1 error) *StorageListPackagesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StorageListPackagesCall) Do(f func(context.Context) (storage.PackageCollection, error)) *StorageListPackagesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StorageListPackagesCall) DoAndReturn(f func(context.Context) (storage.PackageCollection, error)) *StorageListPackagesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LocationPackages mocks base method.
func (m *MockStorage) LocationPackages(arg0 context.Context, arg1 uuid.UUID) (storage.PackageCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocationPackages", arg0, arg1)
	ret0, _ := ret[0].(storage.PackageCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocationPackages indicates an expected call of LocationPackages.
func (mr *MockStorageMockRecorder) LocationPackages(arg0, arg1 any) *StorageLocationPackagesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocationPackages", reflect.TypeOf((*MockStorage)(nil).LocationPackages), arg0, arg1)
	return &StorageLocationPackagesCall{Call: call}
}

// StorageLocationPackagesCall wrap *gomock.Call
type StorageLocationPackagesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StorageLocationPackagesCall) Return(arg0 storage.PackageCollection, arg1 error) *StorageLocationPackagesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StorageLocationPackagesCall) Do(f func(context.Context, uuid.UUID) (storage.PackageCollection, error)) *StorageLocationPackagesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StorageLocationPackagesCall) DoAndReturn(f func(context.Context, uuid.UUID) (storage.PackageCollection, error)) *StorageLocationPackagesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadLocation mocks base method.
func (m *MockStorage) ReadLocation(arg0 context.Context, arg1 uuid.UUID) (*storage.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadLocation", arg0, arg1)
	ret0, _ := ret[0].(*storage.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadLocation indicates an expected call of ReadLocation.
func (mr *MockStorageMockRecorder) ReadLocation(arg0, arg1 any) *StorageReadLocationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadLocation", reflect.TypeOf((*MockStorage)(nil).ReadLocation), arg0, arg1)
	return &StorageReadLocationCall{Call: call}
}

// StorageReadLocationCall wrap *gomock.Call
type StorageReadLocationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StorageReadLocationCall) Return(arg0 *storage.Location, arg1 error) *StorageReadLocationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StorageReadLocationCall) Do(f func(context.Context, uuid.UUID) (*storage.Location, error)) *StorageReadLocationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StorageReadLocationCall) DoAndReturn(f func(context.Context, uuid.UUID) (*storage.Location, error)) *StorageReadLocationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadPackage mocks base method.
func (m *MockStorage) ReadPackage(arg0 context.Context, arg1 uuid.UUID) (*storage.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPackage", arg0, arg1)
	ret0, _ := ret[0].(*storage.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPackage indicates an expected call of ReadPackage.
func (mr *MockStorageMockRecorder) ReadPackage(arg0, arg1 any) *StorageReadPackageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPackage", reflect.TypeOf((*MockStorage)(nil).ReadPackage), arg0, arg1)
	return &StorageReadPackageCall{Call: call}
}

// StorageReadPackageCall wrap *gomock.Call
type StorageReadPackageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StorageReadPackageCall) Return(arg0 *storage.Package, arg1 error) *StorageReadPackageCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StorageReadPackageCall) Do(f func(context.Context, uuid.UUID) (*storage.Package, error)) *StorageReadPackageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StorageReadPackageCall) DoAndReturn(f func(context.Context, uuid.UUID) (*storage.Package, error)) *StorageReadPackageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdatePackageLocationID mocks base method.
func (m *MockStorage) UpdatePackageLocationID(arg0 context.Context, arg1, arg2 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackageLocationID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePackageLocationID indicates an expected call of UpdatePackageLocationID.
func (mr *MockStorageMockRecorder) UpdatePackageLocationID(arg0, arg1, arg2 any) *StorageUpdatePackageLocationIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackageLocationID", reflect.TypeOf((*MockStorage)(nil).UpdatePackageLocationID), arg0, arg1, arg2)
	return &StorageUpdatePackageLocationIDCall{Call: call}
}

// StorageUpdatePackageLocationIDCall wrap *gomock.Call
type StorageUpdatePackageLocationIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StorageUpdatePackageLocationIDCall) Return(arg0 error) *StorageUpdatePackageLocationIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StorageUpdatePackageLocationIDCall) Do(f func(context.Context, uuid.UUID, uuid.UUID) error) *StorageUpdatePackageLocationIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StorageUpdatePackageLocationIDCall) DoAndReturn(f func(context.Context, uuid.UUID, uuid.UUID) error) *StorageUpdatePackageLocationIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdatePackageStatus mocks base method.
func (m *MockStorage) UpdatePackageStatus(arg0 context.Context, arg1 uuid.UUID, arg2 types.PackageStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackageStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePackageStatus indicates an expected call of UpdatePackageStatus.
func (mr *MockStorageMockRecorder) UpdatePackageStatus(arg0, arg1, arg2 any) *StorageUpdatePackageStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackageStatus", reflect.TypeOf((*MockStorage)(nil).UpdatePackageStatus), arg0, arg1, arg2)
	return &StorageUpdatePackageStatusCall{Call: call}
}

// StorageUpdatePackageStatusCall wrap *gomock.Call
type StorageUpdatePackageStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StorageUpdatePackageStatusCall) Return(arg0 error) *StorageUpdatePackageStatusCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StorageUpdatePackageStatusCall) Do(f func(context.Context, uuid.UUID, types.PackageStatus) error) *StorageUpdatePackageStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StorageUpdatePackageStatusCall) DoAndReturn(f func(context.Context, uuid.UUID, types.PackageStatus) error) *StorageUpdatePackageStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
