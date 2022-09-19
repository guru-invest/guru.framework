package http_connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type HttpClient struct {
	Header  http.Header
	Timeout time.Duration
}

func (c *HttpClient) Get(uri string) ([]byte, error) {
	client := &http.Client{}
	if c.Timeout != 0 {
		client.Timeout = c.Timeout
	}
	
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.Header = c.Header
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing get request")
	}

	reqBody, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return reqBody, nil
	}

	return reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}

func (c *HttpClient) Post(uri string, v interface{}) ([]byte, error) {

	requestBody, err := json.Marshal(v)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on parsing request body")
	}
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(requestBody))
	req.Header = c.Header
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

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}

func (c *HttpClient) Delete(uri string, v interface{}) ([]byte, error) {

	requestBody, err := json.Marshal(v)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on parsing request body")
	}
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, uri, bytes.NewBuffer(requestBody))
	req.Header = c.Header

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

func (c *HttpClient) Put(uri string, v interface{}) ([]byte, error) {

	requestBody, err := json.Marshal(v)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on parsing request body")
	}
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(requestBody))
	req.Header = c.Header

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

func (c *HttpClient) Patch(uri string, v interface{}) ([]byte, error) {

	requestBodyBuffered := &bytes.Buffer{}
	if v != nil {
		requestBody, err := json.Marshal(v)
		if err != nil {
			return []byte{}, errors.Wrap(err, "Error on parsing request body")
		}

		requestBodyBuffered = bytes.NewBuffer(requestBody)
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, uri, requestBodyBuffered)
	req.Header = c.Header
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing new request")
	}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error on executing request")
	}

	defer res.Body.Close()

	reqBody, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return reqBody, nil
	}

	return reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}
