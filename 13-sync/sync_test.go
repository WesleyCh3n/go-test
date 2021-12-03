package main

import (
  "sync"
  "testing"
)

func TestCounter(t *testing.T) {

  t.Run("Inc counter 3 time and result is 3", func(t *testing.T) {
    counter := Counter{}
    counter.Inc()
    counter.Inc()
    counter.Inc()

    assertCounter(t, &counter, 3)
  })

  t.Run("Run Concurrently", func(t *testing.T) {
    wantedCount := 1000
    counter := Counter{}

    var wg sync.WaitGroup
    wg.Add(wantedCount)

    go func ()  {
      for i := 0; i < wantedCount; i++ {
        counter.Inc()
        wg.Done()
      }
    }()
    wg.Wait()

    assertCounter(t, &counter, wantedCount)
  })
}

func assertCounter(t testing.TB, got *Counter, want int) {
  t.Helper()

  if got.Value() != want {
    t.Errorf("got %d but want %d", got.Value(), want)
  }

}
