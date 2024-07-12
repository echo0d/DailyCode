package work115

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	// s, sep := "a:b:c", ":"
	// words := strings.Split(s, sep)
	tests := []struct {
		s    string
		sep  string
		want []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a:b:c", ",", []string{"a:b:c"}},
		{"a:b:c", " ", []string{"a:b:c"}},
		{"abc", "", []string{"a", "b", "c"}},
		{"a b c ", " ", []string{"a", "b", "c", ""}},
		{"::::,,,", ",", []string{"::::", "", "", ""}},
	}

	for _, test := range tests {
		got := strings.Split(test.s, test.sep)
		want := test.want
		if len(got) != len(want) {

			t.Errorf("Split(%q, %q) returned %q words, want %q", test.s, test.sep, got, want)
		}
	}
}
