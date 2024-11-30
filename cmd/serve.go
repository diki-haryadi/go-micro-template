package cmd

import (
	"fmt"
	"github.com/diki-haryadi/go-micro-template/app"
	"github.com/diki-haryadi/go-micro-template/config"
	"github.com/diki-haryadi/ztools/logger"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:              "serve",
		Short:            "A API for publish audio to kafka",
		Long:             "A API for publish audio to kafka",
		PersistentPreRun: servePreRun,
		RunE:             runServe,
	}
)

func ServeCmd() *cobra.Command {
	return serveCmd
}

func servePreRun(cmd *cobra.Command, args []string) {
	config.LoadConfig()
}

func runServe(cmd *cobra.Command, args []string) error {
	err := app.New().Init().Run()
	if err != nil {
		fmt.Println(err)
		logger.Zap.Sugar().Fatal(err)
	}

	return nil
}
