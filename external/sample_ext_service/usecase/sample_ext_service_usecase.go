package sampleExtServiceUseCase

import (
	"context"

	articleV1 "github.com/diki-haryadi/protobuf-template/go-micro-template/article/v1"

	sampleExtServiceDomain "github.com/diki-haryadi/go-micro-template/external/sample_ext_service/domain"
	grpcError "github.com/diki-haryadi/ztools/error/grpc"
	"github.com/diki-haryadi/ztools/grpc"
)

type sampleExtServiceUseCase struct {
	grpcClient grpc.Client
}

func NewSampleExtServiceUseCase(grpcClient grpc.Client) sampleExtServiceDomain.SampleExtServiceUseCase {
	return &sampleExtServiceUseCase{
		grpcClient: grpcClient,
	}
}

func (esu *sampleExtServiceUseCase) CreateArticle(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error) {
	articleGrpcClient := articleV1.NewArticleServiceClient(esu.grpcClient.GetGrpcConnection())

	res, err := articleGrpcClient.CreateArticle(ctx, req)
	if err != nil {
		return nil, grpcError.ParseExternalGrpcErr(err)
	}

	return res, nil
}
