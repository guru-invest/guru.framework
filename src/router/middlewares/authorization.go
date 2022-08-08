package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guru-invest/guru.framework/src/security/auth"
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

			authenticatedUser := auth.AuthenticatedUserModel{}
			err := authenticatedUser.Validate(token, t.SigningKey, t.Issuer)
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
