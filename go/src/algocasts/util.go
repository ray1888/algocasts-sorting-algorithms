package algocasts

func swap(array []int, a int, b int) {
	tmp := array[b]
	array[b] = array[a]
	array[a] = tmp
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
