package main

import "testing"

func TestSum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{14, 14}, 28},
		test{[]int{4, 4}, 8},
		test{[]int{15, 14}, 29},
		test{[]int{2, 3, 6}, 11},
	}

	for _, ts := range tests {
		if v := sum(ts.data...); v != ts.answer {
			t.Error("Expected", ts.answer, "Got", v)

		}
	}

}

func TestMul(t *testing.T) {
	if v := mul(7, 7); v != 49 {
		t.Error("Expected", 49, "Got", v)

	}

}
