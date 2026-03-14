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
	linkRegexp      = regexp.MustCompile(`\[\[(.+?)\]\]`)
)

const listItemPrefix = "\t\u2022 "

func applyInlineFormatting(s string) string {
	s = html.EscapeString(s)

	s = boldRegexp.ReplaceAllString(s, "<strong>$1</strong>")
	s = italicRegexp.ReplaceAllString(s, "<em>$1</em>")
	s = underlineRegexp.ReplaceAllString(s, "<u>$1</u>")
	s = highlightRegexp.ReplaceAllString(s, "<mark>$1</mark>")
	s = linkRegexp.ReplaceAllString(s, `<a href="$1">$1</a>`)

	return s
}

func isListBlock(block string) bool {
	return strings.HasPrefix(block, listItemPrefix)
}

func renderList(block string, buf *strings.Builder) {
	lines := strings.Split(block, "\n")

	buf.WriteString("<ul>")

	first := true
	for _, line := range lines {
		if strings.HasPrefix(line, listItemPrefix) {
			if !first {
				buf.WriteString("</li>")
			}
			first = false
			content := applyInlineFormatting(strings.TrimPrefix(line, listItemPrefix))
			buf.WriteString("<li>")
			buf.WriteString(content)
		} else {
			line = strings.TrimPrefix(line, "\t")
			buf.WriteString("<br />")
			buf.WriteString(applyInlineFormatting(line))
		}
	}

	if !first {
		buf.WriteString("</li>")
	}
	buf.WriteString("</ul>")
}

func renderParagraph(block string, buf *strings.Builder) {
	p := applyInlineFormatting(block)

	p = strings.ReplaceAll(p, "\n", "<br />")
	p = strings.ReplaceAll(p, "\u2028", "<br />")
	p = strings.ReplaceAll(p, "\u0085", "<br />")

	buf.WriteString("<p>")
	buf.WriteString(p)
	buf.WriteString("</p>")
}

func ToHTML(s string) string {
	if "" == s {
		return ""
	}

	paragraphs := paragraphSplitter.Split(s, -1)

	var buf strings.Builder
	for _, p := range paragraphs {
		p = strings.TrimRight(p, "\n")
		if "" == p {
			continue
		}
		if isListBlock(p) {
			renderList(p, &buf)
		} else {
			renderParagraph(p, &buf)
		}
	}

	return buf.String()
}
