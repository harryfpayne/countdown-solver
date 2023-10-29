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

	OperationLoop:
		for opGen.Next() {
			value := permutation[0]
			ops := opGen.Get()
			var ok bool
			for i, op := range ops {
				value, ok = Operate(value, permutation[i+1], op)
				if !ok { // Invalid operation (bad division), so skip to next operation set
					continue OperationLoop
				}
			}
			if value == target {
				solutions = append(solutions, NewSolution(permutation, ops))
			}
		}
	}
	return solutions
}
