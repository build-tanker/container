package config_test

import (
	"testing"

	"github.com/gojekfarm/tanker-builds/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigValues(t *testing.T) {
	conf := config.NewConfig([]string{"./testutil"})
	assert.Equal(t, "dbname=tankerBuilds user=tankerBuilds password='tankerBuilds' host=localhost port=5432 sslmode=disable", conf.Database().ConnectionString())
	assert.Equal(t, "postgres://tankerBuilds:tankerBuilds@localhost:5432/tankerBuilds?sslmode=disable", conf.Database().ConnectionURL())
}
