package functions

import "testing"

func TestIsLetter(t *testing.T) {

	type pair struct {
		s string
		b bool
	}
	test := []pair{
		{"A", true},
		{"8", false},
		{"N", true},
		{",", false},
	}
	for _, v := range test {
		if v.b != IsLetter(v.s) {
			t.Error("For", v.s, "expected", v.b, "got:", IsLetter(v.s))
		}
	}
}

func TestIsNumeral(t *testing.T) {

	type pair struct {
		s string
		b bool
	}
	test := []pair{
		{"0", true},
		{"A", false},
		{"8", true},
		{",", false},
	}
	for _, v := range test {
		if v.b != IsNumeral(v.s) {
			t.Error("For", v.s, "expected", v.b, "got:", IsNumeral(v.s))
		}
	}
}

func TestStrToInt(t *testing.T) {

	type pair struct {
		s string
		i int
	}
	test := []pair{
		{"_", 1},
		{"_164", 164},
		{"_0808", 808},
		{"_0", 0},
		{"_179826", 179826},
	}
	for _, v := range test {
		if v.i != StrToInt(v.s) {
			t.Error("For", v.s, "expected", v.i, "got:", StrToInt(v.s))
		}
	}
}

func TestCharToInt(t *testing.T) {

	type pair struct {
		ch byte
		i  int
	}
	test := []pair{
		{'0', 0},
		{'9', 9},
		{'6', 6},
	}
	for _, v := range test {
		if v.i != CharToInt(v.ch) {
			t.Error("For", string(v.ch), "expected", v.i, "got:", CharToInt(v.ch))
		}
	}
}
