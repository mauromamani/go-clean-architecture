package entity

import "time"

type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int64     `json:"user_id"`
}
