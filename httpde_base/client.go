package httpde_base

import (
	"net/http"

	"github.com/koalaone/mock/base"
)

type HttpClient struct {
	base.Report
	client *http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
	}
}

func (hc *HttpClient) Do(req *http.Request) (*http.Response, error) {
	if hc.client == nil {
		hc.client = &http.Client{}
	}

	resp, err := hc.client.Do(req)
	if err != nil {
		hc.LogError("HttpClient Do error:%v", err.Error())
		return nil, err
	}

	return resp, nil
}

func (hc *HttpClient) Client() *http.Client {
	if hc.client == nil {
		hc.client = &http.Client{}
	}

	return hc.client
}
