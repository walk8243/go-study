package main

import (
    "fmt"
    "time"
)

func say() {
    for i := 0; i < 6; i++ {
        time.Sleep(1000 * time.Millisecond)
        fmt.Println(i)
    }
}

func main() {
    go say()
    fmt.Println("この文章が数字より先に表示されます")
    time.Sleep(5000 * time.Millisecond)
    fmt.Println("この文章は5秒後に強制的に表示されます")
}
