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
*/
