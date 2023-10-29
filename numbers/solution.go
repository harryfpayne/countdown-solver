package numbers

import "fmt"

type Solution struct {
	NumberSequence    []int
	OperationSequence []Operation
	BracketSequence   []BracketFull
}

func NewSolution(numberSequence []int, operationSequence []Operation, bracketSequence []BracketFull) Solution {
	return Solution{
		NumberSequence:    numberSequence,
		OperationSequence: operationSequence,
		BracketSequence:   bracketSequence,
	}
}

func (s Solution) String() string {
	out := fmt.Sprintf("%d", s.NumberSequence[0])
	for i, op := range s.OperationSequence {
		if op == NoOp {
			continue
		}
		number := s.NumberSequence[i+1]
		bracket := s.BracketSequence[i+1]
		switch bracket.Type {
		case NoBracket:
			switch op {
			case Plus, Minus:
				out = fmt.Sprintf("(%s %s %d)", out, op, number)
			default:
				out = fmt.Sprintf("%s %s %d", out, op, number)
			}
		case BracketOpen:
			out = fmt.Sprintf("%s %s (%d", out, op, number)
		case BracketClose:
			out = fmt.Sprintf("%s %s %d)", out, op, number)
		}
	}
	return out
}

func (s Solution) StringMultiLine() string {
	return s.String()
}

// Niceness - shorter solutions are better
// solutions using division are less nice
func (s Solution) Niceness() int {
	niceness := 0
	for _, operation := range s.OperationSequence {
		if operation == NoOp {
			niceness++
		}
		if operation == Div {
			niceness--
		}
	}
	return niceness
}
