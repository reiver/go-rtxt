package rtxt

import (
	"testing"
)

func TestRenderRunesToHTML(t *testing.T) {
	tests := []struct {
		Name     string
		Runes    string
		Initial  []byte
		Expected string
	}{
		{
			Name:     "empty string",
			Runes:    "",
			Expected: "",
		},

		{
			Name:     "plain text",
			Runes:    "hello",
			Expected: "hello",
		},
		{
			Name:     "single ampersand",
			Runes:    "&",
			Expected: "&amp;",
		},
		{
			Name:     "single less-than",
			Runes:    "<",
			Expected: "&lt;",
		},
		{
			Name:     "single greater-than",
			Runes:    ">",
			Expected: "&gt;",
		},

		{
			Name:     "text with ampersand",
			Runes:    "AT&T",
			Expected: "AT&amp;T",
		},
		{
			Name:     "text with less-than and greater-than",
			Runes:    "<hello>",
			Expected: "&lt;hello&gt;",
		},
		{
			Name:     "all special characters",
			Runes:    "<>&",
			Expected: "&lt;&gt;&amp;",
		},
		{
			Name:     "html tag",
			Runes:    "<script>alert('xss')</script>",
			Expected: "&lt;script&gt;alert('xss')&lt;/script&gt;",
		},

		// Unicode
		{
			Name:     "persian text",
			Runes:    "درود",
			Expected: "درود",
		},
		{
			Name:     "korean text",
			Runes:    "안녕하세요",
			Expected: "안녕하세요",
		},
		{
			Name:     "persian with special characters",
			Runes:    "سیب & موز & گیلاس",
			Expected: "سیب &amp; موز &amp; گیلاس",
		},
		{
			Name:     "emoji",
			Runes:    "🍎🍌🍒",
			Expected: "🍎🍌🍒",
		},

		// Append to existing buffer
		{
			Name:     "append to existing buffer",
			Runes:    " world",
			Initial:  []byte("hello"),
			Expected: "hello world",
		},
		{
			Name:     "append special characters to existing buffer",
			Runes:    " <>&",
			Initial:  []byte("chars:"),
			Expected: "chars: &lt;&gt;&amp;",
		},
	}

	for testNumber, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := string(renderRunesToHTML(test.Initial, test.Runes))

			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual HTML from renderRunesToHTML is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("RUNES:    %q", test.Runes)
				return
			}
		})
	}
}
