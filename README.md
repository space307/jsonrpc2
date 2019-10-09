# jsonrpc2

It is simple jsonrpc2 client.

Usage:
```go
	client := NewClient(serverRpcUrl)
	var val int
	err := client.Call(context.Background(), "Divide", struct {
		A,B int
	}{4,2}, &val)
```