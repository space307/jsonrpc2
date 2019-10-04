package jsonrpc2

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
)

var defaultCodec = NewHTTPCodec()

func NewHTTPRoundTripper(url string) RoundTripper {
	return &httpRoundTripper{
		url:    url,
		client: http.DefaultClient,
		codec:  defaultCodec,
	}
}

type httpRoundTripper struct {
	url    string
	client *http.Client
	codec  Codec
}

func (h *httpRoundTripper) RoundTrip(ctx context.Context, request *Request) (*Response, error) {
	encodedRequest, err := h.codec.EncodeRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, h.url, bytes.NewReader(encodedRequest))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentTypeApplicationJSON)

	resp, err := h.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return h.codec.DecodeResponse(ctx, body)
}