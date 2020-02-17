package main

import (
	"fmt"
	"math/rand"
)

// Julie proposes the following wager to her sister Louise.
// She suggests they play a game of chance to determine who will wash the dishes.
// For this game, they shall use a generator of independent random numbers uniformly distributed between 0 and 1.
// The game starts with S = 0.
// The first player, Louise, adds to S different random numbers from the generator until S > 1 and records her last random number 'x'.
// The second player, Julie, continues adding to S different random numbers from the generator until S > 2 and records her last random number 'y'.
// The player with the highest number wins and the loser washes the dishes, i.e. if y > x the second player wins.

// For example, if the first player draws 0.62 and 0.44, the first player turn ends since 0.62+0.44 > 1 and x = 0.44.
// If the second players draws 0.1, 0.27 and 0.91, the second player turn ends since 0.62+0.44+0.1+0.27+0.91 > 2 and y = 0.91. Since y > x, the second player wins.

// Louise thinks about it for a second, and objects: "That's not fair".
// What is the probability that the second player wins?
// Give your answer rounded to 10 places behind the decimal point in the form 0.abcdefghij

func main() {
	results := make([]float64, 1)
	iterations := 10000000.0
	res := 0.0
	total := 0.0
	for iteration := 0.0; iteration < 10000000000; iteration++ {
		current := rand.Float64()
		res = res + current
		if res > 1.0 {
			results = append(results, current)
			total = total + current
			res = 0.0
			iteration = iteration + 1
		}
		if len(results) > int(iterations) {
			break
		}
	}
	fmt.Println("average is ", total/iterations)

	//1.3591296257282346

}
