package cmd

import (
	"auth-microservice/src/config"
	"auth-microservice/src/log"
	"auth-microservice/src/services/db"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

const path = "migrations"

var mgrCmd = &cobra.Command{
	Use: "migration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewEnvConfig()
		logger:= log.NewLogger()
		database := db.GetDB(cfg,logger)
		defer db.CloseDB(database,logger)
		if err := goose.Run(args[0], database.DB(), path, args[1:]...); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(mgrCmd)
}
