package model

type Like struct {
	ID     int  `json:"id"`
	UserID int  `json:"user_id"`
	User   User `json:"user"`
	PostID int  `json:"post_id"`
	Post   Post `json:"post"`
}
