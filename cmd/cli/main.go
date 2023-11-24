package main

import (
	"fmt"
	"github.com/harryfpayne/countdown-solver/config"
	"github.com/harryfpayne/countdown-solver/letters"
	"github.com/harryfpayne/countdown-solver/numbers"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"sort"
	"time"
)

var cfg config.Config

func init() {
	rootCmd.PersistentFlags().DurationVarP(&cfg.Timeout, "timeout", "t", 28*time.Second, "timeout for solving")
	rootCmd.PersistentFlags().BoolVarP(&cfg.UseAllNumbers, "all-numbers", "a", false, "timeout for solving")
	rootCmd.PersistentFlags().BoolVarP(&cfg.Debug, "debug", "d", false, "enable logs")
}

var rootCmd = &cobra.Command{
	Use:   "countdown",
	Short: "A command line tool for solving Countdown letters and numbers rounds",
	Long: `A command line tool for solving Countdown letters and numbers rounds.

		It's expecting a list of numbers, the final one being the target, e.g.
			./countdown 1 2 3 4 5 6 1080

		Or a list of letters (either space delimited or not), e.g.
			./countdown abcdefghi
			./countdown a b c d e f g h i
	`,
	Run: func(cmd *cobra.Command, args []string) {
		toSolve, err := ParseArgs(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}

		if cfg.Debug {
			slog.SetDefault(slog.New(
				slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
					Level: slog.LevelDebug,
				}),
			))
		}

		if toSolve.Numbers != nil {
			numbersRound(toSolve)
		} else {
			lettersRound(toSolve)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func lettersRound(args Args) {
	returnChan := make(chan string)
	fmt.Println("Trying to find words using:", string(args.Letters))
	t := time.Now()
	slog.Debug("Beginning search")
	go letters.Solve(cfg, args.Letters, returnChan)

	var solutionsMap = make(map[string]struct{})
	for word := range returnChan {
		if len(solutionsMap)%10 == 0 {
			slog.Debug("Found", "count", len(solutionsMap))
		}
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
		slog.Debug("Getting info for", word)
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

func numbersRound(args Args) {
	nums, target := args.Numbers, args.Target
	fmt.Println("Trying to get to", target, "using", nums)
	returnChan := make(chan numbers.Expression)
	t := time.Now()
	slog.Debug("Beginning search")
	go numbers.Solve(cfg, nums, target, returnChan)
	go func(returnChan chan numbers.Expression) { // Exit with timeout
		slog.Debug("Waiting for timeout: ", "timeout", cfg.Timeout.String())
		<-time.After(cfg.Timeout)
		slog.Debug("Timeout reached")
		close(returnChan)
	}(returnChan)

	var solutions []numbers.Expression
	for solution := range returnChan {
		if len(solutions) == 0 {
			fmt.Println("First solution found:")
			fmt.Println(solution.WorkingOut)
		}
		if len(solutions)%10 == 0 {
			slog.Debug("Found", "count", len(solutions))
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
