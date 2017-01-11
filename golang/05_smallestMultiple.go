// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?


package main

import (
	"fmt"
	"os"
)


func isDivisable(num int) bool {
	for i := 1; i <= 20; i++ {
		if num % i  != 0 {
			return false
		}
	}
	return  true
}


func main() {


	var smallest int
	i := 1
	for i < 10000000000 {
		if isDivisable(i) == true {
			smallest = i
			fmt.Println(smallest)
			os.Exit(3)

		}
		i++
	}

}
