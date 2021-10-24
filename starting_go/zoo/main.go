package main;

import (
  "fmt"
   "zoo/animals"
)

func main() {
  fmt.Println("Hello world!")
  fmt.Println(AppName());
  fmt.Println(animals.ElephantFeed());
  fmt.Println(animals.MonkeyFeed());
  fmt.Println(animals.RabbitFeed());
}
