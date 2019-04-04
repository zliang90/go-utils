package http

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	ClientTimeout uint = 30
)

func Request(method, url string, headers map[string]string, body []byte) (resByte []byte, statusCode int, err error) {
	if method == "" {
		method = "GET"
	}

	if url == "" {
		err = errors.New("http request url can not be empty")
		return
	}

	var (
		req  *http.Request
		resp *http.Response
	)

	// http request
	req, err = http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return
	}

	// request headers
	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	client := http.Client{
		Timeout: time.Duration(ClientTimeout) * time.Second,
	}

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	resByte, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	return resByte, resp.StatusCode, nil
}
