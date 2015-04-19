package gads

import (
	"crypto/rand"
	"golang.org/x/net/context"
	"testing"
)

func rand_str(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func rand_word(str_size int) string {
	alphanum := "abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func testAuthSetup(t *testing.T) Auth {
	config, err := NewCredentials(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	config.Auth.Testing = t
	return config.Auth
}
