package main

import "errors"

type Dictionary map[string]string

var (
  errNoKey = errors.New("could not find the word")
  errWordExist = errors.New("word existing in dictionary")
)

func (d Dictionary) Search(key string) (string, error) {
  definition, ok := d[key]
  if !ok {
    return "", errNoKey
  }
  return definition, nil
}

func (d Dictionary) Add(key, value string) error {
  _, err := d.Search(key)

  switch err {
  case errNoKey:
    d[key] = value
  case nil:
    return errWordExist
  default:
    return err
  }

  return nil
}
