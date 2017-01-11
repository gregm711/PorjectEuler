/*
Euler's Totient function, φ(n) [sometimes called the phi function], is used to determine the number of numbers less than n which are relatively prime to n. For example, as 1, 2, 4, 5, 7, and 8, are all less than nine and relatively prime to nine, φ(9)=6.

n	Relatively Prime	φ(n)	n/φ(n)
2	1	1	2
3	1,2	2	1.5
4	1,3	2	2
5	1,2,3,4	4	1.25
6	1,5	2	3
7	1,2,3,4,5,6	6	1.1666...
8	1,3,5,7	4	2
9	1,2,4,5,7,8	6	1.5
10	1,3,7,9	4	2.5
It can be seen that n=6 produces a maximum n/φ(n) for n ≤ 10.

Find the value of n ≤ 1,000,000 for which n/φ(n) is a maximum.
*/

package main

import "fmt"

//func gcd(x int, y int) int {
//	for y != 0  {
//		x, y  = y, x %y
//
//	}
//	return x
//}

func primes(n int) (primes []int) {
	set := make([]bool, n)
	for x := 2; x <n; x++ {
		if set[x] == true{
			continue
		}
		primes = append(primes, x)
		for k := x*x; k < n; k += x {
			set[k] = true
		}

	}
	return primes
}



func main() {
	x := primes(1000000)

	prodPrimes := 1

	for y := 0; y < len(x); y++ {

		if prodPrimes * x[y] > 1000000 {
			break
		}
	}
	fmt.Println(prodPrimes)


}
