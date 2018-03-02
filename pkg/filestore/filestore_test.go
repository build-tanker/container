package filestore

import (
	"os"

	"github.com/build-tanker/container/pkg/appcontext"
	"github.com/build-tanker/container/pkg/config"
	"github.com/build-tanker/container/pkg/logger"
)

var testContext *appcontext.AppContext

func NewTestContext() *appcontext.AppContext {

	if testContext == nil {
		conf := config.NewConfig([]string{".", "../config/testutil"})
		log := logger.NewLogger(conf, os.Stdout)
		testContext = appcontext.NewAppContext(conf, log)
	}
	return testContext
}
