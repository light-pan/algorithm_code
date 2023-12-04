package intelligentdriving

import (
	"testing"
)

func TestGerOriginOli(t *testing.T) {
	cost := [][]int{
		{10, 20},
		{30, 40},
	}

	res := gerOriginOli(cost)
	expect := 70

	if res != expect {
		t.Errorf("error result:%d,expect %d", res, expect)
	}
}

func TestGerOriginOli1(t *testing.T) {
	cost := [][]int{
		{10, 0},
		{0, 40},
	}

	res := gerOriginOli(cost)
	expect := -1

	if res != expect {
		t.Errorf("error result:%d,expect %d", res, expect)
	}
}

func TestGerOriginOli2(t *testing.T) {
	cost := [][]int{
		{10, -1},
		{-1, 40},
	}

	res := gerOriginOli(cost)
	expect := 10

	if res != expect {
		t.Errorf("error result:%d,expect %d", res, expect)
	}
}

func TestGerOriginOli3(t *testing.T) {
	cost := [][]int{
		{20, 5, 10, 5, 50},
		{30, 20, 20, -1, 50},
		{20, 60, 50, 60, 50},
		{5, 60, 30, 40, 50},
		{-1, 5, 5, 5, 50},
	}
	res := gerOriginOli(cost)
	expect := 40

	if res != expect {
		t.Errorf("error result:%d,expect %d", res, expect)
	}
}
