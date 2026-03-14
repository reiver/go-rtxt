package rtxt

import (
	"testing"
)

func TestMarkedIndexes(t *testing.T) {
	tests := []struct {
		Name           string
		Line           string
		Marker         string
		ExpectedOpened int
		ExpectedClosed int
	}{
		{
			Name:           "bold marker in sentence",
			Line:           "apple **banana** cherry",
			Marker:         "**",
			ExpectedOpened: 6,
			ExpectedClosed: 14,
		},
		{
			Name:           "italic marker in sentence",
			Line:           "apple //banana// cherry",
			Marker:         "//",
			ExpectedOpened: 6,
			ExpectedClosed: 14,
		},
		{
			Name:           "underline marker in sentence",
			Line:           "apple __banana__ cherry",
			Marker:         "__",
			ExpectedOpened: 6,
			ExpectedClosed: 14,
		},
		{
			Name:           "highlight marker in sentence",
			Line:           "apple ||banana|| cherry",
			Marker:         "||",
			ExpectedOpened: 6,
			ExpectedClosed: 14,
		},
		{
			Name:           "link marker no closing match",
			Line:           "visit [[https://example.com]] now",
			Marker:         "[[",
			ExpectedOpened: -1,
			ExpectedClosed: -1,
		},

		{
			Name:           "no marker present",
			Line:           "apple banana cherry",
			Marker:         "**",
			ExpectedOpened: -1,
			ExpectedClosed: -1,
		},
		{
			Name:           "only opening marker",
			Line:           "apple **banana cherry",
			Marker:         "**",
			ExpectedOpened: -1,
			ExpectedClosed: -1,
		},

		{
			Name:           "marker at start",
			Line:           "**banana** cherry",
			Marker:         "**",
			ExpectedOpened: 0,
			ExpectedClosed: 8,
		},
		{
			Name:           "marker at end",
			Line:           "apple **banana**",
			Marker:         "**",
			ExpectedOpened: 6,
			ExpectedClosed: 14,
		},
		{
			Name:           "entire line is marked",
			Line:           "**banana**",
			Marker:         "**",
			ExpectedOpened: 0,
			ExpectedClosed: 8,
		},

		{
			Name:           "empty content between markers",
			Line:           "apple **** cherry",
			Marker:         "**",
			ExpectedOpened: 6,
			ExpectedClosed: 8,
		},
		{
			Name:           "empty line",
			Line:           "",
			Marker:         "**",
			ExpectedOpened: -1,
			ExpectedClosed: -1,
		},

		// Unicode
		{
			Name:           "persian text with bold marker",
			Line:           "درود **دوست** عزیز",
			Marker:         "**",
			ExpectedOpened: 9,
			ExpectedClosed: 19,
		},
		{
			Name:           "korean text with bold marker",
			Line:           "이것은 **굵은** 글씨",
			Marker:         "**",
			ExpectedOpened: 10,
			ExpectedClosed: 18,
		},
	}

	for testNumber, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actualOpened, actualClosed := markedIndexes(test.Line, test.Marker)

			if test.ExpectedOpened != actualOpened || test.ExpectedClosed != actualClosed {
				t.Errorf("For test #%d, the actual marked-indexes are not what was expected.", testNumber)
				t.Logf("EXPECTED OPENED: %d", test.ExpectedOpened)
				t.Logf("ACTUAL   OPENED: %d", actualOpened)
				t.Logf("EXPECTED CLOSED: %d", test.ExpectedClosed)
				t.Logf("ACTUAL   CLOSED: %d", actualClosed)
				t.Logf("LINE:   %q", test.Line)
				t.Logf("MARKER: %q", test.Marker)
				return
			}
		})
	}
}
