package dtos

type AuthCreateRequestDTO struct {
	Value string `json:"value" binding:"required,email"`
}
