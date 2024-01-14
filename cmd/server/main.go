package main

import (
	"encoding/json"
	"fmt"
	"github.com/harryfpayne/countdown-solver/config"
	"github.com/harryfpayne/countdown-solver/letters"
	"github.com/harryfpayne/countdown-solver/numbers"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Port = 8080

func main() {
	fs := http.FileServer(http.Dir("./client/build"))

	http.Handle("/", fs)
	http.HandleFunc("/api/letters", lettersHandler)
	http.HandleFunc("/api/numbers", numbersHandler)
	fmt.Println("Listening on port", Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
	if err != nil {
		panic(err)
	}
}

func lettersHandler(w http.ResponseWriter, r *http.Request) {
	cfg := config.Default

	lettersStr := r.URL.Query().Get("letters")
	lettersStr = strings.ToLower(lettersStr)
	lets := make([]rune, len(lettersStr))
	for i, letter := range lettersStr {
		if !(letter >= 'a' && letter <= 'z') {
			json.NewEncoder(w).Encode(invalidInput)
			return
		}

		lets[i] = letter
	}

	returnChan := make(chan string)
	go letters.Solve(cfg, lets, returnChan)
	go func() {
		<-time.After(cfg.Timeout)
		close(returnChan)
	}()

	var solutionsMap = make(map[string]struct{})
	for solution := range returnChan {
		solutionsMap[solution] = struct{}{}
	}

	if len(solutionsMap) == 0 {
		err := json.NewEncoder(w).Encode(noSolutions)
		if err != nil {
			slog.Error("Letters no solutions", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	bestSolution := ""
	var solutions []string
	for solution := range solutionsMap {
		solutions = append(solutions, solution)
		if len(solution) > len(bestSolution) {
			bestSolution = solution
		}
	}
	var lettersSolution LettersSolution
	lettersSolution.BestSolution = bestSolution
	lettersSolution.Solutions = solutions
	err := json.NewEncoder(w).Encode(lettersSolution)
	if err != nil {
		slog.Error("Letters solutions", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func numbersHandler(w http.ResponseWriter, r *http.Request) {
	cfg := config.Default

	numbersSlice := r.URL.Query().Get("numbers")
	numbersSlice = strings.ToLower(numbersSlice)
	fmt.Println(numbersSlice)
	numbStrings := strings.Split(numbersSlice, ",")
	numbs := make([]int, len(numbStrings))
	for i, numbString := range numbStrings {
		numb, err := strconv.Atoi(numbString)
		if err != nil {
			json.NewEncoder(w).Encode(invalidInput)
			return
		}
		numbs[i] = numb
	}

	targetStr := r.URL.Query().Get("target")
	fmt.Println(targetStr)
	target, err := strconv.Atoi(targetStr)
	if err != nil || target < 100 || len(numbs) < 2 {

		json.NewEncoder(w).Encode(invalidInput)
		return
	}

	returnChan := make(chan numbers.Expression)
	go numbers.Solve(cfg, numbs, target, returnChan)
	go func() {
		<-time.After(cfg.Timeout)
		close(returnChan)
	}()

	// TODO make results json
	solution := <-returnChan
	fmt.Println("we back", solution)
	if len(solution.Numbers) == 0 {
		err := json.NewEncoder(w).Encode(noSolutions)
		if err != nil {
			slog.Error("Numbers no solutions", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	var numbersSolution NumbersSolution
	numbersSolution.BestSolution = solution.String()
	err = json.NewEncoder(w).Encode(numbersSolution)
	if err != nil {
		slog.Error("Numbers solutions", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
