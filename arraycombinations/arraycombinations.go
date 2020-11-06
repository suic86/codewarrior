package arraycombinations

func Solve(data [][]int) int {
	r := 1
	for _, v := range data {
		r *= uniqCount(v)
	}
	return r
}

func uniqCount(arr []int) int {
	m := map[int]bool{}
	var c int
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = true
		c++
	}
	return c
}
