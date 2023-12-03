package maxsubstring

import (
	"regexp"
	"strings"
)

func getMaxsub(s string) int {

	lFirstcharStrings := getFirstcharStrings("l", s)
	xFirstcharStrings := getFirstcharStrings("x", s)
	oFirstcharStrings := getFirstcharStrings("o", s)

	lEvenFlag := len(lFirstcharStrings)%2 == 0
	xEvenFlag := len(xFirstcharStrings)%2 == 0
	oEvenFlag := len(oFirstcharStrings)%2 == 0

	// 全是偶数，直接返回当前字符串长度
	if lEvenFlag && xEvenFlag && oEvenFlag {
		return len(s)
	}

	// 只有一个字符是奇数，这是只要去除那个字符就可以满足都是偶数的要求[因为字符串是个环]
	if (!lEvenFlag && xEvenFlag && oEvenFlag) ||
		(lEvenFlag && !xEvenFlag && oEvenFlag) ||
		(lEvenFlag && xEvenFlag && !oEvenFlag) {
		return len(s) - 1
	}

	// 下面三种是有2个字符为奇数的情况
	if !lEvenFlag && !xEvenFlag && oEvenFlag {
		return getTwocharMaxlength("l", lFirstcharStrings, "x", xFirstcharStrings)
	}

	if !lEvenFlag && xEvenFlag && !oEvenFlag {
		return getTwocharMaxlength("l", lFirstcharStrings, "o", oFirstcharStrings)
	}

	if lEvenFlag && !xEvenFlag && !oEvenFlag {
		return getTwocharMaxlength("o", oFirstcharStrings, "x", xFirstcharStrings)
	}

	// 最后剩下的就是三个字符都是奇数的情况
	return getThreecharMaxLength(lFirstcharStrings, xFirstcharStrings, oFirstcharStrings)
}

func getFirstcharStrings(char, s string) []string {
	reg := regexp.MustCompile(char)
	indexes := reg.FindAllIndex([]byte(s), len(s))

	var realIndexes []int
	for i := 0; i < len(indexes); i++ {
		realIndexes = append(realIndexes, indexes[i][0])
	}

	var firstcharStrings []string

	for _, i := range realIndexes {
		firstcharStrings = append(firstcharStrings, s[i:]+s[:i])
	}

	return firstcharStrings
}

func getTwocharMaxlength(char1 string, char1FirstStrings []string, char2 string, char2FirstStrings []string) int {
	maxLen := 0

	for _, str := range char1FirstStrings {
		currentLen := len(str) - strings.Index(str, char2) - 1
		maxLen = max(maxLen, currentLen)
	}

	for _, str := range char2FirstStrings {
		currentLen := len(str) - strings.Index(str, char1) - 1
		maxLen = max(maxLen, currentLen)
	}

	return maxLen
}

func getThreecharMaxLength(lFirstcharStrings, xFirstcharStrings, oFirstcharStrings []string) int {
	maxLen := 0
	l1 := getThreecharOneByOneMaxLength(lFirstcharStrings)
	l2 := getThreecharOneByOneMaxLength(xFirstcharStrings)
	l3 := getThreecharOneByOneMaxLength(oFirstcharStrings)
	maxLen = max(maxLen, l1)
	maxLen = max(maxLen, l2)
	maxLen = max(maxLen, l3)
	return maxLen
}

func getThreecharOneByOneMaxLength(firstcharStrings []string) int {
	maxLen := 0
	l := len(firstcharStrings[0])
	for _, str := range firstcharStrings {
		lCount := 0
		xCount := 0
		oCount := 0
		i := 0
		for lCount%2 == 0 || xCount%2 == 0 || oCount%2 == 0 {
			if str[i] == 'l' {
				lCount++
			} else if str[i] == 'x' {
				xCount++
			} else if str[i] == 'o' {
				oCount++
			}
			i++
		}
		currentLen := l - i
		maxLen = max(maxLen, currentLen)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
