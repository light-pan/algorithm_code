package parkvisitpath

import "testing"

func TestGetPaths(t *testing.T) {
	parks := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	res := getPaths(parks)
	expect := 2
	if res != expect {
		t.Errorf("error result:%d, expect %d", res, expect)
	}
}
