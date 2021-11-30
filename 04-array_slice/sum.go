package main

func Sum(numbers []int) int {
  sum := 0
  for _, number := range numbers {
    sum += number
  }
  return sum
}

func SumAll(slices ...[]int) (sum []int) {
  for _, slice := range slices{
    sum = append(sum, Sum(slice))
  }
  return
}

func SumAllTails(slices ...[]int) (sum []int) {
  for _, slice := range slices {
    if len(slice) == 0 {
      sum = append(sum, 0)
    } else {
      sum = append(sum, Sum(slice[1:]))
    }
  }
  return
}
