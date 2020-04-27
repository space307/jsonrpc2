package jsonrpc2

import (
	"context"
	"encoding/json"
)

// key is the unique and unexported ctx key for the jsonrpc package.
type key int

// keyRequestID is particular ctx key for request id field.
var keyRequestID key

// WithRequestID returns a new Context that carries id.
func WithRequestID(ctx context.Context, id *json.RawMessage) context.Context {
	return context.WithValue(ctx, keyRequestID, id)
}

// RequestIDFromContext returns the id value stored in ctx, if any.
func RequestIDFromContext(ctx context.Context) (*json.RawMessage, bool) {
	u, ok := ctx.Value(keyRequestID).(*json.RawMessage)
	return u, ok
}
