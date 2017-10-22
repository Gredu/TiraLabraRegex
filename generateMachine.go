/*
TiraLabraRegex is a regular expression intepreter
*/
package main

// generateMachine generates automata.
// It returns State which is the first State of the automata.
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

// generateTransition generates transitions for states.
// It returns a reference to the state which is also the 'current state'
func generateTransition(token Token, currentState *State, last bool) *State {

	switch token.typeOperator {
	case "literal", "meta", "dot":
		if len(currentState.transitions) > 0 && currentState.transitions[0].token.typeOperator == "questionmark" {
			transition := Transition{&State{accept: last}, token}

			currentState.transitions[0].state.transitions = append(currentState.transitions[0].state.transitions, &transition)
			currentState.transitions = append(currentState.transitions, &transition)
		} else {
			transition := Transition{&State{accept: last}, token}
			currentState.transitions = append(currentState.transitions, &transition)
			currentState = transition.state
		}

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

	case "questionmark":
		transition := Transition{&State{accept: last}, token}
		currentState.transitions = append(currentState.transitions, &transition)

	}

	return currentState
}
