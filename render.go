package rtxt

import (
//	"strings"
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
