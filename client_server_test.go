package jsonrpc2

import (
	"context"
	"github.com/semrush/zenrpc"
)

//go:generate zenrpc
type TestService struct {
	zenrpc.Service
}

func (TestService) Sum(ctx context.Context, a int, b int) (result int) {
	return a + b
}

func (TestService) Divide(ctx context.Context, a int, b int) (result float32, err error) {
	if b == 0 {
		return 0, zenrpc.NewStringError(-32500, "Could not divide to zero")
	}
	return float32(a)/float32(b), nil
}

