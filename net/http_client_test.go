package net

import (
	"fmt"
	"testing"
)

func TestHttpRequestPost(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json;UTF-8",
	}

	body := []byte(`
		{
			"name": "omp",
			"url": "http://127.0.0.1/omp"
		}`)

	res, _, err := HttpRequest("POST", "http://127.0.0.1:8082/api/v1/menus", headers, body)
	if err != nil {
		t.Error()
	}

	fmt.Println("Response Body: ", string(res))
}

func TestHttpRequestPut(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json;UTF-8",
	}
	body := []byte(`
		{
			"attr": {
				"name": "test1",
				"url": "http://127.0.0.1/testxxx"
			}
		}`)

	res, _, err := HttpRequest("PUT", "http://127.0.0.1:8082/api/v1/menus/6", headers, body)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Response Body: ", string(res))
}

func TestHttpRequestGET(t *testing.T) {
	HttpClientTimeout = 5

	_, _, err := HttpRequest("GET", "http://127.0.0.1:8082/api/v1/menus", nil, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestHttpRequestDelete(t *testing.T) {
	resp, _, err := HttpRequest("DELETE", "http://127.0.0.1:8082/api/v1/menus/6", nil, nil)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(resp))
}
