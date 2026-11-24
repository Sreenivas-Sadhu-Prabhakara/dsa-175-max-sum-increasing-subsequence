package main

import "testing"

func TestMaxSumIS(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 101, 2, 3, 100, 4, 5}, 106}, // 1+2+3+100
		{[]int{3, 4, 5, 10}, 22},
		{[]int{10, 5, 4, 3}, 10}, // descending: best single element
		{[]int{1, 1, 1}, 1},      // strictness blocks chaining equals
		{[]int{}, 0},
		{[]int{-1, -2, 5}, 5}, // [5] alone beats -1+5: negative prefixes only hurt
	}
	for _, c := range cases {
		if got := MaxSumIS(c.in); got != c.want {
			t.Errorf("MaxSumIS(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
