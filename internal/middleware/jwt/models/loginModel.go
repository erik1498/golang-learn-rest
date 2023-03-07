package models

type Login struct {
	Username string `json:"username" validation:"required"`
	Password string `json:"Password" validation:"required"`
}
