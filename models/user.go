package models

type RegisterParameter struct {
	Name            string `json:"name"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
