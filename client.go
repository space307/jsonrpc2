package jsonrpc2

//go:generate easyjson

import (
	"context"
	"encoding/json"
	"net/http"

	uuid "github.com/gofrs/uuid"
)

const (
	// ErrCodeServer indicates that client could not unmarshal response from server.
	ErrCodeParseError = -32700

	protocolVersionStr         = "2.0"
	contentTypeApplicationJSON = "application/json"
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
	Version string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  interface{}      `json:"params"`
	Id      *json.RawMessage `json:"id"`
}

// Request represents a jsonrpc2 response.
//easyjson:json
type Response struct {
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *json.RawMessage `json:"error"`
	Id      *json.RawMessage `json:"id"`
}

type (
	// Client represents jsonrpc caller interface.
	// It have only one method call which satisfies simple case of jsonrpc2 usage.
	Client interface {
		Call(ctx context.Context, methodName string, params interface{}, result interface{}) error
	}

	client struct {
		url        string
		httpClient HTTPClient
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
	requestID, _ := RequestIDFromContext(ctx)

	if requestID == nil {
		reqUUID, err := uuid.NewV4()
		if err != nil {
			return err
		}

		// We do not use json.Marshal because we know the json representation of a string.
		requestIDVal := json.RawMessage("\"" + reqUUID.String() + "\"")
		requestID = &requestIDVal
	}

	encodedReq, err := encodeRequest(requestID, methodName, params)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Post(ctx, c.url, encodedReq)
	if err != nil {
		return err
	}

	return decodeResponse(resp, result)
}

func encodeRequest(id *json.RawMessage, method string, args interface{}) ([]byte, error) {
	return json.Marshal(&Request{
		Version: protocolVersionStr,
		Method:  method,
		Params:  args,
		Id:      id,
	})
}

func decodeResponse(r []byte, reply interface{}) error {
	var resp Response

	err := json.Unmarshal(r, &resp)
	if err != nil {
		return err
	}

	if resp.Error != nil {
		jsonRpcError := &ResponseError{}
		if err := json.Unmarshal(*resp.Error, jsonRpcError); err != nil {
			return &ResponseError{
				Code:    ErrCodeParseError,
				Message: string(*resp.Error),
			}
		}
		return jsonRpcError
	}

	if resp.Result == nil {
		return nil
	}

	return json.Unmarshal(*resp.Result, reply)
}
