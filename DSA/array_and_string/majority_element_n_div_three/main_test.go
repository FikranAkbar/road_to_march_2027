package main

import (
	"slices"
	"testing"
)

// Spec: return every element that appears strictly more than n/3 times.
// At most two elements can qualify. Result order is not specified — map
// iteration is random in Go — so the tests sort both sides before comparing,
// and duplicate values in the result would still be caught by the length check.

type majThreeCase struct {
	name string
	in   []int
	want []int
}

// solve is the only line you touch if the signature differs.
func solve(arr []int) []int {
	return majorityElementNDivThree(arr)
}

var majThreeCases = []majThreeCase{
	// canonical examples
	{"statement example 1", []int{3, 2, 3}, []int{3}},
	{"own manual check", []int{2, 2, 3, 1, 3, 2, 1, 1}, []int{1, 2}},
	{"two elements qualify", []int{1, 2, 3, 1, 2, 3, 1, 2}, []int{1, 2}},

	// one element
	{"single element", []int{1}, []int{1}},
	{"single zero", []int{0}, []int{0}},

	// two elements
	{"two same", []int{1, 1}, []int{1}},
	{"two different", []int{1, 2}, []int{1, 2}},

	// three elements: n/3 = 1, so 2 occurrences are needed
	{"three distinct", []int{1, 2, 3}, []int{}},
	{"two of three", []int{1, 1, 2}, []int{1}},
	{"all same", []int{0, 0, 0}, []int{0}},

	// exact-threshold traps: exactly n/3 is not enough
	{"four distinct", []int{1, 2, 3, 4}, []int{}},
	{"two pairs", []int{1, 1, 2, 2}, []int{1, 2}},
	{"three pairs of six", []int{1, 1, 2, 2, 3, 3}, []int{}},
	{"pair of four", []int{0, 1, 0, 1}, []int{0, 1}},

	// odd lengths
	{"three of four", []int{1, 1, 1, 2}, []int{1}},
	{"two of five", []int{5, 5, 6, 6, 7}, []int{5, 6}},
	{"two of five, uneven", []int{1, 1, 2, 2, 3}, []int{1, 2}},

	// n/3 vs n/2 — three occurrences can qualify here even when the array is long
	{"three of six", []int{1, 1, 1, 2, 2, 3}, []int{1}},
	{"three of seven, second place", []int{2, 2, 1, 1, 1, 2, 2}, []int{1, 2}},
	{"three of seven, only one", []int{1, 2, 2, 3, 3, 3}, []int{3}},

	// no elements qualify
	{"six distinct", []int{1, 2, 3, 4, 5, 6}, []int{}},

	// one value everywhere
	{"all identical", []int{1, 1, 1, 1}, []int{1}},
}

func TestMajorityElementNDivThree(t *testing.T) {
	for _, c := range majThreeCases {
		t.Run(c.name, func(t *testing.T) {
			got := slices.Clone(solve(c.in)) // clone so sorting never touches the function's own slice
			want := slices.Clone(c.want)

			slices.Sort(got)
			slices.Sort(want)

			if !slices.Equal(got, want) {
				t.Errorf("majorityElementNDivThree(%v) = %v, want %v (any order)", c.in, got, want)
			}
		})
	}
}
