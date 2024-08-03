package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// A buffer is a temporary storage area used to hold data while it is being moved from one place to another. 
	// Think of it as a waiting room where data can sit until it's ready to be processed or transferred. 
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}