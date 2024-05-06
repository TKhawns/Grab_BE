package request

type RequestSignUp struct {
	FullName string `json:"full_name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,pwd"`
	Gender   string `json:"gender"`
}
