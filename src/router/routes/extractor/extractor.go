package extractor

import (
	"github.com/gin-gonic/gin"
	"github.com/guru-invest/guru.framework/src/helpers/messages"
	"io/ioutil"
)

func Extract(v interface{}, c *gin.Context) []byte{
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		c.AbortWithStatusJSON(messages.HttpCode.BadRequest, "")
	}
	return reqBody
}
