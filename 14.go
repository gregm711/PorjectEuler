package main

/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

import (
	"fmt"
)

func sequence(x int, len int) (final int) {
	if x == 1 {
		final := len
	}
	if x%2 != 0 {

		sequence(3*x+1, len+1)
	}
	if x%2 == 0 {
		sequence(x/2, len+1)
	}
	return

}

func main() {
	//highest := 0
	//for x := 1000000; x > 1; x-- {
	//
	//	if sequence(x, 1) > highest {
	//		highest = x
	//	}
	//}
	final := sequence(13, 11)
	fmt.Println(final)

}
