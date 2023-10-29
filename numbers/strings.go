package numbers

import "fmt"

func (s Expression) String() string {
	return s.StringInline()
}

func (s Expression) StringInline() string {
	out := fmt.Sprintf("%d", s.Numbers[0])
	for i, op := range s.Operations {
		if op == NoOp {
			continue
		}
		number := s.Numbers[i+1]
		bracket := s.Brackets[i+1]
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

func PrintOperation(acc int, n int, op Operation) string {
	next, _ := Operate(acc, n, op)
	if op == NoOp {
		return ""
	}
	str := fmt.Sprintf("%d %s %d = %d\n", acc, op, n, next)
	return str
}

func (s Expression) Raw() string {
	return fmt.Sprintf("%v %v %v", s.Numbers, s.Operations, s.Brackets)
}

type EvaluateConfig struct {
	withWorkingOut bool
}
type EvaluateOption func(config *EvaluateConfig)

func WithWorkingOut(config *EvaluateConfig) {
	config.withWorkingOut = true
}
