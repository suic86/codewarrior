package mixbonacci

type generator struct {
	state       []int64
	computeNext func([]int64) int64
}

func (g *generator) next() int64 {
	r := g.state[0]
	n := g.computeNext(g.state)
	g.state = append(g.state[1:], n)
	return r
}

func Mixbonacci(pattern []string, length int) []int64 {
	if length == 0 || len(pattern) == 0 {
		return []int64{}
	}
	generators := map[string]*generator{
		"fib": {[]int64{0, 1}, func(a []int64) int64 { return a[0] + a[1] }},
		"pad": {[]int64{1, 0, 0}, func(a []int64) int64 { return a[0] + a[1] }},
		"jac": {[]int64{0, 1}, func(a []int64) int64 { return 2*a[0] + a[1] }},
		"pel": {[]int64{0, 1}, func(a []int64) int64 { return a[0] + 2*a[1] }},
		"tri": {[]int64{0, 0, 1}, func(a []int64) int64 { return a[0] + a[1] + a[2] }},
		"tet": {[]int64{0, 0, 0, 1}, func(a []int64) int64 { return a[0] + a[1] + a[2] + a[3] }},
	}
	result := make([]int64, 0, length)
	pLength := len(pattern)
	for i := 0; i < length; i++ {
		result = append(result, generators[pattern[i%pLength]].next())
	}
	return result
}
