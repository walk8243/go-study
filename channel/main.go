package main

import (
    "fmt"
    "time"
)

func say(c chan bool) {
    for i := 0; i < 5; i++ {
        time.Sleep(1000 * time.Millisecond)
        fmt.Println(i)
    }
    c <- true
}

func main() {
	channels := []chan bool{make(chan bool), make(chan bool)}
	for _, v := range channels {
		go say(v)
	}
	
	fmt.Println("この文章が数字より先に表示されます")
	
	for _, v := range channels {
		b := <- v
		fmt.Println(b)
	}
}
