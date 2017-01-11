package main

// What is the largest prime factor of the number 600851475143 ?

import (
	"fmt"
	"math"
)

func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func exists(N int, arr []int) bool {
	i := 0
	for range arr {
		if N == i {
			return true
		}
		i++
	}
	return false
}

func main() {
	final := 0
	primes := sieveOfEratosthenes(int(math.Sqrt(600851475143)))
	for i := len(primes) - 1; i > 0; i-- {
		v := primes[i]
		if 600851475143 % v == 0 {
		final = v	
		break
		}
	}
	fmt.Println(final)

}
