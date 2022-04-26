package bm

import (
	"fmt"
)

func Main() {
	fmt.Println(bmMatch("ABCD", "BC"))
}

func bmMatch(str string, pattern string) bool {
	last := buildLast(pattern)
	n := len(str)
	m := len(pattern)
	i := m - 1

	if i > n-1 {
		return false
	}

	j := m - 1
	for {
		if pattern[j] == str[i] {
			if j == 0 {
				return true
			} else {
				i--
				j--
			}
		} else {
			lo := last[str[i]]
			i = i + m - min(j, 1+lo)
			j = m - 1
		}
		if !(i <= n-1) {
			break
		}
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func buildLast(pattern string) [90]int {
	var last [90]int

	for i := 0; i < 90; i++ {
		last[i] = -1
	}

	for i := 0; i < len(pattern); i++ {
		last[pattern[i]] = i
	}
	return last
}
