package auth_test

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/guru-invest/guru.framework/src/security/auth"
	"github.com/guru-invest/guru.framework/src/security/claims"
)

func createValidToken() auth.TokenStructure {
	DocumentNumber := "5kCTnuVt"
	Email := "tom@guru.com.vc"
	Fullname := "Tom teste de geração de token"
	authBackend := auth.InitJWTAuthenticationBackend(DocumentNumber)
	m := make(map[string]string)
	m["id"] = DocumentNumber
	m["email"] = Email
	m["fullname"] = Fullname
	return authBackend.GenerateToken(m, 260000)
}

func createInvalidToken() auth.TokenStructure {
	DocumentNumber := "5kCTnuVt"
	Email := "tom@guru.com.vc"
	Fullname := "Tom teste de geração de token"
	authBackend := auth.InitJWTAuthenticationBackend("InvalidToken")
	m := make(map[string]string)
	m["id"] = DocumentNumber
	m["email"] = Email
	m["fullname"] = Fullname
	return authBackend.GenerateToken(m, 260000)
}

func TestTokenGeneration(t *testing.T) {

	token := createValidToken()
	if token.Token == "" {
		fmt.Print("Null token")
		t.Fail()
	} else {
		fmt.Print(token.Token + "\n")
	}
}

func TestTokenValidation(t *testing.T) {

	tokenString := createValidToken().Token
	if tokenString != "" {
		token, err := jwt.ParseWithClaims(tokenString, &claims.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(token.Claims.(*claims.Claims).DocumentNumber), nil
		})

		if _, ok := token.Claims.(*claims.Claims); ok && token.Valid {
			fmt.Print("Token is valid! Welcome! \n")
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("Couldn't handle this token: That's not even a token")
				t.Fail()
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println("Couldn't handle this token: Timing is everything")
				t.Fail()
			} else {
				fmt.Println("Couldn't handle this token:", err)
				t.Fail()
			}
		} else {
			fmt.Println("Couldn't handle this token:", err)
			t.Fail()
		}
	}
}
