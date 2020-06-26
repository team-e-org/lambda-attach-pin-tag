package view

import (
	"hello/models"
	"time"
)

type Pin struct {
	ID          int        `json:"id"`
	UserID      int        `json:"userId"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	URL         *string    `json:"url,omitempty"`
	ImageURL    string     `json:"imageUrl"`
	IsPrivate   bool       `json:"isPrivate"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func NewPin(pin *models.Pin) *Pin {
	p := &Pin{
		ID:          pin.ID,
		UserID:      *pin.UserID,
		Title:       pin.Title,
		Description: *pin.Description,
		URL:         pin.URL,
		ImageURL:    pin.ImageURL,
		IsPrivate:   pin.IsPrivate,
	}

	return p
}
