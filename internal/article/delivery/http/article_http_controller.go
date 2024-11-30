package articleHttpController

import (
	"github.com/diki-haryadi/go-micro-template/pkg/response"
	"github.com/labstack/echo/v4"

	articleDomain "github.com/diki-haryadi/go-micro-template/internal/article/domain"
	articleDto "github.com/diki-haryadi/go-micro-template/internal/article/dto"
)

type controller struct {
	useCase articleDomain.UseCase
}

func NewController(uc articleDomain.UseCase) articleDomain.HttpController {
	return &controller{
		useCase: uc,
	}
}

func (c controller) CreateArticle(ctx echo.Context) error {
	res := response.NewJSONResponse()
	aDto := new(articleDto.CreateArticleRequestDto)
	if err := ctx.Bind(aDto); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	if err := aDto.ValidateCreateArticleDto(); err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	article, err := c.useCase.CreateArticle(ctx.Request().Context(), aDto)
	if err != nil {
		res.SetError(response.ErrBadRequest).SetMessage(err.Error()).Send(ctx.Response().Writer)
		return nil
	}

	res.APIStatusSuccess().SetData(article).Send(ctx.Response().Writer)
	return nil
}
