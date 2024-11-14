package sampleExtServiceDomain

import (
	"context"

	articleV1 "github.com/diki-haryadi/protobuf-template/go-micro-template/article/v1"
)

type SampleExtServiceUseCase interface {
	CreateArticle(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error)
}
