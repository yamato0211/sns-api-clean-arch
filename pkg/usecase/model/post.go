package model

import "clean-sns-api/pkg/domain/model"

type Post struct {
	ID      int             `json:"id"`
	Content string          `json:"content"`
	UserID  int             `json:"user_id"`
	User    model.User      `json:"user"`
	Likes   model.LikeSlice `json:"likes"`
}

type PostSlice []*Post

func PostFromDomainModel(m *model.Post) *Post {
	p := &Post{
		ID:      m.ID,
		Content: m.Content,
		UserID:  m.UserID,
		User:    *m.R.User,
		Likes:   m.R.Likes,
	}

	return p
}
