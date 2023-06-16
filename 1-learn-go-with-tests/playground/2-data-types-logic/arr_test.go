package main

import "testing"

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15
	if got != want {
		t.Errorf("got %d want %d given %v", got, want, nums)
	}

	nums = []int{1, 2, 3, 4}
	got = Sum(nums)
	want = 10
	if got != want {
		t.Errorf("got %d want %d given %v", got, want, nums)
	}
}
