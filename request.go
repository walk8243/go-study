package main

import (
	"os"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.yahoo.co.jp/")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Hello, World!", resp)
}
