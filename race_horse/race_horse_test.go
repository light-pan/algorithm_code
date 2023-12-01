package racehorse

import "testing"

func TestGetOptimalNum(t *testing.T) {
	nums1 := []int{3, 5, 9, 15, 20, 37}
	nums2 := []int{1, 2, 4, 17, 29, 35}
	ans := 6
	total := getOptimalNum(nums1, nums2)

	if ans != total {
		t.Errorf("bad result")
	}
}

func TestGetOptimalNum1(t *testing.T) {
	nums1 := []int{11, 8, 20}
	nums2 := []int{10, 13, 7}
	ans := 1
	total := getOptimalNum(nums1, nums2)
	if ans != total {
		t.Errorf("bad result")
	}

}

func TestIsVliad(t *testing.T) {
	nums1 := []int{9, 15, 5, 20}
	nums2 := []int{1, 2, 13, 17}
	if isVliad(nums1, nums2) {
		t.Errorf("error result")
	}
}
