package main

import (
  "context"
  "errors"
  "net/http"
  "net/http/httptest"
  "testing"
  "time"
)

type SpyStore struct {
  response string
  t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
  data := make(chan string, 1)

  go func() {
    var result string

    // simulate slow process by appending char by char
    for _, c := range s.response {
      select {
      case <-ctx.Done():
        s.t.Log("spy store got canceled")
        return
      default:
        time.Sleep(100 * time.Millisecond)
        result += string(c)
      }
    }
    data <- result
  }()

  select {
  case <-ctx.Done():
    return "", ctx.Err()
  case res := <-data:
    return res, nil
  }
}

type SpyResponseWrite struct {
  written bool
}

func (s *SpyResponseWrite) Header() http.Header {
  s.written = true
  return nil
}

func (s *SpyResponseWrite) Write([]byte) (int, error) {
  s.written = true
  return 0, errors.New("not implemented")
}

func (s *SpyResponseWrite) WriteHeader(statusCode int) {
  s.written = true
}

func TestServer(t *testing.T) {
  data := "hello world"

  t.Run("return data from store", func(t *testing.T) {
    store := SpyStore{response: data, t: t}
    svr := Server(&store)

    request := httptest.NewRequest(http.MethodGet, "/", nil)
    response := httptest.NewRecorder()

    svr.ServeHTTP(response, request)

    if response.Body.String() != data {
      t.Errorf("got %q but want %q", response.Body.String(), data)
    }
  })

  t.Run("tell store to cancel work if request is cancelled", func(t *testing.T) {
    store := SpyStore{response: data, t: t}
    svr := Server(&store)

    request := httptest.NewRequest(http.MethodGet, "/", nil)

    cancellingCtx, cancel := context.WithCancel(request.Context())
    time.AfterFunc(5 * time.Millisecond, cancel)
    request = request.WithContext(cancellingCtx)

    response := &SpyResponseWrite{}

    svr.ServeHTTP(response, request)

    if response.written {
      t.Error("should not be written")
    }
  })
}
