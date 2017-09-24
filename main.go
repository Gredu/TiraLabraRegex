package main

import "fmt"

func main() {

	var token Token
	token.typeOperator = "literal"
	token.value = "t"

	evalToken(token)

	var regexp = "abc"

	tokens := parseRegexp(regexp)

	fmt.Println(tokens)

	fmt.Println("---")

	var machine = generateMachine(tokens)
	fmt.Println(machine.transitions[0].state.transitions[0].state.transitions[0].token.value)

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
		case "\\":
			tokens[j] = Token{typeOperator: "meta", value: fmt.Sprintf("%c", regexp[i+1])}
			i++
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
