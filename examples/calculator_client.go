package examples

import (
	"fmt"

	"github.com/cv21/jsonrpc2"
)

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Result int `json:"result"`
}

func Example() {

	// You can provide options along with rpc endpoint url parameter.
	jr2 := jsonrpc2.NewClient("http://localhost/rpc")

	resp := &SumResponse{}
	err := jr2.Call("Hello", &SumRequest{1, 2}, resp)
	fmt.Println(err)
}
