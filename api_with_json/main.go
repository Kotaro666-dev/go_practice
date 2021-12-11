package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Menu struct {
	XMLName  xml.Name   `xml:"menu"`
	Contents []Contents `xml:"contents"`
}

type Contents struct {
	Index int    `xml:"index"`
	Text  string `xml:"text"`
	Url   string `xml:"url"`
}

var MenuData = Menu{
	Contents: []Contents{
		Contents{
			Index: 0,
			Text:  "Menu00",
			Url:   "/test00",
		},
		Contents{
			Index: 1,
			Text:  "Menu01",
			Url:   "/test01",
		},
		Contents{
			Index: 2,
			Text:  "Menu02",
			Url:   "/test03",
		},
		Contents{
			Index: 3,
			Text:  "Menu03",
			Url:   "/test03",
		},
		Contents{
			Index: 4,
			Text:  "Menu04",
			Url:   "/test04",
		},
	},
}

func HandlerXml(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	encoder := xml.NewEncoder(&buffer)
	encoder.Indent("", "\t") // 各行にインデントを追加
	if err := encoder.Encode(&MenuData); err != nil {
		log.Fatal(err)
	}
	output := xml.Header + buffer.String()
	_, err := fmt.Fprint(w, output)
	if err != nil {
		fmt.Println("Error encoding XML: ", err)
		return
	}
	fmt.Println("Returning XML data successfully.")
}

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

func HandlerJson(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/json", HandlerJson)
	http.HandleFunc("/xml", HandlerXml)
	log.Println("Start listening server...")
	http.ListenAndServe(":8080", nil)
}
