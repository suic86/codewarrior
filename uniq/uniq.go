package uniq

func Uniq(a []string) []string {
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
