package main

import "testing"

type testvalues struct {
	tokens            []Token
	expectedValues    []string
	expectedOperators []string
}

var tests = []testvalues{
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

// func TestGenerateMachine {
//
// }
//
// type testvalues struct {
// 	machines []State
// 	expectedValues []State
// }
//
// type parsedtokens struct {
// 	tokens []Token
// }
//
// var tokens = []{
// 	{
// 		parseRegexp("abc")
// 	},
// }
//
// var tests = []testvalues{
// 	{
// 		generateMachine("abc"),
//
// 	}
// }
//
