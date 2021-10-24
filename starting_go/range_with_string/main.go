package main

import "fmt"

func main() {
	const str = "ABCDEFG"

	for i, c := range str {
		fmt.Printf("str[%d] = %d\n", i, c)
		fmt.Printf("str[%d] = %s\n", i, string(c))
	}

}
