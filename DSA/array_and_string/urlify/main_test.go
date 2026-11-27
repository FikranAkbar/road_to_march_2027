package main

import "testing"

// Spec: replace every ASCII space with "%20". Nothing else is touched — tabs,
// newlines and non-ASCII characters pass through unchanged.

type urlifyCase struct {
	name string
	in   string
	want string
}

// solve is the only line you touch if the signature differs.
func solve(s string) string {
	return urlifyString(s)
}

var urlifyCases = []urlifyCase{
	// the caller in main() and the classic example
	{"own manual check", "i love programming", "i%20love%20programming"},
	{"classic example", "Mr John Smith", "Mr%20John%20Smith"},

	// no spaces at all
	{"empty string", "", ""},
	{"single word", "abc", "abc"},

	// spaces only
	{"single space", " ", "%20"},
	{"three spaces", "   ", "%20%20%20"},

	// position of the spaces
	{"leading space", " a", "%20a"},
	{"trailing space", "a ", "a%20"},
	{"space in the middle", "a b", "a%20b"},
	{"double space", "a  b", "a%20%20b"},
	{"spaces both ends", "  hi  ", "%20%20hi%20%20"},

	// words
	{"two words", "hello world", "hello%20world"},
	{"three words", "one two three", "one%20two%20three"},
	{"ends with a word and space", "done ", "done%20"},

	// characters that must NOT be replaced
	{"tab is not a space", "a\tb", "a\tb"},
	{"newline is not a space", "a\nb", "a\nb"},
	{"percent is untouched", "100%", "100%"},
	{"existing %20 stays as is", "a%20b", "a%20b"},

	// non-ASCII (the probe measures runes, not bytes)
	{"accents", "héllo wörld", "héllo%20wörld"},
	{"emoji", "😀 😀", "😀%20😀"},
}

func TestURLifyString(t *testing.T) {
	for _, c := range urlifyCases {
		t.Run(c.name, func(t *testing.T) {
			if got := solve(c.in); got != c.want {
				t.Errorf("urlifyString(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
