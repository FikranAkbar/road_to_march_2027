package main

import "testing"

// Spec under test: ASCII, case-sensitive, punctuation counted as a character.
// If the GfG problem defines it differently (ignore spaces/case), change this
// table first — otherwise it grades the wrong contract.

type palindromeCase struct {
	name string
	word string
	want bool
}

func TestIsPalindrome(t *testing.T) {
	cases := []palindromeCase{
		// smallest boundaries
		{"empty string", "", true},
		{"single char", "a", true},
		{"two same", "aa", true},
		{"two different", "ab", false},

		// odd and even length
		{"odd length", "aba", true},
		{"even length", "abba", true},
		{"middle element is free", "abcba", true},

		// trap: ends match but the middle does not
		{"ends match, middle breaks it", "abca", false},
		{"near miss at 4th char", "abcda", false},

		// passes even if you compare the wrong indices -> keep it next to the traps above
		{"all identical", "aaaa", true},

		// digits and words
		{"digits palindrome", "12321", true},
		{"digits not palindrome", "12345", false},
		{"word", "level", true},
		{"long palindrome", "abcdefghijihgfedcba", true},
		{"case-sensitive near miss", "abcdefghijihgfedcbA", false},

		// whitespace: not alphanumeric, but this spec counts it as a character
		{"single space", " ", true},
		{"punctuation breaks it", "a.", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := isPalindrome(c.word); got != c.want {
				t.Errorf("isPalindrome(%q) = %v, want %v", c.word, got, c.want)
			}
		})
	}
}

// These cases are deliberately skipped: the current implementation compares bytes,
// so it needs a code change (runes / ToLower / filtering) first. Delete the t.Skip
// line to watch them go red.
func TestIsPalindrome_BeyondASCII(t *testing.T) {
	cases := []palindromeCase{
		{"unicode accent", "été", true},   // 5 bytes, 3 runes
		{"emoji", "😀😀", true},             // 8 bytes, 2 runes
		{"chinese", "人人为我我为人人", true},      // 6 runes, 18 bytes
		{"case insensitive", "Madam", true}, // needs ToLower
		{"punctuation ignored", "A man, a plan, a canal: Panama", true},
		{"phrase not palindrome", "race a car", false}, // negative control
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Skip("needs runes/ToLower/filtering — current implementation compares bytes")
			if got := isPalindrome(c.word); got != c.want {
				t.Errorf("isPalindrome(%q) = %v, want %v", c.word, got, c.want)
			}
		})
	}
}
