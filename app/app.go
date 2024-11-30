package app

import (
	"context"
	"github.com/diki-haryadi/ztools/config"
	"github.com/diki-haryadi/ztools/env"
	"github.com/diki-haryadi/ztools/logger"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	articleConfigurator "github.com/diki-haryadi/go-micro-template/internal/article/configurator"
	healthCheckConfigurator "github.com/diki-haryadi/go-micro-template/internal/health_check/configurator"
	externalBridge "github.com/diki-haryadi/ztools/external_bridge"
	iContainer "github.com/diki-haryadi/ztools/infra_container"
)

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Init() *App {
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

	return a
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	container := iContainer.IContainer{}
	ic, infraDown, err := container.IContext(ctx).
		ICDown().ICPostgres().ICGrpc().ICEcho().NewIC()
	if err != nil {
		return err
	}
	defer infraDown()

	extBridge, extBridgeDown, err := externalBridge.NewExternalBridge(ctx)
	if err != nil {
		return err
	}
	defer extBridgeDown()

	me := configureModule(ctx, ic, extBridge)
	if me != nil {
		return me
	}

	var serverError error
	go func() {
		if err := ic.GrpcServer.RunGrpcServer(ctx, nil); err != nil {
			ic.Logger.Sugar().Errorf("(s.RunGrpcServer) err: {%v}", err)
			serverError = err
			cancel()
		}
	}()

	go func() {
		if err := ic.EchoHttpServer.RunServer(ctx, nil); err != nil {
			ic.Logger.Sugar().Errorf("(s.RunEchoServer) err: {%v}", err)
			serverError = err
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		ic.Logger.Sugar().Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		ic.Logger.Sugar().Errorf("ctx.Done: %v", done)
	}

	ic.Logger.Sugar().Info("Server Exited Properly")
	return serverError
}

func configureModule(ctx context.Context, ic *iContainer.IContainer, extBridge *externalBridge.ExternalBridge) error {
	err := articleConfigurator.NewConfigurator(ic, extBridge).Configure(ctx)
	if err != nil {
		return err
	}

	err = healthCheckConfigurator.NewConfigurator(ic).Configure(ctx)
	if err != nil {
		return err
	}

	return nil
}
