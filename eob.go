package rtxt

import (
	"strings"

	"github.com/reiver/go-unicode"
)

const (
	crlfcrlf = string(unicode.CR) + string(unicode.LF) + string(unicode.CR) + string(unicode.LF)
	lflf     = string(unicode.LF) + string(unicode.LF)
	nelnel   = string(unicode.NEL) + string(unicode.NEL)
//	ps       = string(unicode.PS)
)

func eobIndex(source string) (index int, eob string) {

//@TODO: this could probably be made more performant if it did a single pass through the `source` rather than multiple.

	index = -1

	i := strings.Index(source, crlfcrlf) // Carriage Return, Line Feed, Carriage Return, Line Feed
	if index < 0 || (0 <= i && i < index) {
		index = i
		eob = crlfcrlf
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, lflf)    // Line Feed, Line Feed
	if index < 0 || (0 <= i && i < index) {
		index = i
		eob = lflf
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, nelnel)   // Next Line, Next Line
	if index < 0 || (0 <= i && i < index) {
		index = i
		eob = nelnel
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	i = strings.Index(source, ps)    // Paragraph Separator
	if index < 0 || (0 <= i && i < index) {
		index = i
		eob = ps
	}
	if 0 <= index {
		source = source[:index] // don't neded to search the whole string now. shorten string so the next strings.Index() is quicker.
	}

	if index < 0 {
		return -1, ""
	}

	return index, eob
}
