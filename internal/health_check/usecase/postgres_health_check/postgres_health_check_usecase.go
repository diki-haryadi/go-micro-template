package postgresHealthCheckUseCase

import (
	healthCheckDomain "github.com/diki-haryadi/go-micro-template/internal/health_check/domain"
	"github.com/diki-haryadi/ztools/postgres"
)

type useCase struct {
	postgres *postgres.Postgres
}

func NewUseCase(postgres *postgres.Postgres) healthCheckDomain.PostgresHealthCheckUseCase {
	return &useCase{
		postgres: postgres,
	}
}

func (uc *useCase) Check() bool {
	if err := uc.postgres.SqlxDB.DB.Ping(); err != nil {
		return false
	}
	return true
}
