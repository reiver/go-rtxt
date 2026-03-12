package rtxt

import (
	"html"
	"regexp"
	"strings"
)

var (
	paragraphSplitter = regexp.MustCompile(`\n[ \t\n]*\n|\x{2029}`)

	boldRegexp      = regexp.MustCompile(`\*\*(.+?)\*\*`)
	italicRegexp    = regexp.MustCompile(`//(.+?)//`)
	underlineRegexp = regexp.MustCompile(`__(.+?)__`)
	highlightRegexp = regexp.MustCompile(`\|\|(.+?)\|\|`)
)

func ToHTML(s string) string {
	if "" == s {
		return ""
	}

	paragraphs := paragraphSplitter.Split(s, -1)

	var buf strings.Builder
	for _, p := range paragraphs {
		p = html.EscapeString(p)

		p = boldRegexp.ReplaceAllString(p, "<strong>$1</strong>")
		p = italicRegexp.ReplaceAllString(p, "<em>$1</em>")
		p = underlineRegexp.ReplaceAllString(p, "<u>$1</u>")
		p = highlightRegexp.ReplaceAllString(p, "<mark>$1</mark>")

		p = strings.ReplaceAll(p, "\n", "<br />")
		p = strings.ReplaceAll(p, "\u2028", "<br />")
		p = strings.ReplaceAll(p, "\u0085", "<br />")

		buf.WriteString("<p>")
		buf.WriteString(p)
		buf.WriteString("</p>")
	}

	return buf.String()
}
