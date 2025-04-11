package multiplicationtable_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/suic86/codewarrior/multiplicationtable"
)

var _ = Describe("Sample Tests", func() {
	It("Sample Tests", func() {
		Expect(MultiplicationTable(0)).To(Equal([][]int{}))
		Expect(MultiplicationTable(1)).To(Equal([][]int{{1}}))
		Expect(MultiplicationTable(2)).To(Equal([][]int{{1, 2}, {2, 4}}))
		Expect(MultiplicationTable(3)).To(Equal([][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}))
	})
})

var _ = Describe("Test Suite", func() {
	It("Fixed Tests", func() {
		Expect(MultiplicationTable(0)).To(Equal([][]int{}))
		Expect(MultiplicationTable(1)).To(Equal([][]int{{1}}))
		Expect(MultiplicationTable(2)).To(Equal([][]int{{1, 2}, {2, 4}}))
		Expect(MultiplicationTable(3)).To(Equal([][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}))
	})
	It("Random Tests", func() {
		for range 100 {
			randomInput := rand.Intn(100)
			expected := reference(randomInput)
			Expect(MultiplicationTable(randomInput)).To(Equal(expected))
		}
	})
})

func reference(size int) [][]int {
	table := make([][]int, size)
	for i := range table {
		table[i] = make([]int, size)
	}

	for i := range size {
		for j := range size {
			table[i][j] = (i + 1) * (j + 1)
		}
	}
	return table
}
