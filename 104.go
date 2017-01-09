package main

/*
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
It turns out that F541, which contains 113 digits, is the first Fibonacci number for which the last nine digits are 1-9 pandigital (contain all the digits 1 to 9, but not necessarily in order). And F2749, which contains 575 digits, is the first Fibonacci number for which the first nine digits are 1-9 pandigital.

Given that Fk is the first Fibonacci number for which the first nine digits AND the last nine digits are 1-9 pandigital, find k.
*/

import (
	"fmt"
	"math/big"
	//"strconv"
)

//func isPandigital(str string, val int) bool {
//
//	for i := 0; i < len(str); i++ {
//		num, _ := strconv.Atoi(string(str[i]))
//		if val == num {
//			if val == 9 {
//				return true
//			}
//			newNum := val + 1
//			return isPandigital(str, newNum)
//		}
//	}
//
//	return false
//
//}
//
func convertToString(n *big.Int) string {
	return n.String()
}

func main() {

	a := big.NewInt(0)
	b := big.NewInt(1)
	k := 1
	//var str string
	// Initialize limit as 10^99, the smallest integer with 100 digits.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(1000000), nil)

	// Loop while a is smaller than 1e100.
	for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		k += 1
		fmt.Println(k)
		//if k >= 329468 {
		//fmt.Println(k)
		convertToString(a)

		//if len(str) >= 9 {
		//	start := 1
		//	firstNine, lastNine := isPandigital(str[:9], start), isPandigital(str[(len(str) - 9):], start)
		//	if firstNine && lastNine {
		//		fmt.Println(str, len(str), k, lastNine)
		//		break
		//	}
		//}
		//// Swap a and b so that b is the next number in the sequence.

		a, b = b, a
	}
	//fmt.Println(a,len(convertToString(a)), k) // 100-digit Fibonacci number

}
