package main

import "fmt"

func generateMachine(tokens []Token) State {
	var firstState State
	var currentState *State = &firstState

	for _, token := range tokens {
		currentState = generateTransition(token, currentState)
	}

	currentState.accept = true
	return firstState
}

func generateTransition(token Token, currentState *State) *State {

	switch token.typeOperator {
	case "literal":
		nextState := State{accept: false}
		transition := Transition{&nextState, token}

		currentState.transitions = []*Transition{&transition}
		currentState = &nextState

	case "meta":
		// TODO
	case "star":
		// TODO
	default:
		fmt.Println("wrong type operator")
	}

	return currentState
}
