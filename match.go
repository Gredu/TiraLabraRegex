package main

func match(input string, currentState State) bool {

	if len(input) == 0 && len(currentState.transitions) == 0 {
		return true
	} else if len(input) > 0 && len(currentState.transitions) > 0 {
		for _, transition := range currentState.transitions {
			switch transition.token.typeOperator {
			case "literal":
				if input[0:1] == transition.token.value {
					return match(input[1:], *transition.state)
				}
			case "star":
				if input[0:1] == transition.token.value {
					return match(input[1:], *transition.state)
				}
			}
		}
	}

	return false
}
