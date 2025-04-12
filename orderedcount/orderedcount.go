package orderedcount

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	counts := make(map[rune]int)
	order := make([]rune, 0, len(text))

	for _, c := range text {
		if _, ok := counts[c]; ok {
			counts[c] += 1
			continue
		}
		counts[c] = 1
		order = append(order, c)
	}

	result := make([]Tuple, 0, len(order))
	for _, c := range order {
		result = append(result, Tuple{c, counts[c]})
	}

	return result
}
