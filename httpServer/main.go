package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World")
}

func main() {
  http.HandleFunc("/", handler)
  fmt.Println("Listening Server....")
  http.ListenAndServe(":8080", nil)
}
