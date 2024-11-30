package cmd

import (
	"context"
	"fmt"
	"github.com/RichardKnop/go-fixtures"
	"github.com/diki-haryadi/ztools/config"
	"github.com/diki-haryadi/ztools/env"
	"github.com/diki-haryadi/ztools/logger"
	"github.com/diki-haryadi/ztools/postgres"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
	"runtime"
	"time"
)

var (
	loadDataCmd = &cobra.Command{
		Use:              "load_data",
		Short:            "Load data into the system",
		Long:             "Load data into the system for initialization or testing purposes",
		PersistentPreRun: loadDataPreRun,
		RunE:             runLoadData,
	}
)

func LoadDataCmd() *cobra.Command {
	return loadDataCmd
}

func loadDataPreRun(cmd *cobra.Command, args []string) {
	_, callerDir, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Error generating env dir")
	}

	// Define the possible paths to the .env file
	envPaths := []string{
		filepath.Join(filepath.Dir(callerDir), "..", "envs/.env"),
	}

	// Load the .env file from the provided paths
	env.LoadEnv(envPaths...) // Use ... to expand the slice
	config.NewConfig()

	loggerPath := []string{
		filepath.Join(filepath.Dir(callerDir), "..", "tmp/logs"),
	}
	logger.NewLogger(loggerPath...)

}

func runLoadData(cmd *cobra.Command, args []string) error {
	pg, err := postgres.NewConnection(context.Background(), &postgres.Config{
		Host:    config.BaseConfig.Postgres.Host,
		Port:    config.BaseConfig.Postgres.Port,
		User:    config.BaseConfig.Postgres.User,
		Pass:    config.BaseConfig.Postgres.Pass,
		DBName:  config.BaseConfig.Postgres.DBName,
		SslMode: config.BaseConfig.Postgres.SslMode,
	})
	defer pg.SqlxDB.Close()
	if err != nil {
		return err
	}

	_, callerDir, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Error generating env dir")
	}
	envPaths := []string{
		filepath.Join(filepath.Dir(callerDir), "..", "db/fixtures/article.yml"),
		filepath.Join(filepath.Dir(callerDir), "..", "db/fixtures/scopes.yml"),
		filepath.Join(filepath.Dir(callerDir), "..", "db/fixtures/test_access_tokens.yml"),
		filepath.Join(filepath.Dir(callerDir), "..", "db/fixtures/test_clients.yml"),
		filepath.Join(filepath.Dir(callerDir), "..", "db/fixtures/test_users.yml"),
	}

	bar := progressbar.NewOptions(len(envPaths),
		progressbar.OptionSetDescription("Loading fixtures..."),
		progressbar.OptionShowCount(),
		progressbar.OptionShowDescriptionAtLineEnd(),
	)

	for _, path := range envPaths {
		description := fmt.Sprintf("Processing: %s", filepath.Base(path))
		bar.Describe(description)
		err = fixtures.LoadFiles(envPaths, pg.SqlxDB.DB, "postgres")
		if err != nil {
			return err
		}

		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
	bar.Describe("Finished")
	fmt.Println("\nFinished loading data")
	return nil
}
