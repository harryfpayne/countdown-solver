package numbers

import "fmt"

type Solution struct {
	NumberSequence    []int
	OperationSequence []Operation
}

func NewSolution(numberSequence []int, operationSequence []Operation) Solution {
	return Solution{
		NumberSequence:    numberSequence,
		OperationSequence: operationSequence,
	}
}

func (s Solution) String() string {
	out := fmt.Sprintf("%d", s.NumberSequence[0])
	for i, op := range s.OperationSequence {
		if op == NoOp {
			continue
		}
		out = fmt.Sprintf("(%s %s %d )", out, op, s.NumberSequence[i+1])
	}
	return out
}

func (s Solution) StringMultiLine() string {
	numbers := s.NumberSequence
	ops := s.OperationSequence

	out := ""
	value := numbers[0]
	for i, op := range ops {
		if op == NoOp {
			continue
		}
		next, _ := Operate(value, numbers[i+1], op)
		out += fmt.Sprintf("%d %s %d = %d \n", value, op, numbers[i+1], next)
		value = next
	}
	return out
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

// Complexity - solutions which use higher numbers are better
func (s Solution) Complexity() int {
	highestValue := 0
	value := s.NumberSequence[0]
	for i, op := range s.OperationSequence {
		if op == NoOp {
			continue
		}
		value, _ = Operate(value, s.NumberSequence[i+1], op)
		if value > highestValue {
			highestValue = value
		}
	}
	return highestValue
}
