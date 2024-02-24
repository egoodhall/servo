package textutil_test

import (
	"testing"

	"github.com/egoodhall/servo/internal/textutil"
)

type tc struct {
	name     string
	input    string
	expected string
}

func TestDedent(t *testing.T) {
	runTests(t, textutil.Dedent,
		tc{
			name:     "LeadingWhitespace",
			input:    "   abc\n  def",
			expected: " abc\ndef",
		},
	)
}

func TestTrimBlankLines(t *testing.T) {
	runTests(t, textutil.TrimBlankLines,
		tc{
			name:     "LeadingWhitespace",
			input:    "\n\n   a\n\n \t",
			expected: "   a",
		},
		tc{
			name:     "NoTrailingWhitespace",
			input:    "\n\n   a",
			expected: "   a",
		},
		tc{
			name:     "NoLeadingWhitespace",
			input:    "   a\n",
			expected: "   a",
		},
		tc{
			name:     "No whitespace",
			input:    "   a",
			expected: "   a",
		},
	)
}

func TestTrimDedent(t *testing.T) {
	runTests(t, textutil.TrimDedent,
		tc{
			name:     "IndentedAndLeadingWhitespace",
			input:    "\n   abc\n  def",
			expected: " abc\ndef",
		},
	)
}

func runTests(t *testing.T, fn func(string) string, cases ...tc) {
	t.Helper()
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := fn(testCase.input)
			if actual != testCase.expected {
				t.Fatalf("\n----- Input\n%s\n----- Expected\n%s\n----- Actual\n%s\n-----\n", testCase.input, testCase.expected, actual)
			}
		})
	}
}
