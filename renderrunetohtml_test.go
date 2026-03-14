package rtxt

import (
	"testing"
)

func TestRenderRuneToHTML(t *testing.T) {
	tests := []struct {
		Name     string
		Rune     rune
		Initial  []byte
		Expected string
	}{
		{
			Name:     "ampersand",
			Rune:     '&',
			Expected: "&amp;",
		},
		{
			Name:     "less-than",
			Rune:     '<',
			Expected: "&lt;",
		},
		{
			Name:     "greater-than",
			Rune:     '>',
			Expected: "&gt;",
		},

		{
			Name:     "letter a",
			Rune:     'a',
			Expected: "a",
		},
		{
			Name:     "letter Z",
			Rune:     'Z',
			Expected: "Z",
		},
		{
			Name:     "digit 0",
			Rune:     '0',
			Expected: "0",
		},
		{
			Name:     "space",
			Rune:     ' ',
			Expected: " ",
		},
		{
			Name:     "single quote",
			Rune:     '\'',
			Expected: "'",
		},
		{
			Name:     "double quote",
			Rune:     '"',
			Expected: "\"",
		},
		{
			Name:     "newline",
			Rune:     '\n',
			Expected: "\n",
		},
		{
			Name:     "tab",
			Rune:     '\t',
			Expected: "\t",
		},

		// Unicode
		{
			Name:     "persian letter",
			Rune:     'د',
			Expected: "د",
		},
		{
			Name:     "korean letter",
			Rune:     '안',
			Expected: "안",
		},
		{
			Name:     "bullet U+2022",
			Rune:     '•',
			Expected: "•",
		},
		{
			Name:     "emoji",
			Rune:     '🍎',
			Expected: "🍎",
		},

		// Append to existing buffer
		{
			Name:     "append ampersand to existing",
			Rune:     '&',
			Initial:  []byte("hello"),
			Expected: "hello&amp;",
		},
		{
			Name:     "append less-than to existing",
			Rune:     '<',
			Initial:  []byte("hello"),
			Expected: "hello&lt;",
		},
		{
			Name:     "append greater-than to existing",
			Rune:     '>',
			Initial:  []byte("hello"),
			Expected: "hello&gt;",
		},
		{
			Name:     "append letter to existing",
			Rune:     'x',
			Initial:  []byte("hello"),
			Expected: "hellox",
		},
	}

	for testNumber, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := string(renderRuneToHTML(test.Initial, test.Rune))

			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual HTML from renderRuneToHTML is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("RUNE:     %q", test.Rune)
				return
			}
		})
	}
}
