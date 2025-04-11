package uniq_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/suic86/codewarrior/uniq"
)

var _ = Describe("Test Suite", func() {
	It("Sample Tests", func() {
		Expect(Uniq([]string{"a", "a", "b", "b", "c", "a", "b", "c", "c"})).To(Equal([]string{"a", "b", "c", "a", "b", "c"}))
		Expect(Uniq([]string{"a", "a", "a", "b", "b", "b", "c", "c", "c"})).To(Equal([]string{"a", "b", "c"}))
		Expect(Uniq([]string{})).To(Equal([]string{}))
		Expect(Uniq([]string{"foo"})).To(Equal([]string{"foo"}))
		Expect(Uniq([]string{"bar"})).To(Equal([]string{"bar"}))
		Expect(Uniq([]string{""})).To(Equal([]string{""}))
		Expect(Uniq([]string{"a", "a"})).To(Equal([]string{"a"}))
	})

	It("Fixed Tests", func() {
		Expect(Uniq([]string{"a", "a", "b", "b", "c", "a", "b", "c", "c"})).To(Equal([]string{"a", "b", "c", "a", "b", "c"}))
		Expect(Uniq([]string{"a", "a", "a", "b", "b", "b", "c", "c", "c"})).To(Equal([]string{"a", "b", "c"}))
		Expect(Uniq([]string{})).To(Equal([]string{}))
		Expect(Uniq([]string{"foo"})).To(Equal([]string{"foo"}))
		Expect(Uniq([]string{"bar"})).To(Equal([]string{"bar"}))
		Expect(Uniq([]string{""})).To(Equal([]string{""}))
		Expect(Uniq([]string{"a", "a"})).To(Equal([]string{"a"}))
	})

	It("Random Tests", func() {
		var ri, expected []string
		for range 100 {
			ri = randomInput()
			expected = reference(ri)
			Expect(Uniq(ri), expected)
		}
	})
})

func randomInput() []string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	alength := len(alphabet)
	var result []string
	for range rand.Intn(15) {
		w := alphabet[rand.Intn(alength)]
		for range 1 + rand.Intn(5) {
			result = append(result, w)
		}
	}
	return result
}

func reference(a []string) []string {
	if len(a) == 0 {
		return []string{}
	}
	if len(a) == 1 {
		return a
	}
	result := make([]string, 0, len(a))
	current := a[0]
	for _, s := range a[1:] {
		if current != s {
			result, current = append(result, current), s
		}
	}
	if len(result) == 0 || current != result[len(result)-1] {
		result = append(result, current)
	}
	return result
}
