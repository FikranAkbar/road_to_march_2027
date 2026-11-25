package main

import "testing"

// Spec: a majority element appears strictly MORE than n/2 times. Return that
// element, or -1 when no element qualifies. Elements are assumed non-negative,
// otherwise the -1 sentinel would be ambiguous with a real value.
// If the problem guarantees a majority always exists, the -1 cases below should
// be dropped before using this table as a score.

type majorityCase struct {
	name string
	in   []int
	want int
}

// solve is the only line you touch if the signature differs.
func solve(arr []int) int {
	return findMajorityElement(arr)
}

var majorityCases = []majorityCase{
	// examples / the caller you already have in main()
	{"statement example", []int{2, 2, 1, 1, 1, 2, 2}, 2},
	{"three elements", []int{3, 3, 4}, 3},
	{"own manual check", []int{1, 1, 2, 1, 3, 5, 1}, 1},

	// one element
	{"single element", []int{1}, 1},
	{"single zero", []int{0}, 0},

	// two elements
	{"two same", []int{1, 1}, 1},
	{"two different", []int{1, 2}, -1},

	// the exact-threshold trap: n/2 is NOT more than n/2
	{"exactly half each", []int{1, 1, 2, 2}, -1},
	{"exactly half, no tie", []int{1, 1, 2, 3}, -1},
	{"exactly half of six", []int{5, 5, 5, 1, 2, 3}, -1},
	{"exactly half of eight", []int{1, 1, 1, 1, 2, 2, 3, 4}, -1},
	{"half plus one of eight", []int{1, 1, 1, 1, 1, 2, 2, 3}, 1},
	{"two thirds of six", []int{2, 2, 2, 2, 3, 3}, 2},
	{"half plus one in an even array", []int{2, 2, 1, 1, 2}, 2},
	{"three of seven is not enough", []int{1, 1, 1, 2, 2, 2, 3}, -1},
	{"four of seven is enough", []int{2, 2, 2, 2, 1, 1, 1}, 2},
	{"four of seven, other side", []int{2, 2, 2, 1, 1, 1, 1}, 1},

	// odd lengths
	{"two of three", []int{1, 2, 2}, 2},
	{"three of five", []int{2, 2, 1, 1, 2}, 2},
	{"four of seven scattered", []int{1, 2, 1, 3, 1, 4, 1}, 1},
	{"four of seven scattered, larger value", []int{4, 4, 4, 4, 1, 2, 3}, 4},

	// no majority
	{"all distinct", []int{1, 2, 3}, -1},
	{"five distinct", []int{1, 2, 3, 4, 5}, -1},
	{"mostly one value but not a majority", []int{1, 1, 2, 3, 4}, -1},

	// single value repeated
	{"all identical", []int{5, 5, 5, 5}, 5},
	{"all identical, three", []int{7, 7, 7}, 7},

	// zeros and larger values
	{"zeros as the majority", []int{0, 0, 0, 1}, 0},
	{"large value majority", []int{9, 9, 1, 9, 2, 9, 3}, 9},
}

func TestFindMajorityElement(t *testing.T) {
	for _, c := range majorityCases {
		t.Run(c.name, func(t *testing.T) {
			if got := solve(c.in); got != c.want {
				t.Errorf("findMajorityElement(%v) = %d, want %d", c.in, got, c.want)
			}
		})
	}
}
