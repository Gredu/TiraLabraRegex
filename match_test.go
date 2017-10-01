package main

import "testing"

func TestMatch(t *testing.T) {

	machine := generateMachine(parseRegexp("abc"))

	if !match("abc", machine) {
		t.Error("Word 'abc' should match, but did not")
	}
	if match("abcd", machine) {
		t.Error("Word 'abcd' should not match, but it did")
	}

}
