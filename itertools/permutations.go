package itertools

// PermutationGenerator inspired by: https://stackoverflow.com/a/30230552
type PermutationGenerator[S ~[]E, E any] struct {
	arr S
	p   []int
}

func NewPermutationGenerator[S ~[]E, E any](elements S) PermutationGenerator[S, E] {
	return PermutationGenerator[S, E]{
		arr: elements,
		p:   make([]int, len(elements)),
	}
}

func (p *PermutationGenerator[S, E]) Next() bool {
	for i := len(p.p) - 1; i >= 0; i-- {
		if i == 0 || p.p[i] < len(p.p)-i-1 {
			p.p[i]++
			goto Check
		}
		p.p[i] = 0
	}

Check:
	return p.p[0] < len(p.p)
}

func (p *PermutationGenerator[S, E]) Get() S {
	result := make(S, len(p.arr))
	copy(result, p.arr)
	for i, v := range p.p {
		result[i],
			result[i+v] =
			result[i+v],
			result[i]
	}
	return result
}
