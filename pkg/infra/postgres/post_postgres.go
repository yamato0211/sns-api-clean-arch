package postgres

import (
	"clean-sns-api/pkg/domain/model"
	"clean-sns-api/pkg/domain/repository"
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type postRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) repository.IPostRepository {
	return &postRepository{
		DB: db,
	}
}

func (pr *postRepository) SelectAllPosts(ctx context.Context) (model.PostSlice, error) {
	return model.Posts(
		qm.Load("User"),
		qm.Load("Likes"),
	).All(ctx, pr.DB)
}

func (pr *postRepository) SelectPostByID(ctx context.Context, postID int) (*model.Post, error) {
	whereID := fmt.Sprintf("%s = ?", model.PostColumns.ID)
	return model.Posts(
		qm.Where(whereID, postID),
		qm.Load("User"),
		qm.Load("Likes"),
	).One(ctx, pr.DB)
}
