package numbers

type Bracket string

type BracketFull struct {
	Type                      Bracket
	CorrespondingBracketIndex int
}

const (
	BracketOpen  Bracket = "("
	BracketClose Bracket = ")"
	NoBracket    Bracket = ""
)

var Brackets = []Bracket{BracketOpen, BracketClose, NoBracket}

func IsValidBracketSequence(brackets []Bracket) bool {
	var stack []Bracket
	for _, bracket := range brackets {
		if bracket == BracketOpen {
			stack = append(stack, bracket)
		} else if bracket == BracketClose {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func GetBracketSequence(brackets []Bracket) []BracketFull {
	output := make([]BracketFull, len(brackets))
	var stack []BracketFull
	for i, bracket := range brackets {
		if bracket == BracketOpen {
			output[i] = BracketFull{
				Type:                      bracket,
				CorrespondingBracketIndex: -1,
			}
			stack = append(stack, BracketFull{
				Type:                      bracket,
				CorrespondingBracketIndex: i, // misusing this, it's storing the index of this bracket instead
			})
		} else if bracket == BracketClose {
			if len(stack) == 0 {
				panic("Invalid bracket sequence")
			}
			correspondingOpenBracketIndex := stack[len(stack)-1].CorrespondingBracketIndex
			stack = stack[:len(stack)-1]
			output[i] = BracketFull{
				Type:                      bracket,
				CorrespondingBracketIndex: correspondingOpenBracketIndex,
			}
			output[correspondingOpenBracketIndex].CorrespondingBracketIndex = i
		} else {
			output[i] = BracketFull{
				Type:                      bracket,
				CorrespondingBracketIndex: -1,
			}
		}
	}
	return output
}
