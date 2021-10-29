package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now();
  unix := t.Unix();
  fmt.Println(unix);

  time.Sleep(5 * time.Second);
  fmt.Println("I fell asleep...");
}
