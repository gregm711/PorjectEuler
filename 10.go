package main

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

import (
	"fmt"
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

func main() {
	final := 0
	primes := sieveOfEratosthenes(2000000)
	for i := 0; i < len(primes); i++ {
		final = final + primes[i]
	}
	fmt.Println(final)

}

// 142913828922
