package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type BaseClient struct {
	BaseUrl    string
	HTTPClient *http.Client
}

func (c *BaseClient) post(ctx context.Context, apiPath string, payload interface{}, response interface{}) error {
	fullUrl := c.BaseUrl + apiPath
	return WrappedPost(ctx, c.HTTPClient, fullUrl, payload, response)
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
