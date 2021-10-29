package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
	  if _, err := os.Open("test"); err != nil {
	    log.Fatal(err);
	  }
	*/

	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// fmt.Println(file);

	buffer := make([]byte, 64)
	for {
		bytes, err := file.Read(buffer)
		if bytes == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(string(buffer))
	}
}
