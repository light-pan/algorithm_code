package splicingstring

import (
	"regexp"
)

func getStringCount(s string, length int) int {
	m := len(s)
	if m == 0 || m > 30 {
		return 0
	}

	if length == 0 || length > 5 {
		return 0
	}

	reg := regexp.MustCompile("[^a-z]")
	if reg.FindString(s) != "" {
		return 0
	}

	dfs(s, "", length)
	return count
}

var count = 0

func dfs(s string, prevChar string, length int) {
	if length == 0 {
		count++
		return
	}

	if len(s) == 0 {
		return
	}

	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char == prevChar {
			continue
		}

		newStr := ""
		if i == 0 {
			newStr = s[1:]
		} else if i == len(s)-1 {
			newStr = s[0 : len(s)-1]
		} else {
			newStr = s[0:i] + s[i+1:len(s)]
		}
		dfs(newStr, char, length-1)
	}
}
