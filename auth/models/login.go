package models

// LoginRequest struct
type LoginRequest struct {
	MSISDN   string `json:"msisdn" validate:"required"`
	Password string `json:"password" validate:"required"`
}
