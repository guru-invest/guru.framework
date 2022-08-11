package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	http_connector "github.com/guru-invest/guru.framework/src/infrastructure/http-connector"
	"github.com/guru-invest/guru.framework/src/models"
)

type GenialAuthentication struct {
	URLBase string
}

type GenialTokenResponse struct {
	Token string `json:"token,omitempty"`
}

func (t GenialAuthentication) AuthorizeWithGinContext(c *gin.Context) {
	access, err := t.getToken()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	models.EXTERNALAUTHENTICATION.GenialToken = access.Token
	c.Next()
}

func (t GenialAuthentication) getToken() (*GenialTokenResponse, error) {
	token := GenialTokenResponse{}
	uri := t.URLBase + "/genialtoken"
	client := http_connector.HttpClient{
		Header: http.Header{
			"Content-type": []string{"application/json"},
		},
	}

	getToken, err := client.Get(uri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(getToken, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
