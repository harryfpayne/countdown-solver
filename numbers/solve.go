package numbers

import (
	"fmt"
	"github.com/harryfpayne/countdown-solver/itertools"
)

var log = false
var useAllNumbers = false

func Solve(numbers []int, target int, returnChan chan Solution) {
	gen := itertools.NewPermutationGenerator(numbers)
	operations := Operations
	if useAllNumbers {
		operations = OperationsWithoutNoOp
	}

	for gen.Next() {
		permutation := gen.Get()
		opGen := itertools.NewCombinationGenerator(operations, len(numbers)-1)

		for opGen.Next() {
			ops := opGen.Get()

			bracketGen := itertools.NewCombinationGenerator(Brackets, len(numbers)-1)
			for bracketGen.Next() {
				brackets := bracketGen.Get()
				if !IsValidBracketSequence(brackets) {
					continue
				}

				bracketSequence := GetBracketSequence(
					append(
						[]Bracket{NoBracket}, // redundant if expression starts with bracket
						brackets...,
					),
				)

				value, ok := Calculate(permutation, ops, bracketSequence)
				if !ok {
					continue
				}
				if value == target {
					returnChan <- NewSolution(permutation, ops, bracketSequence)
				}
			}
		}
	}
	close(returnChan)
}

func Calculate(numbers []int, operations []Operation, brackets []BracketFull) (int, bool) {
	return CalculateRec(numbers, operations, brackets, 0)
}

func CalculateRec(numbers []int, operations []Operation, brackets []BracketFull, indexOffset int) (int, bool) {
	value := numbers[0]
	var ok bool

	if log {
		fmt.Println("Starting expression: ", numbers, operations, brackets)
	}

	for i := 0; i < len(operations); i++ {
		op := operations[i]
		bracket := brackets[i+1]
		number := numbers[i+1]
		if log {
			fmt.Println("current op: ", number, op, bracket)
		}

		if bracket.Type == BracketOpen {
			if log {
				fmt.Println("bracket open")
			}
			// calculate sub expression
			closeBracketIndex := bracket.CorrespondingBracketIndex - indexOffset
			numbersSub := numbers[i+1 : closeBracketIndex+1]
			operationsSub := operations[i+1 : closeBracketIndex]
			bracketsSub := brackets[i+1 : closeBracketIndex+1]
			if log {
				fmt.Println("sub expression: ", numbersSub, operationsSub, bracketsSub)
			}
			subVal, ok := CalculateRec(numbersSub, operationsSub, bracketsSub, i+1)
			if !ok {
				return 0, false
			}
			if log {
				fmt.Println("sub value: ", subVal)
			}
			value, ok = Operate(value, subVal, op)
			if !ok { // Invalid operation (bad division), so skip to next operation set
				if log {
					fmt.Println("invalid operation")
				}
				return 0, false
			}
			if log {
				fmt.Println("new value: ", value)
			}
			i = closeBracketIndex - 1
			if log {
				fmt.Println("skipping to: ", i)
			}
			continue
		}

		value, ok = Operate(value, number, op)
		if !ok { // Invalid operation (bad division), so skip to next operation set
			if log {
				fmt.Println("invalid operation")
			}
			return 0, false
		}
		if log {
			fmt.Println("new value: ", value)
		}
	}
	return value, true
}
