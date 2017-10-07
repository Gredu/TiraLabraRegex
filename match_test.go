package main

import "testing"

func TestMatch(t *testing.T) {

	type testValues struct {
		regexp         string
		inputs         []string
		expectedValues []bool
	}

	var tests = []testValues{
		{
			"abc",
			[]string{"", "a", "ab", "abc", "abcd"},
			[]bool{false, false, false, true, false},
		},
		{
			"ab*c",
			[]string{"", "a", "ab", "abc", "abcd", "bbc", "abbbbc", "abbb", "ab", "ac"},
			[]bool{false, false, false, true, false, false, true, false, false, true},
		},
		{
			"ab*cd*e",
			[]string{"", "a", "ab", "abc", "abcd", "bbc", "abbbbc", "abbb", "ab"},
			[]bool{false, false, false, false, false, false, false, false, false},
		},
		{
			"ab*cd*e",
			[]string{"abcde", "ace", "abbbbbbcdddddde", "acdddde", "abbbbbce", "abbbbbbc"},
			[]bool{true, true, true, true, true, false},
		},
	}

	for _, test := range tests {
		machine := generateMachine(parseRegexp(test.regexp))
		for i, _ := range test.inputs {
			if match(test.inputs[i], machine) != test.expectedValues[i] {
				if test.expectedValues[i] {
					t.Error("Word", test.inputs[i], "should match match with regexp", test.regexp)
				} else {
					t.Error("Word", test.inputs[i], "should not match with regexp", test.regexp)
				}
			}
		}
	}

}
