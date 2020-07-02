package view

type Pin struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	ImageURL    string `json:"imageUrl"`
	IsPrivate   bool   `json:"isPrivate"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}
