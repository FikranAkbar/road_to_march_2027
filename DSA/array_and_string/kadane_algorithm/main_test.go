package main

import "testing"

// Spec from the problem statement: find the contiguous subarray holding at least
// one element with the maximum sum, and return that sum. Input always has >= 1 element.

type kadaneCase struct {
	name string
	in   []int
	want int
}

// solve is the only line you touch if your function's signature differs.
// If yours returns the subarray instead of the sum, replace this body with
// return sum(maximumSubarraySum(arr)) and add a tiny sum helper.
func solve(arr []int) int {
	return maximumSubarraySum(arr)
}

var kadaneCases = []kadaneCase{
	// the three examples from the statement
	{"statement example 1", []int{2, 3, -8, 7, -1, 2, 3}, 11},
	{"statement example 2", []int{-2, -4}, -2},
	{"statement example 3", []int{5, 4, 1, 7, 8}, 25},

	// one element
	{"single positive", []int{1}, 1},
	{"single negative", []int{-5}, -5},
	{"single zero", []int{0}, 0},

	// arrays made only of one sign
	{"all positive", []int{1, 2, 3, 4, 5}, 15},
	{"all negative", []int{-1, -2, -3, -4}, -1},
	{"all negative, small values", []int{-3, -3, -1, -3}, -1},
	{"all zeros", []int{0, 0, 0}, 0},

	// two elements
	{"positive then negative", []int{1, -1}, 1},
	{"negative then positive", []int{-1, 1}, 1},
	{"zero then negative", []int{-1, 0}, 0},

	// mixed signs
	{"drop in the middle", []int{5, -10, 5}, 5},
	{"small dip between positives", []int{2, -1, 2}, 3},
	{"zeros around a negative", []int{0, -1, 0}, 0},
	{"large values around a tiny dip", []int{100, -1, 100}, 199},
	{"same value both ends", []int{7, -1, 7}, 13},
	{"positive run then dip", []int{3, -2, 5, -1}, 6},
	{"ends with the best element", []int{1, -2, 3}, 3},

	// the canonical multi-dip case
	{"alternating signs, several candidates", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
}

func TestMaximumSubarraySum(t *testing.T) {
	for _, c := range kadaneCases {
		t.Run(c.name, func(t *testing.T) {
			if got := solve(c.in); got != c.want {
				t.Errorf("maximumSubarraySum(%v) = %d, want %d", c.in, got, c.want)
			}
		})
	}
}
