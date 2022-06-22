package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	i := flag.Int("int", 0, "integer to check for flags")
	flag.Parse()
	if *i < 1 {
		fmt.Println("integer must be greater than one")
	}
	primes := getPrimes(*i)
	fmt.Println(strings.Trim(
		strings.Join(
			strings.Fields(
				fmt.Sprint(primes)), ","), "[]",
	))
}

func getPrimes(num int) []int {
	primes := make([]int, 0)
	// 3 is the first prime number
	for j := 3; j <= num; j++ {
		if isPrime(j) {
			primes = append(primes, j)
		}
	}
	return primes
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
