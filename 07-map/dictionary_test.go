package main

import "testing"

func TestSearch(t *testing.T) {
  dictionary := Dictionary{"test": "this is just a test"}

  t.Run("known word", func(t *testing.T) {
    key := "test"
    got, _ := dictionary.Search(key)
    want := "this is just a test"

    assertString(t, got, want, key)
  })

  t.Run("unknown word", func(t *testing.T) {
    _, err := dictionary.Search("unknown")

    assertError(t, err, errNoKey)
  })
}

func TestAdd(t *testing.T) {

  t.Run("add new word", func(t *testing.T) {
    dictionary := Dictionary{}
    key := "test"
    value := "this is just a test"

    err := dictionary.Add(key, value)

    assertError(t, err, nil)
    assertDefinition(t, dictionary, key, value)
  })

  t.Run("add existing word", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionary := Dictionary{key: value}

    err := dictionary.Add(key, value)

    assertError(t, err, errWordExist)
    assertDefinition(t, dictionary, key, value)
  })
}

func TestUpdate(t *testing.T) {
  t.Run("existing key", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionary := Dictionary{key: value}
    newValue := "this is new definition"

    err := dictionary.Update(key, newValue)

    assertError(t, err, nil)
    assertDefinition(t, dictionary, key, newValue)
  })

  t.Run("none existing key", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionary := Dictionary{}

    err := dictionary.Update(key, value)

    assertError(t, err, errWordNotExist)
  })
}

func TestDelete(t *testing.T) {
  key := "test"
  dictionary := Dictionary{key: "key definition"}
  dictionary.Delete(key)

  _, err := dictionary.Search(key)
  if err != errNoKey {
    t.Errorf("expect %q to be deleted", key)
  }
}

func assertString(t testing.TB, got, want, key string) {
  t.Helper()
  if got != want {
    t.Errorf("got %q want %q, given %q", got, want, key)
  }
}

func assertError(t testing.TB, got, want error) {
  t.Helper()
  if got != want {
    t.Errorf("got error %q want error %q", got, want)
  }
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
  t.Helper()
  got, err := dictionary.Search(key)
  if err != nil {
    t.Fatal("should find added word: ", err)
  }

  if got != value {
    t.Errorf("got %q want %q given %q", got, value, "test")
  }
}
