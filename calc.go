package main

func f(a int, b int) (int, int) {
	c := a + b
	d := a * b
	return c, d
}

func getPrimeNumber(maxNum int) ([]int) {
	primes := [] int{}
	for i := 2; i < maxNum; i++ {
		if isPrimeNumber(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrimeNumber(number int) (bool) {
	max := number / 2
	flag := true
	for i := 2; i <= max; i++ {
		if number % i == 0 {
			flag = false
			break
		}
	}
	return flag
}
