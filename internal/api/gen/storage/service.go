// Code generated by goa v3.8.3, DO NOT EDIT.
//
// storage service
//
// Command:
// $ goa-v3.8.3 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package storage

import (
	"context"

	storageviews "github.com/artefactual-sdps/enduro/internal/api/gen/storage/views"
	goa "goa.design/goa/v3/pkg"
)

// The storage service manages the storage of packages.
type Service interface {
	// Start the submission of a package
	Submit(context.Context, *SubmitPayload) (res *SubmitResult, err error)
	// Signal the storage service that an upload is complete
	Update(context.Context, *UpdatePayload) (err error)
	// Download package by AIPID
	Download(context.Context, *DownloadPayload) (res []byte, err error)
	// List locations
	Locations(context.Context) (res StoredLocationCollection, err error)
	// Add a storage location
	AddLocation(context.Context, *AddLocationPayload) (res *AddLocationResult, err error)
	// Move a package to a permanent storage location
	Move(context.Context, *MovePayload) (err error)
	// Retrieve the status of a permanent storage location move of the package
	MoveStatus(context.Context, *MoveStatusPayload) (res *MoveStatusResult, err error)
	// Reject a package
	Reject(context.Context, *RejectPayload) (err error)
	// Show package by AIPID
	Show(context.Context, *ShowPayload) (res *StoredStoragePackage, err error)
	// Show location by UUID
	ShowLocation(context.Context, *ShowLocationPayload) (res *StoredLocation, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "storage"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [10]string{"submit", "update", "download", "locations", "add_location", "move", "move_status", "reject", "show", "show-location"}

// AddLocationPayload is the payload type of the storage service add_location
// method.
type AddLocationPayload struct {
	Name        string
	Description *string
	Source      string
	Purpose     string
	Config      interface {
		configVal()
	}
}

// AddLocationResult is the result type of the storage service add_location
// method.
type AddLocationResult struct {
	UUID string
}

// DownloadPayload is the payload type of the storage service download method.
type DownloadPayload struct {
	AipID string
}

// MovePayload is the payload type of the storage service move method.
type MovePayload struct {
	AipID    string
	Location string
}

// MoveStatusPayload is the payload type of the storage service move_status
// method.
type MoveStatusPayload struct {
	AipID string
}

// MoveStatusResult is the result type of the storage service move_status
// method.
type MoveStatusResult struct {
	Done bool
}

// RejectPayload is the payload type of the storage service reject method.
type RejectPayload struct {
	AipID string
}

type S3Config struct {
	Bucket    string
	Region    string
	Endpoint  *string
	PathStyle *bool
	Profile   *string
	Key       *string
	Secret    *string
	Token     *string
}

// ShowLocationPayload is the payload type of the storage service show-location
// method.
type ShowLocationPayload struct {
	UUID string
}

// ShowPayload is the payload type of the storage service show method.
type ShowPayload struct {
	AipID string
}

// Storage location not found.
type StorageLocationNotfound struct {
	// Message of error
	Message string
	// Identifier of missing location
	UUID string
}

// Storage package not found.
type StoragePackageNotfound struct {
	// Message of error
	Message string
	// Identifier of missing package
	AipID string
}

// StoredLocation is the result type of the storage service show-location
// method.
type StoredLocation struct {
	// ID is the unique id of the location.
	ID uint
	// Name of location
	Name string
	// Description of the location
	Description *string
	// Data source of the location
	Source string
	// Purpose of the location
	Purpose string
	UUID    *string
}

// StoredLocationCollection is the result type of the storage service locations
// method.
type StoredLocationCollection []*StoredLocation

// StoredStoragePackage is the result type of the storage service show method.
type StoredStoragePackage struct {
	ID    uint
	Name  string
	AipID string
	// Status of the package
	Status    string
	ObjectKey string
	Location  *string
}

// SubmitPayload is the payload type of the storage service submit method.
type SubmitPayload struct {
	AipID string
	Name  string
}

// SubmitResult is the result type of the storage service submit method.
type SubmitResult struct {
	URL string
}

// UpdatePayload is the payload type of the storage service update method.
type UpdatePayload struct {
	AipID string
}

// Error returns an error description.
func (e *StorageLocationNotfound) Error() string {
	return "Storage location not found."
}

// ErrorName returns "StorageLocationNotfound".
func (e *StorageLocationNotfound) ErrorName() string {
	return e.Message
}

// Error returns an error description.
func (e *StoragePackageNotfound) Error() string {
	return "Storage package not found."
}

// ErrorName returns "StoragePackageNotfound".
func (e *StoragePackageNotfound) ErrorName() string {
	return e.Message
}
func (*S3Config) configVal() {}

// MakeNotAvailable builds a goa.ServiceError from an error.
func MakeNotAvailable(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not_available", false, false, false)
}

// MakeNotValid builds a goa.ServiceError from an error.
func MakeNotValid(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not_valid", false, false, false)
}

// MakeFailedDependency builds a goa.ServiceError from an error.
func MakeFailedDependency(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "failed_dependency", false, false, false)
}

// NewStoredLocationCollection initializes result type StoredLocationCollection
// from viewed result type StoredLocationCollection.
func NewStoredLocationCollection(vres storageviews.StoredLocationCollection) StoredLocationCollection {
	return newStoredLocationCollection(vres.Projected)
}

// NewViewedStoredLocationCollection initializes viewed result type
// StoredLocationCollection from result type StoredLocationCollection using the
// given view.
func NewViewedStoredLocationCollection(res StoredLocationCollection, view string) storageviews.StoredLocationCollection {
	p := newStoredLocationCollectionView(res)
	return storageviews.StoredLocationCollection{Projected: p, View: "default"}
}

// NewStoredStoragePackage initializes result type StoredStoragePackage from
// viewed result type StoredStoragePackage.
func NewStoredStoragePackage(vres *storageviews.StoredStoragePackage) *StoredStoragePackage {
	return newStoredStoragePackage(vres.Projected)
}

// NewViewedStoredStoragePackage initializes viewed result type
// StoredStoragePackage from result type StoredStoragePackage using the given
// view.
func NewViewedStoredStoragePackage(res *StoredStoragePackage, view string) *storageviews.StoredStoragePackage {
	p := newStoredStoragePackageView(res)
	return &storageviews.StoredStoragePackage{Projected: p, View: "default"}
}

// NewStoredLocation initializes result type StoredLocation from viewed result
// type StoredLocation.
func NewStoredLocation(vres *storageviews.StoredLocation) *StoredLocation {
	return newStoredLocation(vres.Projected)
}

// NewViewedStoredLocation initializes viewed result type StoredLocation from
// result type StoredLocation using the given view.
func NewViewedStoredLocation(res *StoredLocation, view string) *storageviews.StoredLocation {
	p := newStoredLocationView(res)
	return &storageviews.StoredLocation{Projected: p, View: "default"}
}

// newStoredLocationCollection converts projected type StoredLocationCollection
// to service type StoredLocationCollection.
func newStoredLocationCollection(vres storageviews.StoredLocationCollectionView) StoredLocationCollection {
	res := make(StoredLocationCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredLocation(n)
	}
	return res
}

// newStoredLocationCollectionView projects result type
// StoredLocationCollection to projected type StoredLocationCollectionView
// using the "default" view.
func newStoredLocationCollectionView(res StoredLocationCollection) storageviews.StoredLocationCollectionView {
	vres := make(storageviews.StoredLocationCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredLocationView(n)
	}
	return vres
}

// newStoredLocation converts projected type StoredLocation to service type
// StoredLocation.
func newStoredLocation(vres *storageviews.StoredLocationView) *StoredLocation {
	res := &StoredLocation{
		Description: vres.Description,
		UUID:        vres.UUID,
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Source != nil {
		res.Source = *vres.Source
	}
	if vres.Purpose != nil {
		res.Purpose = *vres.Purpose
	}
	if vres.Source == nil {
		res.Source = "unspecified"
	}
	if vres.Purpose == nil {
		res.Purpose = "unspecified"
	}
	return res
}

// newStoredLocationView projects result type StoredLocation to projected type
// StoredLocationView using the "default" view.
func newStoredLocationView(res *StoredLocation) *storageviews.StoredLocationView {
	vres := &storageviews.StoredLocationView{
		Name:        &res.Name,
		Description: res.Description,
		Source:      &res.Source,
		Purpose:     &res.Purpose,
		UUID:        res.UUID,
	}
	return vres
}

// newStoredStoragePackage converts projected type StoredStoragePackage to
// service type StoredStoragePackage.
func newStoredStoragePackage(vres *storageviews.StoredStoragePackageView) *StoredStoragePackage {
	res := &StoredStoragePackage{
		Location: vres.Location,
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.AipID != nil {
		res.AipID = *vres.AipID
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	if vres.ObjectKey != nil {
		res.ObjectKey = *vres.ObjectKey
	}
	if vres.Status == nil {
		res.Status = "unspecified"
	}
	return res
}

// newStoredStoragePackageView projects result type StoredStoragePackage to
// projected type StoredStoragePackageView using the "default" view.
func newStoredStoragePackageView(res *StoredStoragePackage) *storageviews.StoredStoragePackageView {
	vres := &storageviews.StoredStoragePackageView{
		Name:      &res.Name,
		AipID:     &res.AipID,
		Status:    &res.Status,
		ObjectKey: &res.ObjectKey,
		Location:  res.Location,
	}
	return vres
}
