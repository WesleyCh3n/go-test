package main

import "testing"

func TestHello(t *testing.T) {
  assertCorrectMessage := func(t testing.TB, got, want string){
    t.Helper()
    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  }

  t.Run("Saying hello to people", func(t *testing.T){
    got := Hello("Wesley", "")
    want := "Hello Wesley"
    assertCorrectMessage(t, got, want)
  })
  t.Run("Say 'Hello World' when empty string is supply", func(t *testing.T){
    got := Hello("", "")
    want := "Hello World"
    assertCorrectMessage(t, got, want)
  })
  t.Run("Saying hello to people in Spanish", func(t *testing.T){
    got := Hello("Wesley", "Spanish")
    want := "Hola Wesley"
    assertCorrectMessage(t, got, want)
  })
  t.Run("Saying hello to people in Franch", func(t *testing.T){
    got := Hello("Wesley", "Franch")
    want := "Bonjour Wesley"
    assertCorrectMessage(t, got, want)
  })
}
