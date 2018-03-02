package builds

import (
	"github.com/build-tanker/container/pkg/appcontext"
	"github.com/build-tanker/container/pkg/filestore"
	"github.com/jmoiron/sqlx"
)

// Service - handles business logic for builds
type Service interface {
	Add(fileName, shipper, bundle, platform, extension string) (string, error)
}

type service struct {
	ctx       *appcontext.AppContext
	fs        filestore.FileStore
	datastore Datastore
}

// NewService - create a new service for builds
func NewService(ctx *appcontext.AppContext, db *sqlx.DB) Service {
	datastore := NewDatastore(ctx, db)
	s := &service{
		ctx:       ctx,
		datastore: datastore,
	}
	s.init()
	return s
}

func (s *service) init() {
	log := s.ctx.GetLogger()
	fileStore := s.ctx.GetConfig().FileStore()
	switch fileStore {
	case "googlecloud":
		s.fs = filestore.NewGoogleCloudStorageFileStore(s.ctx)
		err := s.fs.Setup()
		if err != nil {
			log.Fatalln("Could not setup GoogleCloudStorage", err.Error())
		}
	case "s3":
		log.Fatalln("This FileStore is not supported:", fileStore)
	case "local":
		log.Fatalln("This FileStore is not supported:", fileStore)
	default:
		log.Fatalln("This FileStore is not supported:", fileStore)
	}
}

func (s *service) Add(fileName, shipper, bundle, platform, extension string) (string, error) {
	// Does two things
	// Get a url from the google cloud package and return it
	url, err := s.fs.GetWriteURL()
	if err != nil {
		return "", err
	}
	// Create an entry in the database
	_, err = s.datastore.Add(fileName, shipper, bundle, platform, extension)
	if err != nil {
		return "", err
	}
	return url, nil
}
