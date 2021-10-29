package main

import (
  "fmt"
  "time"
)

func main() {
  ch := time.Tick(3 * time.Second);
  for {
    t := <-ch;
    fmt.Println(t);
  }
}
