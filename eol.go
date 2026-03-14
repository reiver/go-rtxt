package rtxt

import (
	"strings"
)

func eolIndex(source string) int {
	var index int = -1

	// This check for "\r\n" MUSt come before the check for "\n".
	// Since "\n" is within "\r\n".
	i := strings.Index(source, "\r\n")   // Carriage Return, Line Feed
	if index < 0 || (0 <= i && i < index) {
		index = i
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, "\n")     // Line Feed
	if index < 0 || (0 <= i && i < index) {
		index = i
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, "\u0085") // Next Line
	if index < 0 || (0 <= i && i < index) {
		index = i
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, "\u2028") // Line Separator
	if index < 0 || (0 <= i && i < index) {
		index = i
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, "\u2029") // Paragraph Separator
	if index < 0 || (0 <= i && i < index) {
		index = i
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	if index < 0 {
		return len(source)
	}

	return index
}
