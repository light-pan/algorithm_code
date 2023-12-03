package maxsubstring

import "testing"

func TestGetMaxsub1(t *testing.T) {
	s := "alolobo"
	l := getMaxsub(s)
	if l != 6 {
		t.Errorf("error result")
	}
}

func TestGetMaxsub2(t *testing.T) {
	s := "alolobox"
	l := getMaxsub(s)
	res := len("alolob")
	if l != res {
		t.Errorf("error result:%d,expect %d", l, res)
	}
}

func TestGetMaxsub3(t *testing.T) {
	s := "aloloboacnnmmdlsddddxaaaa"
	l := getMaxsub(s)
	res := len("loboacnnmmdlsdddd")
	if l != res {
		t.Errorf("error result:%d,expect %d", l, res)
	}
}
