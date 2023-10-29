package main

import (
	"fmt"
	"github.com/harryfpayne/countdown-solver/numbers"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	nums, target := ReadArgs()
	fmt.Println("Trying to get to", target, "using", nums)
	returnChan := make(chan numbers.Solution)
	t := time.Now()
	go numbers.Solve(nums, target, returnChan)
	go func(returnChan chan numbers.Solution) { // Exit with timeout
		<-time.After(28 * time.Second)
		close(returnChan)
	}(returnChan)

	var solutions []numbers.Solution
	for solution := range returnChan {
		if len(solutions) == 0 {
			fmt.Println("First solution found:")
			fmt.Println(solution)
		}
		solutions = append(solutions, solution)
	}
	if len(solutions) == 0 {
		fmt.Println("No solutions found")
		return
	}
	fmt.Println("Found", len(solutions), "in", time.Since(t))

	nicestSolution := solutions[0]
	for _, solution := range solutions {
		solution := solution
		if nicestSolution.Niceness() < solution.Niceness() {
			nicestSolution = solution
		}
	}
	fmt.Println("Nicest solution:")
	fmt.Println(nicestSolution)
}

var AllowedNumbers = [...]int{
	1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10,
	25, 50, 75, 100,
}

func ReadArgs() ([]int, int) {
	if len(os.Args) == 1 {
		fmt.Println("No arguments provided, using random numbers")
		return RandomNumbers()
	}
	argsWithoutProg := os.Args[1:]
	nums := make([]int, len(argsWithoutProg)-1)
	var err error
	for i, arg := range argsWithoutProg[:len(argsWithoutProg)-1] {
		// convert arg to int
		nums[i], err = strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Error parsing number", arg)
			os.Exit(1)
		}
	}

	target, err := strconv.Atoi(argsWithoutProg[len(argsWithoutProg)-1])
	if err != nil {
		fmt.Println("Error parsing target", argsWithoutProg[len(argsWithoutProg)-1])
		os.Exit(1)
	}
	return nums, target
}

func RandomNumbers() ([]int, int) {
	nums := make([]int, 6)
ILoop:
	for i := range nums {
		// Random number from AllowedNumbers
		// No repeats
		nums[i] = AllowedNumbers[rand.Intn(len(AllowedNumbers))]
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				i--
				continue ILoop
			}
		}
	}
	return nums, rand.Intn(900) + 100
}
