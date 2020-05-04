package mixbonacci_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "codewarrior/mixbonacci"
)

var _ = Describe("Test Suite", func() {
	It("Edge Cases", func() {
		Expect(Mixbonacci([]string{}, 10)).Should(Equal([]int{}))
		Expect(Mixbonacci([]string{"fib"}, 0)).Should(Equal([]int{}))
	})
	It("Single-element Patterns", func() {
		Expect(Mixbonacci([]string{"fib"}, 10)).Should(Equal([]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}))
		Expect(Mixbonacci([]string{"pad"}, 10)).Should(Equal([]int{1, 0, 0, 1, 0, 1, 1, 1, 2, 2}))
		Expect(Mixbonacci([]string{"jac"}, 10)).Should(Equal([]int{0, 1, 1, 3, 5, 11, 21, 43, 85, 171}))
		Expect(Mixbonacci([]string{"pel"}, 10)).Should(Equal([]int{0, 1, 2, 5, 12, 29, 70, 169, 408, 985}))
		Expect(Mixbonacci([]string{"tri"}, 10)).Should(Equal([]int{0, 0, 1, 1, 2, 4, 7, 13, 24, 44}))
		Expect(Mixbonacci([]string{"tet"}, 10)).Should(Equal([]int{0, 0, 0, 1, 1, 2, 4, 8, 15, 29}))
	})
	It("Multi-element Patterns", func() {
		Expect(Mixbonacci([]string{"fib", "tet"}, 10)).Should(Equal([]int{0, 0, 1, 0, 1, 0, 2, 1, 3, 1}))
		Expect(Mixbonacci([]string{"jac", "jac", "pel"}, 10)).Should(Equal([]int{0, 1, 0, 1, 3, 1, 5, 11, 2, 21}))
	})
})
