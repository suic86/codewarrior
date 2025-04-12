package findwithinarray

func FindInArray(array []any, predicate func(any, int) bool) int {
	for i, v := range array {
		if predicate(v, i) {
			return i
		}
	}
	return -1
}
