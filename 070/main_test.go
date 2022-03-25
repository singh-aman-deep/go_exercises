package main

import "testing"

func TestSum(t *testing.T) {
	if v := sum(7, 7); v != 14 {
		t.Error("Expected", 14, "Got", v)

	}

}

func TestMul(t *testing.T) {
	if v := mul(7, 7); v != 49 {
		t.Error("Expected", 49, "Got", v)

	}

}
