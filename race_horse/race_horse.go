package racehorse

import (
	"sort"
)

func getOptimalNum(nums1, nums2 []int) int {

	// 先给两个数组排序
	sort.Ints(nums1)
	sort.Ints(nums2)

	m := len(nums1)

	// nums1最小值和最大值的索引
	left := 0
	right := m - 1

	// 定义一个最优解的数组一个组合,里面的元素满足a[i] > b[i]
	optimalNums := []int{}
	for i := m - 1; i >= 0; i-- {
		if nums2[i] < nums1[right] {
			optimalNums = append(optimalNums, nums1[right])
			right--
		} else {
			left++
			optimalNums = nil
		}
	}
	total := 0
	combs = [][]int{}

	getComb(optimalNums, nil)
	nums2 = nums2[:len(optimalNums)]
	for i := 0; i < len(combs); i++ {
		if isVliad(combs[i], nums2) {
			total++
		}
	}
	return total
}

func isVliad(nums1, nums2 []int) bool {
	for i := 0; i < len(nums1); i++ {
		if nums1[i] < nums2[i] {
			return false
		}
	}
	return true
}

// 数组的所有排列
var combs [][]int

// 回溯法列出所有的数组组合
func getComb(nums, comb []int) {
	if len(nums) == 0 {
		combs = append(combs, comb)
	}

	for i := 0; i < len(nums); i++ {
		nextNum := make([]int, len(nums)-1)
		if i == 0 {
			copy(nextNum, nums[1:])
		} else if i == len(nums)-1 {
			copy(nextNum, nums[:len(nums)-1])
		} else {
			temp := []int{}
			temp = append(temp, nums[0:i]...)
			temp = append(temp, nums[i+1:]...)
			copy(nextNum, temp)
		}

		getComb(nextNum, append(comb, nums[i]))
	}
}
