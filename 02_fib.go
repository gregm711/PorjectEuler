package main

//By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
//Fibonacci = sum of two preceding numbers
import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	t1 := time.Now()
	for i, j := 0, 1; i < 4000000; i, j = i+j, i {
		if i%2 == 0 {
			sum = sum + i
		}

	}
	fmt.Println(time.Since(t1))
	fmt.Println(sum)

}
