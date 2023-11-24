package main

type LettersSolution struct {
	BestSolution string
	Solutions    []string
}

type NumbersSolution struct {
	BestSolution string
}

type Error struct {
	Message string `json:"error"`
}

var noSolutions = Error{Message: "No solutions found"}
var invalidInput = Error{Message: "Invalid input"}
