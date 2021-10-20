package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for _, lang := range []string{"Golang", "Scala", "PHP"} {
		// 正の整数であればブロックしているゴルーチンが存在していること
		wg.Add(1)
		go func(lang string) {
			// この関数が終了した時に呼ばれる
			// カウンタに対して作用し、カウンタの値を１減算する
			defer wg.Done()
			fmt.Println(lang)
		}(lang)
		// カウンタが０になるまで待つ
		wg.Wait()
	}
}

/*
もしgoroutines内にて、defer wg.Done()をコメントアウトした場合
つまり、wg.Wait()で待ち受けているものの、いっこうにカウンタが0にならずに常に待ち構えている状態となっている

Golang
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc00009a008)
        /Users/Kotaro/.goenv/versions/1.11.4/src/runtime/sema.go:56 +0x39
sync.(*WaitGroup).Wait(0xc00009a000)
        /Users/Kotaro/.goenv/versions/1.11.4/src/sync/waitgroup.go:130 +0x64
main.main()
        /Users/Kotaro/go_practice/goroutines_waitGroup/main.go:22 +0xe3
*/