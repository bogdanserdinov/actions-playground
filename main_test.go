package main

import "testing"

func TestAdding(t *testing.T) {
	sum := 1 + 1
	if sum != 2 {
		t.Errorf("1 + 1 did not equal 2")
	}
}
