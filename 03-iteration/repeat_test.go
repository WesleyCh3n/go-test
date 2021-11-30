package iteration

import (
  "fmt"
  "testing"
)

func TestRepeat(t *testing.T) {
  result := Repeat("a", 8)
  expected := "aaaaaaaa"

  if result != expected {
    t.Errorf("expected %q but got %q", expected, result)
  }
}

func ExampleRepeat() {
  result := Repeat("a", 8)
  fmt.Println(result)
}

func BenchmarkRepeat(b *testing.B) {
  for i:= 0; i < b.N; i++ {
    Repeat("a", 2)
  }
}
