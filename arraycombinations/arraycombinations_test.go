package arraycombinations_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/suic86/codewarrior/arraycombinations"
)

var _ = Describe("Sample Tests", func() {
	It("should work with sample tests", func() {
		Expect(Solve([][]int{{1, 2}, {4}, {5, 6}})).To(Equal(4))
		Expect(Solve([][]int{{1, 2}, {4, 4}, {5, 6, 6}})).To(Equal(4))
		Expect(Solve([][]int{{1, 2}, {3, 4}, {5, 6}})).To(Equal(8))
		Expect(Solve([][]int{{1, 2, 3}, {3, 4, 6, 6, 7}, {8, 9, 10, 12, 5, 6}})).To(Equal(72))
	})
})

var _ = Describe("Fixed Tests", func() {
	It("should work with fixed tests", func() {
		Expect(Solve([][]int{{1, 2}, {4}, {5, 6}})).To(Equal(4))
		Expect(Solve([][]int{{1, 2}, {4, 4}, {5, 6, 6}})).To(Equal(4))
		Expect(Solve([][]int{{1, 2}, {3, 4}, {5, 6}})).To(Equal(8))
		Expect(Solve([][]int{{1, 2, 3}, {3, 4, 6, 6, 7}, {8, 9, 10, 12, 5, 6}})).To(Equal(72))
	})
})

var _ = Describe("Random Tests", func() {
	It("should work with random tests", func() {
		for range 100 {
			ra := randomArray()
			exp := solution(ra)
			Expect(Solve(ra)).To(Equal(exp))
		}
	})
})

func randomArray() [][]int {
	var r [][]int
	for i := 1; i < 1+rand.Intn(5); i++ {
		var a []int
		for j := 1; j < 1+rand.Intn(5); i++ {
			a = append(a, rand.Intn(10))
		}
		r = append(r, a)
	}
	return r
}

func solution(data [][]int) int {
	r := 1
	for _, v := range data {
		r *= uniqCount(v)
	}
	return r
}

func uniqCount(arr []int) int {
	m := map[int]bool{}
	var c int
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = true
		c++
	}
	return c
}
