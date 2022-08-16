// Code generated by goa v3.8.1, DO NOT EDIT.
//
// storage views
//
// Command:
// $ goa-v3.8.1 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// StoredLocationCollection is the viewed result type that is projected based
// on a view.
type StoredLocationCollection struct {
	// Type to project
	Projected StoredLocationCollectionView
	// View to render
	View string
}

// StoredStoragePackage is the viewed result type that is projected based on a
// view.
type StoredStoragePackage struct {
	// Type to project
	Projected *StoredStoragePackageView
	// View to render
	View string
}

// StoredLocationCollectionView is a type that runs validations on a projected
// type.
type StoredLocationCollectionView []*StoredLocationView

// StoredLocationView is a type that runs validations on a projected type.
type StoredLocationView struct {
	// ID is the unique id of the location.
	ID *uint
	// Name of location
	Name *string
	// Data source of the location
	Source *string
	// Purpose of the location
	Purpose *string
	UUID    *string
}

// StoredStoragePackageView is a type that runs validations on a projected type.
type StoredStoragePackageView struct {
	ID    *uint
	Name  *string
	AipID *string
	// Status of the package
	Status    *string
	ObjectKey *string
	Location  *string
}

var (
	// StoredLocationCollectionMap is a map indexing the attribute names of
	// StoredLocationCollection by view name.
	StoredLocationCollectionMap = map[string][]string{
		"default": {
			"name",
			"source",
			"purpose",
			"uuid",
		},
	}
	// StoredStoragePackageMap is a map indexing the attribute names of
	// StoredStoragePackage by view name.
	StoredStoragePackageMap = map[string][]string{
		"default": {
			"name",
			"aip_id",
			"status",
			"object_key",
			"location",
		},
	}
	// StoredLocationMap is a map indexing the attribute names of StoredLocation by
	// view name.
	StoredLocationMap = map[string][]string{
		"default": {
			"name",
			"source",
			"purpose",
			"uuid",
		},
	}
)

// ValidateStoredLocationCollection runs the validations defined on the viewed
// result type StoredLocationCollection.
func ValidateStoredLocationCollection(result StoredLocationCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredLocationCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateStoredStoragePackage runs the validations defined on the viewed
// result type StoredStoragePackage.
func ValidateStoredStoragePackage(result *StoredStoragePackage) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredStoragePackageView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateStoredLocationCollectionView runs the validations defined on
// StoredLocationCollectionView using the "default" view.
func ValidateStoredLocationCollectionView(result StoredLocationCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredLocationView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredLocationView runs the validations defined on
// StoredLocationView using the "default" view.
func ValidateStoredLocationView(result *StoredLocationView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Source == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("source", "result"))
	}
	if result.Purpose == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("purpose", "result"))
	}
	if result.UUID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("uuid", "result"))
	}
	if result.Source != nil {
		if !(*result.Source == "unspecified" || *result.Source == "minio") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.source", *result.Source, []interface{}{"unspecified", "minio"}))
		}
	}
	if result.Purpose != nil {
		if !(*result.Purpose == "unspecified" || *result.Purpose == "aip_store") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.purpose", *result.Purpose, []interface{}{"unspecified", "aip_store"}))
		}
	}
	return
}

// ValidateStoredStoragePackageView runs the validations defined on
// StoredStoragePackageView using the "default" view.
func ValidateStoredStoragePackageView(result *StoredStoragePackageView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.AipID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("aip_id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.ObjectKey == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("object_key", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "unspecified" || *result.Status == "in_review" || *result.Status == "rejected" || *result.Status == "stored" || *result.Status == "moving") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"unspecified", "in_review", "rejected", "stored", "moving"}))
		}
	}
	return
}
