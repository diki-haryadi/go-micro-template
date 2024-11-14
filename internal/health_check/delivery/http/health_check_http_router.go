package healthCheckHttp

import (
	"github.com/labstack/echo/v4"

	healthCheckDomain "github.com/diki-haryadi/go-micro-template/internal/health_check/domain"
)

type Router struct {
	controller healthCheckDomain.HttpController
}

func NewRouter(controller healthCheckDomain.HttpController) *Router {
	return &Router{
		controller: controller,
	}
}

func (r *Router) Register(e *echo.Group) {
	e.GET("/health", r.controller.Check)
}
