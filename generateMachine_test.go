package main

import "testing"

func TestGenerateMachine(t *testing.T) {

	var currentState = generateMachine(parseRegexp("abc"))

	testCheckHelper(t, currentState, false, "a", "literal")
	currentState = *currentState.transitions[0].state

	testCheckHelper(t, currentState, false, "b", "literal")
	currentState = *currentState.transitions[0].state

	testCheckHelper(t, currentState, true, "c", "literal")
}

func testCheckHelper(t *testing.T, s State, expectedBool bool, expectedTokenValue string, expectedTokenOperator string) {
	if s.accept {
		t.Error("Expected", false, "but got", s.accept)
	}
	if s.transitions[0].token.value != expectedTokenValue {
		t.Error("Expected", expectedTokenValue, "but got", s.transitions[0].token.value)
	}
	if s.transitions[0].token.typeOperator != expectedTokenOperator {
		t.Error("Expected", expectedTokenOperator, "but got", s.transitions[0].token.typeOperator)
	}
}
