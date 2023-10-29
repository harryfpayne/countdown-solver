package main

import (
	"fmt"
	"github.com/harryfpayne/countdown-solver/numbers"
	"time"
)

func main() {
	nums := []int{50, 75, 100, 25, 2, 1}
	target := 943
	t := time.Now()
	solutions := numbers.Solve(nums, target, false)
	fmt.Println("Time taken:", time.Since(t))

	if len(solutions) == 0 {
		fmt.Println("No solutions found")
		return
	}

	nicestSolution := solutions[0]
	complexestSolution := solutions[0]
	for _, solution := range solutions {
		solution := solution
		if nicestSolution.Niceness() < solution.Niceness() {
			nicestSolution = solution
		}
		if complexestSolution.Complexity() < solution.Complexity() {
			complexestSolution = solution
		}
	}

	fmt.Println("Found", len(solutions), "solutions")
	fmt.Println("Nicest solution:", nicestSolution)
	fmt.Println("Complexest solution:", complexestSolution.StringMultiLine())
}

/*
	Can't solve all problems e.g.:
		30 * 12 = 360
		9 + 9 = 18
		360 / 18 = 20

		Need to handle brackets to solve this but think it's quite hard to do efficiently:
		naive but close approx to number of checks is 6! * 5^5 * 51 = 114,750,000
		Atm does all solutions in 367ms, if no extra compute overhead including the brackets then would be 18s
		so would fit in 30s time but very close to the limit, ~600ms is our limit to fit under 30s

		n.b. 51 comes from:
			each digit can either open, close, or be no bracket
			e.g. '(3' or '3)', '3'
			computing all valid combinations of brackets for 6 digits there are only 51 solutions
*/
