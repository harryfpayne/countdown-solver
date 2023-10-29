package main

import (
	"fmt"
	"github.com/harryfpayne/countdown-solver/letters"
	"github.com/harryfpayne/countdown-solver/numbers"
	"github.com/harryfpayne/countdown-solver/utils"
	"sort"
	"time"
)

func main() {
	args := utils.ReadArgs()
	if args.Numbers != nil {
		numbersRound(args)
	} else {
		lettersRound(args)
	}
}

func lettersRound(args utils.Args) {
	returnChan := make(chan string)
	fmt.Println("Trying to find words using:", string(args.Letters))
	t := time.Now()
	go letters.Solve(args.Letters, returnChan)

	var solutionsMap = make(map[string]struct{})
	for word := range returnChan {
		solutionsMap[word] = struct{}{}
	}
	if len(solutionsMap) == 0 {
		fmt.Println("No solutions found")
		return
	}
	var solutions []string
	for word := range solutionsMap {
		solutions = append(solutions, word)
	}
	fmt.Println("Found", len(solutions), "in", time.Since(t))

	sort.Slice(solutions, func(i, j int) bool {
		return len(solutions[i]) > len(solutions[j])
	})

	var longestFound int
	var firstFailure bool
	for _, word := range solutions {
		info, ok := letters.GetWordInfo(word)
		if !ok {
			if !firstFailure {
				fmt.Println("\nGot", word, "but can't find it's meaning")
				firstFailure = true
			}
			continue
		}
		fmt.Println("\n", info)
		if longestFound == 0 {
			longestFound = len(word)
		}
		if len(word) != longestFound {
			break
		}
	}
}

func numbersRound(args utils.Args) {
	nums, target := args.Numbers, args.Target
	fmt.Println("Trying to get to", target, "using", nums)
	returnChan := make(chan numbers.Expression)
	t := time.Now()
	go numbers.Solve(nums, target, returnChan)
	go func(returnChan chan numbers.Expression) { // Exit with timeout
		<-time.After(28 * time.Second)
		close(returnChan)
	}(returnChan)

	var solutions []numbers.Expression
	for solution := range returnChan {
		if len(solutions) == 0 {
			fmt.Println("First solution found:")
			fmt.Println(solution.WorkingOut)
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
	fmt.Println(nicestSolution.WorkingOut)
}
