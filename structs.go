package main

type Token struct {
	typeOperator string
	value        string
	token        *Token
}

type State struct {
	transitions []*Transition
	accept      bool
}

type Transition struct {
	state *State
	token Token
}
