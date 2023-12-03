package splicingstring

import "testing"

func TestGetStringCount(t *testing.T) {
	s := "asdsdcc1212"
	length := 1

	res := getStringCount(s, length)

	if res != 0 {
		t.Errorf("error result:%d,expect 0", res)
	}

}

func TestGetStringCount1(t *testing.T) {
	s := "abc"
	length := 2

	res := getStringCount(s, length)

	if res != 6 {
		t.Errorf("error result:%d,expect 3", res)
	}

}
