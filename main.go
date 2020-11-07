package main

import (
	"fmt"
)

func f(a int, b int) (int, int) {
	c := a + b
	d := a * b
	return c, d
}

func main() {
	a := 1
	b := 2
	c, d := f(a, b)
	fmt.Println("Hello, World!", c, d)
}
