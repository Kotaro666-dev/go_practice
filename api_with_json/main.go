package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CustomerInfo struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	AccountNumber string `json:"accountNumber"`
	IsLoggedIn    bool   `json:"isLoggedIn"`
	DepositAmount int    `json:"depositAmount"`
}

var CustomerInfoData = CustomerInfo{
	Id:            1,
	Name:          "斎藤健太",
	AccountNumber: "1234 5678 1111",
	IsLoggedIn:    true,
	DepositAmount: 1500000,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	if err := encoder.Encode(&CustomerInfoData); err != nil {
		log.Fatal(err)
	}
	_, err := fmt.Fprint(w, buffer.String())
	if err != nil {
		fmt.Println("Error encoding JSON: ", err)
		return
	}
  fmt.Println("Returning JSON data successfully.")
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start listening server...")
	http.ListenAndServe(":8080", nil)
}
