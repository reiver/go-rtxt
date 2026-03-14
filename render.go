package rtxt

import (
	"strings"
	"unicode/utf8"
)

// renderRuneToHTML renders a single run to HTML.
//
//	'&' ➔ "&amp;"
//	'<' ➔ "&lt;"
//	'>' ➔ "&gt;"
//	everything else stays as is.
func renderRuneToHTML(p []byte, r rune) []byte {
	switch r {
	case '&':
		p = append(p, "&amp;"...)
	case '<':
		p = append(p, "&lt;"...)
	case '>':
		p = append(p, "&gt;"...)
	default:
		p = utf8.AppendRune(p, r)
	}

	return p
}

// renderRunesToHTML is similar to [renderRuneToHTML] except does for many runes.
func renderRunesToHTML(p []byte, runes string) []byte {
	for _, r := range runes {
		p = renderRuneToHTML(p, r)
	}

	return p
}

// renderMarkedSubLineToHTML expected to have the `marker` at the begining and the ending of the `subline`.
// So, for example, if the `marker` was "**", then it would expect:
//
//	"**banana**"
//
// If that is not the case (that the marker is at the beginning and end of the `subline`) then
// renderMarkedSubLineToHTML just acts like [renderRunesToHTML].
//
// However, if the marker is indeed at the beginning and end of the `subline`, then it renders to HTML.
// So, continuing our example, we would get:
//
//	"<strong>banana</strong>"
func renderMarkedSubLineToHTML(p []byte, subline string, marker string, openedHTML string, closedHTML string) []byte {
	if !strings.HasPrefix(subline, marker) || !strings.HasSuffix(subline, marker) {
		return renderRunesToHTML(p, subline)
	}
	if len(subline) < len(marker) {
		return renderRunesToHTML(p, subline)
	}

	inner := subline[len(marker):len(subline)-len(marker)]

	p = append(p, openedHTML...)
	p = renderRunesToHTML(p, inner)
	p = append(p, closedHTML...)

	return p
}

// markerIndexes return the indexes of the opening-marker and the ending-marker in a line.
//
// So, for example, if the `marker` was "**" and the line was:
//
//	"apple **banana** cherry"
//
// Then the markerIndexes would return:
//
//	opened == 6
//	closed == 14
//
// Because:
//
//	           1111111111222
//	 01234567890123456789012
//	"apple **banana** cherry"
//	       🠙       🠙
//
// If both the opening-marker or the closing-marker doesn't
// exist in `line`, then markerIndexes returns:
//
//	opened == -1
//	closed == -1
//
// Note that this function assumes it has a line.
// I.e., that there are no EOL (end-of-line) characters in it.
// It will not itself check if there are EOL (end-of-line)
// characters in `line`. This matters because marked-text is
// NOT supposed to span multiple lines.
func markerIndexes(line string, openingMarker string, closingMarker string) (opened int, closed int) {
	opened = strings.Index(line, openingMarker)
	if opened < 0 {
		return -1, -1
	}

	var skip int = opened + len(openingMarker)
	closed = strings.Index(line[skip:], closingMarker)
	if closed < 0 {
		return -1, -1
	}
	closed += skip

	return opened, closed
}
