package jsonrpc2

//go:generate easyjson

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
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

func (e *ResponseError) Error() string {
	return e.Message
}

//easyjson:json
type Request struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      uint64      `json:"id"`
}

//easyjson:json
type Response struct {
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *json.RawMessage `json:"error"`
}

type (
	Client interface {
		Call(methodName string, req interface{}, res interface{}) error
	}

	client struct {
		url        string
		httpClient *http.Client
	}

	clientOption func(c *client)
)

func WithHttpClient(httpClient *http.Client) clientOption {
	return func(c *client) {
		c.httpClient = httpClient
	}
}

func NewClient(rpcEndpointURL string, options ...clientOption) Client {
	c := &client{
		url:        rpcEndpointURL,
		httpClient: http.DefaultClient,
	}

	for _, opt := range options {
		opt(c)
	}

	return c
}

func (c *client) Call(methodName string, req interface{}, resp interface{}) error {
	encodedReq, err := encodeRequest(methodName, req)
	if err != nil {
		return err
	}

	rawResp, err := c.httpClient.Post(c.url, contentTypeApplicationJSON, bytes.NewBuffer(encodedReq))
	if err != nil {
		return err
	}

	return decodeResponse(rawResp.Body, resp)
}

func encodeRequest(method string, args interface{}) ([]byte, error) {
	return json.Marshal(&Request{
		Version: protocolVersionStr,
		Method:  method,
		Params:  args,
		Id:      uint64(rand.Int63()),
	})
}

func decodeResponse(r io.Reader, reply interface{}) error {
	var resp Response
	if err := json.NewDecoder(r).Decode(&resp); err != nil {
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
