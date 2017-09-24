package main

import "fmt"

func main() {

	var token Token
	token.typeOperator = "literal"
	token.value = "t"

	evalToken(token)

	var regexp = "foo*ba?"

	tokens := parseRegexp(regexp)

	fmt.Println(tokens)

}

func parseRegexp(regexp string) []Token {

	tokens := make([]Token, len(regexp))

	for i, j := 0, 0; i < len(tokens); i, j = i+1, j+1 {

		r := regexp[i]
		c := fmt.Sprintf("%c", r)

		switch c {
		case "*":
			tokens[j-1] = Token{typeOperator: "star", value: tokens[j-1].value}
			j--
		case ".":
			tokens[j] = Token{typeOperator: "dot", value: "."}
		case "?":
			tokens[j-1] = Token{typeOperator: "questionmark", value: tokens[j-1].value}
			j--
		default:
			tokens[j] = Token{typeOperator: "literal", value: c}
		}

	}
	return tokens
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
	var firstState State
	var currentState *State = &firstState

	for i := 0; i < len(regexp); i++ {

		var token Token
		token.typeOperator = "literal"
		token.value = string(regexp[i])

		var nextState State
		if i >= len(regexp)-1 {
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
