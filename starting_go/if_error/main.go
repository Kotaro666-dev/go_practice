package main

import (
	"fmt"
	"net/http"
)

const url = "https://this_doesnt_exist.com"

func main() {
	if _, err := http.Get(url); err != nil {
		fmt.Println("http error")
		return
	}
	fmt.Println("http success")
}
