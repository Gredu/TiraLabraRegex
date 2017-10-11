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
		transition := Transition{&State{accept: last}, token}

		currentState.transitions = append(currentState.transitions, &transition)
		currentState = transition.state

	case "star":
		s := Token{"star", token.token.value, nil}
		transition := Transition{currentState, s}
		currentState.accept = last
		currentState.transitions = append(currentState.transitions, &transition)

	case "plus":

		nextTransition := Transition{&State{}, Token{"literal", token.value, nil}}
		currentState.transitions = append(currentState.transitions, &nextTransition)
		currentState = nextTransition.state

		lastTransition := Transition{currentState, Token{"star", token.value, nil}}
		currentState.transitions = append(currentState.transitions, &lastTransition)

	default:
		fmt.Println("wrong type operator")
	}

	return currentState
}
