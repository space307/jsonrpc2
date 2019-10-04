package jsonrpc2

import (
	"context"
	"encoding/json"
)

func NewHTTPCodec() Codec {
	return &httpCodec{}
}

type httpCodec struct {}

func (j *httpCodec) EncodeRequest(ctx context.Context, request *Request) ([]byte, error) {
	return json.Marshal(request)
}

func (j *httpCodec) DecodeResponse(ctx context.Context, response []byte) (*Response, error) {
	resp := &Response{}
	return resp, json.Unmarshal(response, resp)
}