package service

import "time"

type ShortedLink struct {
	Id        string    `json:"id"`
	LongLink  string    `json:"long_link"`
	ShortLink string    `json:"short_link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    string    `json:"user_id"`
}

type ShortLinkInput struct {
	Link string `json:"link" binding:"required"`
}
