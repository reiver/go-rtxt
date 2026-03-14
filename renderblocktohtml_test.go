package rtxt

import (
	"testing"
)

func TestRenderBlockToHTML(t *testing.T) {
	tests := []struct {
		Name     string
		Lines    string
		Initial  []byte
		Expected string
	}{
		// Empty
		{
			Name:     "empty string",
			Lines:    "",
			Expected: "",
		},

		// Single line, no EOL
		{
			Name:     "single line plain text",
			Lines:    "Hello world",
			Expected: "Hello world",
		},
		{
			Name:     "single line with bold",
			Lines:    "Hello **world**",
			Expected: "Hello <strong>world</strong>",
		},

		// Single line with trailing EOL
		{
			Name:     "single line with trailing LF",
			Lines:    "Hello world\n",
			Expected: "Hello world",
		},
		{
			Name:     "single line with trailing CRLF",
			Lines:    "Hello world\r\n",
			Expected: "Hello world",
		},
		{
			Name:     "single line with trailing U+0085",
			Lines:    "Hello world\u0085",
			Expected: "Hello world",
		},
		{
			Name:     "single line with trailing U+2028",
			Lines:    "Hello world\u2028",
			Expected: "Hello world",
		},
		{
			Name:     "single line with trailing U+2029",
			Lines:    "Hello world\u2029",
			Expected: "Hello world",
		},

		// Two lines with LF
		{
			Name:     "two lines with LF",
			Lines:    "Hello!\nHow do you do?",
			Expected: "Hello!\nHow do you do?",
		},
		{
			Name:     "two lines with CRLF",
			Lines:    "Hello!\r\nHow do you do?",
			Expected: "Hello!\nHow do you do?",
		},
		{
			Name:     "two lines with U+0085",
			Lines:    "Hello!\u0085How do you do?",
			Expected: "Hello!\nHow do you do?",
		},
		{
			Name:     "two lines with U+2028",
			Lines:    "Hello!\u2028How do you do?",
			Expected: "Hello!\n<br />\nHow do you do?",
		},
		{
			Name:     "two lines with U+2029",
			Lines:    "Hello!\u2029How do you do?",
			Expected: "Hello!\nHow do you do?",
		},

		// Formatting across lines
		{
			Name:     "formatting on multiple lines with LF",
			Lines:    "Hello!\nHow //do// **you** do?",
			Expected: "Hello!\nHow <em>do</em> <strong>you</strong> do?",
		},
		{
			Name:     "formatting on multiple lines with CRLF",
			Lines:    "Hello!\r\nHow //do// **you** do?",
			Expected: "Hello!\nHow <em>do</em> <strong>you</strong> do?",
		},
		{
			Name:     "formatting on multiple lines with U+2028",
			Lines:    "Hello!\u2028How //do// **you** do?",
			Expected: "Hello!\n<br />\nHow <em>do</em> <strong>you</strong> do?",
		},

		// HTML escaping
		{
			Name:     "html escaping across lines",
			Lines:    "AT&T\n2 < 3",
			Expected: "AT&amp;T\n2 &lt; 3",
		},

		// Three lines
		{
			Name:     "three lines with LF",
			Lines:    "one\ntwo\nthree",
			Expected: "one\ntwo\nthree",
		},

		// Unicode content
		{
			Name:     "persian text across lines",
			Lines:    "درود\nخداحافظ",
			Expected: "درود\nخداحافظ",
		},
		{
			Name:     "korean text across lines",
			Lines:    "안녕하세요\n감사합니다",
			Expected: "안녕하세요\n감사합니다",
		},

		// Link on a line
		{
			Name:     "link on a line",
			Lines:    "visit [[https://example.com]]\ndone",
			Expected: "visit <a href=\"https://example.com\">LINK</a>\ndone",
		},

		// Append to existing buffer
		{
			Name:     "append to existing buffer",
			Lines:    "**bold**\nnormal",
			Initial:  []byte("prefix "),
			Expected: "prefix <strong>bold</strong>\nnormal",
		},

		// Trailing EOL on multi-line
		{
			Name:     "two lines with trailing LF",
			Lines:    "Hello!\nHow do you do?\n",
			Expected: "Hello!\nHow do you do?",
		},
		{
			Name:     "two lines with trailing CRLF",
			Lines:    "Hello!\r\nHow do you do?\r\n",
			Expected: "Hello!\nHow do you do?",
		},



		{
			Name:     "more complex example",
			Lines:
				`H||e||l||l||o!`         +"\u0085"+
				`How //do// **you** do?` +"\u2028"+
				`H||e||l||l||o!`         +"\u0085"+
				`How //do// **you** do?` +"\u2028"+
				"",
			Expected:
				`H<mark>e</mark>l<mark>l</mark>o!`         +"\n"+
				`How <em>do</em> <strong>you</strong> do?` +"\n"+
				`<br />`                                   +"\n"+
				`H<mark>e</mark>l<mark>l</mark>o!`         +"\n"+
				`How <em>do</em> <strong>you</strong> do?`,
		},
	}

	for testNumber, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := string(renderBlockToHTML(test.Initial, test.Lines))

			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual HTML from renderBlockToHTML is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("LINES:    %q", test.Lines)
				return
			}
		})
	}
}
