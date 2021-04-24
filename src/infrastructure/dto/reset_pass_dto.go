package dto

type ResetPassDTO struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	VerificationCode string`json:"code" validate:"required"`
}
