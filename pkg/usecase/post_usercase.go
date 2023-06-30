package usecase

import (
	"clean-sns-api/pkg/domain/service"
	"clean-sns-api/pkg/usecase/model"
	"context"
)

type IPostUsecase interface {
	FindAllPosts(ctx context.Context) (model.PostSlice, error)
	FindPostByID(ctx context.Context, id int) (*model.Post, error)
}

type postUsecase struct {
	svc service.IPostService
}

func NewPostUsecase(ps service.IPostService) *postUsecase {
	return &postUsecase{
		svc: ps,
	}
}

func (pu *postUsecase) FindAllPosts(ctx context.Context) (model.PostSlice, error) {
	mpSlice, err := pu.svc.FindAllPosts(ctx)
	if err != nil {
		return nil, err
	}

	pSlice := make(model.PostSlice, 0, len(mpSlice))
	for _, mp := range mpSlice {
		pSlice = append(pSlice, model.PostFromDomainModel(mp))
	}

	return pSlice, nil
}

func (su *postUsecase) FindPostByID(ctx context.Context, id int) (*model.Post, error) {
	ms, err := su.svc.FindPostByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.PostFromDomainModel(ms), nil
}
