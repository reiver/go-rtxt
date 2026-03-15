package rtxt

import (
	"testing"
)

func TestAppendRenderedHTML(t *testing.T) {
	tests := []struct {
		Name     string
		Source   string
		Initial  []byte
		Expected string
	}{
		{
			Name:     "empty source",
			Source:   "",
			Expected: "",
		},



		{
			Name:     "plain text",
			Source:   "Hello world!",
			Expected: "Hello world!",
		},


		// HTML entities
		{
			Name:     "ampersand",
			Source:   "&",
			Expected: "&amp;",
		},
		{
			Name:     "less-than",
			Source:   "<",
			Expected: "&lt;",
		},
		{
			Name:     "greater-than",
			Source:   ">",
			Expected: "&gt;",
		},



		// marked
		{
			Name:     "bold",
			Source:   "**bold**",
			Expected: "<strong>bold</strong>",
		},
		{
			Name:     "italic",
			Source:   "//italic//",
			Expected: "<em>italic</em>",
		},
		{
			Name:     "underline",
			Source:   "__underline__",
			Expected: "<u>underline</u>",
		},
		{
			Name:     "highlight",
			Source:   "||highlight||",
			Expected: "<mark>highlight</mark>",
		},



		// paragraphs
		{
			Name:     "paragraphs",
			Source:
				`one two`                   +"\n"+
				`three four five`           +"\n"+
				`six seven`                 +"\n"+
				`eight nine ten`            +"\n"+
				""                          +"\n"+
				`eleven twelve`             +"\n"+
				`thirteen fourteen fifteen` +"\n"+
				"",
			Expected:
				`<p>`                       +"\n"+
				`one two`                   +"\n"+
				`three four five`           +"\n"+
				`six seven`                 +"\n"+
				`eight nine ten`            +"\n"+
				`</p>`                      +"\n"+
				`<p>`                       +"\n"+
				`eleven twelve`             +"\n"+
				`thirteen fourteen fifteen` +"\n"+
				`</p>`                      +"\n"+
				"",
		},
		{
			Name:     "paragraphs",
			Source:
				`one two`                   +"\n"+
				`three four five`           +"\n"+
				`six seven`                 +"\n"+
				`eight nine ten`            +"\n"+
				""                          +"\n"+
				`eleven twelve`             +"\n"+
				`thirteen fourteen fifteen` +"\n"+
				""                          +"\n"+
				`sixteen seventeen`         +"\n"+
				`eighteen ninteen twenty`   +"\n"+
				"",
			Expected:
				`<p>`                       +"\n"+
				`one two`                   +"\n"+
				`three four five`           +"\n"+
				`six seven`                 +"\n"+
				`eight nine ten`            +"\n"+
				`</p>`                      +"\n"+
				`<p>`                       +"\n"+
				`eleven twelve`             +"\n"+
				`thirteen fourteen fifteen` +"\n"+
				`</p>`                      +"\n"+
				`<p>`                       +"\n"+
				`sixteen seventeen`         +"\n"+
				`eighteen ninteen twenty`   +"\n"+
				`</p>`                      +"\n"+
				"",
		},
	}

	for testNumber, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := string(AppendRenderedHTML(test.Initial, test.Source))

			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual HTML from AppendRenderedHTML is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("SOURCE:   %q", test.Source)
				return
			}
		})
	}
}
