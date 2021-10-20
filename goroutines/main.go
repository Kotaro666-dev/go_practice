package  main

import (
    "fmt"
    "time"
)

func my_func(value string) {
    for i := 0; i < 3; i++ {
        fmt.Println(value)
        time.Sleep(3 * time.Second)
    }
}

func main() {
    go my_func("goroutineを使って実行")
    my_func("普通に実行")
    fmt.Println("done")
}
