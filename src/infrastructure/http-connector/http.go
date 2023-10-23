package http_connector

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type HttpClient struct {
	Header  http.Header
	Timeout time.Duration
}

func (c *HttpClient) Get(uri string) (int, []byte, error) {
	client := &http.Client{}
	if c.Timeout != 0 {
		client.Timeout = c.Timeout
	}

	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.Header = c.Header
	res, err := client.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			return http.StatusRequestTimeout, []byte{}, errors.Wrap(err, "Error on executing get request")
		}
		if res == nil {
			return http.StatusInternalServerError, []byte{}, errors.Wrap(err, "Error on executing get request")
		}
		return res.StatusCode, []byte{}, errors.Wrap(err, "Error on executing get request")
	}

	reqBody, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return res.StatusCode, reqBody, nil
	}

	return res.StatusCode, reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}

func (c *HttpClient) Post(uri string, v interface{}) (int, []byte, error) {

	requestBody, err := json.Marshal(v)
	if err != nil {
		return 500, []byte{}, errors.Wrap(err, "Error on parsing request body")
	}
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(requestBody))
	req.Header = c.Header
	res, err := client.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			return http.StatusRequestTimeout, []byte{}, errors.Wrap(err, "Error on executing get request")
		}
		if res == nil {
			return http.StatusInternalServerError, []byte{}, errors.Wrap(err, "Error on executing get request")
		}
		return res.StatusCode, []byte{}, errors.Wrap(err, "Error on executing get request")
	}

	reqBody, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return res.StatusCode, reqBody, nil
	}

	return res.StatusCode, reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}

func (c *HttpClient) Delete(uri string, v interface{}) (int, []byte, error) {

	requestBody, err := json.Marshal(v)
	if err != nil {
		return http.StatusInternalServerError, []byte{}, errors.Wrap(err, "Error on parsing request body")
	}
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, uri, bytes.NewBuffer(requestBody))
	req.Header = c.Header

	res, err := client.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			return http.StatusRequestTimeout, []byte{}, errors.Wrap(err, "Error on executing get request")
		}
		if res == nil {
			return http.StatusInternalServerError, []byte{}, errors.Wrap(err, "Error on executing get request")
		}
		return res.StatusCode, []byte{}, errors.Wrap(err, "Error on executing get request")
	}

	reqBody, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return res.StatusCode, reqBody, nil
	}

	return res.StatusCode, reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}

func (c *HttpClient) Put(uri string, v interface{}) (int, []byte, error) {

	requestBody, err := json.Marshal(v)
	if err != nil {
		return http.StatusInternalServerError, []byte{}, errors.Wrap(err, "Error on parsing request body")
	}
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(requestBody))
	req.Header = c.Header
	res, err := client.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			return http.StatusRequestTimeout, []byte{}, errors.Wrap(err, "Error on executing get request")
		}

		if res == nil {
			return http.StatusInternalServerError, []byte{}, errors.Wrap(err, "Error on executing get request")
		}
		return res.StatusCode, []byte{}, errors.Wrap(err, "Error on executing get request")
	}

	reqBody, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return res.StatusCode, reqBody, nil
	}

	return res.StatusCode, reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}

func (c *HttpClient) Patch(uri string, v interface{}) (int, []byte, error) {

	requestBodyBuffered := &bytes.Buffer{}
	if v != nil {
		requestBody, err := json.Marshal(v)
		if err != nil {
			return http.StatusInternalServerError, []byte{}, errors.Wrap(err, "Error on parsing request body")
		}

		requestBodyBuffered = bytes.NewBuffer(requestBody)
	}
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPatch, uri, requestBodyBuffered)
	req.Header = c.Header

	res, err := client.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			return http.StatusRequestTimeout, []byte{}, errors.Wrap(err, "Error on executing get request")
		}

		if res == nil {
			return http.StatusInternalServerError, []byte{}, errors.Wrap(err, "Error on executing get request")
		}
		return res.StatusCode, []byte{}, errors.Wrap(err, "Error on executing get request")
	}

	defer res.Body.Close()

	reqBody, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return res.StatusCode, reqBody, nil
	}

	return res.StatusCode, reqBody, errors.Wrap(errors.New(strconv.Itoa(res.StatusCode)), res.Status)
}
