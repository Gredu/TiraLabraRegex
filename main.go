package main

import "fmt"

func main() {

	var token Token
	token.typeOperator = "literal"
	token.value = "t"

	evalToken(token)

	var firstState = generateMachine("abc")
	fmt.Println("---")
	fmt.Println("firstState has value of accept: ", firstState.accept)
	fmt.Println(firstState)

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
	var currentState State
	currentState.accept = false

	var addressOfFirstState = &currentState

	for i := 0; i < regexpLength; i++ {

		var token Token
		token.typeOperator = "literal"
		token.value = string(regexp[i])

		var state State
		if i >= regexpLength-1 {
			state.accept = true
		} else {
			state.accept = false
		}

		var transition Transition
		transition.state = &state
		transition.token = token

		var transitions = []*Transition{&transition}

		currentState.transitions = transitions
		currentState = state

	}

	fmt.Println(firstState.accept)

	return *firstState
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
