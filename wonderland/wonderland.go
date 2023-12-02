package wonderland

func getCost(costs, days []int) int {
	m := len(days)
	dp := make([]int, m+1)
	for i := 1; i <= m; i++ {
		// 默认增加一张一日票
		dp[i] = dp[i-1] + costs[0]
		for j := i - 1; j >= 0; j-- {
			// 计算前面的每天与当前值买各种类型的票的最小花费，超过30天或者一天内只有买一日票的可能，不用额外计算
			if days[i-1]-days[j] > 30 {
				continue
			} else if days[i-1]-days[j] > 7 {
				dp[i] = min(dp[i], dp[j]+costs[3])
			} else if days[i-1]-days[j] > 3 {
				dp[i] = min(dp[i], dp[j]+costs[2])
			} else if days[i-1]-days[j] > 1 {
				dp[i] = min(dp[i], dp[j]+costs[1])
			}
		}
	}
	return dp[m]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
