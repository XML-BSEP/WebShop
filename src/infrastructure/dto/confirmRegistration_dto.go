package dto

type ConfirmRegistrationDTO struct {
	Email string `json:"email" validate:"required,email"`
	VerificationCode string`json:"code" validate:"required"`
}