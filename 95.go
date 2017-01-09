/*
The proper divisors of a number are all the divisors excluding the number itself. For example, the proper divisors of 28 are 1, 2, 4, 7, and 14. As the sum of these divisors is equal to 28, we call it a perfect number.

Interestingly the sum of the proper divisors of 220 is 284 and the sum of the proper divisors of 284 is 220, forming a chain of two numbers. For this reason, 220 and 284 are called an amicable pair.

Perhaps less well known are longer chains. For example, starting with 12496, we form a chain of five numbers:

12496 → 14288 → 15472 → 14536 → 14264 (→ 12496 → ...)

Since this chain returns to its starting point, it is called an amicable chain.

Find the smallest member of the longest amicable chain with no element exceeding one million.
*/

package main

import (
	"fmt"
	"math"
	"runtime"
)

func sieveFactors(N int) (sum int) {
	limit := int(math.Sqrt(float64(N)))
	for i := 2; i < limit; i++ {
		if N%i == 0 {
			sum += i
			sum += N / i

		}
	}
	return sum + 1
}

func inChain(num int, chain []int) bool {
	for x := 1; x < len(chain); x++ {
		if num == chain[x] {
			return true
		}
	}
	return false
}

func findChain(num int, currentChain []int) []int {
	newInt := sieveFactors(num)
	if newInt < 1000000 && inChain(newInt, currentChain) == false {
		currentChain = append(currentChain, newInt)
		if newInt == currentChain[0] {
			return currentChain

		} else {
			return findChain(newInt, currentChain)
		}

	}
	return []int{0}

}

func main() {
	runtime.GOMAXPROCS(6)

	start := 12496
	oldChain := []int{start}
	for x := start; x < 1000000; x++ {
		newChain := findChain(x, []int{x})
		if len(newChain) >= len(oldChain) {
			oldChain = newChain
		}
	}
	fmt.Println("done", oldChain)
}
