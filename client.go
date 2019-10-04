package jsonrpc2

//go:generate easyjson

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
)

const (
	// ErrCodeServer indicates that client could not unmarshal response from server.
	ErrCodeParseError = -32700

	protocolVersionStr         = "2.0"
	contentTypeApplicationJSON = "application/json"
)

var (
	ErrCouldNotEncodeRequest = errors.New("could not encode request")
	ErrCouldNotDecodeResponse = errors.New("could not decode response")
)

// ResponseError is a struct which represents a typical jsonrpc2 error according to specification.
//easyjson:json
type ResponseError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Error returns an error message of ResponseError.
func (e *ResponseError) Error() string {
	return e.Message
}

// Request represents a jsonrpc2 request.
//easyjson:json
type Request struct {
	Id      uint64      `json:"id"`
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

// Request represents a jsonrpc2 response.
//easyjson:json
type Response struct {
	Id string `json:"id"`
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *ResponseError `json:"error"`
}

type (
	// Client represents jsonrpc caller interface.
	// It have only one method call which satisfies simple case of jsonrpc2 usage.
	Client interface {
		Call(ctx context.Context, methodName string, params interface{}) error
	}

	client struct {
		url        string
		httpClient HTTPClient

		roundTripper RoundTripper
	}

	clientOption func(c *client)
)

//  WithHttpClient is an option which sets http client implementation for jsonrpc2 client.
func WithHttpClient(httpClient HTTPClient) clientOption {
	return func(c *client) {
		c.httpClient = httpClient
	}
}

// NewClient returns jsonrpc2 client.
func NewClient(rpcEndpointURL string, options ...clientOption) Client {
	c := &client{
		url:        rpcEndpointURL,
		httpClient: NewHttpClient(http.DefaultClient),
	}

	for _, opt := range options {
		opt(c)
	}

	return c
}

// Call makes and does jsonrpc2 request.
func (c *client) Call(ctx context.Context, methodName string, params interface{}, result interface{}) error {
	response, err := c.roundTripper.RoundTrip(ctx, &Request{
		Version: protocolVersionStr,
		Method:  methodName,
		Params:  params,
		Id:      uint64(rand.Int63()),
	})
	if err != nil {
		return err
	}

	if response.Error != nil {
		return response.Error
	}

	if response.Result == nil {
		return nil
	}

	return json.Unmarshal(*response.Result, result)
}