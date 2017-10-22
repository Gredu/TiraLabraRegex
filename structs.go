package trex

// Token is a single piece of regular expression. generateMachine() uses []Token for input to generate automata.
type Token struct {
	typeOperator string
	value        string
	token        *Token
}

// State is a state in automata.
type State struct {
	transitions []*Transition
	accept      bool
}

// Transition describes how to get to certain state with a certain input
type Transition struct {
	state *State
	token Token
}
