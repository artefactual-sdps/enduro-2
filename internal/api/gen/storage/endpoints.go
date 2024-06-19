// Code generated by goa v3.15.2, DO NOT EDIT.
//
// storage endpoints
//
// Command:
// $ goa gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package storage

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "storage" service endpoints.
type Endpoints struct {
	Create           goa.Endpoint
	Submit           goa.Endpoint
	Update           goa.Endpoint
	Download         goa.Endpoint
	Move             goa.Endpoint
	MoveStatus       goa.Endpoint
	Reject           goa.Endpoint
	Show             goa.Endpoint
	Locations        goa.Endpoint
	AddLocation      goa.Endpoint
	ShowLocation     goa.Endpoint
	LocationPackages goa.Endpoint
}

// NewEndpoints wraps the methods of the "storage" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Create:           NewCreateEndpoint(s, a.JWTAuth),
		Submit:           NewSubmitEndpoint(s, a.JWTAuth),
		Update:           NewUpdateEndpoint(s, a.JWTAuth),
		Download:         NewDownloadEndpoint(s, a.JWTAuth),
		Move:             NewMoveEndpoint(s, a.JWTAuth),
		MoveStatus:       NewMoveStatusEndpoint(s, a.JWTAuth),
		Reject:           NewRejectEndpoint(s, a.JWTAuth),
		Show:             NewShowEndpoint(s, a.JWTAuth),
		Locations:        NewLocationsEndpoint(s, a.JWTAuth),
		AddLocation:      NewAddLocationEndpoint(s, a.JWTAuth),
		ShowLocation:     NewShowLocationEndpoint(s, a.JWTAuth),
		LocationPackages: NewLocationPackagesEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "storage" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.Submit = m(e.Submit)
	e.Update = m(e.Update)
	e.Download = m(e.Download)
	e.Move = m(e.Move)
	e.MoveStatus = m(e.MoveStatus)
	e.Reject = m(e.Reject)
	e.Show = m(e.Show)
	e.Locations = m(e.Locations)
	e.AddLocation = m(e.AddLocation)
	e.ShowLocation = m(e.ShowLocation)
	e.LocationPackages = m(e.LocationPackages)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "storage".
func NewCreateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreatePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:package:create"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.Create(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPackage(res, "default")
		return vres, nil
	}
}

// NewSubmitEndpoint returns an endpoint function that calls the method
// "submit" of service "storage".
func NewSubmitEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*SubmitPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:package:submit"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Submit(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "storage".
func NewUpdateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdatePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:package:submit"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Update(ctx, p)
	}
}

// NewDownloadEndpoint returns an endpoint function that calls the method
// "download" of service "storage".
func NewDownloadEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DownloadPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:package:download"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Download(ctx, p)
	}
}

// NewMoveEndpoint returns an endpoint function that calls the method "move" of
// service "storage".
func NewMoveEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*MovePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:package:move"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Move(ctx, p)
	}
}

// NewMoveStatusEndpoint returns an endpoint function that calls the method
// "move_status" of service "storage".
func NewMoveStatusEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*MoveStatusPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:package:move"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.MoveStatus(ctx, p)
	}
}

// NewRejectEndpoint returns an endpoint function that calls the method
// "reject" of service "storage".
func NewRejectEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*RejectPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:package:review"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Reject(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "storage".
func NewShowEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ShowPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:package:read"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPackage(res, "default")
		return vres, nil
	}
}

// NewLocationsEndpoint returns an endpoint function that calls the method
// "locations" of service "storage".
func NewLocationsEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*LocationsPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:location:list"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.Locations(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedLocationCollection(res, "default")
		return vres, nil
	}
}

// NewAddLocationEndpoint returns an endpoint function that calls the method
// "add_location" of service "storage".
func NewAddLocationEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*AddLocationPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:location:create"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.AddLocation(ctx, p)
	}
}

// NewShowLocationEndpoint returns an endpoint function that calls the method
// "show_location" of service "storage".
func NewShowLocationEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ShowLocationPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:location:read"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.ShowLocation(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedLocation(res, "default")
		return vres, nil
	}
}

// NewLocationPackagesEndpoint returns an endpoint function that calls the
// method "location_packages" of service "storage".
func NewLocationPackagesEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*LocationPackagesPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"package:list", "package:listActions", "package:move", "package:read", "package:review", "package:upload", "storage:location:create", "storage:location:list", "storage:location:listPackages", "storage:location:read", "storage:package:create", "storage:package:download", "storage:package:move", "storage:package:read", "storage:package:review", "storage:package:submit"},
			RequiredScopes: []string{"storage:location:listPackages"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.LocationPackages(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPackageCollection(res, "default")
		return vres, nil
	}
}
