package maxvowelsubstring

import "strings"

const vowel = "aeiouAEIOU"

func getMaxString(flaw int, s string) int {

	maxLen := 0
	left := -1
	right := 0
	currentFlaw := flaw
	for i := 0; i < len(s); i++ {
		str := string(s[i])
		if strings.Contains(vowel, str) {
			if left == -1 {
				// 只有重新计数的时候才重置瑕疵度
				left = i
				currentFlaw = flaw
			} else {
				right = i
			}
		} else {
			if currentFlaw == 0 && left != -1 {
				current := right - left + 1
				maxLen = max(maxLen, current)
				left = -1
			} else {
				currentFlaw--
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
