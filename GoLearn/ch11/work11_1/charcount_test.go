package main

import (
	"bufio"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharcount(t *testing.T) {
	input := "Hello, 世界！"
	reader := strings.NewReader(input)
	bufReader := bufio.NewReader(reader)
	expectedCounts := map[rune]int{
		'H': 1,
		'e': 1,
		'l': 2,
		'o': 1,
		',': 1,
		' ': 1,
		'世': 1,
		'界': 1,
		'！': 1,
	}
	expectedUTFlen := [utf8.UTFMax + 1]int{1: 7, 2: 0, 3: 3, 4: 0}

	counts, utflen, invalid := Charcount(bufReader)

	// Verify counts of Unicode characters
	for c, expectedCount := range expectedCounts {
		if counts[c] != expectedCount {
			t.Errorf("Unexpected count for character %q. Expected: %d, Got: %d", c, expectedCount, counts[c])
		}
	}

	// Verify lengths of UTF-8 encodings
	for i, expectedLen := range expectedUTFlen {
		if utflen[i] != expectedLen {
			t.Errorf("Unexpected count for length %d. Expected: %d, Got: %d", i, expectedLen, utflen[i])
		}
	}

	// Verify count of invalid UTF-8 characters
	if invalid != 0 {
		t.Errorf("Unexpected count of invalid UTF-8 characters. Expected: 3, Got: %d", invalid)
	}
}
