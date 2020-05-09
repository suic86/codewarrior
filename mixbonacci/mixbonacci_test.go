package mixbonacci_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/suic86/codewarrior/mixbonacci"
)

var _ = Describe("Test Suite", func() {
	It("Edge Cases", func() {
		Expect(Mixbonacci([]string{}, 10)).Should(Equal([]int64{}))
		Expect(Mixbonacci([]string{"fib"}, 0)).Should(Equal([]int64{}))
	})
	It("Single-element Patterns", func() {
		Expect(Mixbonacci([]string{"fib"}, 10)).Should(Equal([]int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}))
		Expect(Mixbonacci([]string{"pad"}, 10)).Should(Equal([]int64{1, 0, 0, 1, 0, 1, 1, 1, 2, 2}))
		Expect(Mixbonacci([]string{"jac"}, 10)).Should(Equal([]int64{0, 1, 1, 3, 5, 11, 21, 43, 85, 171}))
		Expect(Mixbonacci([]string{"pel"}, 10)).Should(Equal([]int64{0, 1, 2, 5, 12, 29, 70, 169, 408, 985}))
		Expect(Mixbonacci([]string{"tri"}, 10)).Should(Equal([]int64{0, 0, 1, 1, 2, 4, 7, 13, 24, 44}))
		Expect(Mixbonacci([]string{"tet"}, 10)).Should(Equal([]int64{0, 0, 0, 1, 1, 2, 4, 8, 15, 29}))
	})
	It("Multi-element Patterns", func() {
		Expect(Mixbonacci([]string{"fib", "tet"}, 10)).Should(Equal([]int64{0, 0, 1, 0, 1, 0, 2, 1, 3, 1}))
		Expect(Mixbonacci([]string{"jac", "jac", "pel"}, 10)).Should(Equal([]int64{0, 1, 0, 1, 3, 1, 5, 11, 2, 21}))
	})
	It("Random Tests", func() {
		for i := 0; i < 100; i++ {
			pattern, length := randomInput()
			expected := reference(pattern, length)
			Expect(Mixbonacci(pattern, length)).Should(Equal(expected))
		}
	})
})

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

func reference(pattern []string, length int) []int64 {
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

func randomInput() ([]string, int) {
	sequences := []string{"fib", "pad", "jac", "pel", "tri", "tet"}
	patternLength := rand.Intn(20)
	var pattern []string
	for i := 0; i < patternLength; i++ {
		pattern = append(pattern, sequences[rand.Intn(6)])
	}
	return pattern, rand.Intn(51)
}
