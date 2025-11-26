package utils

import (
	"testing"

	"github.com/Genekkion/gogogadgets/pkg/test"
)

func TestSplitByColumns(t *testing.T) {
	type Data struct {
		name     string
		input    string
		expected []string
	}

	testData := []Data{
		{
			name:     "simple ASCII",
			input:    "abc",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "ANSI colored",
			input:    "\x1b[31ma\x1b[0mb", // red 'a', normal 'b'
			expected: []string{"\x1b[31ma", "\x1b[0mb"},
		},
		{
			name:     "emoji and ANSI",
			input:    "\x1b[32m😀\x1b[0mA", // green emoji + ASCII
			expected: []string{"\x1b[32m😀", "\x1b[0mA"},
		},
		{
			name:     "CJK wide char",
			input:    "你A",
			expected: []string{"你", "A"}, // '你' width 2, treated as one column
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			t.Parallel()

			got := SplitColumns(data.input)
			test.AssertEqual(t, "Unexpected length of output", len(data.expected), len(got))

			for i := range got {
				expected := data.expected[i]
				test.AssertEqual(t, "Unexpected column value", expected, got[i])
			}
		})
	}
}
