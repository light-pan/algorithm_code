package minmatrixwidth

func getMinWidth(matrix [][]int, nums []int) int {
	n := len(matrix)

	if n == 0 {
		return -1
	}

	m := len(matrix[0])

	k := len(nums)

	//  k个nums最少需要遍历几列数据
	minWidth, mod := k/m, k%m
	if mod != 0 {
		minWidth += 1
	}
	for minWidth <= m {
		for j := 0; j+minWidth <= m; j++ {
			numSet := make(map[int]int)
			for o := 0; o < k; o++ {
				numSet[nums[o]]++
			}
			for i := 0; i < n; i++ {
				for l := j; l < j+minWidth; l++ {
					if v, ok := numSet[matrix[i][l]]; ok {
						if v == 1 {
							delete(numSet, matrix[i][l])
							if len(numSet) == 0 {
								return minWidth
							}
						} else {
							numSet[matrix[i][l]]--
						}

					}
				}
			}
		}
		minWidth++
	}

	return -1
}
