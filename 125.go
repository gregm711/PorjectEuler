//The palindromic number 595 is interesting because it can be written as the sum of consecutive squares: 62 + 72 + 82 + 92 + 102 + 112 + 122.
//
//There are exactly eleven palindromes below one-thousand that can be written as consecutive square sums, and the sum of these palindromes is 4164. Note that 1 = 02 + 12 has not been included as this problem is concerned with the squares of positive integers.
//
//Find the sum of all the numbers less than 108 that are both palindromic and can be written as the sum of consecutive squares.

package main

import (
	"fmt"
	//"math"
	//"strconv"
	"math"
	"strconv"
)

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-(i+1)] {
			return false
		}
	}
	return true

}


func removeDuplicates(unfiltered []int) []int {
	filtered := []int{}

	for x := 0; x < len(unfiltered); x++ {
		exists := false
		for i := x+1; i < len(unfiltered); i ++ {
			if unfiltered[x] == unfiltered[i] {
				exists = true
				break
			}

		}
		if exists == false {
			filtered = append(filtered, unfiltered[x])
		}
	}
	return filtered
}


func main() {
	limit := int(math.Sqrt(math.Pow(10,8)))
	canidates := []int{}
	for x := 2; x <= limit; x++ {
		num := x * x
		for i := x + 1; i<= limit; i++ {
			num += i*i
			if num > int(math.Pow(10,8)) {
				break
			}
			if isPalindrome(strconv.Itoa(num)) == true {
				canidates = append(canidates, num)

			}
		}
	}
	filtered := removeDuplicates(canidates)
	fmt.Println(filtered)
	sum := 0

	for x := 0; x < len(filtered); x++ {
		sum += filtered[x]
	}
	fmt.Println(sum)

}
