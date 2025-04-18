package orderedcount_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/suic86/codewarrior/orderedcount"
)

var _ = Describe("Test Suite", func() {
	It("Fixed Tests", func() {
		Expect(OrderedCount("abracadabra")).Should(Equal([]Tuple{{'a', 5}, {'b', 2}, {'r', 2}, {'c', 1}, {'d', 1}}))
		Expect(OrderedCount("Code Wars")).Should(Equal([]Tuple{{'C', 1}, {'o', 1}, {'d', 1}, {'e', 1}, {' ', 1}, {'W', 1}, {'a', 1}, {'r', 1}, {'s', 1}}))
		Expect(OrderedCount("")).Should(Equal([]Tuple{}))
	})
	It("Random Tests", func() {
		var randomInput string
		for range 100 {
			randomInput = randomText()
			Expect(OrderedCount(randomInput)).Should(Equal(reference(randomInput)))
		}
	})
})

// Reference solution

func reference(text string) []Tuple {
	counts := make(map[rune]int)
	order := make([]rune, 0, len(text))

	for _, c := range text {
		v, ok := counts[c]
		if ok {
			counts[c] = v + 1
		} else {
			counts[c] = 1
			order = append(order, c)
		}
	}

	result := make([]Tuple, 0, len(order))
	for _, c := range order {
		result = append(result, Tuple{c, counts[c]})
	}

	return result
}

// Random input generator

func randomText() string {
	return randStringRunes(rand.Intn(1000))
}

// Random string generator source: https://stackoverflow.com/a/31832326

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~0123456789     ")
var length = len(letterRunes)

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(length)]
	}
	return string(b)
}
