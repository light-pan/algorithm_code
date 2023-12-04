package chinesewordsegmentationsimulator

import "testing"

func TestGetWords(t *testing.T) {
	s := "ilovechina"
	dictString := "i,love,china,ch,na,ve,lo,this,is,the,word"
	res := getWords(s, dictString)
	expect := "i,love,china"
	if res != expect {
		t.Errorf("error result:%s, expect %s", res, expect)
	}
}

func TestGetWords1(t *testing.T) {
	s := "ilovechina"
	dictString := "i,ilove,lo,love,ch,china,lovechina"
	res := getWords(s, dictString)
	expect := "ilove,china"
	if res != expect {
		t.Errorf("error result:%s, expect %s", res, expect)
	}
}
