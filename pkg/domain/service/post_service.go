package service

import (
	"clean-sns-api/pkg/domain/model"
	"clean-sns-api/pkg/domain/repository"
	"context"
)

type IPostService interface {
	FindAllPosts(ctx context.Context) (model.PostSlice, error)
	FindPostByID(ctx context.Context, id int) (*model.Post, error)
}

type postService struct {
	repo repository.IPostRepository
}

func NewPostService(pr repository.IPostRepository) IPostService {
	return &postService{
		repo: pr,
	}
}

func (ps *postService) FindAllPosts(ctx context.Context) (model.PostSlice, error) {
	return ps.repo.SelectAllPosts(ctx)
}

func (ps *postService) FindPostByID(ctx context.Context, id int) (*model.Post, error) {
	return ps.repo.SelectPostByID(ctx, id)
}
