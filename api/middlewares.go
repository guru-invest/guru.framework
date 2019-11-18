package api

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Extract(v interface{}, c *gin.Context) []byte {
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		Error400(err, c)
	}
	return reqBody
}

func Error404(err error, c *gin.Context) {
	msg := make(map[string]interface{})
	msg["error"] = "Not found - " + err.Error()
	c.AbortWithStatusJSON(404, msg)
}

func Error400(err error, c *gin.Context) {
	msg := make(map[string]interface{})
	msg["error"] = "Invalid Format - " + err.Error()
	c.AbortWithStatusJSON(400, msg)
}

func Error500(err error, c *gin.Context) {
	msg := make(map[string]interface{})
	msg["error"] = err.Error()
	c.AbortWithStatusJSON(500, msg)
}
