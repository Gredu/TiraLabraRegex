package main

import "fmt"

func main() {

	var token Token
	token.typeOperator = "literal"
	token.value = "t"

	evalToken(token)

}

func evalToken(t Token) {

	switch t.typeOperator {
	case "literal":
		// TODO
	case "meta":
		// TODO
	case "repeater":
		// TODO
	default:
		fmt.Println("wrong type operator")
	}

}

func generateMachine(regexp string) State {

	var regexpLength int = len([]rune(regexp))
	var firstState State
	var currentState *State = &firstState

	for i := 0; i < regexpLength; i++ {

		var token Token
		token.typeOperator = "literal"
		token.value = string(regexp[i])

		var nextState State
		if i >= regexpLength-1 {
			nextState.accept = true
		} else {
			nextState.accept = false
		}

		var transition Transition
		transition.state = &nextState
		transition.token = token

		currentState.transitions = []*Transition{&transition}
		currentState = &nextState

	}

	return firstState
}

type Token struct {
	typeOperator string
	value        string
}

type State struct {
	transitions []*Transition
	accept      bool
}

type Transition struct {
	state *State
	token Token
}
