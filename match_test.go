package main

import "testing"

func TestMatch(t *testing.T) {

	machine := generateMachine(parseRegexp("abc"))

	if match("", machine) {
		t.Error("Word '' should not match, but it did")
	}
	if match("a", machine) {
		t.Error("Word 'a' should not match, but it did")
	}
	if match("ab", machine) {
		t.Error("Word 'ab' should not match, but it did")
	}
	if !match("abc", machine) {
		t.Error("Word 'abc' should match, but did not")
	}
	if match("abcd", machine) {
		t.Error("Word 'abcd' should not match, but it did")
	}

}
