/*
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
 */


package main

import (
	"fmt"
)


func main() {
	sumSquares := 0
	squareSums := 0
	for i := 1; i <= 100; i++ {
		sumSquares = sumSquares + (i * i)
		squareSums = squareSums + i

	}
	squareSums = squareSums * squareSums

	fmt.Println(squareSums - sumSquares)
}
