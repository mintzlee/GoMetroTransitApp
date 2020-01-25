package main

import "testing"

func TestHello(t *testing.T) {
	var returnedString = hello()

	if returnedString != "hello, world\n" {
		t.Errorf("Expected \"Hello, world\\n\", but received: %s", returnedString)
	}
}
