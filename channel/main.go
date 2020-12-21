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
	channel := make(chan bool)
	count := 5

	for i := 0; i < count; i++ {
		go say(channel)
	}
	
	for i := 0; i < count; i++ {
		fmt.Println(<- channel)
	}
}
