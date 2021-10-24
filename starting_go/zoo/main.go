package main;

import (
  "fmt"
   "./animals"
)

func main() {
  fmt.Println("Hello world!");
  fmt.Println(animals.ElephantFeed());
  fmt.Println(animals.MonkeyFeed());
  fmt.Println(animals.RabbitFeed());
}
