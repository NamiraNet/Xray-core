package kcp_test

import (
	"testing"

	"github.com/NamiraNet/xray-core/common"
	. "github.com/NamiraNet/xray-core/transport/internet/kcp"
	"github.com/google/go-cmp/cmp"
)

func TestSimpleAuthenticator(t *testing.T) {
	cache := make([]byte, 512)

	payload := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	auth := NewSimpleAuthenticator()
	b := auth.Seal(cache[:0], nil, payload, nil)
	c, err := auth.Open(cache[:0], nil, b, nil)
	common.Must(err)
	if r := cmp.Diff(c, payload); r != "" {
		t.Error(r)
	}
}

func TestSimpleAuthenticator2(t *testing.T) {
	cache := make([]byte, 512)

	payload := []byte{'a', 'b'}

	auth := NewSimpleAuthenticator()
	b := auth.Seal(cache[:0], nil, payload, nil)
	c, err := auth.Open(cache[:0], nil, b, nil)
	common.Must(err)
	if r := cmp.Diff(c, payload); r != "" {
		t.Error(r)
	}
}
