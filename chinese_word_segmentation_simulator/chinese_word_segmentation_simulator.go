package chinesewordsegmentationsimulator

import (
	"regexp"
	"sort"
	"strings"
)

type wordStruct struct {
	word  string
	index int
}

var wordStrings []wordStruct

func getWords(s, dictString string) string {
	dictStrings := strings.Split(dictString, ",")
	dict := make(map[string]struct{})
	for _, v := range dictStrings {
		dict[v] = struct{}{}
	}
	wordStrings = []wordStruct{}
	// 根据字母顺序开始匹配字符串
	for i := 'a'; i <= 'z'; i++ {
		s = getWordByChar(string(i), s, dict)
	}

	sort.Slice(wordStrings, func(i, j int) bool {
		return wordStrings[i].index < wordStrings[j].index
	})

	var strs []string

	for i := 0; i < len(wordStrings); i++ {
		strs = append(strs, wordStrings[i].word)
	}
	return strings.Join(strs, ",")
}

func getWordByChar(char, s string, dict map[string]struct{}) string {
	reg := regexp.MustCompile(char)
	indexes := reg.FindAllIndex([]byte(s), len(s))
	m := len(s)
	for len(indexes) > 0 {
		for i := 0; i < len(indexes); i++ {
			index := indexes[i][0]
			for j := m; j >= index; j-- {
				word := s[index:j]
				if _, ok := dict[word]; ok {
					space := strings.Repeat(" ", len(word))
					if j == m {
						s = s[0:index] + space
					} else {
						s = s[0:index] + space + s[j:]
					}
					wordStrings = append(wordStrings, wordStruct{
						word:  word,
						index: index,
					})
				}
			}

		}

		newIndexes := reg.FindAllIndex([]byte(s), len(s))
		if len(newIndexes) == len(indexes) {
			return s
		} else {
			indexes = newIndexes
		}
	}

	return s
}
