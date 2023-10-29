package main

import (
	"fmt"
	"github.com/harryfpayne/countdown-solver/numbers"
	"time"
)

func main() {
	nums := []int{29, 1, 12, 9, 9, 2}
	target := 20
	t := time.Now()
	solutions := numbers.Solve(nums, target, true)
	fmt.Println("Time taken:", time.Since(t))

	if len(solutions) == 0 {
		fmt.Println("No solutions found")
		return
	}

	nicestSolution := solutions[0]
	for _, solution := range solutions {
		solution := solution
		if nicestSolution.Niceness() < solution.Niceness() {
			nicestSolution = solution
		}
	}

	fmt.Println("Found", len(solutions), "solutions")
	fmt.Println("Nicest solution:", nicestSolution)
}
