package model

type Follow struct {
	ID         int  `json:"id"`
	FollowerID int  `json:"follower_id"`
	Follower   User `json:"follower"`
	FolloweeID int  `json:"followee_id"`
	Followee   User `json:"followee"`
}
