package config_test

import (
	"testing"

	"github.com/build-tanker/container/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigValues(t *testing.T) {
	conf := config.NewConfig([]string{"./testutil"})
	assert.Equal(t, "dbname=container user=container password='container' host=localhost port=5432 sslmode=disable", conf.Database().ConnectionString())
	assert.Equal(t, "postgres://container:container@localhost:5432/container?sslmode=disable", conf.Database().ConnectionURL())
}
