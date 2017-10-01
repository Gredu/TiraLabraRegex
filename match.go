package main

var accepted bool = false

func match(input string, currentState State) bool {

	for _, transition := range currentState.transitions {
		switch transition.token.typeOperator {
		case "literal":
			if input[0:1] == transition.token.value {
				if len(input) == 1 && transition.state.accept {
					accepted = true
					return accepted
				}
				match(input[1:], *transition.state)
			}
		}
	}

	return accepted
}
