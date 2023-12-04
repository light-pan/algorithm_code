package findnumber

import "testing"

func TestGetNumber(t *testing.T) {
	var n int64 = 10
	res := getNumber(n)
	expect := 12

	if res != int64(expect) {
		t.Errorf("error result:%d,expect %d", res, expect)
	}
}

func TestGetNumber1(t *testing.T) {
	var n int64 = 2
	res := getNumber(n)
	expect := 4

	if res != int64(expect) {
		t.Errorf("error result:%d,expect %d", res, expect)
	}
}

func TestGetNumber2(t *testing.T) {
	var n int64 = 3
	res := getNumber(n)
	expect := 5

	if res != int64(expect) {
		t.Errorf("error result:%d,expect %d", res, expect)
	}
}

func TestGetNumber4(t *testing.T) {
	var n int64 = 1000000000
	res := getNumber(n)
	expect := 1000000512

	if res != int64(expect) {
		t.Errorf("error result:%d,expect %d", res, expect)
	}
}

func TestGetNumber5(t *testing.T) {
	var n int64 = 1
	res := getNumber(n)
	expect := 2

	if res != int64(expect) {
		t.Errorf("error result:%d,expect %d", res, expect)
	}
}
