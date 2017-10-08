package main

func match(input string, currentState State) bool {

	if len(input) == 0 && currentState.accept {
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
			case "dot":
				return match(input[1:], *transition.state)
			case "meta":
				switch transition.token.value {
				case "d":
					if isDigit(input[0:1]) {
						return match(input[1:], *transition.state)
					}
				}
			}
		}
	}

	return false
}

func isDigit(input string) bool {
	switch input {
	case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
		return true
	}
	return false
}
