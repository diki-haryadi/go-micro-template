package healthCheckConfigurator

import (
	"context"

	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"

	kafkaHealthCheckUseCase "github.com/diki-haryadi/go-micro-template/internal/health_check/usecase/kafka_health_check"
	postgresHealthCheckUseCase "github.com/diki-haryadi/go-micro-template/internal/health_check/usecase/postgres_health_check"
	tmpDirHealthCheckUseCase "github.com/diki-haryadi/go-micro-template/internal/health_check/usecase/tmp_dir_health_check"

	healthCheckGrpc "github.com/diki-haryadi/go-micro-template/internal/health_check/delivery/grpc"
	healthCheckHttp "github.com/diki-haryadi/go-micro-template/internal/health_check/delivery/http"
	healthCheckDomain "github.com/diki-haryadi/go-micro-template/internal/health_check/domain"
	healthCheckUseCase "github.com/diki-haryadi/go-micro-template/internal/health_check/usecase"
	infraContainer "github.com/diki-haryadi/ztools/infra_container"
)

type configurator struct {
	ic *infraContainer.IContainer
}

func NewConfigurator(ic *infraContainer.IContainer) healthCheckDomain.Configurator {
	return &configurator{ic: ic}
}

func (c *configurator) Configure(ctx context.Context) error {
	postgresHealthCheckUc := postgresHealthCheckUseCase.NewUseCase(c.ic.Postgres)
	kafkaHealthCheckUc := kafkaHealthCheckUseCase.NewUseCase()
	tmpDirHealthCheckUc := tmpDirHealthCheckUseCase.NewUseCase()

	healthCheckUc := healthCheckUseCase.NewUseCase(postgresHealthCheckUc, kafkaHealthCheckUc, tmpDirHealthCheckUc)

	// grpc
	grpcController := healthCheckGrpc.NewController(healthCheckUc, postgresHealthCheckUc, kafkaHealthCheckUc, tmpDirHealthCheckUc)
	grpcHealthV1.RegisterHealthServer(c.ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	// http
	httpRouterGp := c.ic.EchoHttpServer.GetEchoInstance().Group(c.ic.EchoHttpServer.GetBasePath())
	httpController := healthCheckHttp.NewController(healthCheckUc)
	healthCheckHttp.NewRouter(httpController).Register(httpRouterGp)

	return nil
}
