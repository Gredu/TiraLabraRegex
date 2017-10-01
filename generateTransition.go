package main

import "fmt"

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
