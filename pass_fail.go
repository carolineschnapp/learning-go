// pass_fail reports whether a grade is passing or failing.
package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func main() {
  fmt.Print("Enter a grade: ")
  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }
  input = strings.TrimSpace(input)
  score, err := strconv.ParseFloat(input, 64)
  if err != nil {
    log.Fatal("You have to enter a number!")
  }
  printResult(score)
}


func printResult(score float64) {
  if score == 100 {
    fmt.Println("Perfect!")
  } else if score >= 60 {
    fmt.Println("You pass.")
  } else {
    fmt.Println("You fail!")
  }
}
