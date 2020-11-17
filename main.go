package main

import (
	"fmt"
	"math"
)

func main() {
	digit := 3

	number := int(math.Pow(10, float64(digit))) - 1
	primes := getPrimeNumber(number)
	fmt.Println(primes)
}
