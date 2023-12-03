package maxvowelsubstring

import "testing"

func TestGetMaxString1(t *testing.T) {
	s := "asdbuiodevauufgh"
	flaw := 0
	maxLen := getMaxString(flaw, s)
	res := len("uio")

	if maxLen != res {
		t.Errorf("errer result: %d,expect %d", maxLen, res)
	}
}

func TestGetMaxString2(t *testing.T) {
	s := "asdbuiodevauufgh"
	flaw := 3
	maxLen := getMaxString(flaw, s)
	res := len("asdbuio")

	if maxLen != res {
		t.Errorf("errer result: %d,expect %d", maxLen, res)
	}
}
