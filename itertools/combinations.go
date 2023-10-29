package itertools

import (
	"fmt"
	"math"
	"strconv"
)

type CombinationGenerator[S ~[]E, E any] struct {
	elements S   // Available elements
	choose   int // How many elements to choose
	i        int
	limit    int
}

func NewCombinationGenerator[S ~[]E, E any](elements S, choose int) CombinationGenerator[S, E] {
	return CombinationGenerator[S, E]{
		elements: elements,
		i:        -1,
		choose:   choose,
		limit:    int(math.Pow(float64(len(elements)), float64(choose))),
	}
}

func (p *CombinationGenerator[S, E]) Next() bool {
	if p.i >= p.limit-1 {
		return false
	}
	p.i++
	return true
}

func (p *CombinationGenerator[S, E]) Get() S {
	if p.i == -1 {
		return nil
	}

	res := make(S, p.choose)
	iInBase := fmt.Sprintf("%0*s", p.choose, strconv.FormatInt(int64(p.i), len(p.elements)))
	//fmt.Println(iInBase)
	for i := 0; i < p.choose; i++ {
		if i >= len(iInBase) {
			res[i] = p.elements[0]
			continue
		}
		j, _ := strconv.Atoi(string(iInBase[i]))
		res[i] = p.elements[j]
	}
	return res
}

/*
	elements = [0, 1, 2]
 	choose = 4

	0 0 0 0
    0 0 0 1
    0 0 0 2
    0 0 1 0
    ...
*/
