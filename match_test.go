package main

import (
	"testing"
)

func TestMatch(t *testing.T) {

	type testValues struct {
		regexp         string
		inputs         []string
		expectedValues []bool
	}

	var tests = []testValues{
		{
			"abc",
			[]string{"", "a", "ab", "abc", "abcd"},
			[]bool{false, false, false, true, false},
		},
		{
			"ab*c",
			[]string{"", "a", "ab", "abc", "abcd", "bbc", "abbbbc", "abbb", "ab", "ac"},
			[]bool{false, false, false, true, false, false, true, false, false, true},
		},
		{
			"ab*cd*e",
			[]string{"", "a", "ab", "abc", "abcd", "bbc", "abbbbc", "abbb", "ab"},
			[]bool{false, false, false, false, false, false, false, false, false},
		},
		{
			"ab*cd*e",
			[]string{"abcde", "ace", "abbbbbbcdddddde", "acdddde", "abbbbbce", "abbbbbbc"},
			[]bool{true, true, true, true, true, false},
		},
		{
			"ab*c*d",
			[]string{"ad", "abd", "acd", "abbbbbbbccccccd", "aaacccbbddd", "bbbbbbcccc", "bc"},
			[]bool{true, true, true, true, false, false, false},
		},
		{
			"ab*",
			[]string{"a", "ab", "abbbbbbb", "bbbb", ""},
			[]bool{true, true, true, false, false},
		},
		{
			"a\\db",
			[]string{"ab", "a2b", "3", "a8b", "a3", "9b"},
			[]bool{false, true, false, true, false, false},
		},
		{
			"a\\d\\db",
			[]string{"a12b", "12", "ab"},
			[]bool{true, false, false},
		},
		{
			"a.b",
			[]string{"abc", "a1b", "abb", "axb", "a", "", "ab"},
			[]bool{false, true, true, true, false, false, false},
		},
		{
			".b",
			[]string{"b", "4b", "xb"},
			[]bool{false, true, true},
		},
		{
			"a.",
			[]string{"a", "abc", "ab", "a3"},
			[]bool{false, false, true, true},
		},
		{
			".",
			[]string{"a", "b", "", "4", "55"},
			[]bool{true, true, false, true, false},
		},
		{
			"ab+c",
			[]string{"abc", "abbbc", "ab"},
			[]bool{true, true, false},
		},
		{
			"ab?c",
			[]string{"ac", "abc", "abbc", "aabbcc", "acc", "acbc"},
			[]bool{true, true, false, false, false, false},
		},
	}

	for _, test := range tests {
		machine := generateMachine(parseRegexp(test.regexp))
		for i, _ := range test.inputs {
			if match(test.inputs[i], machine) != test.expectedValues[i] {
				if test.expectedValues[i] {
					t.Error("Word", test.inputs[i], "should match match with regexp", test.regexp)
				} else {
					t.Error("Word", test.inputs[i], "should not match with regexp", test.regexp)
				}
			}
		}
	}

}

func BenchmarkMatchStars(b *testing.B) {

	type testValues struct {
		regexp         string
		inputs         []string
		expectedValues []bool
	}

	var tests = []testValues{
		{
			"ab*c*d",
			[]string{"ad", "abd", "acd", "abbbbbbbccccccd", "aaacccbbddd", "bbbbbbcccc", "bc"},
			[]bool{true, true, true, true, false, false, false},
		},
	}

	for _, test := range tests {
		machine := generateMachine(parseRegexp(test.regexp))
		for i, _ := range test.inputs {
			match(test.inputs[i], machine)
		}
	}

}

// func Fib(n int) int {
// 	if n < 2 {
// 		return n
// 	}
// 	return Fib(n-1) + Fib(n-2)
// }
//
// func benchmarkFib(i int, b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		Fib(i)
// 	}
// }
//
// func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
// func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
// func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
// func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
// func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
// func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
