package repository

import "time"

type CreateBody struct {
	RedirectURL string `json:"redirect_url" form:"url"`
}

type URL struct {
	ID          int       `json:"id"`
	URN         string    `json:"urn"`
	RedirectURL string    `json:"redirect_url"`
	CreatedAt   time.Time `json:"created_at"`
}
