package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/guru-invest/guru.framework/src/models"
)

type Authorization struct {
	URLBase    string
	SigningKey string
	Issuer     string
}

func (t Authorization) AuthorizeWithGinContext(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token != "" {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", t.URLBase+"/authorize", nil)
		req.Header.Add("Authorization", token)
		res, err := client.Do(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
			return
		}

		if res.StatusCode == http.StatusOK {
			err := t.validate(token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
				return
			}
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Authorization Not Found"})
}

func (t *Authorization) validate(tokenString string) error {
	authenticatedUserClaims := models.AuthenticatedUser{}
	tokenString = strings.ReplaceAll(strings.ReplaceAll(tokenString, "bearer ", ""), "Bearer ", "")

	token, err := jwt.ParseWithClaims(tokenString, &authenticatedUserClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.SigningKey), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*models.AuthenticatedUser); ok && token.Valid {
		if claims.StandardClaims.VerifyIssuer(t.Issuer, true) {
			models.AUTHENTICATEDUSER = *claims
			models.AUTHENTICATEDUSER.JwtToken = tokenString
			return nil
		}
	}
	return err
}
