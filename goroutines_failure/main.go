package main

import "fmt"

func sayHello() {
	fmt.Println("say hello")
}

func main() {
	// go sayHello() によりゴルーチンが生成(fork)されるがメインゴルーチンの処理が終わるまでにその実行タイミングが無い
	go sayHello()
	fmt.Println("say hello from main goroutine")
}

