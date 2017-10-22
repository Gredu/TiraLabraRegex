/*
TiraLabraRegex is a regular expression intepreter
*/
package main

import "fmt"

// parseRegexp parses the regular expression.
// It returns a slice of Tokens.
func parseRegexp(regexp string) []Token {

	tokens := make([]Token, len(regexp))
	sliceRight := 0

	for i, j := 0, 0; i < len(tokens); i, j = i+1, j+1 {

		r := regexp[i]
		c := fmt.Sprintf("%c", r)

		switch c {
		case "*":
			token := tokens[j-1]
			tokens[j-1] = Token{"star", "", &token}
			j--
			sliceRight++
		case ".":
			tokens[j] = Token{typeOperator: "dot", value: "."}
		case "?":
			tokens[j-1] = Token{typeOperator: "questionmark", value: tokens[j-1].value}
			j--
			sliceRight++
		case "+":
			tokens[j-1] = Token{typeOperator: "plus", value: tokens[j-1].value}
			j--
			sliceRight++
		case "\\":
			tokens[j] = Token{typeOperator: "meta", value: fmt.Sprintf("%c", regexp[i+1])}
			i++
			sliceRight++
		default:
			tokens[j] = Token{typeOperator: "literal", value: c}
		}

	}

	tokens = tokens[:len(tokens)-sliceRight]
	return tokens
}
