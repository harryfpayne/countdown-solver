package numbers

import (
	"github.com/harryfpayne/countdown-solver/itertools"
)

func Solve(numbers []int, target int, useAllNumbers bool) []Solution {
	gen := itertools.NewPermutationGenerator(numbers)
	operations := Operations
	if useAllNumbers {
		operations = OperationsWithoutNoOp
	}

	var solutions []Solution
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
					solutions = append(solutions, NewSolution(permutation, ops, bracketSequence))
				}
			}
		}
	}
	return solutions
}

func Calculate(numbers []int, operations []Operation, brackets []BracketFull) (int, bool) {
	return CalculateRec(numbers, operations, brackets, 0)
}

func CalculateRec(numbers []int, operations []Operation, brackets []BracketFull, indexOffset int) (int, bool) {
	value := numbers[0]
	var ok bool

	for i := 0; i < len(operations); i++ {
		op := operations[i]
		bracket := brackets[i+1]

		if bracket.Type == BracketOpen {
			// calculate sub expression
			closeBracketIndex := bracket.CorrespondingBracketIndex - indexOffset
			numbersSub := numbers[i+1 : closeBracketIndex+1]
			operationsSub := operations[i+1 : closeBracketIndex]
			bracketsSub := brackets[i+1 : closeBracketIndex+1]
			subVal, ok := CalculateRec(numbersSub, operationsSub, bracketsSub, i+1)
			if !ok {
				return 0, false
			}
			value, ok = Operate(value, subVal, op)
			i = closeBracketIndex - 1
			continue
		}

		value, ok = Operate(value, numbers[i+1], op)
		if !ok { // Invalid operation (bad division), so skip to next operation set
			return 0, false
		}
	}
	return value, true
}
