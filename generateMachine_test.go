package main

import "testing"

func TestGenerateMachine(t *testing.T) {

	var currentState State = generateMachine(parseRegexp("abc"))
	expectedValues := []string{"a", "b", "c"}

	for i := 0; i < 3; i++ {
		if i == 2 {
			if !currentState.accept {
				// t.Error("Expected last state's accept to be true, but got", currentState.accept)
			} else {
				if currentState.accept {
					t.Error("Expected", false, "but got", currentState.accept)
				}
			}
		}
		if currentState.transitions[0].token.value != expectedValues[i] {
			t.Error("Expected", expectedValues[i], "but got", currentState.transitions[0].token.value)
		}
		if currentState.transitions[0].token.typeOperator != "literal" {
			t.Error("Expected literal but got", currentState.transitions[0].token.typeOperator)
		}

		currentState = *currentState.transitions[0].state
	}
}
