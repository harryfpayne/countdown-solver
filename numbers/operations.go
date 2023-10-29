package numbers

type Operation string

const (
	Plus  Operation = "+"
	Minus Operation = "-"
	Times Operation = "*"
	Div   Operation = "/"
	NoOp  Operation = "?"
)

var Operations = []Operation{NoOp, Plus, Minus, Times, Div}
var OperationsWithoutNoOp = []Operation{Plus, Minus, Times, Div}

// Operate takes an accumulator, an operand, and an operation, and returns the result of the operation.
// If operation is NoOp then the accumulator is returned.
func Operate(acc, a int, op Operation) (int, bool) {
	switch op {
	case Plus:
		return acc + a, true
	case Minus:
		return acc - a, true
	case Times:
		return acc * a, true
	case Div:
		if a == 0 || acc%a != 0 {
			return 0, false
		}
		return acc / a, true
	}
	return acc, true
}
