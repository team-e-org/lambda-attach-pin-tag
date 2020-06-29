package models

type Pin struct {
	ID          int
	UserID      *int
	Title       string
	Description *string
	URL         *string
	ImageURL    string
	IsPrivate   bool
	CreatedAt   string
	UpdatedAt   string
}
