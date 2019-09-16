package jsonrpc2

import (
	"context"
	"io"
	"net/http"
)

type (
	// HTTPClient is a convenient interface for http worker.
	// It allows to use a custom http transport level implementation.
	HTTPClient interface {
		Post(ctx context.Context, url string, req io.Reader) (*http.Response, error)
	}

	httpClient struct {
		c *http.Client
	}
)

// Post makes http request with method POST to url with body.
func (h *httpClient) Post(ctx context.Context, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentTypeApplicationJSON)

	return h.c.Do(req.WithContext(ctx))
}

// NewHttpClient returns wiring of standard http.Client.
func NewHttpClient(c *http.Client) HTTPClient {
	return &httpClient{
		c: c,
	}
}
