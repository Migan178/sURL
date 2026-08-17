package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

var (
	ErrInvalid = fmt.Errorf("invalid request")
)

type CreateBody struct {
	RedirectURL string `form:"url"`
}

func (b *CreateBody) Normalize() {
	b.RedirectURL = strings.TrimSpace(b.RedirectURL)
}

func (b *CreateBody) Validate() error {
	if err := validate.Var(b.RedirectURL, "url"); err != nil {
		return err
	}

	return nil
}

type URL struct {
	ID          int       `json:"id"`
	URN         string    `json:"urn"`
	RedirectURL string    `json:"redirect_url"`
	CreatedAt   time.Time `json:"created_at"`
}
