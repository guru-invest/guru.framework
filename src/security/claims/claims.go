package claims


import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string `json:"email"`
	DocumentNumber string `json:"document"`
	Fullname string  `json:"fullname"`
	Message string `json:"locale"`
	jwt.StandardClaims
}