package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func request() {
	resp, err := http.Get("https://www.yahoo.co.jp/")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
