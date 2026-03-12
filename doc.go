/*
Package rtxt converts rich text to HTML.

It supports the following inline formatting:

	**bold**       → <strong>bold</strong>
	//italic//     → <em>italic</em>
	__underline__  → <u>underline</u>
	||highlight||  → <mark>highlight</mark>
	[[url]]        → <a href="url">url</a>

Unordered lists are created by starting lines with a tab, U+2022 (bullet), and a space:

	\t• apples
	\t• bananas
	\t• cherries

Lines without the bullet prefix are continuation lines joined to the previous item with a <br />.

Paragraphs are separated by blank lines (double newline) or the Unicode paragraph separator (U+2029).
Single newlines, next-line characters (U+0085), and Unicode line separators (U+2028) within a paragraph become <br /> tags.
HTML special characters are escaped.

	input := "Hello **world**\n\nSecond paragraph"
	output := rtxt.ToHTML(input)
	// "<p>Hello <strong>world</strong></p><p>Second paragraph</p>"
*/
package rtxt
