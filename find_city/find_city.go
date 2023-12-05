package findcity

import (
	"math"
	"sort"
)

type unionFind struct {
	ids []int
}

func newUnionFind(n int) unionFind {
	uf := unionFind{}
	uf.ids = make([]int, n+1)

	for i := range uf.ids {
		uf.ids[i] = i
	}

	return uf
}

func (uf unionFind) union(i, j int) {
	ids := uf.ids
	idi := ids[i]
	idj := ids[j]

	for i := range ids {
		if ids[i] == idi {
			ids[i] = idj
		}
	}
}

func (uf unionFind) getCurrentDp() int {
	count := make(map[int]int)

	for i := 1; i < len(uf.ids); i++ {
		count[uf.ids[i]]++
	}

	var values []int
	for _, v := range count {
		values = append(values, v)
	}

	sort.Ints(values)

	return values[len(values)-1]

}

func getMinDp(relations [][]int, n int) []int {
	currentIds := []int{}
	minDp := math.MaxInt

	for i := 1; i <= n; i++ {

		uf := newUnionFind(n)
		for _, v := range relations {
			if v[0] == i || v[1] == i {
				continue
			}
			uf.union(v[0], v[1])
		}

		dp := uf.getCurrentDp()

		if dp == minDp {
			currentIds = append(currentIds, i)
		} else if dp < minDp {
			currentIds = nil
			currentIds = append(currentIds, i)
			minDp = dp
		}
	}

	return currentIds
}
