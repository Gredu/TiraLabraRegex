package main

func generateMachine(tokens []Token) State {
	var firstState State
	var currentState *State = &firstState

	for _, token := range tokens {

		var nextState State
		nextState.accept = false

		var transition Transition
		transition.state = &nextState
		transition.token = token

		currentState.transitions = []*Transition{&transition}
		currentState = &nextState

	}

	currentState.accept = true
	return firstState
}
