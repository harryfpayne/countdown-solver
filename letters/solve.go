package letters

import (
	"bufio"
	"github.com/harryfpayne/countdown-solver/itertools"
	"os"
	"slices"
)

func Solve(letters []rune, resultChan chan string) {
	words := openfile()

	permutationGen := itertools.NewPermutationGenerator(letters)
	for permutationGen.Next() {
		word := permutationGen.Get()
		_, ok := slices.BinarySearch(words, string(word))
		if ok {
			resultChan <- string(word)
			continue
		}

		for i := 1; i < len(word); i++ {
			subword := word[:i]
			_, ok := slices.BinarySearch(words, string(subword))
			if ok {
				resultChan <- string(subword)
			}
		}

	}
	close(resultChan)
}

func openfile() []string {
	f, err := os.Open("words_alpha.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}
