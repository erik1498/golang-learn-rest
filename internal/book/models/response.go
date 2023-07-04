package models

type BookResponse struct {
	Title     string `json:"title"`
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
