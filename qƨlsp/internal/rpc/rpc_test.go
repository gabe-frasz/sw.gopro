package rpc_test

import (
	"testing"

	"github.com/gabe-frasz/qslsp/internal/rpc"
)

type EncodingExample struct {
	Method string `json:"method"`
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 18\r\n\r\n{\"method\":\"value\"}"
	actual := rpc.EncodeMessage(&EncodingExample{Method: "value"})

	if expected != actual {
		t.Fatalf("Expected:\n%s, got:\n%s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incoming := "Content-Length: 18\r\n\r\n{\"method\":\"value\"}"
	method, content, err := rpc.DecodeMessage([]byte(incoming))
	if err != nil {
		t.Fatal(err)
	}

	contentLength := len(content)

	if contentLength != 18 {
		t.Fatalf("Expected content length 18, got %d", contentLength)
	}

	if method != "value" {
		t.Fatalf("Expected method 'value', got '%s'", method)
	}
}
