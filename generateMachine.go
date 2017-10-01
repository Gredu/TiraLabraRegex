package main

import "fmt"

func generateMachine(tokens []Token) State {

	lastToken := func(tokens []Token, i int) bool {
		if len(tokens) == i+1 {
			return true
		} else {
			return false
		}
	}

	var firstState State
	var currentState *State = &firstState

	for i, token := range tokens {
		currentState = generateTransition(token, currentState, lastToken(tokens, i))
	}

	// currentState.accept = true
	return firstState
}

func generateTransition(token Token, currentState *State, last bool) *State {

	switch token.typeOperator {
	case "literal":
		nextState := State{accept: last}
		transition := Transition{&nextState, token}

		currentState.transitions = append(currentState.transitions, &transition)
		currentState = &nextState

	case "meta":
		// TODO
	case "star":
		transition := Transition{currentState, token}
		currentState.accept = last
		currentState.transitions = append(currentState.transitions, &transition)

	default:
		fmt.Println("wrong type operator")
	}

	return currentState
}
