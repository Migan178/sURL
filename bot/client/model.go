package client

import "time"

type URL struct {
	ID          int       `json:"id"`
	URN         string    `json:"urn"`
	RedirectURL string    `json:"redirect_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type errResponse struct {
	Message string `json:"message"`
}
