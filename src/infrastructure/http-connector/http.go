package http_connector

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type HttpClient struct {
}

func (c *HttpClient) Get(uri string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
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

func (c *HttpClient) Post(uri string, v interface{}, headers map[string]string) ([]byte, error) {

	requestBody, err := json.Marshal(v)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on parsing request body")
	}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(requestBody))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing request")
	}

	reqBody, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return reqBody, nil
	}

	return reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}
