package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type BaseClient struct {
	BaseUrl    string
	HTTPClient *http.Client
}

func NewClient(baseUrl string) *BaseClient {
	return &BaseClient{
		BaseUrl: baseUrl,
		HTTPClient: &http.Client{
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		},
	}
}

func (c *BaseClient) post(ctx context.Context, apiPath string, payload interface{}, response interface{}) error {
	fullUrl := c.BaseUrl + apiPath
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, fullUrl, bytes.NewBuffer(bytePayload))
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(request)
	if err != nil {
		return err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseBody, response)
	if err != nil {
		return err
	}

	return nil
}

func (c *BaseClient) get(ctx context.Context, apiPath string, query url.Values, response interface{}) error {
	fullUrl := c.BaseUrl + apiPath
	if query != nil {
		url, err := url.Parse(fullUrl)
		if err != nil {
			return err
		}
		url.RawQuery = query.Encode()
		fullUrl = url.String()
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fullUrl, bytes.NewBuffer(make([]byte, 0)))
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(request)
	if err != nil {
		return err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseBody, response)
	if err != nil {
		return err
	}

	return nil
}
