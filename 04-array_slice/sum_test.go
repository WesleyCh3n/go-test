package main

import (
  "reflect"
  "testing"
)

func TestSum(t *testing.T){
  t.Run("collection of size 5", func(t *testing.T){
    numbers := []int{1,2,3,4,5}
    got := Sum(numbers)
    want := 15

    if got != want {
      t.Errorf("want '%d' but got '%d', given %v", want, got, numbers)
    }
  })

  t.Run("collection of any size", func(t *testing.T){
    numbers := []int{1,2,3,4,5,6}
    got := Sum(numbers)
    want := 21

    if got != want {
      t.Errorf("want '%d' but got '%d', given %v", want, got, numbers)
    }
  })
}

func TestSumAll(t *testing.T) {
  got := SumAll([]int{0,2}, []int{8,2})
  want := []int{2,10}
  if !reflect.DeepEqual(got, want) {
    t.Errorf("want '%d' but got '%d'", want, got)
  }
}

func TestSumAllTails(t *testing.T) {
  checkSum := func(t testing.TB, got, want []int){
    t.Helper()
    if !reflect.DeepEqual(got, want) {
      t.Errorf("want '%d' but got '%d'", want, got)
    }
  }
  t.Run("make the sum of some slices", func(t *testing.T){
    got := SumAllTails([]int{0,2}, []int{2,8})
    want := []int{2,8}
    checkSum(t, got, want)
  })
  t.Run("saftly sum empty slices", func(t *testing.T){
    got := SumAllTails([]int{}, []int{2,8,12})
    want := []int{0,20}
    checkSum(t, got, want)
  })
}
