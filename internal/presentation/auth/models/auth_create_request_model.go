package models

type AuthCreateRequestModel struct {
	Value string `json:"value" binding:"required,email"`
}
