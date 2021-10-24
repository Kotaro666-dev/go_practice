package main

import "fmt"

func return_multiple() (string, string, string) {
  first := "First";
  second := "Second";
  third := "Third";

  return first, second, third;
}

func main() {
  _, _, third := return_multiple();

  // fmt.Println(first);
  // fmt.Println(second);
  fmt.Println(third);
}
