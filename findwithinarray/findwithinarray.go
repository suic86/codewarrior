package findwithinarray

func FindInArray(array []interface{}, predicate func(interface{}, int) bool) int {
	for i, v := range array {
		if predicate(v, i) {
			return i
		}
	}
	return -1
}
