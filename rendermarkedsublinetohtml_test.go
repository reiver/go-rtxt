package rtxt

import (
	"testing"
)

func TestRenderMarkedSubLineToHTML(t *testing.T) {
	tests := []struct {
		Name       string
		SubLine    string
		Marker     string
		OpenedHTML string
		ClosedHTML string
		Initial    []byte
		Expected   string
	}{
		// Bold
		{
			Name:       "bold",
			SubLine:    "**banana**",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Expected:   "<strong>banana</strong>",
		},

		// Italic
		{
			Name:       "italic",
			SubLine:    "//banana//",
			Marker:     "//",
			OpenedHTML: "<em>",
			ClosedHTML: "</em>",
			Expected:   "<em>banana</em>",
		},

		// Underline
		{
			Name:       "underline",
			SubLine:    "__banana__",
			Marker:     "__",
			OpenedHTML: "<u>",
			ClosedHTML: "</u>",
			Expected:   "<u>banana</u>",
		},

		// Highlight
		{
			Name:       "highlight",
			SubLine:    "||banana||",
			Marker:     "||",
			OpenedHTML: "<mark>",
			ClosedHTML: "</mark>",
			Expected:   "<mark>banana</mark>",
		},

		// Fallback — no markers present
		{
			Name:       "no markers falls back to plain rendering",
			SubLine:    "banana",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Expected:   "banana",
		},

		// Fallback — only opening marker
		{
			Name:       "only opening marker falls back",
			SubLine:    "**banana",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Expected:   "**banana",
		},

		// Fallback — only closing marker
		{
			Name:       "only closing marker falls back",
			SubLine:    "banana**",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Expected:   "banana**",
		},

		// HTML escaping inside markers
		{
			Name:       "html escaping inside bold",
			SubLine:    "**<b>**",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Expected:   "<strong>&lt;b&gt;</strong>",
		},
		{
			Name:       "ampersand inside bold",
			SubLine:    "**AT&T**",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Expected:   "<strong>AT&amp;T</strong>",
		},

		// HTML escaping in fallback
		{
			Name:       "html escaping in fallback",
			SubLine:    "AT&T",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Expected:   "AT&amp;T",
		},

		// Empty content between markers
		{
			Name:       "empty content between markers",
			SubLine:    "****",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Expected:   "<strong></strong>",
		},

		// Unicode
		{
			Name:       "persian text in bold",
			SubLine:    "**دوست**",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Expected:   "<strong>دوست</strong>",
		},
		{
			Name:       "korean text in italic",
			SubLine:    "//굵은//",
			Marker:     "//",
			OpenedHTML: "<em>",
			ClosedHTML: "</em>",
			Expected:   "<em>굵은</em>",
		},

		// Append to existing buffer
		{
			Name:       "append bold to existing buffer",
			SubLine:    "**banana**",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Initial:    []byte("hello "),
			Expected:   "hello <strong>banana</strong>",
		},
		{
			Name:       "append fallback to existing buffer",
			SubLine:    "banana",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Initial:    []byte("hello "),
			Expected:   "hello banana",
		},

		// unpaired marker
		{
			Name:       "only opened marker",
			SubLine:    "**once & twice & thrice & fource",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Initial:    []byte("hello "),
			Expected:   "hello **once &amp; twice &amp; thrice &amp; fource",
		},
		{
			Name:       "only closed marker",
			SubLine:    "once & twice & thrice & fource**",
			Marker:     "**",
			OpenedHTML: "<strong>",
			ClosedHTML: "</strong>",
			Initial:    []byte("hello "),
			Expected:   "hello once &amp; twice &amp; thrice &amp; fource**",
		},
	}

	for testNumber, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := string(renderMarkedSubLineToHTML(test.Initial, test.SubLine, test.Marker, test.OpenedHTML, test.ClosedHTML))

			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual HTML from renderMarkedSubLineToHTML is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("SUBLINE:  %q", test.SubLine)
				t.Logf("MARKER:   %q", test.Marker)
				return
			}
		})
	}
}
