package intelligentdriving

import (
	"math"
)

func gerOriginOli(cost [][]int) int {

	m := len(cost)

	n := len(cost[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt
		}
	}

	r := record{
		isOil:   false,
		realUse: 0,
		left:    0,
	}
	dfs(cost, 0, 0, dp, r)

	if dp[m-1][n-1] == math.MaxInt {
		return -1
	}

	return dp[m-1][n-1]
}

type record struct {
	isOil   bool //是否经历过加油站
	realUse int  //值为未经过加油站时的油耗[此值是到终点之后真正需要输出的值]
	left    int  //邮箱还剩多少油，指加过油之后
}

func dfs(cost [][]int, i, j int, dp [][]int, r record) {
	m := len(cost)
	n := len(cost[0])

	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}

	if cost[i][j] == 0 {
		return
	} else if cost[i][j] == -1 {
		r.isOil = true
		r.left = 100
	} else {
		if r.isOil {
			r.left -= cost[i][j]
		} else {
			r.realUse += cost[i][j]
		}
	}

	// 加满油这个点也走不完
	if r.isOil && r.left < 0 {
		return
	}

	// 不加油跑不到
	if !r.isOil && r.realUse > 100 {
		return
	}

	// 如果当前节点已经很小了，证明已经有路径走过了，不用再重复遍历了
	if r.realUse >= dp[i][j] {
		return
	}

	dp[i][j] = min(r.realUse, dp[i][j])

	dfs(cost, i-1, j, dp, r) //上
	dfs(cost, i+1, j, dp, r) //下
	dfs(cost, i, j-1, dp, r) //左
	dfs(cost, i, j+1, dp, r) //右
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
