package models

import (
	"fmt"
	"strings"
)

// RegisterRequest struct
type RegisterRequest struct {
	MSISDN   string `json:"msisdn" validate:"required"`
	Username string `json:"username" validate:"required"`
	Name     string `json:"name"`
	Password string `json:"password" validate:"required"`
}

// Auth creates auth struct from RegisterRequest
func (m RegisterRequest) Auth() Auth {
	if m.MSISDN[0:1] == "0" {
		m.MSISDN = strings.Replace(m.MSISDN, "0", "62", 1)
	}
	if m.MSISDN[0:1] == "8" {
		m.MSISDN = fmt.Sprintf("%s%s", "62", m.MSISDN)
	}
	return Auth{
		MSISDN:   m.MSISDN,
		Username: m.Username,
		Name:     m.Name,
		Password: m.Password,
	}
}
