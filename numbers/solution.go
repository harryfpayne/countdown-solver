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
	return s.StringMultiLine()
}

func (s Solution) StringInline() string {
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
	return Print(s.NumberSequence, s.OperationSequence, s.BracketSequence, 0)
}

func Print(numbers []int, operations []Operation, brackets []BracketFull, indexOffset int) string {
	value := numbers[0]
	var output string
	var subStr string

	for i := 0; i < len(operations); i++ {
		op := operations[i]
		bracket := brackets[i+1]
		number := numbers[i+1]

		if bracket.Type == BracketOpen {
			closeBracketIndex := bracket.CorrespondingBracketIndex - indexOffset
			numbersSub := numbers[i+1 : closeBracketIndex+1]
			operationsSub := operations[i+1 : closeBracketIndex]
			bracketsSub := brackets[i+1 : closeBracketIndex+1]
			subExpr := Print(numbersSub, operationsSub, bracketsSub, i+1)
			subVal, _ := CalculateRec(numbersSub, operationsSub, bracketsSub, i+1)
			output += fmt.Sprintf("%s", subExpr)

			subStr, value = PrintOperation(value, subVal, op)
			output += fmt.Sprintf("%s", subStr)
			i = closeBracketIndex - 1
			continue
		}

		subStr, value = PrintOperation(value, number, op)
		output += fmt.Sprintf("%s", subStr)
	}
	return output
}

func PrintOperation(acc int, n int, op Operation) (string, int) {
	next, _ := Operate(acc, n, op)
	if op == NoOp {
		return "", next
	}
	str := fmt.Sprintf("%d %s %d = %d\n", acc, op, n, next)
	return str, next
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
