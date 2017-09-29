package main

import "testing"

func TestParseRegexp(t *testing.T) {
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

type testvalues struct {
	tokens            []Token
	expectedValues    []string
	expectedOperators []string
}

var tests = []testvalues{
	{
		parseRegexp("abc*defg"),
		[]string{"a", "b", "c", "d", "e", "f", "g"},
		[]string{"literal", "literal", "star", "literal", "literal", "literal", "literal"},
	},
	{
		parseRegexp("abc"),
		[]string{"a", "b", "c"},
		[]string{"literal", "literal", "literal"},
	},
}
