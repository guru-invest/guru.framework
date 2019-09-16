package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func Extract(v interface{}, c *gin.Context) (interface{}){
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		Error400(err, c)
	}
	err = json.Unmarshal(reqBody, &v)
	if err != nil{
		Error400(err, c)
	}
	return v
}

func Error404(err error, c *gin.Context){
	msg := make(map[string]interface{})
	msg["error"] = "Not found - " + err.Error()
	c.AbortWithStatusJSON(400, msg)
}

func Error400(err error, c *gin.Context){
	msg := make(map[string]interface{})
	msg["error"] = "Invalid Format - " + err.Error()
	c.AbortWithStatusJSON(400, msg)
}

func Error500(err error, c *gin.Context){
	msg := make(map[string]interface{})
	msg["error"] = err.Error()
	c.AbortWithStatusJSON(400, msg)
}