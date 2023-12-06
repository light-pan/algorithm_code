package employeedispatch

import "testing"

func TestGetPaths(t *testing.T) {
	x, y, cntx, cnty := 2, 3, 3, 1

	res := getMin(x, y, cntx, cnty)
	expect := 5
	if res != expect {
		t.Errorf("error result:%d, expect %d", res, expect)
	}
}

func TestGetPaths1(t *testing.T) {
	x, y, cntx, cnty := 2, 3, 5, 6

	res := getMin(x, y, cntx, cnty)
	expect := 13
	if res != expect {
		t.Errorf("error result:%d, expect %d", res, expect)
	}
}

func TestGetPaths2(t *testing.T) {
	x, y, cntx, cnty := 2, 3, 8, 6

	res := getMin(x, y, cntx, cnty)
	expect := 16
	if res != expect {
		t.Errorf("error result:%d, expect %d", res, expect)
	}
}

func TestGetPaths3(t *testing.T) {
	x, y, cntx, cnty := 2, 3, 8, 7

	res := getMin(x, y, cntx, cnty)
	expect := 17
	if res != expect {
		t.Errorf("error result:%d, expect %d", res, expect)
	}
}
