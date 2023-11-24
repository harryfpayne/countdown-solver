package letters

import (
	_ "embed"
	"github.com/harryfpayne/countdown-solver/config"
	"github.com/harryfpayne/countdown-solver/itertools"
	"log/slog"
	"slices"
	"sort"
	"strings"
)

//go:embed words_alpha.txt
var words string

func Solve(cfg config.Config, letters []rune, resultChan chan string) {
	words := parseFile()
	slog.Debug("Loaded", "count", len(words))

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

func parseFile() []string {
	var lines []string
	for _, word := range strings.Split(words, "\n") {
		lines = append(lines, word)
	}
	sortWords(lines)
	return lines
}

func sortWords(words []string) {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
}
