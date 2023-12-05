package findcity

import "testing"

func TestUnionFind(t *testing.T) {
	relations := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	uf := newUnionFind(5)
	for i := 0; i < len(relations); i++ {
		uf.union(relations[i][0], relations[i][1])
	}

	t.Log(uf.ids)
}

func TestGetMinDp(t *testing.T) {
	relations := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	res := getMinDp(relations, 5)
	expect := 3
	if res[0] != expect {
		t.Errorf("error result:%d,expect %d", res[0], expect)
	}
}
