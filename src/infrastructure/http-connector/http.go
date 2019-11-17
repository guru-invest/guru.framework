package http_connector

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func Get(uri string) ([]byte,error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing get request")
	} else {
		if res.Status == "200 OK" {
			reqBody, _ := ioutil.ReadAll(res.Body)
			return reqBody, nil
		}
	}
	return []byte{}, errors.Wrap(err, "Error on executing get request")
}

func Post(uri string, v interface{}) ([]byte,error) {

	requestBody, err := json.Marshal(v)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on parsing request body")
	}
	res, err := http.Post(uri, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing request")
	} else {
		if res.Status == "200 OK" {
			reqBody, _ := ioutil.ReadAll(res.Body)
			return reqBody, nil
		}
	}
	return []byte{}, errors.Wrap(err, "Error on executing request")
}