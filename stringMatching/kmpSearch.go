package stringMatching

func KmpMatch(text string, pattern string) int {
	var n int = len(text)
	var m int = len(pattern)
	var fail []int = ComputeFail(pattern)

	var i int = 0
	var j int = 0

	for i < n {
		if text[i] == pattern[j] {
			if j == m-1 {
				return i - m + 1
			}
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			i++
		}
	}
	return -1
}

func ComputeFail(pattern string) []int {
	var m int = len(pattern)
	var fail []int = make([]int, m)
	fail[0] = 0

	var j int = 0
	var i int = 1
	for x := 1; x < m; x++ {
		if pattern[i] == pattern[j] {
			fail[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			fail[i] = 0
			i++
		}
	}
	return fail
}
