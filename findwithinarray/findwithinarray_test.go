package findwithinarray_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/suic86/codewarrior/findwithinarray"
)

var _ = Describe("Sample Tests", func() {
	It("Integer functions", func() {
		Expect(FindInArray([]interface{}{}, isEven)).Should(Equal(-1))
		Expect(FindInArray([]interface{}{1, 3, 5, 6, 7}, isEven)).Should(Equal(3))
		Expect(FindInArray([]interface{}{2, 4, 6, 8}, isEven)).Should(Equal(0))
		Expect(FindInArray([]interface{}{2, 4, 6, 8}, neverTrue)).Should(Equal(-1))
		Expect(FindInArray([]interface{}{13, 5, 3, 1, 4, 5}, isValueEqualToIndex)).Should(Equal(4))
		Expect(FindInArray([]interface{}{-3, -9, 8, 7, -9, 7, 4, 7, 1}, isValueEqualToIndex)).Should(Equal(7))
	})
	It("String Functions", func() {
		Expect(FindInArray([]interface{}{"one", "two", "three", "four", "five", "six"}, isLengthEqualToIndex)).Should(Equal(4))
		Expect(FindInArray([]interface{}{"bc", "af", "d", "e"}, isLengthEqualToIndex)).Should(Equal(-1))
	})
})

func isEven(v interface{}, index int) bool {
	return int(v.(int))%2 == 0
}

func alwaysTrue(v interface{}, index int) bool {
	return true
}

func neverTrue(v interface{}, index int) bool {
	return false
}

func isValueEqualToIndex(v interface{}, index int) bool {
	return v.(int) == index
}

func isLengthEqualToIndex(v interface{}, index int) bool {
	return len(v.(string)) == index
}
