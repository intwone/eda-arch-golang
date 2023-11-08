package inputs

type AuthCreateInput struct {
	Value string `json:"value" validate:"required,email"`
}
