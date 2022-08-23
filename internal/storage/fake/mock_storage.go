// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-sdps/enduro/internal/storage (interfaces: Service)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	storage "github.com/artefactual-sdps/enduro/internal/api/gen/storage"
	storage0 "github.com/artefactual-sdps/enduro/internal/storage"
	types "github.com/artefactual-sdps/enduro/internal/storage/types"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	blob "gocloud.dev/blob"
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

// AddLocation mocks base method.
func (m *MockService) AddLocation(arg0 context.Context, arg1 *storage.AddLocationPayload) (*storage.AddLocationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLocation", arg0, arg1)
	ret0, _ := ret[0].(*storage.AddLocationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLocation indicates an expected call of AddLocation.
func (mr *MockServiceMockRecorder) AddLocation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLocation", reflect.TypeOf((*MockService)(nil).AddLocation), arg0, arg1)
}

// Delete mocks base method.
func (m *MockService) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), arg0, arg1)
}

// Download mocks base method.
func (m *MockService) Download(arg0 context.Context, arg1 *storage.DownloadPayload) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockServiceMockRecorder) Download(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockService)(nil).Download), arg0, arg1)
}

// Location mocks base method.
func (m *MockService) Location(arg0 context.Context, arg1 uuid.UUID) (storage0.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location", arg0, arg1)
	ret0, _ := ret[0].(storage0.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Location indicates an expected call of Location.
func (mr *MockServiceMockRecorder) Location(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockService)(nil).Location), arg0, arg1)
}

// Locations mocks base method.
func (m *MockService) Locations(arg0 context.Context) (storage.StoredLocationCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Locations", arg0)
	ret0, _ := ret[0].(storage.StoredLocationCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Locations indicates an expected call of Locations.
func (mr *MockServiceMockRecorder) Locations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Locations", reflect.TypeOf((*MockService)(nil).Locations), arg0)
}

// Move mocks base method.
func (m *MockService) Move(arg0 context.Context, arg1 *storage.MovePayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Move indicates an expected call of Move.
func (mr *MockServiceMockRecorder) Move(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockService)(nil).Move), arg0, arg1)
}

// MoveStatus mocks base method.
func (m *MockService) MoveStatus(arg0 context.Context, arg1 *storage.MoveStatusPayload) (*storage.MoveStatusResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveStatus", arg0, arg1)
	ret0, _ := ret[0].(*storage.MoveStatusResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MoveStatus indicates an expected call of MoveStatus.
func (mr *MockServiceMockRecorder) MoveStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveStatus", reflect.TypeOf((*MockService)(nil).MoveStatus), arg0, arg1)
}

// PackageReader mocks base method.
func (m *MockService) PackageReader(arg0 context.Context, arg1 *storage.StoredStoragePackage) (*blob.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackageReader", arg0, arg1)
	ret0, _ := ret[0].(*blob.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackageReader indicates an expected call of PackageReader.
func (mr *MockServiceMockRecorder) PackageReader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackageReader", reflect.TypeOf((*MockService)(nil).PackageReader), arg0, arg1)
}

// ReadPackage mocks base method.
func (m *MockService) ReadPackage(arg0 context.Context, arg1 string) (*storage.StoredStoragePackage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPackage", arg0, arg1)
	ret0, _ := ret[0].(*storage.StoredStoragePackage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPackage indicates an expected call of ReadPackage.
func (mr *MockServiceMockRecorder) ReadPackage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPackage", reflect.TypeOf((*MockService)(nil).ReadPackage), arg0, arg1)
}

// Reject mocks base method.
func (m *MockService) Reject(arg0 context.Context, arg1 *storage.RejectPayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reject indicates an expected call of Reject.
func (mr *MockServiceMockRecorder) Reject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reject", reflect.TypeOf((*MockService)(nil).Reject), arg0, arg1)
}

// Show mocks base method.
func (m *MockService) Show(arg0 context.Context, arg1 *storage.ShowPayload) (*storage.StoredStoragePackage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", arg0, arg1)
	ret0, _ := ret[0].(*storage.StoredStoragePackage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockServiceMockRecorder) Show(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockService)(nil).Show), arg0, arg1)
}

// ShowLocation mocks base method.
func (m *MockService) ShowLocation(arg0 context.Context, arg1 *storage.ShowLocationPayload) (*storage.StoredLocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowLocation", arg0, arg1)
	ret0, _ := ret[0].(*storage.StoredLocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowLocation indicates an expected call of ShowLocation.
func (mr *MockServiceMockRecorder) ShowLocation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowLocation", reflect.TypeOf((*MockService)(nil).ShowLocation), arg0, arg1)
}

// Submit mocks base method.
func (m *MockService) Submit(arg0 context.Context, arg1 *storage.SubmitPayload) (*storage.SubmitResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Submit", arg0, arg1)
	ret0, _ := ret[0].(*storage.SubmitResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Submit indicates an expected call of Submit.
func (mr *MockServiceMockRecorder) Submit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Submit", reflect.TypeOf((*MockService)(nil).Submit), arg0, arg1)
}

// Update mocks base method.
func (m *MockService) Update(arg0 context.Context, arg1 *storage.UpdatePayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), arg0, arg1)
}

// UpdatePackageLocationID mocks base method.
func (m *MockService) UpdatePackageLocationID(arg0 context.Context, arg1 uuid.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackageLocationID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePackageLocationID indicates an expected call of UpdatePackageLocationID.
func (mr *MockServiceMockRecorder) UpdatePackageLocationID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackageLocationID", reflect.TypeOf((*MockService)(nil).UpdatePackageLocationID), arg0, arg1, arg2)
}

// UpdatePackageStatus mocks base method.
func (m *MockService) UpdatePackageStatus(arg0 context.Context, arg1 types.PackageStatus, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackageStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePackageStatus indicates an expected call of UpdatePackageStatus.
func (mr *MockServiceMockRecorder) UpdatePackageStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackageStatus", reflect.TypeOf((*MockService)(nil).UpdatePackageStatus), arg0, arg1, arg2)
}
