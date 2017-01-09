package main

/*

Working from left-to-right if no digit is exceeded by the digit to its left it is called an increasing number; for example, 134468.

Similarly if no digit is exceeded by the digit to its right it is called a decreasing number; for example, 66420.

We shall call a positive integer that is neither increasing nor decreasing a "bouncy" number; for example, 155349.

Clearly there cannot be any bouncy numbers below one-hundred, but just over half of the numbers below one-thousand (525) are bouncy. In fact, the least number for which the proportion of bouncy numbers first reaches 50% is 538.

Surprisingly, bouncy numbers become more and more common and by the time we reach 21780 the proportion of bouncy numbers is equal to 90%.

Find the least number for which the proportion of bouncy numbers is exactly 99%.
*/
import (
	"fmt"
	"strconv"
)



func isBouncy(str string) bool {
	firstVal, _ := strconv.Atoi(string(str[0]))
	secondVal, _ := strconv.Atoi(string(str[1]))
	increasing := (firstVal <= secondVal)
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(string(str[i]))
		nextNum, _ := strconv.Atoi(string(str[i+1]))
		fmt.Println(num, nextNum, increasing)
		if !((nextNum > num) == increasing) {
			return false
		}
		if nextNum == num {
			continue
		} else {
			increasing = !increasing
		}

	}
	return true
}

func main() {
	num := 155349
	str := strconv.Itoa(num)
	fmt.Println(isBouncy(str))

}
