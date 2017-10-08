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
	case "literal", "meta", "dot":
		transition := Transition{&State{}, token}

		currentState.transitions = append(currentState.transitions, &transition)
		currentState = transition.state
		// currentState = &nextState

	case "star":
		transition := Transition{currentState, token}
		// currentState.accept = last
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
