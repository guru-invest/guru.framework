package models

import "github.com/golang-jwt/jwt"

var AUTHENTICATEDUSER AuthenticatedUser

type AuthenticatedUser struct {
	JwtToken string `json:"-"`
	DeviceID string `json:"device_id"`
	Customer struct {
		CustomerCode   string `json:"customer_code"`
		Name           string `json:"name"`
		NickName       string `json:"nick_name"`
		Email          string `json:"email"`
		DocumentNumber string `json:"document_number"`
	} `json:"customer"`
	External struct {
		CustomerID         string            `json:"external_id"`
		PartnerIds         map[string]string `json:"partner_ids"`
		Status             string            `json:"status"`
		SuitabilityProfile string            `json:"suitability"`
	} `json:"external"`
	jwt.StandardClaims
}
