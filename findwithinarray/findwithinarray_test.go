package findwithinarray_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/suic86/codewarrior/findwithinarray"
)

var _ = Describe("Sample Tests", func() {
	It("Integer functions", func() {
		Expect(FindInArray([]any{}, isEven)).Should(Equal(-1))
		Expect(FindInArray([]any{1, 3, 5, 6, 7}, isEven)).Should(Equal(3))
		Expect(FindInArray([]any{2, 4, 6, 8}, isEven)).Should(Equal(0))
		Expect(FindInArray([]any{2, 4, 6, 8}, neverTrue)).Should(Equal(-1))
		Expect(FindInArray([]any{13, 5, 3, 1, 4, 5}, isValueEqualToIndex)).Should(Equal(4))
		Expect(FindInArray([]any{-3, -9, 8, 7, -9, 7, 4, 7, 1}, isValueEqualToIndex)).Should(Equal(7))
	})
	It("String Functions", func() {
		Expect(FindInArray([]any{"one", "two", "three", "four", "five", "six"}, isLengthEqualToIndex)).Should(Equal(4))
		Expect(FindInArray([]any{"bc", "af", "d", "e"}, isLengthEqualToIndex)).Should(Equal(-1))
	})
})

var _ = Describe("Fixed Tests", func() {
	It("Integer functions", func() {
		Expect(FindInArray([]any{}, isEven)).Should(Equal(-1))
		Expect(FindInArray([]any{1, 3, 5, 6, 7}, isEven)).Should(Equal(3))
		Expect(FindInArray([]any{2, 4, 6, 8}, isEven)).Should(Equal(0))
		Expect(FindInArray([]any{2, 4, 6, 8}, neverTrue)).Should(Equal(-1))
		Expect(FindInArray([]any{13, 5, 3, 1, 4, 5}, isValueEqualToIndex)).Should(Equal(4))
		Expect(FindInArray([]any{-3, -9, 8, 7, -9, 7, 4, 7, 1}, isValueEqualToIndex)).Should(Equal(7))
	})
	It("String Functions", func() {
		Expect(FindInArray([]any{"one", "two", "three", "four", "five", "six"}, isLengthEqualToIndex)).Should(Equal(4))
		Expect(FindInArray([]any{"bc", "af", "d", "e"}, isLengthEqualToIndex)).Should(Equal(-1))
	})
})

var _ = Describe("Random Tests", func() {
	It("Random Tests", func() {
		for range 100 {
			a, p := randomInput()
			expected := solution(a, p)
			Expect(FindInArray(a, p)).Should(Equal(expected))
		}
	})
})

func init() {
	for i := 2; i < 8; i++ {
		integerPredicates = append(integerPredicates, isMultiplyOfN(i))
	}
}

func randomInput() ([]any, func(any, int) bool) {
	minLength, maxLength := 0, 30

	sps := len(stringPredicates)
	ips := len(integerPredicates)
	var arr []any
	var predicate func(any, int) bool

	if rand.Float32() < 0.8 {
		ri := randomIntegers(minLength, maxLength)
		arr = make([]any, len(ri))
		for i, v := range ri {
			arr[i] = v
		}
		predicate = integerPredicates[rand.Intn(ips)]
	} else {
		rw := randomWords(minLength, maxLength)
		arr = make([]any, len(rw))
		for i, w := range rw {
			arr[i] = w
		}
		predicate = stringPredicates[rand.Intn(sps)]
	}

	return arr, predicate
}

func solution(array []any, predicate func(any, int) bool) int {
	for i, v := range array {
		if predicate(v, i) {
			return i
		}
	}
	return -1
}

func randomIntegers(minLength int, maxLength int) []int {
	var ints []int
	for i := minLength; i < rand.Intn(maxLength); i++ {
		ints = append(ints, rand.Int())
	}
	return ints
}

func randomWords(minLength int, maxLength int) []string {
	var words []string
	for i := minLength; i < rand.Intn(maxLength); i++ {
		words = append(words, randStringRunes(1+rand.Intn(10)))
	}
	return words
}

// Random string generator source: https://stackoverflow.com/a/31832326

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var length = len(letterRunes)

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(length)]
	}
	return string(b)
}

// predicates

var stringPredicates []func(any, int) bool = []func(any, int) bool{
	alwaysTrue,
	neverTrue,
	isLengthEqualToIndex,
	isLengthGreaterThanIndex,
	isLengthLessThanIndex,
}

var integerPredicates []func(any, int) bool = []func(any, int) bool{
	alwaysTrue,
	neverTrue,
	isEven,
	isOdd,
	isValueEqualToIndex,
	isValueGreaterThanIndex,
}

func alwaysTrue(v any, index int) bool {
	return true
}

func neverTrue(v any, index int) bool {
	return false
}

// integer predicates

func isEven(v any, index int) bool {
	return int(v.(int))%2 == 0
}

func isOdd(v any, index int) bool {
	return int(v.(int))%2 == 1
}

func isValueGreaterThanIndex(v any, index int) bool {
	return v.(int) > index
}

func isValueEqualToIndex(v any, index int) bool {
	return v.(int) == index
}

func isMultiplyOfN(n int) func(any, int) bool {
	return func(v any, index int) bool { return v.(int)%n == 0 }
}

// string predicates

func isLengthEqualToIndex(v any, index int) bool {
	return len(v.(string)) == index
}

func isLengthGreaterThanIndex(v any, index int) bool {
	return len(v.(string)) > index
}

func isLengthLessThanIndex(v any, index int) bool {
	return len(v.(string)) < index
}
