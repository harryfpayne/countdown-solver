package numbers

import (
	"github.com/harryfpayne/countdown-solver/config"
	"github.com/harryfpayne/countdown-solver/itertools"
)

func Solve(cfg config.Config, numbers []int, target int, returnChan chan Expression) {
	gen := itertools.NewPermutationGenerator(numbers)
	operations := Operations
	if cfg.UseAllNumbers {
		operations = OperationsWithoutNoOp
	}

	for gen.Next() {
		permutation := gen.Get()
		opGen := itertools.NewCombinationGenerator(operations, len(numbers)-1)

		for opGen.Next() {
			ops := opGen.Get()

			bracketGen := itertools.NewCombinationGenerator(Brackets, len(numbers)-1)
		BracketLoop:
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

				expr := NewExpression(permutation, ops, bracketSequence)
				value, ok := expr.Evaluate(WithWorkingOut)
				if !ok {
					continue
				}
				if value == target {
					returnChan <- expr
					// Found a solution, don't need to check more brackets
					break BracketLoop
				}
			}
		}
	}
	close(returnChan)
}
