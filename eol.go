package rtxt

import (
	"strings"
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
// If eolIndex does NOT find any, it returns an `index` of -1, and a `length` of 0.
//
// Else if eolIndex does find one of them, then it returns an `index` of it, and its `length`.
//
// For example, a `source` of:
//
//	"Hello!\nHow do you do?\n"
//
// Would yield an `index` of 6, and a `length of 1.
//
// Also, for example, a `source` of:
//
//	"Hello!\r\nHow do you do?\r\n"
//
// Would yield an `index` of 6, and a `length of 2.
// (The length is 2 here because "\r\n" take up 2 bytes.)
//
// And finally, for example, a `source` of:
//
//	"Hello!\u2028How do you do?\u2028"
//
// Would yield an `index` of 6, and a `length of 3.
// (The length is 3 here because U+2028 take up 3 bytes in its UTF-8 encoding: 0xE2 0x80 0xA8.)
func eolIndex(source string) (index int, length int) {

//@TODO: this could probably be made more performant if it did a single pass through the `source` rather than multiple.

	index = -1

	// This check for "\r\n" MUSt come before the check for "\n".
	// Since "\n" is within "\r\n".
	i := strings.Index(source, "\r\n")   // Carriage Return, Line Feed
	if index < 0 || (0 <= i && i < index) {
		index = i
		length = len("\r\n")
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, "\n")     // Line Feed
	if index < 0 || (0 <= i && i < index) {
		index = i
		length = len("\n")
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, "\u0085") // Next Line
	if index < 0 || (0 <= i && i < index) {
		index = i
		length = len("\u0085")
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, "\u2028") // Line Separator
	if index < 0 || (0 <= i && i < index) {
		index = i
		length = len("\u2028")
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, "\u2029") // Paragraph Separator
	if index < 0 || (0 <= i && i < index) {
		index = i
		length = len("\u2029")
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	if index < 0 {
		return -1, 0
	}

	return index, length
}
