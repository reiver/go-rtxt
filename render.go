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

// renderMarkedSubLineToHTML expected to have the `openingMarker` at the begining and the `closingMarker` at the ending of the `subline`.
// So, for example, if the `openingMarker` was "**" and the `closingMarker` was "**", then it would expect:
//
//	"**banana**"
//
// Or, if the `openingMarker` was "[[" and the `closingMarker` was "]]", then it would expect:
//
//	"[[https://example.com]]"
//
// If that is not the case (that the opening marker is at the beginning and the closing marker is at the end of the `subline`) then
// renderMarkedSubLineToHTML just acts like [renderRunesToHTML].
//
// However, if the markers are indeed at the beginning and end of the `subline`, then it renders to HTML.
// So, continuing our example, we would get:
//
//	"<strong>banana</strong>"
//
// And:
//
//	"<a href=\"http://example.com\">LINK<a>"
func renderMarkedSubLineToHTML(p []byte, subline string, openingMarker string, closingMarker string, openingHTML string, closingHTML string) []byte {
	if !strings.HasPrefix(subline, openingMarker) || !strings.HasSuffix(subline, closingMarker) {
		return renderRunesToHTML(p, subline)
	}
	if len(subline) < len(openingMarker)+len(closingMarker) {
		return renderRunesToHTML(p, subline)
	}

	inner := subline[len(openingMarker):len(subline)-len(closingMarker)]

	p = append(p, openingHTML...)
	p = renderRunesToHTML(p, inner)
	p = append(p, closingHTML...)

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

// renderLineToHTML renders a line to HTML.
//
// For example, if the `line` was:
//
//	"Hello, //how// do **you** do?"
//
// Then (logically) we it would get rendered as:
//
//	"Hello, <em>how</em> do <strong>you</strong> do?"
//
// And, the following string:
//
//	"2 < 3 & 5 > 4"
//
// Would (logically) get rendered as:
//
//	"2 &lt; 3 &amp; 5 &gt;4"
//
// Note that this function assumes it has a line.
// I.e., that there are no EOL (end-of-line) characters in it.
// It will not itself check if there are EOL (end-of-line)
// characters in `line`. This matters because marked-text is
// NOT supposed to span multiple lines.
func renderLineToHTML(p []byte, line string) []byte {

	var openingMarkers = [...]string{
		"**",
		"//",
		"__",
		"||",
		"[[",
	}

	var closingMarkers = [...]string{
		"**",
		"//",
		"__",
		"||",
		"]]",
	}

	var openingHTMLs = [...]string{
		`<strong>`,
		`<em>`,
		`<u>`,
		`<mark>`,
		`<a href="`,
	}

	var closingHTMLs = [...]string{
		`</strong>`,
		`</em>`,
		`</u>`,
		`</mark>`,
		`">LINK</a>`,
	}

	var openingIndex int = -1
	var closingIndex int = -1
	var openingMarker string
	var closingMarker string
	var openingHTML string
	var closingHTML string
	for i, oMarker := range openingMarkers {
		cMarker := closingMarkers[i]

		opened, closed := markerIndexes(line, oMarker, cMarker)
		if opened < 0 || closed < 0 {
			continue
		}

		if openingIndex < 0 || opened < openingIndex {
			openingIndex  = opened
			closingIndex  = closed
			openingMarker = oMarker
			closingMarker = cMarker
			openingHTML   = openingHTMLs[i]
			closingHTML   = closingHTMLs[i]
		}
	}
	if openingIndex < 0 {
		return renderRunesToHTML(p, line)
	}

	// Render what is before the opening marker.
	p = renderRunesToHTML(p, line[:openingIndex])
	line = line[openingIndex:]
	closingIndex -= openingIndex
	openingIndex = 0 // openingIndex -= openingIndex

	// Render the marked part.
	subline := line[:closingIndex+len(closingMarker)]
	p = renderMarkedSubLineToHTML(p, subline, openingMarker, closingMarker, openingHTML, closingHTML)

	// Render what is after the marked part.
	line = line[len(subline):]
//@TODO: Will this blow the stack if the text is very large?
	return renderLineToHTML(p, line)
}

// renderLinesToHTML render (potentially) multiple lines in a block to HTML.
// A block may be the lines of a paragraph, or could be the lines of an item
// in an unordered list, etc.
//
// If the block was:
//
//	"Hello!"                 +"\n"+
//	"How //do// **you** do?" +"\n"+
//
// Then that would (logically) be rendered to:
//
//	"Hello!"                                   +"\n"+
//	"How <em>do</em> <strong>you</strong> do?" +"\n"+
//
// Or, if the block was:
//
//	"Hello!"                 +"\r\n"+
//	"How //do// **you** do?" +"\r\n"+
//
// Then that would (logically) be rendered to:
//
//	"Hello!"                                   +"\n"+
//	"How <em>do</em> <strong>you</strong> do?" +"\n"+
//
// Also, if the block was:
//
//	"Hello!"                 +"\u2028"+
//	"How //do// **you** do?" +"\u2028"+
//
// Then that would (logically) be rendered to:
//
//	"Hello!"                                   +"\n"+
//	"How <em>do</em> <strong>you</strong> do?" +"\n"+
func renderBlockToHTML(p []byte, lines string) []byte {

	for {
		eolindex, eol := eolIndex(lines)
		if eolindex < 0 {
			return renderLineToHTML(p, lines)
		}
		if 0 == eolindex && "" == eol {
			// This should never happen.
			return renderRunesToHTML(p, lines)
		}

		line := lines[:eolindex]
		p = renderLineToHTML(p, line)

		skip := eolindex + len(eol)
		lines = lines[skip:]

		if len(lines) <= 0 {
			break
		}

		switch eol {
		case ls:
			p = append(p, "\n<br />\n"...)
		case "":
			// nothing here.
		default:
			p = append(p, '\n')
		}
	}

	return p
}
