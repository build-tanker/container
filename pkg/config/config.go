package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config - structure to hold the configuration for container
type Config struct {
	fileStore             string
	port                  string
	enableDelayMiddleware bool
	gcsJSONConfig         string
	gcsBucket             string
	database              DatabaseConfig
}

// NewConfig - create a new configuration
func NewConfig(paths []string) *Config {
	config := &Config{}

	viper.AutomaticEnv()

	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	viper.SetConfigName("container")
	viper.SetConfigType("toml")

	viper.SetDefault("application.fileStore", "googlecloud")
	viper.SetDefault("server.port", "4000")
	viper.SetDefault("server.enableDelayMiddleware", false)
	viper.SetDefault("googlecloud.jsonConfig", "")
	viper.SetDefault("googlecloud.bucket", "shrieking-cat")

	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file %s was edited, reloading config\n", e.Name)
		config.readLatestConfig()
	})

	config.readLatestConfig()

	return config
}

// FileStore - get the filestore from config
func (c *Config) FileStore() string {
	return c.fileStore
}

// Port - get the port from config
func (c *Config) Port() string {
	return c.port
}

// EnableDelayMiddleware - get if delay middlware is enabled or not from the config
func (c *Config) EnableDelayMiddleware() bool {
	return c.enableDelayMiddleware
}

// Database - load the database config
func (c *Config) Database() DatabaseConfig {
	return c.database
}

// GCSJSONConfig - get the google cloud storage json path from config
func (c *Config) GCSJSONConfig() string {
	return c.gcsJSONConfig
}

// GCSBucket - get the google cloud storage bucket from config
func (c *Config) GCSBucket() string {
	return c.gcsBucket
}

func (c *Config) readLatestConfig() {
	c.fileStore = viper.GetString("application.fileStore")
	c.port = viper.GetString("server.port")
	c.enableDelayMiddleware = viper.GetBool("server.enableDelayMiddleware")
	c.gcsJSONConfig = viper.GetString("googlecloud.jsonConfig")
	c.gcsBucket = viper.GetString("googlecloud.bucket")
	c.database = NewDatabaseConfig()
}
