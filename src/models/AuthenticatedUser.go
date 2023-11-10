package models

import "github.com/golang-jwt/jwt"

var AUTHENTICATEDUSER AuthenticatedUser

type AuthenticatedUser struct {
	JwtToken string `json:"-"`
	DeviceID string `json:"device_id"`
	Customer struct {
		CustomerCode         string `json:"customer_code"`
		Name                 string `json:"name"`
		NickName             string `json:"nick_name"`
		Email                string `json:"email"`
		DocumentNumber       string `json:"document_number"`
		HasPassword          string `json:"has_password"`
		CreatedAt            string `json:"created_at"`
		HasActiveRLP         bool   `json:"has_active_rlp"`
		UserAcceptCryptoTerm bool   `json:"user_accept_crypto_term"`
		Bonus                struct {
			ContaIncentivada string `json:"conta_incentivada"`
		} `json:"bonus"`
		HasAcceptTerms             bool   `json:"has_accept_terms"`
		ProductInvestmentPeference string `json:"product_invest_preference"`
	} `json:"customer"`
	External struct {
		CustomerID         string            `json:"external_id"`
		PartnerIDs         map[string]string `json:"partner_ids"`
		PartnerTokens      map[string]string `json:"partner_tokens"`
		Status             string            `json:"status"`
		SuitabilityProfile string            `json:"suitability"`
	} `json:"external"`
	Smartportfolio struct {
		SelectedRisk int `json:"selected_risk"`
	} `json:"smartportfolio"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}
