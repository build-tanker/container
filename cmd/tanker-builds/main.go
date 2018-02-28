package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/gojekfarm/tanker-builds/pkg/appcontext"
	"github.com/gojekfarm/tanker-builds/pkg/config"
	"github.com/gojekfarm/tanker-builds/pkg/logger"
	"github.com/gojekfarm/tanker-builds/pkg/postgres"
	"github.com/gojekfarm/tanker-builds/pkg/server"
)

func main() {
	config := config.NewConfig([]string{".", "..", "../.."})
	logger := logger.NewLogger(config, os.Stdout)
	ctx := appcontext.NewAppContext(config, logger)
	db := postgres.NewPostgres(logger, config.Database().ConnectionURL(), config.Database().MaxPoolSize())
	server := server.NewServer(ctx, db)

	logger.Infoln("Starting tanker-builds")

	app := cli.NewApp()
	app.Name = "tanker-builds"
	app.Version = "0.0.1"
	app.Usage = "this service saves files and makes them available for distribution"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the service",
			Action: func(c *cli.Context) error {
				return server.Start()
			},
		},
		{
			Name:  "migrate",
			Usage: "run database migrations",
			Action: func(c *cli.Context) error {
				return postgres.RunDatabaseMigrations(ctx)
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback the latest database migration",
			Action: func(c *cli.Context) error {
				return postgres.RollbackDatabaseMigration(ctx)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
