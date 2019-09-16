package jsonrpc2

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
)

type (
	// HTTPClient is a convenient interface for http worker.
	// It allows to use a custom http transport level implementation.
	HTTPClient interface {
		Post(ctx context.Context, url string, body []byte) ([]byte, error)
	}

	httpClient struct {
		c *http.Client
	}
)

// Post makes http request with method POST to url with body.
func (h *httpClient) Post(ctx context.Context, url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentTypeApplicationJSON)

	resp, err := h.c.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// NewHttpClient returns wiring of standard http.Client.
func NewHttpClient(c *http.Client) HTTPClient {
	return &httpClient{
		c: c,
	}
}
