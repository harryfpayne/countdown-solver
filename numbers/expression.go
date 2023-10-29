package numbers

import "fmt"

type Expression struct {
	Numbers    []int
	Operations []Operation
	Brackets   []BracketFull
	WorkingOut string
}

func NewExpression(
	numberSequence []int,
	operationSequence []Operation,
	bracketSequence []BracketFull,
) Expression {
	return Expression{
		Numbers:    numberSequence,
		Operations: operationSequence,
		Brackets:   bracketSequence,
	}
}

func (s *Expression) Evaluate(opts ...EvaluateOption) (int, bool) {
	cfg := EvaluateConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}

	value := s.Numbers[0]
	var ok bool

	for i := 0; i < len(s.Operations); i++ {
		op := s.Operations[i]
		bracket := s.Brackets[i+1]
		number := s.Numbers[i+1]

		if bracket.Type == BracketOpen {
			subExpr := s.BuildSubExpression(i)
			subVal, ok := subExpr.Evaluate(opts...)
			if cfg.withWorkingOut {
				s.WorkingOut += fmt.Sprintf("%s", subExpr.WorkingOut)
			}
			if !ok {
				return 0, false
			}

			if cfg.withWorkingOut {
				subStr := PrintOperation(value, subVal, op)
				s.WorkingOut += fmt.Sprintf("%s", subStr)
			}
			value, ok = Operate(value, subVal, op)
			if !ok { // Invalid operation (bad division), so skip to next operation set
				return 0, false
			}
			i = bracket.CorrespondingBracketIndex - 1
			continue
		}

		if cfg.withWorkingOut {
			subStr := PrintOperation(value, number, op)
			s.WorkingOut += fmt.Sprintf("%s", subStr)
		}
		value, ok = Operate(value, number, op)
		if !ok { // Invalid operation (bad division), so skip to next operation set
			return 0, false
		}
	}
	return value, true
}

func (s *Expression) BuildSubExpression(operationIndex int) Expression {
	bracketIndex := operationIndex + 1
	bracket := s.Brackets[bracketIndex]
	if bracket.Type != BracketOpen {
		fmt.Println(s.Raw(), bracketIndex)
		panic("Error: trying to build subexpression from non-open bracket")
	}
	closeBracketIndex := bracket.CorrespondingBracketIndex
	numbersSub := s.Numbers[bracketIndex : closeBracketIndex+1]
	operationsSub := s.Operations[bracketIndex:closeBracketIndex]
	bracketsSub := make([]BracketFull, closeBracketIndex-bracketIndex+1)
	copy(bracketsSub, s.Brackets[bracketIndex:closeBracketIndex+1])

	for i := range bracketsSub {
		// Correct indexes
		bracketsSub[i].CorrespondingBracketIndex = bracketsSub[i].CorrespondingBracketIndex - operationIndex
	}

	return NewExpression(numbersSub, operationsSub, bracketsSub)
}

// Niceness - shorter solutions are better
// solutions using division are less nice
func (s Expression) Niceness() int {
	niceness := 0
	for _, operation := range s.Operations {
		if operation == NoOp {
			niceness++
		}
		if operation == Div {
			niceness--
		}
	}
	return niceness
}
