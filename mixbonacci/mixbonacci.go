package mixbonacci

type generator struct {
	state       []int
	computeNext func([]int) int
}

func (g *generator) next() int {
	r := g.state[0]
	n := g.computeNext(g.state)
	g.state = append(g.state[1:], n)
	return r
}

func Mixbonacci(pattern []string, length int) []int {
	if length == 0 || len(pattern) == 0 {
		return []int{}
	}
	generators := map[string]*generator{
		"fib": &generator{[]int{0, 1}, func(a []int) int { return a[0] + a[1] }},
		"pad": &generator{[]int{1, 0, 0}, func(a []int) int { return a[0] + a[1] }},
		"jac": &generator{[]int{0, 1}, func(a []int) int { return 2*a[0] + a[1] }},
		"pel": &generator{[]int{0, 1}, func(a []int) int { return a[0] + 2*a[1] }},
		"tri": &generator{[]int{0, 0, 1}, func(a []int) int { return a[0] + a[1] + a[2] }},
		"tet": &generator{[]int{0, 0, 0, 1}, func(a []int) int { return a[0] + a[1] + a[2] + a[3] }},
	}
	result := make([]int, 0, length)
	pLength := len(pattern)
	for i := 0; i < length; i++ {
		result = append(result, generators[pattern[i%pLength]].next())
	}
	return result
	return []int{}
}
