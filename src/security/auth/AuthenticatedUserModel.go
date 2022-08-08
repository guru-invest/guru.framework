package auth

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
)

var AUTHENTICATEDUSERMODEL AuthenticatedUserModel

type AuthenticatedUserModel struct {
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
		CustomerID         string `json:"external_id"`
		GenialID           string `json:"genial_id"`
		Status             string `json:"status"`
		SuitabilityProfile string `json:"suitability"`
	} `json:"external"`
	jwt.StandardClaims
}

func (t *AuthenticatedUserModel) Validate(tokenString string, signingKey string, issuer string) error {
	authenticatedUserClaims := AuthenticatedUserModel{}
	tokenString = strings.ReplaceAll(strings.ReplaceAll(tokenString, "bearer ", ""), "Bearer ", "")

	token, err := jwt.ParseWithClaims(tokenString, &authenticatedUserClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*AuthenticatedUserModel); ok && token.Valid {
		if claims.StandardClaims.VerifyIssuer(issuer, true) {
			AUTHENTICATEDUSERMODEL = *claims
			AUTHENTICATEDUSERMODEL.JwtToken = tokenString
			return nil
		}
	}
	return err
}
