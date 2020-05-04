package kata

import (
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
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
