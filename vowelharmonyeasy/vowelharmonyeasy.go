package vowelharmonyeasy

import "errors"

// Dative returns the dative case for a given Hungarian word.
func Dative(word string) (string, error) {
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
