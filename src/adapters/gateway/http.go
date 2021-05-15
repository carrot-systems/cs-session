package gateway

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func buildClient() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	return client
}

func buildRequest(method string, url string, body io.Reader, header map[string]string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		request = nil
		return nil, err
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("Accept", "*/*")
	for key, val := range header {
		request.Header.Add(key, val)
	}

	return request, nil
}

func httpGET(url string, body interface{}, header map[string]string) (*http.Response, error) {
	client := buildClient()

	marshal, err := json.Marshal(body)
	bReader := bytes.NewBuffer(marshal)

	request, err := buildRequest("GET", url, bReader, header)

	if err != nil {
		client = nil
		request = nil
		return nil, err
	}

	response, err := client.Do(request)
	request = nil
	client = nil

	return response, err
}
