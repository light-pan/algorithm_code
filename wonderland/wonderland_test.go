package wonderland

import "testing"

func TestGetCost(t *testing.T) {
	costs := []int{5, 14, 30, 100}
	days := []int{1, 2, 3, 15, 20, 21, 200, 202, 230}

	cost := getCost(costs, days)
	if cost != 44 {
		t.Errorf("error cost")
	}
}
