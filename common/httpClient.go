package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClient interface {
	CallApi(method, url, resource, id string) ([]byte, error)
}

type httpClient struct {
	client *http.Client
}

func NewHttpClient() *httpClient {
	newClient := &httpClient{
		client: &http.Client{},
	}
	return newClient
}

// CallApi is the single method exposed for all http requests with http.Client
func (h *httpClient) CallApi(method, url, resource, id string) ([]byte, error) {
	var response *http.Response
	var err error
	uri := h.buildURI(url, resource, id)
	switch method {
	case http.MethodGet:
		response, err = h.callGET(uri)
	default:
		return nil, fmt.Errorf("bad method")
	}
	if err != nil {

	}
	return h.processResponse(response)
}

func (h *httpClient) buildURI(url, resource, id string) string {
	httpPlaceholder := "http://%s/%s/%s"
	return fmt.Sprintf(httpPlaceholder, url, resource, id)
}

func (h *httpClient) callGET(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &http.Response{}, fmt.Errorf("get request failed: %s", err)
	}

	response, err := h.client.Do(request)
	if err != nil {
		return &http.Response{}, fmt.Errorf("get request failed: %s", err)
	}
	return response, nil
}

func (h *httpClient) processResponse(response *http.Response) ([]byte, error) {
	if response == nil || response.Body == nil {
		return []byte{}, fmt.Errorf("empty or nil response")
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read response")
	}
	if len(body) > 0 {
		return body, nil
	}
	return []byte{}, fmt.Errorf("empty or nil response")
}
