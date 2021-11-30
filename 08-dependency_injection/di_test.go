package main

import (
  "bytes"
  "testing"
)

func TestGreet(t *testing.T) {
  buffer := bytes.Buffer{}
  Greet(&buffer, "Wesley")

  got := buffer.String()
  want := "Hello Wesley"

  if got != want {
    t.Errorf("got %q but want %q", got, want)
  }
}
