package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	http_connector "github.com/guru-invest/guru.framework/src/infrastructure/http-connector"
	"github.com/guru-invest/guru.framework/src/models"
)

type B3Authentication struct {
	URLBase string
}

type B3TokenResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	EXTExpiresIn int64  `json:"ext_expires_in"`
	AccessToken  string `json:"token"`
}

func (t B3Authentication) AuthorizeWithGinContext(c *gin.Context) {
	access, err := t.getToken()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	models.EXTERNALAUTHENTICATION.B3Token = access.AccessToken
	c.Next()
}

func (t B3Authentication) getToken() (*B3TokenResponse, error) {
	investorAuth := B3TokenResponse{}
	uri := t.URLBase + "/b3token"
	client := http_connector.HttpClient{
		Header: http.Header{
			"Content-type": []string{"application/json"},
		},
	}

	getToken, err := client.Get(uri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(getToken, &investorAuth)
	if err != nil {
		return nil, err
	}

	return &investorAuth, nil
}
