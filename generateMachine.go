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
	case "literal", "meta", "dot":
		transition := Transition{&State{}, token}

		currentState.transitions = append(currentState.transitions, &transition)
		currentState = transition.state

	case "star":
		transition := Transition{currentState, token}
		currentState.transitions = append(currentState.transitions, &transition)

	case "plus":

		nextTransition := Transition{&State{}, Token{"literal", token.value}}
		currentState.transitions = append(currentState.transitions, &nextTransition)
		currentState = nextTransition.state

		lastTransition := Transition{currentState, Token{"star", token.value}}
		currentState.transitions = append(currentState.transitions, &lastTransition)

	default:
		fmt.Println("wrong type operator")
	}

	return currentState
}
