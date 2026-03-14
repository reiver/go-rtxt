package rtxt

import (
	"testing"
)

func TestRenderLineToHTML(t *testing.T) {
	tests := []struct {
		Name     string
		Line     string
		Initial  []byte
		Expected string
	}{
		// Plain text
		{
			Name:     "empty string",
			Line:     "",
			Expected: "",
		},
		{
			Name:     "plain text",
			Line:     "hello world",
			Expected: "hello world",
		},

		// HTML escaping
		{
			Name:     "ampersand escaped",
			Line:     "AT&T",
			Expected: "AT&amp;T",
		},
		{
			Name:     "less-than escaped",
			Line:     "2 < 3",
			Expected: "2 &lt; 3",
		},
		{
			Name:     "greater-than escaped",
			Line:     "5 > 4",
			Expected: "5 &gt; 4",
		},
		{
			Name:     "mixed special characters",
			Line:     "2 < 3 & 5 > 4",
			Expected: "2 &lt; 3 &amp; 5 &gt; 4",
		},

		// Bold
		{
			Name:     "bold word at start of line",
			Line:     "**banana**",
			Expected: "<strong>banana</strong>",
		},
		{
			Name:     "bold followed by plain text",
			Line:     "**banana** cherry",
			Expected: "<strong>banana</strong> cherry",
		},
		{
			Name:     "bold preceded by plain text",
			Line:     "apple **banana**",
			Expected: "apple <strong>banana</strong>",
		},
		{
			Name:     "bold surrounded by plain text",
			Line:     "apple **banana** cherry",
			Expected: "apple <strong>banana</strong> cherry",
		},

		// Italic
		{
			Name:     "italic word at start of line",
			Line:     "//banana//",
			Expected: "<em>banana</em>",
		},
		{
			Name:     "italic followed by plain text",
			Line:     "//banana// cherry",
			Expected: "<em>banana</em> cherry",
		},
		{
			Name:     "italic preceded by plain text",
			Line:     "apple //banana//",
			Expected: "apple <em>banana</em>",
		},
		{
			Name:     "italic surrounded by plain text",
			Line:     "apple //banana// cherry",
			Expected: "apple <em>banana</em> cherry",
		},

		// Underline
		{
			Name:     "underline word at start of line",
			Line:     "__banana__",
			Expected: "<u>banana</u>",
		},
		{
			Name:     "underline followed by plain text",
			Line:     "__banana__ cherry",
			Expected: "<u>banana</u> cherry",
		},
		{
			Name:     "underline preceded by plain text",
			Line:     "apple __banana__",
			Expected: "apple <u>banana</u>",
		},
		{
			Name:     "underline surrounded by plain text",
			Line:     "apple __banana__ cherry",
			Expected: "apple <u>banana</u> cherry",
		},

		// Highlight
		{
			Name:     "highlight word at start of line",
			Line:     "||banana||",
			Expected: "<mark>banana</mark>",
		},
		{
			Name:     "highlight followed by plain text",
			Line:     "||banana|| cherry",
			Expected: "<mark>banana</mark> cherry",
		},
		{
			Name:     "highlight preceded by plain text",
			Line:     "apple ||banana||",
			Expected: "apple <mark>banana</mark>",
		},
		{
			Name:     "highlight surrounded by plain text",
			Line:     "apple ||banana|| cherry",
			Expected: "apple <mark>banana</mark> cherry",
		},

		// Link
		{
			Name:     "link at start of line",
			Line:     "[[http://example.com]]",
			Expected: `<a href="http://example.com">LINK</a>`,
		},
		{
			Name:     "link followed by plain text",
			Line:     "[[http://example.com]] now",
			Expected: `<a href="http://example.com">LINK</a> now`,
		},
		{
			Name:     "link preceded by plain text",
			Line:     "tap [[http://example.com]]",
			Expected: `tap <a href="http://example.com">LINK</a>`,
		},
		{
			Name:     "link surrounded by plain text",
			Line:     "tap [[http://example.com]] now",
			Expected: `tap <a href="http://example.com">LINK</a> now`,
		},

		// HTML escaping with formatting
		{
			Name:     "ampersand inside bold",
			Line:     "**AT&T**",
			Expected: "<strong>AT&amp;T</strong>",
		},

		// No closing marker
		{
			Name:     "unpaired opening bold",
			Line:     "apple **banana cherry",
			Expected: "apple **banana cherry",
		},
		{
			Name:     "unpaired opening italic",
			Line:     "apple //banana cherry",
			Expected: "apple //banana cherry",
		},
		{
			Name:     "unpaired opening underline",
			Line:     "apple __banana cherry",
			Expected: "apple __banana cherry",
		},
		{
			Name:     "unpaired opening highlight",
			Line:     "apple ||banana cherry",
			Expected: "apple ||banana cherry",
		},
		{
			Name:     "unpaired opening link",
			Line:     "tap [[http://example.com now",
			Expected: "tap [[http://example.com now",
		},
		{
			Name:     "unpaired close link",
			Line:     "tap http://example.com]] now",
			Expected: "tap http://example.com]] now",
		},

		// Empty content between markers
		{
			Name:     "empty bold",
			Line:     "****",
			Expected: "<strong></strong>",
		},

		// Unicode at start
		{
			Name:     "persian bold at start of line",
			Line:     "**دوست**",
			Expected: "<strong>دوست</strong>",
		},
		{
			Name:     "korean italic at start of line",
			Line:     "//굵은//",
			Expected: "<em>굵은</em>",
		},

		// Append to existing buffer
		{
			Name:     "append to existing buffer",
			Line:     "**bold**",
			Initial:  []byte("prefix "),
			Expected: "prefix <strong>bold</strong>",
		},
		{
			Name:     "append plain text to existing buffer",
			Line:     "hello",
			Initial:  []byte("prefix "),
			Expected: "prefix hello",
		},



		{
			Name:     "lots",
			Line:     "Hello, //how// do **you** do? Did ||you|| __know__ that 2 < 3 & 5 > 4?",
			Initial:  []byte("Hi! "),
			Expected: "Hi! Hello, <em>how</em> do <strong>you</strong> do? Did <mark>you</mark> <u>know</u> that 2 &lt; 3 &amp; 5 &gt; 4?",
		},
	}

	for testNumber, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := string(renderLineToHTML(test.Initial, test.Line))

			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual HTML from renderLineToHTML is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("LINE:     %q", test.Line)
				return
			}
		})
	}
}
