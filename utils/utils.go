package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var AllowedNumbers = [...]int{
	1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10,
	25, 50, 75, 100,
}

type Args struct {
	Numbers []int
	Target  int
	Letters []rune
}

func ReadArgs() Args {
	if len(os.Args) == 1 {
		fmt.Println("No arguments provided, using random numbers")
		nums, target := RandomNumbers()
		return Args{
			Numbers: nums,
			Target:  target,
		}
	}

	nums, target, ok := ReadArgsAsNumbersRound()
	if ok {
		return Args{Numbers: nums, Target: target}
	}

	letters, ok := ReadArgsAsLetters()
	if ok {
		return Args{Letters: letters}
	}

	fmt.Println("Invalid arguments provided")
	os.Exit(1)
	return Args{}
}

func ReadArgsAsLetters() ([]rune, bool) {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 1 {
		var letters []rune
		for _, letter := range argsWithoutProg[0] {
			letters = append(letters, letter)
		}
		return letters, true
	}

	letters := make([]rune, len(argsWithoutProg))
	for i, arg := range argsWithoutProg {
		// convert arg to rune
		letters[i] = []rune(arg)[0]
		// is alpha
		if !(letters[i] >= 'a' && letters[i] <= 'z') {
			return nil, false
		}
	}
	return letters, true

}

func ReadArgsAsNumbersRound() ([]int, int, bool) {
	argsWithoutProg := os.Args[1:]
	nums := make([]int, len(argsWithoutProg)-1)
	var err error
	for i, arg := range argsWithoutProg[:len(argsWithoutProg)-1] {
		// convert arg to int
		nums[i], err = strconv.Atoi(arg)
		if err != nil {
			return nil, 0, false
		}
	}

	target, err := strconv.Atoi(argsWithoutProg[len(argsWithoutProg)-1])
	if err != nil {
		return nil, 0, false
	}
	return nums, target, true
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
