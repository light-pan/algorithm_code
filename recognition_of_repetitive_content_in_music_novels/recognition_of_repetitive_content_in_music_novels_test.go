package recognitionofrepetitivecontentinmusicnovels

import "testing"

func TestIsSimilarity(t *testing.T) {

	str1 := "林汉达上下五千年wu"
	str2 := "林汉达上下5千年伍"
	Similarities := []string{"五 5 ⑤ 伍 wu"}
	flag, res := isSimilarity(str1, str2, Similarities)
	t.Logf("相似:%v,字符对:%v", flag, res)
}

func TestUnionSimilarity(t *testing.T) {
	similarities := []string{
		"、 ；",
		"； ，",
		"wu 五",
	}
	res := unionSimilarity(similarities)

	t.Log(res)
}
