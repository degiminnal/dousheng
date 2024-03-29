package entity

import "time"

type Message struct {
	Id         int64     `json:"id"`
	ToUserId   int64     `json:"to_user_id"`
	FromUserId int64     `json:"from_user_id"`
	Contents   string    `json:"contents"`
	CreateTime time.Time `json:"create_time"`
}
