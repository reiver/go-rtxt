package rtxt

import (
	"testing"
)

func TestMarkerIndexes(t *testing.T) {
	tests := []struct {
		Name           string
		Line           string
		OpeningMarker  string
		ClosingMarker  string
		ExpectedOpened int
		ExpectedClosed int
	}{
		{
			Name:           "bold marker in sentence",
			Line:           "apple **banana** cherry",
			OpeningMarker:  "**",
			ClosingMarker:  "**",
			ExpectedOpened: 6,
			ExpectedClosed: 14,
		},
		{
			Name:           "italic marker in sentence",
			Line:           "apple //banana// cherry",
			OpeningMarker:  "//",
			ClosingMarker:  "//",
			ExpectedOpened: 6,
			ExpectedClosed: 14,
		},
		{
			Name:           "underline marker in sentence",
			Line:           "apple __banana__ cherry",
			OpeningMarker:  "__",
			ClosingMarker:  "__",
			ExpectedOpened: 6,
			ExpectedClosed: 14,
		},
		{
			Name:           "highlight marker in sentence",
			Line:           "apple ||banana|| cherry",
			OpeningMarker:  "||",
			ClosingMarker:  "||",
			ExpectedOpened: 6,
			ExpectedClosed: 14,
		},
		{
			Name:           "link marker no closing match",
			Line:           "visit [[https://example.com]] now",
			OpeningMarker:  "[[",
			ClosingMarker:  "[[",
			ExpectedOpened: -1,
			ExpectedClosed: -1,
		},
		{
			Name:           "link marker no closing match",
			Line:           "visit [[https://example.com]] now",
			OpeningMarker:  "[[",
			ClosingMarker:  "]]",
			ExpectedOpened: 6,
			ExpectedClosed: 27,
		},

		{
			Name:           "no marker present",
			Line:           "apple banana cherry",
			OpeningMarker:  "**",
			ClosingMarker:  "**",
			ExpectedOpened: -1,
			ExpectedClosed: -1,
		},
		{
			Name:           "only opening marker",
			Line:           "apple **banana cherry",
			OpeningMarker:  "**",
			ClosingMarker:  "**",
			ExpectedOpened: -1,
			ExpectedClosed: -1,
		},

		{
			Name:           "marker at start",
			Line:           "**banana** cherry",
			OpeningMarker:  "**",
			ClosingMarker:  "**",
			ExpectedOpened: 0,
			ExpectedClosed: 8,
		},
		{
			Name:           "marker at end",
			Line:           "apple **banana**",
			OpeningMarker:  "**",
			ClosingMarker:  "**",
			ExpectedOpened: 6,
			ExpectedClosed: 14,
		},
		{
			Name:           "entire line is marked",
			Line:           "**banana**",
			OpeningMarker:  "**",
			ClosingMarker:  "**",
			ExpectedOpened: 0,
			ExpectedClosed: 8,
		},

		{
			Name:           "empty content between markers",
			Line:           "apple **** cherry",
			OpeningMarker:  "**",
			ClosingMarker:  "**",
			ExpectedOpened: 6,
			ExpectedClosed: 8,
		},
		{
			Name:           "empty line",
			Line:           "",
			OpeningMarker:  "**",
			ClosingMarker:  "**",
			ExpectedOpened: -1,
			ExpectedClosed: -1,
		},

		// Unicode
		{
			Name:           "persian text with bold marker",
			Line:           "درود **دوست** عزیز",
			OpeningMarker:  "**",
			ClosingMarker:  "**",
			ExpectedOpened: 9,
			ExpectedClosed: 19,
		},
		{
			Name:           "korean text with bold marker",
			Line:           "이것은 **굵은** 글씨",
			OpeningMarker:  "**",
			ClosingMarker:  "**",
			ExpectedOpened: 10,
			ExpectedClosed: 18,
		},
	}

	for testNumber, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actualOpened, actualClosed := markerIndexes(test.Line, test.OpeningMarker, test.ClosingMarker)

			if test.ExpectedOpened != actualOpened || test.ExpectedClosed != actualClosed {
				t.Errorf("For test #%d, the actual marked-indexes are not what was expected.", testNumber)
				t.Logf("EXPECTED OPENED: %d", test.ExpectedOpened)
				t.Logf("ACTUAL   OPENED: %d", actualOpened)
				t.Logf("EXPECTED CLOSED: %d", test.ExpectedClosed)
				t.Logf("ACTUAL   CLOSED: %d", actualClosed)
				t.Logf("LINE:   %q", test.Line)
				t.Logf("OPENING-MARKER: %q", test.OpeningMarker)
				t.Logf("CLOSING-MARKER: %q", test.ClosingMarker)
				return
			}
		})
	}
}
