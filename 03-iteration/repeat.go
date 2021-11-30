package iteration

import "strings"

// const repeatCount = 5
func Repeat(character string, repeatCount int) (string) {
  /*
  * for i := 0; i < repeatCount; i++ {
    *   result += character
    * }
    */
    return strings.Repeat(character, repeatCount)
  }
