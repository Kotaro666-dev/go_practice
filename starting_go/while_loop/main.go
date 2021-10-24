package main

import f "fmt"

func main() {
	i := 0

	for {
		f.Println(i)
		if i == 10 {
			break
		}
		i++
	}
}
