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
			[]string{"a", "b", "", "d", "e"},
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
			if test.tokens[i].typeOperator != test.expectedOperators[i] {
				t.Error("Expected", test.expectedOperators[i], "but got", test.tokens[i].typeOperator)
			}
			if test.tokens[i].typeOperator == "star" {
				if test.tokens[i].token == nil {
					t.Error("Expected star token to have a token, but did not")
				}
			}
			if test.tokens[i].value != test.expectedValues[i] {
				t.Error("Expected", test.expectedValues[i], "but got", test.tokens[i].value)
			}
		}
	}

	// custom tests
	tokens := parseRegexp("ab*c")
	if len(tokens) != 3 {
		t.Error("Expected length of tokens to be 3 in regexp 'ab*c', but got", len(tokens))
	}

	tokens = parseRegexp("ab?c")
	if len(tokens) != 3 {
		t.Error("Expected length of tokens to be 3 in regexp 'ab?c', but got", len(tokens))
	}
	if tokens[0].value != "a" {
		t.Error("Expected value of first token value should be a, but is", tokens[0].value)
	}
	if tokens[1].value != "b" {
		t.Error("Expected value of second token value should be a, but is", tokens[1].value)
	}
	if tokens[1].typeOperator != "questionmark" {
		t.Error("Expected value of second token typeOperator should be a, but is", tokens[1].value)
	}
	if tokens[2].value != "c" {
		t.Error("Expected value of third token value should be a, but is", tokens[2].value)
	}

}
