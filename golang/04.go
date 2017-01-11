//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

//Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"strconv"
	// "math"
)

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-(i+1)] {
			return false
		}
	}
	return true

}

func main() {
	var palindrome int  = 0
	for x := 1000; x >= 0; x-- {
		for y := 1000; y >= 0; y-- {
			prod := y * x
			if isPalindrome(strconv.Itoa(prod)) == true {
				if prod >= palindrome {
					palindrome = prod
				}
			}

		}
	}
	fmt.Println(palindrome)

}
