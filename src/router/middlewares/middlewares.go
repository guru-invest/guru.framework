package middlewares

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/guru-invest/guru.framework/src/helpers/messages"
	"github.com/guru-invest/guru.framework/src/router/returns"

	"github.com/dgrijalva/jwt-go"
	"github.com/guru-invest/guru.framework/src/security/claims"
)

func Interceptor(next http.HandlerFunc, parameter string, anonymous func(string)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		anonymous(parameter)
		next(w, r)
	})
}

func RequireTokenAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractTokenFromRequest(r)
		if token == nil {
			w.WriteHeader(messages.HttpCode.Unauthorized)
			resp, _ := json.Marshal(returns.UnauthorizedError())
			_, _ = w.Write(resp)
			return
		}

		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(messages.HttpCode.Unauthorized)
			resp, _ := json.Marshal(returns.UnauthorizedError())
			_, _ = w.Write(resp)
			return
		}
	})
}

func extractTokenFromRequest(r *http.Request) *jwt.Token {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	if tokenString != "" {
		token, err := jwt.ParseWithClaims(tokenString, &claims.Claims{}, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			sEnc := b64.StdEncoding.EncodeToString([]byte(token.Claims.(*claims.Claims).DocumentNumber))
			return []byte(sEnc), nil
		})

		if claim, ok := token.Claims.(*claims.Claims); ok && token.Valid {
			includeFormData(r, claim)
			token.Valid = true
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("Couldn't handle this token: That's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println("Couldn't handle this token: Timing is everything")
			} else {
				fmt.Println("Couldn't handle this token:", err)
			}
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
		return token
	} else {
		return nil
	}
}

func includeFormData(r *http.Request, claims *claims.Claims) {
	r.Header.Add("userId", claims.DocumentNumber)
	r.Header.Add("email", claims.Email)
}
