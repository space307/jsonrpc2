// Code generated by zenrpc; DO NOT EDIT.

package jsonrpc2

import (
	"context"
	"encoding/json"

	"github.com/semrush/zenrpc"
	"github.com/semrush/zenrpc/smd"
)

var RPC = struct {
	TestService struct{ Sum, Divide string }
}{
	TestService: struct{ Sum, Divide string }{
		Sum:    "sum",
		Divide: "divide",
	},
}

func (TestService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Description: ``,
		Methods: map[string]smd.Service{
			"Sum": {
				Description: ``,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Integer,
				},
			},
			"Divide": {
				Description: ``,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Float,
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s TestService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.TestService.Sum:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		resp.Set(s.Sum(ctx, args.A, args.B))

	case RPC.TestService.Divide:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
			}
		}

		resp.Set(s.Divide(ctx, args.A, args.B))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}
