package extractor

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/guru-invest/guru.framework/src/helpers/messages"
	"github.com/guru-invest/guru.framework/src/router/returns"
)

func Extract(v interface{}, w http.ResponseWriter, r *http.Request) []byte {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(messages.HttpCode.BadRequest)
		resp, _ := json.Marshal(returns.InvalidFormatError(""))
		_, _ = w.Write(resp)
	}
	return reqBody
}
