package vowelharmonyeasy_test

import (
	"errors"
	"math/rand"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/suic86/codewarrior/vowelharmonyeasy"
)

var _ = Describe("Sample Tests", func() {
	It("should work with sample tests", func() {
		Expect(Dative("ablak")).To(Equal("ablaknak"))
		Expect(Dative("tükör")).To(Equal("tükörnek"))
		Expect(Dative("keret")).To(Equal("keretnek"))
		Expect(Dative("otthon")).To(Equal("otthonnak"))
		Expect(Dative("virág")).To(Equal("virágnak"))
		Expect(Dative("tett")).To(Equal("tettnek"))
		Expect(Dative("rokkant")).To(Equal("rokkantnak"))
		Expect(Dative("rossz")).To(Equal("rossznak"))
	})
})

var randomWords = strings.Fields("kalap ház tűz víz méz kéz ember rák máz üveg pohár húr gödör csűr lakás rokon")
var front = strings.Fields("terv kérvény vény kép hit tök őr füst űr")
var back = strings.Fields("rag tár kár zár gondnok mór mókus úr")
var words = append(append(randomWords, front...), back...)

var _ = Describe("Test Suite", func() {
	It("Fixed Tests", func() {
		Expect(Dative("ablak")).To(Equal("ablaknak"))
		Expect(Dative("tükör")).To(Equal("tükörnek"))
		Expect(Dative("keret")).To(Equal("keretnek"))
		Expect(Dative("otthon")).To(Equal("otthonnak"))
		Expect(Dative("virág")).To(Equal("virágnak"))
		Expect(Dative("tett")).To(Equal("tettnek"))
		Expect(Dative("rokkant")).To(Equal("rokkantnak"))
		Expect(Dative("rossz")).To(Equal("rossznak"))
	})
	It("Random Tests", func() {
		rand.Shuffle(len(words), func(i, j int) {
			words[i], words[j] = words[j], words[i]
		})
		for _, word := range words {
			expected, _ := reference(word)
			Expect(Dative(word)).To(Equal(expected))
		}
	})
})

func reference(word string) (string, error) {
	runes := []rune(word)
	length := len(runes)
	for i := 1; i <= length; i++ {
		switch runes[length-i] {
		case 'e', 'é', 'i', 'í', 'ö', 'ő', 'ü', 'ű':
			return word + "nek", nil
		case 'a', 'á', 'o', 'ó', 'u', 'ú':
			return word + "nak", nil
		}
	}
	return "", errors.New("invalid word")
}
