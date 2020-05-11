package multiplicationtable_test

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
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
		for i := 0; i < 100; i++ {
			randomInput := rand.Intn(100)
			expected := reference(randomInput)
			Expect(MultiplicationTable(randomInput)).To(Equal(expected))
		}
	})
})

func init() {
	rand.Seed(time.Now().UnixNano())
}

func reference(size int) [][]int {
	table := make([][]int, size)
	for i := range table {
		table[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			table[i][j] = (i + 1) * (j + 1)
		}
	}
	return table
}
