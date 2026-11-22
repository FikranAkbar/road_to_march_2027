package main

import "testing"

// Spec: input is lowercase English letters; return the first non-repeating
// character, or "$" when every character repeats at least twice.

type fnrCase struct {
	name string
	in   string
	want string
}

var fnrCases = []fnrCase{
	// the three examples from the problem statement
	{"example 1", "geeksforgeeks", "f"},
	{"example 2", "racecar", "e"},
	{"example 3", "aabbccc", "$"},

	// boundaries
	{"empty string", "", "$"},
	{"single char", "a", "a"},
	{"two same", "aa", "$"},
	{"all identical", "zzzz", "$"},

	// traps: a repeat whose rightmost occurrence is the leftmost one
	{"abba", "abba", "$"},
	{"aabb", "aabb", "$"},
	{"abcabc", "abcabc", "$"},

	// the unique character sits late
	{"unique at the end", "aabbc", "c"},
	{"mixed repeats", "zzzzab", "a"},

	// classic interview strings
	{"leetcode", "leetcode", "l"},
	{"loveleetcode", "loveleetcode", "v"},
	{"first char is the answer", "abcabd", "c"},
}

func TestFindFirstNonRepeatingChar(t *testing.T) {
	for _, c := range fnrCases {
		t.Run(c.name, func(t *testing.T) {
			if got := findFirstNonRepeatingChar(c.in); got != c.want {
				t.Errorf("findFirstNonRepeatingChar(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
