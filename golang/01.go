package main

import (
	"fmt"
	"time"
)

//Find the sum of all the multiples of 3 or 5 below 1000.
func main() {
	t1 := time.Now()
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	fmt.Println(time.Since(t1))
	fmt.Println(sum)
}
