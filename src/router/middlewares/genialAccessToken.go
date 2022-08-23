package middlewares

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/guru-invest/guru.framework/src/models"
)

type GenialAccessToken struct {
	URL      string `json:"url"`
	Port     string `json:"port"`
	Database int    `json:"database"`
	Auth     string `json:"auth"`
	Key      string `json:"key"`
}

func (t GenialAccessToken) clientCacheConnector() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", t.URL, t.Port),
		Password: t.Auth,
		DB:       t.Database,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

//returna ja preenchido o objt domain GENIALACCESSTOKEN
func (t GenialAccessToken) GetAccessToken(c *gin.Context) {
	client, err := t.clientCacheConnector()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}
	defer client.Close()

	get, err := client.Get(t.Key).Result()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}
	models.GENIALACCESSTOKEN = models.GenialAccessToken{
		Token: get,
	}
	c.Next()
}
