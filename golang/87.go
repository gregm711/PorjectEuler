/*
The smallest number expressible as the sum of a prime square, prime cube, and prime fourth power is 28. In fact, there are exactly four numbers below fifty that can be expressed in such a way:

28 = 22 + 23 + 24
33 = 32 + 23 + 24
49 = 52 + 23 + 24
47 = 22 + 33 + 24

How many numbers below fifty million can be expressed as the sum of a prime square, prime cube, and prime fourth power?
 */


package main

import (
	"fmt"
	"math"
)


func generatePrimes(limit int) (primes []int) {
	b := make([]bool, limit)
	for i := 2; i < limit; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < limit; k += i {
			b[k] = true
		}
	}
	return
}

func generateExpressions (primes []int, solution int) bool {

	for x:= 0; x <len(primes); x++ {
		if primes[x] > (int(math.Pow(solution, 0.25)) + 1) {
			break
		}
		if x
	}

}

func main() {
	primes := generatePrimes(50)







}