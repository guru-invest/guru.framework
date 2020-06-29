package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/guru-invest/guru.framework/src/security/claims"
)

type JWTAuthenticationStructure struct {
	Key []byte
}

const _expireOffset = 3600

type TokenStructure struct {
	Token string `json:"token"`
}

var authBackendInstance *JWTAuthenticationStructure = nil

func InitJWTAuthenticationBackend(userid string) *JWTAuthenticationStructure {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationStructure{
			Key: []byte(userid),
		}
	}
	return authBackendInstance
}

func (backend *JWTAuthenticationStructure) GenerateToken(c map[string]string, tokenDurationMinute int) TokenStructure {
	token := generateClaims(c, tokenDurationMinute)
	return TokenStructure{
		Token: singTokenString(token, backend),
	}
}

func singTokenString(token *jwt.Token, backend *JWTAuthenticationStructure) string {
	tokenString, err := token.SignedString(backend.Key)
	if err != nil {
		return ""
	}
	return tokenString
}

func generateClaims(c map[string]string, tokenDurationMinute int) *jwt.Token {
	c["locale"] = "U2Ugdm9jw6ogY29uc2VndWl1IGNoZWdhciBhdMOpIGFxdWksIHBvZGUgdGVyIGNlcnRlemEgcXVlIGVzdGFtb3MgaW50ZXJlc3NhZG9zIGVtIHNhYmVyIHF1ZW0gdm9jw6ogw6kuIE1hbmRlIHVtIGVtYWlsIHBhcmEgdG9tQGd1cnUuY29tLnZjIGUgdmFtb3MgYmF0ZXIgdW0gcGFwby4gQWJyYcOnb3Mh"
	//noinspection ALL
	claim := claims.Claims{
		c["email"],
		c["id"],
		c["fullname"],
		c["locale"],
		jwt.StandardClaims{
			Audience:  "Guru-APP",
			Issuer:    "https://guruapi.com/authentication",
			ExpiresAt: time.Now().Add(time.Duration(time.Minute * time.Duration(tokenDurationMinute))).Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
}
