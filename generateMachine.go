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

	return firstState
}

func generateTransition(token Token, currentState *State, last bool) *State {

	switch token.typeOperator {
	case "literal":
		transition := Transition{&State{}, token}

		currentState.transitions = append(currentState.transitions, &transition)
		currentState = transition.state

	case "meta":
		transition := Transition{&State{}, token}

		currentState.transitions = append(currentState.transitions, &transition)
		currentState = transition.state

	case "star":
		transition := Transition{currentState, token}
		currentState.transitions = append(currentState.transitions, &transition)

	default:
		fmt.Println("wrong type operator")
	}

	return currentState
}
