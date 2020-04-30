package jsonrpc2

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/semrush/zenrpc"
	"github.com/stretchr/testify/assert"
)

var serverURL string

func TestMain(m *testing.M) {
	mux := http.NewServeMux()
	server := zenrpc.NewServer(zenrpc.Options{})
	server.Register("", TestService{})

	mux.Handle("/rpc", server)
	srv := httptest.NewServer(mux)
	serverURL = srv.URL

	os.Exit(m.Run())
}

func Test_client_Call_Sum_Ok(t *testing.T) {
	client := NewClient(fmt.Sprintf("%s/rpc", serverURL))
	var val int
	err := client.Call(context.Background(), "Sum", struct {
		A, B int
	}{1, 2}, &val)
	assert.NoError(t, err)
	assert.Equal(t, 3, val)
}

func Test_client_Call_Div_Ok(t *testing.T) {
	client := NewClient(fmt.Sprintf("%s/rpc", serverURL))
	var val int
	err := client.Call(context.Background(), "Divide", struct {
		A, B int
	}{4, 2}, &val)
	assert.NoError(t, err)
	assert.Equal(t, 2, val)
}

func Test_client_Call_Div_Fail(t *testing.T) {
	client := NewClient(fmt.Sprintf("%s/rpc", serverURL))
	var val int
	err := client.Call(context.Background(), "Divide", struct {
		A, B int
	}{4, 0}, &val)
	assert.Error(t, err)
	assert.Zero(t, val)
}
