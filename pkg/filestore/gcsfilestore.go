package filestore

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/build-tanker/container/pkg/appcontext"
	"github.com/build-tanker/disk"
	uuid "github.com/satori/go.uuid"
)

// GoogleCloudStorageFileStore - the filestore interface for GCS
type GoogleCloudStorageFileStore interface {
}

type googleCloudStorageFileStore struct {
	ctx   *appcontext.AppContext
	creds *googleCredentials
	dd    disk.Disk
	gcs   GoogleCloudStorage
}

type googleCredentials struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

// NewGoogleCloudStorageFileStore - initialise a new GCS filestore
func NewGoogleCloudStorageFileStore(ctx *appcontext.AppContext) FileStore {
	dd := disk.NewDisk()
	gcs := NewGoogleCloudStorage()
	return &googleCloudStorageFileStore{
		ctx:   ctx,
		creds: &googleCredentials{},
		dd:    dd,
		gcs:   gcs,
	}
}

func (g *googleCloudStorageFileStore) GetWriteURL() (string, error) {
	bucket := g.ctx.GetConfig().GCSBucket()

	if bucket == "" {
		return "", errors.New("Please define a bucket to upload to")
	}

	duration := 60 * time.Minute
	expiration := time.Now().Add(duration)
	key := uuid.NewV4().String()
	// key := fmt.Sprintf("%s.%s", uuid.NewV4().String(), "pdf")

	signed, err := g.gcs.SignedURL(bucket, key, g.creds.ClientEmail, []byte(g.creds.PrivateKey), http.MethodPut, expiration)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (g *googleCloudStorageFileStore) Setup() error {
	log := g.ctx.GetLogger()
	file := g.ctx.GetConfig().GCSJSONConfig()
	if file == "" {
		errMessage := fmt.Sprintln("Please define the config file google cloud storage. Current setup:", file)
		log.Fatalln(errMessage)
		return errors.New(errMessage)
	}

	data, err := g.dd.ReadCompleteFile(file)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	json.Unmarshal(data, &g.creds)
	return nil
}
