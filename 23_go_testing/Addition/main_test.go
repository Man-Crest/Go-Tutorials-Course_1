package main

import "testing"

func TestAddfunc(t *testing.T) {

	type test struct {
		data []int
		ans  int
	}

	tests := []test{
		test{[]int{2, 3}, 5},
		test{[]int{10, 20}, 30},
		test{[]int{8, 7, 9}, 24},
		test{[]int{4, 5, 6}, 15},
		test{[]int{-1, 0, 0, -1, -2, 5}, 1},
	}

	for _, val := range tests {
		x := Addfunc(val.data...)
		if x != val.ans {
			t.Error("Expected output is :", val.ans, "  But got :", x)
		}
	}
}
