package stringMatching

func KmpMatch(text string, pattern string) float32 {
	// Algorithm for Kmp String Matching
	var n int = len(text)
	var m int = len(pattern)
	var lps []int = ComputeLps(pattern)

	var i int = 0
	var j int = 0
	var count float32 = 0.0
	var temp float32 = 0.0

	for i < n {
		if text[i] == pattern[j] {
			if j == m-1 {
				count = temp + 1
				return count / float32(m)
			}
			i++
			j++
			temp++
		} else if j > 0 {
			j = lps[j-1]
			temp = float32(j)
		} else {
			temp = 0
			i++
		}
		if temp >= count {
			count = temp
		}
	}
	return count / float32(m)
}

func ComputeLps(pattern string) []int {
	// Compute longest proper prefix which is also a suffix
	var m int = len(pattern)
	var fail []int = make([]int, m)
	fail[0] = 0

	var j int = 0
	var i int = 1
	for i < m {
		if pattern[i] == pattern[j] {
			j += 1
			fail[i] = j
			i += 1
		} else if j > 0 {
			j = fail[j-1]
		} else {
			fail[i] = 0
			i += 1
		}
	}
	return fail
}
