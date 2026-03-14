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

// markedIndexes return the indexes of the opening-marker and the ending-marker in a line.
//
// So, for example, if the `marker` was "**" and the line was:
//
//	"apple **banana** cherry"
//
// Then the markedIndexes would return:
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
// exist in `line`, then markedIndexes returns:
//
//	opened == -1
//	closed == -1
//
// Note that this function assumes it has a line.
// I.e., that there are no EOL (end-of-line) characters in it.
// It will not itself check if there are EOL (end-of-line)
// characters in `line`.
func markedIndexes(line string, marker string) (opened int, closed int) {
	opened = strings.Index(line, marker)
	if opened < 0 {
		return -1, -1
	}

	var skip int = opened + len(marker)
	closed = strings.Index(line[skip:], marker)
	if closed < 0 {
		return -1, -1
	}
	closed += skip

	return opened, closed
}
