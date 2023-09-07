package persistence

import (
	"context"

	goapackage "github.com/artefactual-sdps/enduro/internal/api/gen/package_"
	"github.com/artefactual-sdps/enduro/internal/event"
	"github.com/artefactual-sdps/enduro/internal/package_"
)

type eventManager struct {
	evsvc event.EventService
	inner Service
}

var _ Service = (*eventManager)(nil)

// WithEvents decorates a persistence service implementation with event
// publication to evsvc.
func WithEvents(evsvc event.EventService, inner Service) *eventManager {
	return &eventManager{evsvc: evsvc, inner: inner}
}

// CreatePackage creates and persists a new package using the values from pkg,
// publishes a "package created" event, then returns the updated package.
//
// The input pkg "ID" and "CreatedAt" values are ignored; the stored package
// "ID" is generated by the persistence implementation and "CreatedAt" is always
// set to the current time.
func (m *eventManager) CreatePackage(ctx context.Context, pkg *package_.Package) (*package_.Package, error) {
	pkg, err := m.inner.CreatePackage(ctx, pkg)
	if err != nil {
		return nil, err
	}

	// Publish a "package created" event.
	ev := &goapackage.EnduroPackageCreatedEvent{ID: uint(pkg.ID), Item: pkg.Goa()}
	event.PublishEvent(ctx, m.evsvc, ev)

	return pkg, nil
}

// UpdatePackage updates the persisted package identified by id using the
// updater function, publishes a "package updated" event, then returns the
// updated package.
//
// The package "ID" and "CreatedAt" field values can not be updated with this
// method.
func (m *eventManager) UpdatePackage(ctx context.Context, id uint, updater PackageUpdater) (*package_.Package, error) {
	pkg, err := m.inner.UpdatePackage(ctx, id, updater)
	if err != nil {
		return nil, err
	}

	// Publish a "package updated" event.
	ev := &goapackage.EnduroPackageUpdatedEvent{ID: pkg.ID, Item: pkg.Goa()}
	event.PublishEvent(ctx, m.evsvc, ev)

	return pkg, nil
}
