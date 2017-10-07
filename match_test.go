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
	}

	for _, test := range tests {
		machine := generateMachine(parseRegexp(test.regexp))
		for i, _ := range test.inputs {
			if match(test.inputs[i], machine) != test.expectedValues[i] {
				if test.expectedValues[i] {
					t.Error("Word", test.inputs[i], "should not match")
				} else {
					t.Error("Word", test.inputs[i], "should match")
				}
			}
		}
	}

}
