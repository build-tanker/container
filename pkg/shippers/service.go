package shippers

import (
	"github.com/build-tanker/container/pkg/appcontext"
	"github.com/jmoiron/sqlx"
)

// Service - gets the request from the handler, controls how the service responds
type Service interface {
	Add(appGroup string, expiry int) (string, error)
	Delete(id string) error
	View(id string) (Shipper, error)
	ViewAll() ([]Shipper, error)
}

type service struct {
	ctx       *appcontext.AppContext
	datastore Datastore
}

// NewService - initialise a new service
func NewService(ctx *appcontext.AppContext, db *sqlx.DB) Service {
	datastore := NewDatastore(ctx, db)
	return &service{ctx, datastore}
}

func (s *service) Add(appGroup string, expiry int) (string, error) {
	return s.datastore.Add(appGroup, expiry)
}

func (s *service) Delete(id string) error {
	return s.datastore.Delete(id)
}

func (s *service) View(id string) (Shipper, error) {
	return s.datastore.View(id)
}

func (s *service) ViewAll() ([]Shipper, error) {
	return s.datastore.ViewAll()
}
