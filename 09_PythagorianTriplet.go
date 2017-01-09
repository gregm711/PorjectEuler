/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import (
	"fmt"
)

func main() {

	for a := 1; a <= 1000; a++ {
		for b := 1; b <= 1000; b++ {
			for c := 1; c <= 1000; c++ {
				if (c * c) == (a*a + b*b) {
					if (c + a + b) == 1000 {
						fmt.Println(c, a, b)
						fmt.Println(c*a*b)
					}
				}

			}

		}

	}

}
