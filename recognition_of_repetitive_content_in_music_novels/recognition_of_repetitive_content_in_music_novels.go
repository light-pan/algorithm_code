package recognitionofrepetitivecontentinmusicnovels

import (
	"strings"
)

func isSimilarity(str1, str2 string, similarities []string) (bool, [][]string) {

	similarityFlag := false
	res := [][]string{}
	str := str1

	// 去除所有相同的字符
	for _, v := range str {
		s := string(v)
		if strings.Contains(str2, s) {
			str1 = strings.Replace(str1, s, " ", 1)
			str2 = strings.Replace(str2, s, " ", 1)
		}
	}
	str1 = strings.TrimSpace(str1)
	str1 = strings.ReplaceAll(str1, "  ", " ")
	str2 = strings.TrimSpace(str2)
	str2 = strings.ReplaceAll(str2, "  ", " ")
	str1s := strings.Split(str1, " ")
	str2s := strings.Split(str2, " ")

	unions := unionSimilarity(similarities)
	for i := 0; i < len(str1s); i++ {

		s1, ok := unions[str1s[i]]
		for ok {
			if s1 == str2s[i] {
				similarityFlag = true
				break
			}
			s1, ok = unions[s1]
		}

		res = append(res, []string{str1s[i], str2s[i]})
	}

	return similarityFlag, res
}

func unionSimilarity(similarities []string) map[string]string {
	unions := make(map[string]string)

	for i := 0; i < len(similarities); i++ {
		strs := strings.Split(similarities[i], " ")
		for j := 1; j < len(strs); j++ {
			unions[strs[j-1]] = strs[j]
		}
	}

	return unions
}
