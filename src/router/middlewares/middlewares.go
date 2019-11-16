package middlewares

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guru-invest/guru.framework/src/security/claims"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func Interceptor(next gin.HandlerFunc, parameter string, anonymous func(string)) (gin.HandlerFunc){
	return gin.HandlerFunc(func(c *gin.Context){
		anonymous(parameter)
		next(c)
	})
}

func RequireTokenAuthentication(next gin.HandlerFunc) gin.HandlerFunc{
	return gin.HandlerFunc(func(c *gin.Context) {
		token := extractTokenFromRequest(c.Writer, c.Request)
		if token != nil{
			if token.Valid {
				next(c)
			}else{
				c.AbortWithStatusJSON(400, "Invalid Token. You are no authorized to perform this action.")
			}
		} else {
			c.AbortWithStatusJSON(200, "You are not authorized to perform this action.")
		}
	})
}

func extractTokenFromRequest(w http.ResponseWriter, r *http.Request) *jwt.Token{
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ","",1)
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
	}else{
		return nil
	}
}


func includeFormData(r *http.Request, claims *claims.Claims){
	r.Header.Add("userId", claims.DocumentNumber)
	r.Header.Add("email", claims.Email)
}