package transporttime

import "testing"

func TestGetTime(t *testing.T) {

	cars := []float64{3, 2}
	var distance float64 = 11

	res := getTime(cars, distance)

	expect := 5.5

	if res != expect {
		if res != 2 {
			t.Errorf("error result:%f,expect %f", res, expect)
		}
	}
}

func TestGetTime1(t *testing.T) {

	cars := []float64{3, 2, 4, 5, 3, 4, 5, 6, 7, 8, 9, 3}
	var distance float64 = 21

	res := getTime(cars, distance)

	expect := 7.00

	if res != expect {
		if res != 2 {
			t.Errorf("error result:%f,expect %f", res, expect)
		}
	}
}

func TestGetTime2(t *testing.T) {

	cars := []float64{3, 2, 4, 5, 3, 4, 5, 6, 7, 8, 9, 4}
	var distance float64 = 21

	res := getTime(cars, distance)

	expect := 5.25

	if res != expect {
		if res != 2 {
			t.Errorf("error result:%f,expect %f", res, expect)
		}
	}
}

func TestGetTime3(t *testing.T) {

	cars := []float64{3, 2, 4, 5, 3, 4, 5, 6, 1, 8, 9, 4}
	var distance float64 = 21

	res := getTime(cars, distance)

	expect := 18.00

	if res != expect {
		if res != 2 {
			t.Errorf("error result:%f,expect %f", res, expect)
		}
	}
}
