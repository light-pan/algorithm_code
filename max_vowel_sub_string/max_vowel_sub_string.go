package maxvowelsubstring

import (
	"strings"
)

const vowel = "aeiouAEIOU"

func getMaxString(flaw int, s string) int {

	maxLen := 0
	m := len(s)
	for i := 0; i < m; i++ {
		str := string(s[i])
		if strings.Contains(vowel, str) {
			currentFlaw := flaw
			right := 0
			for j := i + 1; j < m; j++ {
				str1 := string(s[j])
				if strings.Contains(vowel, str1) {
					right = j
					// 如果从当前i一直到字符串结束都是满足条件的子串，那么i也不需要继续往下执行了，可以直接返回
					if j == m-1 {
						current := right - i + 1
						maxLen = max(maxLen, current)
						return maxLen
					}
				} else {
					if currentFlaw == 0 {
						current := right - i + 1
						maxLen = max(maxLen, current)
						break
					} else {
						currentFlaw--
					}
				}
			}
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
