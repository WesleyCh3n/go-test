package main

type Dictionary map[string]string
type DictionaryErr string

const (
  errNoKey = DictionaryErr("could not find the word")
  errWordExist = DictionaryErr("word existing in dictionary")
  errWordNotExist = DictionaryErr("word not existing in dictionary")
)

func (e DictionaryErr) Error() string {
  return string(e)
}

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

func (d Dictionary) Update(key, value string) error {
  _, err := d.Search(key)

  switch err {
  case errNoKey:
    return errWordNotExist
  case nil:
    d[key] = value
  default:
    return err
  }
  return nil
}

func (d Dictionary) Delete(key string) {
  delete(d, key)
}
