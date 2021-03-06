package security

import (
	b64 "encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/guru-invest/guru.framework/src/security/auth"
	"github.com/guru-invest/guru.framework/src/security/claims"
)

func TestTokenGeneration(t *testing.T) {
	DocumentNumber := "5kCTnuVt"
	Email := "tom@guru.com.vc"
	Fullname := "Tom teste de geração de token"
	authBackend := auth.InitJWTAuthenticationBackend("")
	m := make(map[string]string)
	m["id"] = DocumentNumber
	m["email"] = Email
	m["fullname"] = Fullname
	token := authBackend.GenerateToken(m, 260000)
	if token.Token == "" {
		fmt.Print("Null token")
		t.Fail()
	} else {
		fmt.Print(token.Token + "\n")
	}
}

func TestTokenValidation(t *testing.T) {
	tokenString := "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImRvY3VtZW50IjoibURrcEQzM2ciLCJmdWxsbmFtZSI6Ikd1cnUgYXBwIiwibG9jYWxlIjoiVTJVZ2RtOWp3Nm9nWTI5dWMyVm5kV2wxSUdOb1pXZGhjaUJoZE1PcElHRnhkV2tzSUhCdlpHVWdkR1Z5SUdObGNuUmxlbUVnY1hWbElHVnpkR0Z0YjNNZ2FXNTBaWEpsYzNOaFpHOXpJR1Z0SUhOaFltVnlJSEYxWlcwZ2RtOWp3Nm9ndzZrdUlFMWhibVJsSUhWdElHVnRZV2xzSUhCaGNtRWdkRzl0UUdkMWNuVXVZMjl0TG5aaklHVWdkbUZ0YjNNZ1ltRjBaWElnZFcwZ2NHRndieTRnUVdKeVljT25iM01oIiwiZXhwIjoxNTY3NzA2MDU5fQ.DWhcVVJ5eQ6Sj4EwDLBUhrGQ8Hn-5mTQD3596jHHQQVVVpepYojSRYOFMq1b8LMrFkLNl77IRcFdQFLby_2q7w"
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	if tokenString != "" {
		token, err := jwt.ParseWithClaims(tokenString, &claims.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			sEnc := b64.StdEncoding.EncodeToString([]byte(token.Claims.(*claims.Claims).DocumentNumber))
			return []byte(sEnc), nil
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
