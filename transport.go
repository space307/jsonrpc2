package jsonrpc2

import "context"

type (
	// RoundTripper allows to use a custom http transport level implementation.
	RoundTripper interface {
		RoundTrip(ctx context.Context, request *Request) (*Response, error)
	}

	Codec interface {
		EncodeRequest(ctx context.Context, request *Request) ([]byte, error)
		DecodeResponse(ctx context.Context, response []byte) (*Response, error)
	}
)