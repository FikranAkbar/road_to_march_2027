package main

import "testing"

// Spec: array of positive, negative and zero ints; return the maximum product
// obtainable from any contiguous subarray holding at least one element.

type maxProdCase struct {
	name string
	in   []int
	want int
}

// solve is the only line you touch if the signature differs from the expected one.
func solve(arr []int) int {
	return maximumProductSubarray(arr)
}

var maxProdCases = []maxProdCase{
	// the three examples from the statement
	{"statement example 1", []int{-2, 6, -3, -10, 0, 2}, 180},
	{"statement example 2", []int{-1, -3, -10, 0, 6}, 30},
	{"statement example 3", []int{2, 3, 4}, 24},

	// one element
	{"single positive", []int{2}, 2},
	{"single negative", []int{-2}, -2},
	{"single zero", []int{0}, 0},

	// all positive
	{"ones only", []int{1, 1, 1}, 1},
	{"positive run", []int{1, 2, 3, 4}, 24},

	// only negatives: the answer depends on how many there are
	{"two negatives", []int{-1, -1}, 1},
	{"three negatives", []int{-1, -1, -1}, 1},
	{"three negatives, growing", []int{-4, -3, -2}, 12},
	{"five negatives", []int{-5, -4, -3, -2, -1}, 120},

	// zeros
	{"zeros only", []int{0, 0}, 0},
	{"zero then positive", []int{0, 2}, 2},
	{"zero at the end", []int{0, 0, 3}, 3},
	{"zero between negatives", []int{0, -2, 0}, 0},
	{"negatives split by a zero", []int{-2, -3, 0, -4}, 6},
	{"mixed around a zero", []int{-3, 0, 2, -4}, 2},

	// sign changes
	{"small dip", []int{2, -1, 2}, 2},
	{"negative in the middle", []int{3, -1, 4}, 4},
	{"two sign changes", []int{-2, 3, -4}, 24},
	{"positive run broken twice", []int{2, 3, -2, 4}, 6},
	{"four elements, two negatives", []int{-1, 2, -3, 4}, 24},
	{"five elements, mixed", []int{2, -5, -2, -4, 3}, 24},
	{"best segment before the zero", []int{7, -2, -3, 0, 5}, 42},
}

func TestMaximumProductSubarray(t *testing.T) {
	for _, c := range maxProdCases {
		t.Run(c.name, func(t *testing.T) {
			if got := solve(c.in); got != c.want {
				t.Errorf("maximumProductSubarray(%v) = %d, want %d", c.in, got, c.want)
			}
		})
	}
}

// Same table, run against the O(n) implementation. Kept separate so a failure
// names the implementation, not just the case.
func TestMaximumProductSubarrayOptimum(t *testing.T) {
	for _, c := range maxProdCases {
		t.Run(c.name, func(t *testing.T) {
			if got := maximumProductSubarrayOptimum(c.in); got != c.want {
				t.Errorf("maximumProductSubarrayOptimum(%v) = %d, want %d", c.in, got, c.want)
			}
		})
	}
}
