package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(c chan bool) {
	wg.Add(1)
    for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
        fmt.Println(i)
	}
	defer wg.Done()
	c <- true
}

func main() {
	channel := make(chan bool)
	count := 5

	for i := 0; i < count; i++ {
		go say(channel)
	}
	
	for i := 0; i < count - 1; i++ {
		fmt.Println(<- channel)
	}

	fmt.Println("Hello")
	fmt.Println(<- channel)
	defer close(channel)

	wg.Wait()

	fmt.Println("End")
}
