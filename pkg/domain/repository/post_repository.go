package repository

import (
	"clean-sns-api/pkg/domain/model"
	"context"
)

type IPostRepository interface {
	SelectAllPosts(ctx context.Context) (model.PostSlice, error)
	SelectPostByID(ctx context.Context, id int) (*model.Post, error)
}
