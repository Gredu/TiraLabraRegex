package main

type Token struct {
	typeOperator string
	value        string
}

type State struct {
	transitions []*Transition
}

type Transition struct {
	state *State
	token Token
}
