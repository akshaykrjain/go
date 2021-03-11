package main

import "testing"

func TestSum(t *testing.T) {
	got := add(5, 3)
	if got != 8 {
		t.Errorf("sum(5,3) = %d; expected 8", got)
	}
}
