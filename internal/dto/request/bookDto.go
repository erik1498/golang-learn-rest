package dto

type UpdateBook struct {
	Title string `json:"title" validate:"required,min=5,max=20"`
}
