package main

import "fmt"

func generateMachine(tokens []Token) State {

	var firstState State
	var currentState *State = &firstState

	for _, token := range tokens {
		currentState = generateTransition(token, currentState)
	}

	return firstState
}

func generateTransition(token Token, currentState *State) *State {

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
