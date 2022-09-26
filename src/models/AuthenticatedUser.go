package models

import "github.com/golang-jwt/jwt"

var AUTHENTICATEDUSER AuthenticatedUser

type AuthenticatedUser struct {
	JwtToken string `json:"-"`
	DeviceID string `json:"device_id"`
	Customer struct {
		CustomerCode        string `json:"customer_code"`
		Name                string `json:"name"`
		NickName            string `json:"nick_name"`
		Email               string `json:"email"`
		DocumentNumber      string `json:"document_number"`
		HasPassword         string `json:"has_password"`
		UserAcceptMigration *bool  `json:"user_accept_migration"`
		HasActiveRlp        *bool  `json:"has_active_rlp"`
	} `json:"customer"`
	External struct {
		CustomerID         string            `json:"external_id"`
		PartnerIDs         map[string]string `json:"partner_ids"`
		PartnerTokens      map[string]string `json:"partner_tokens"`
		Status             string            `json:"status"`
		SuitabilityProfile string            `json:"suitability"`
	} `json:"external"`
	jwt.StandardClaims
}
