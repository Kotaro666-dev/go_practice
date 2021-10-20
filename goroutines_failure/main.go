package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("say hello")
}

func main() {
	// go sayHello() によりゴルーチンが生成(fork)されるがメインゴルーチンの処理が終わるまでにその実行タイミングが無い
	go sayHello()
	fmt.Println("say hello from main goroutine")
}

// 以下は失敗パターンを成功させるパターン。goroutines時にwaitさせることで、goroutines処理が終わるをまってから、プログラムを終了する
//func sayHello(wg *sync.WaitGroup) {
//	defer wg.Done()
//	fmt.Println("say hello")
//}
//
//func main() {
//	// go sayHello() によりゴルーチンが生成(fork)されるがメインゴルーチンの処理が終わるまでにその実行タイミングが無い
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go sayHello(&wg)
//	wg.Wait()
//	fmt.Println("say hello from main goroutine")
//}

