package minmatrixwidth

import "testing"

func TestGetMinWidth(t *testing.T) {
	martrix := [][]int{
		{1, 2, 2, 3, 1},
		{2, 3, 2, 3, 2},
	}

	nums := []int{1, 2, 3}

	res := getMinWidth(martrix, nums)

	if res != 2 {
		t.Errorf("error result:%d,expect 2", res)
	}
}

func TestGetMinWidth1(t *testing.T) {
	martrix := [][]int{
		{1, 2, 2, 3, 1},
		{2, 3, 2, 3, 2},
	}

	nums := []int{1, 2, 5}

	res := getMinWidth(martrix, nums)

	if res != -1 {
		t.Errorf("error result:%d,expect 2", res)
	}
}
