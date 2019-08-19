package security

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/guru-invest/guru.framework/security/auth"
	"github.com/guru-invest/guru.framework/security/claims"
	"strings"
	"testing"
)
func TestTokenGeneration(t *testing.T) {
	DocumentNumber := "36163719883"
	Email := "tom@guru.com.vc"
	Fullname := "Tom teste de geração de token"
	sEnc := b64.StdEncoding.EncodeToString([]byte(DocumentNumber))
	authBackend := auth.InitJWTAuthenticationBackend(sEnc)
	m := make(map[string]string)
	m["id"] = DocumentNumber
	m["email"] = Email
	m["fullname"] = Fullname
	token := authBackend.GenerateToken(m)
	if token.Token == "" {
		fmt.Print("Null token")
		t.Fail()
	}else{
		fmt.Print(token.Token + "\n")
	}
}

func TestTokenValidation(t *testing.T) {
	tokenString := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRvbUBndXJ1LmNvbS52YyIsImRvY3VtZW50IjoiMzYxNjM3MTk4ODMiLCJmdWxsbmFtZSI6IlRvbSB0ZXN0ZSBkZSBnZXJhw6fDo28gZGUgdG9rZW4iLCJsb2NhbGUiOiJVMlVnZG05anc2b2dZMjl1YzJWbmRXbDFJR05vWldkaGNpQmhkTU9wSUdGeGRXa3NJSEJ2WkdVZ2RHVnlJR05sY25SbGVtRWdjWFZsSUdWemRHRnRiM01nYVc1MFpYSmxjM05oWkc5eklHVnRJSE5oWW1WeUlIRjFaVzBnZG05anc2b2d3Nmt1SUUxaGJtUmxJSFZ0SUdWdFlXbHNJSEJoY21FZ2RHOXRRR2QxY25VdVkyOXRMblpqSUdVZ2RtRnRiM01nWW1GMFpYSWdkVzBnY0dGd2J5NGdRV0p5WWNPbmIzTWgiLCJleHAiOjE1NjU5NzI2MzR9.eSFpSnaumQvON3ZUcdjytrDCvUmm0l2hUst4HSjqmmA"
	tokenString = strings.Replace(tokenString, "Bearer ","",1)
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