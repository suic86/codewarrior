package kata_test

import (
	. "github.com/suic86/codewarrior/kata"
	"math/rand"
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Suite", func() {
	It("Fixed Tests", func() {
		Expect(EncryptThis("")).Should(Equal(""))
		Expect(EncryptThis("A")).Should(Equal("65"))
		Expect(EncryptThis("A wise old owl lived in an oak")).Should(Equal("65 119esi 111dl 111lw 108dvei 105n 97n 111ka"))
		Expect(EncryptThis("The more he saw the less he spoke")).Should(Equal("84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp"))
		Expect(EncryptThis("The less he spoke the more he heard")).Should(Equal("84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare"))
		Expect(EncryptThis("Why can we not all be like that wise old bird")).Should(Equal("87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri"))
		Expect(EncryptThis("Thank you Piotr for all your help")).Should(Equal("84kanh 121uo 80roti 102ro 97ll 121ruo 104ple"))
	})
	It("Random Tests", func() {
		var randomInput string
		for i := 0; i < 100; i++ {
			randomInput = randomText()
			Expect(EncryptThis(randomInput)).Should(Equal(reference(randomInput)))
		}
	})
})

// Reference solution

func reference(text string) string {
	words := strings.Split(text, " ")
	var result []string
	var encrypted string
	for _, word := range words {
		if word == "" {
			encrypted = ""
		} else {
			encrypted = strconv.Itoa(int(word[0])) + swap(word[1:])
		}
		result = append(result, encrypted)
	}
	return strings.Join(result, " ")
}

func swap(text string) string {
	l := len(text) - 1
	if l == -1 {
		return ""
	}
	temp := []byte(text)
	temp[0], temp[l] = temp[l], temp[0]
	return string(temp)
}

// Random input generator
func randomText() string {
	var words []string
	for i := 0; i < rand.Intn(20); i++ {
		words = append(words, randStringRunes(1+rand.Intn(5)))
	}
	return strings.Join(words, " ")
}

// Random string generator source: https://stackoverflow.com/a/31832326

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var length = len(letterRunes)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(length)]
	}
	return string(b)
}
