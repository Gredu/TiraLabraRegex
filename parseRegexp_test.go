package main

import "testing"

func TestParseRegexp(t *testing.T) {

	type testValues struct {
		tokens            []Token
		expectedValues    []string
		expectedOperators []string
	}

	var tests = []testValues{
		{
			parseRegexp("abc*de"),
			[]string{"a", "b", "c", "d", "e"},
			[]string{"literal", "literal", "star", "literal", "literal"},
		},
		{
			parseRegexp("abc"),
			[]string{"a", "b", "c"},
			[]string{"literal", "literal", "literal"},
		},
	}

	for _, test := range tests {
		for i := 0; i < len(test.tokens); i++ {
			if test.tokens[i].value != test.expectedValues[i] {
				t.Error("Expected", test.expectedValues[i], "but got", test.tokens[i].value)
			}
			if test.tokens[i].typeOperator != test.expectedOperators[i] {
				t.Error("Expected", test.expectedOperators[i], "but got", test.tokens[i].typeOperator)
			}
		}
	}
}
