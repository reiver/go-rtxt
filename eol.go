package rtxt

import (
	"strings"

	"github.com/reiver/go-unicode"
)

const (
	crlf = string(unicode.CR) + string(unicode.LF)
	lf   = string(unicode.LF)
	nel  = string(unicode.NEL)
	ls   = string(unicode.LS)
	ps   = string(unicode.PS)
)

// eolIndex tries to find an EOL (end-of-line) character in the `source` string.
//
// The EOL characters eolIndex looks for are:
//
//	• "\r\n"
//	• "\n"
//	• "\u0085"
//	• "\u2028"
//	• "\u2029"
//
// If eolIndex does NOT find any, it returns an `index` of -1, and a `eol` of "".
//
// Else if eolIndex does find one of them, then it returns the starting `index` of where it is in `source`, and return the `eol` string it found.
//
// For example, a `source` of:
//
//	"Hello!\nHow do you do?\n"
//
// Would yield an `index` of 6, and an `eol` of "\n".
//
// Also, for example, a `source` of:
//
//	"Hello!\r\nHow do you do?\r\n"
//
// Would yield an `index` of 6, and an `eol` of "\r\n".
//
// And finally, for example, a `source` of:
//
//	"Hello!\u2028How do you do?\u2028"
//
// Would yield an `index` of 6, and a `eol` of "\u2028".
func eolIndex(source string) (index int, eol string) {

//@TODO: this could probably be made more performant if it did a single pass through the `source` rather than multiple.

	index = -1

	// This check for "\r\n" MUSt come before the check for "\n".
	// Since "\n" is within "\r\n".
	i := strings.Index(source, crlf) // Carriage Return, Line Feed
	if index < 0 || (0 <= i && i < index) {
		index = i
		eol = crlf
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, lf)    // Line Feed
	if index < 0 || (0 <= i && i < index) {
		index = i
		eol = lf
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, nel)   // Next Line
	if index < 0 || (0 <= i && i < index) {
		index = i
		eol = nel
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, ls)    // Line Separator
	if index < 0 || (0 <= i && i < index) {
		index = i
		eol = ls
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, ps)    // Paragraph Separator
	if index < 0 || (0 <= i && i < index) {
		index = i
		eol = ps
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	if index < 0 {
		return -1, ""
	}

	return index, eol
}
