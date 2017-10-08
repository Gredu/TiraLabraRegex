package main

import "testing"

func TestGenerateMachine(t *testing.T) {

	var currentState State = generateMachine(parseRegexp("abc"))
	expectedValues := []string{"a", "b", "c"}

	for i := 0; i < 3; i++ {
		// if i == 2 {
		// 	if !currentState.transitions[0].state.accept {
		// 		t.Error("Expected last state's accept to be true, but got", currentState.accept)
		// 	} else {
		// 		if currentState.accept {
		// 			t.Error("Expected", false, "but got", currentState.accept)
		// 		}
		// 	}
		// }
		if currentState.transitions[0].token.value != expectedValues[i] {
			t.Error("Expected", expectedValues[i], "but got", currentState.transitions[0].token.value)
		}
		if currentState.transitions[0].token.typeOperator != "literal" {
			t.Error("Expected literal but got", currentState.transitions[0].token.typeOperator)
		}

		currentState = *currentState.transitions[0].state
	}

	currentState = generateMachine(parseRegexp("ab*c"))
	expectedValues = []string{"a", "b"}

	for i := 0; i < 2; i++ {
		if currentState.accept {
			t.Error("Expected", false, "but got", currentState.accept)
		}
		if currentState.transitions[0].token.value != expectedValues[i] {
			t.Error("Expected", expectedValues[i], "but got", currentState.transitions[0].token.value)
		}
		if i == 1 {
			if currentState.transitions[0].token.typeOperator != "star" {
				t.Error("Expected 'literal' but got", currentState.transitions[0].token.typeOperator)
			}
			if currentState.accept {
				t.Error("Expected state accept value to be false , but got", currentState.accept)
			}
			if len(currentState.transitions) != 2 {
				t.Error("Expected length of transitions to be 2, but got", len(currentState.transitions))
			}
			if !currentState.transitions[1].state.accept {
				t.Error("Expected state accept value to be true, but got", currentState.transitions[1].state.accept)
			}
		} else {
			if currentState.transitions[0].token.typeOperator != "literal" {
				t.Error("Expected literal but got", currentState.transitions[0].token.typeOperator)
			}
		}

		currentState = *currentState.transitions[0].state
	}

	currentState = generateMachine(parseRegexp("ab+c"))
	if currentState.transitions[0].token.value != "a" {
		t.Error("Expected token value to be a, but got", currentState.transitions[0].token.value)
	}

	currentState = *currentState.transitions[0].state
	if currentState.transitions[0].token.value != "b" {
		t.Error("Expected token value to be b, but got", currentState.transitions[0].token.value)
	}
	if currentState.transitions[0].token.typeOperator != "literal" {
		t.Error("Expected typeOperator value to be 'literal', but got", currentState.transitions[0].token.typeOperator)
	}
	if len(currentState.transitions) != 1 {
		t.Error("Expected length of transitions to be 1, but got", len(currentState.transitions))
	}

	currentState = *currentState.transitions[0].state
	if len(currentState.transitions) != 2 {
		t.Error("Expected length of transitions to be 2, but got", len(currentState.transitions))
	}
	if currentState.transitions[0].token.value != "b" {
		t.Error("Expected token value to be b, but got", currentState.transitions[0].token.value)
	}
	if currentState.transitions[0].token.typeOperator != "star" {
		t.Error("Expected token value to be 'star', but got", currentState.transitions[0].token.typeOperator)
	}
	if currentState.transitions[1].token.value != "c" {
		t.Error("Expected token value to be c, but got", currentState.transitions[1].token.value)
	}
	if currentState.transitions[1].token.typeOperator != "literal" {
		t.Error("Expected token value to be 'star', but got", currentState.transitions[1].token.typeOperator)
	}

}
