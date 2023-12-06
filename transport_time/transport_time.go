package transporttime

import "math"

func getTime(cars []float64, distance float64) float64 {

	m := len(cars)
	lastIndex := m - 1
	lastSpeed := cars[lastIndex]
	maxTime := distance / lastSpeed

	for i := 0; i < m-1; i++ {
		if cars[i] >= lastSpeed {
			continue
		}

		// 最后一辆车追上当前车辆需要花费的时间
		t := cars[i] * float64(lastIndex-i) / (lastSpeed - cars[i])

		// 此时运行的总路程
		total := t * lastSpeed

		// 如果行程结束后才追上，就证明在行驶过程中不会阻塞最后车辆的通行
		if total >= distance {
			continue
		}
		totalTime := t + (distance-total)/cars[i]
		maxTime = math.Max(totalTime, maxTime)
	}

	return maxTime
}
