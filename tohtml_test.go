package rtxt

import (
	"testing"
)

func TestToHTML(t *testing.T) {
	tests := []struct {
		Name     string
		Input    string
		Expected string
	}{
		{
			Name:     "empty string",
			Input:    "",
			Expected: "",
		},
		{
			Name:     "plain text",
			Input:    "Hello world",
			Expected: "<p>Hello world</p>",
		},
		{
			Name:     "two paragraphs",
			Input:    "First paragraph\n\nSecond paragraph",
			Expected: "<p>First paragraph</p><p>Second paragraph</p>",
		},
		{
			Name:     "paragraphs separated by whitespace between newlines",
			Input:    "First paragraph\n \t\nSecond paragraph",
			Expected: "<p>First paragraph</p><p>Second paragraph</p>",
		},
		{
			Name:     "bold",
			Input:    "**bold**",
			Expected: "<p><strong>bold</strong></p>",
		},
		{
			Name:     "italic",
			Input:    "//italic//",
			Expected: "<p><em>italic</em></p>",
		},
		{
			Name:     "underline",
			Input:    "__underline__",
			Expected: "<p><u>underline</u></p>",
		},
		{
			Name:     "highlight",
			Input:    "||highlight||",
			Expected: "<p><mark>highlight</mark></p>",
		},
		{
			Name:     "mixed formatting in one paragraph",
			Input:    "This is **bold** and //italic// and __underlined__ and ||highlighted||",
			Expected: "<p>This is <strong>bold</strong> and <em>italic</em> and <u>underlined</u> and <mark>highlighted</mark></p>",
		},
		{
			Name:     "single newline within paragraph becomes br",
			Input:    "Line one\nLine two",
			Expected: "<p>Line one<br />Line two</p>",
		},
		{
			Name:     "HTML special chars escaped",
			Input:    "<script>alert('xss')</script>",
			Expected: "<p>&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;</p>",
		},
		{
			Name:     "ampersand escaped",
			Input:    "AT&T",
			Expected: "<p>AT&amp;T</p>",
		},
		{
			Name:     "ampersand in words",
			Input:    "apples & bananas",
			Expected: "<p>apples &amp; bananas</p>",
		},
		{
			Name:     "three paragraphs",
			Input:    "One\n\nTwo\n\nThree",
			Expected: "<p>One</p><p>Two</p><p>Three</p>",
		},
		{
			Name:     "formatting across multiple paragraphs",
			Input:    "**bold text**\n\n//italic text//",
			Expected: "<p><strong>bold text</strong></p><p><em>italic text</em></p>",
		},
		{
			Name:     "single newline and double newline mixed",
			Input:    "Line one\nLine two\n\nNew paragraph",
			Expected: "<p>Line one<br />Line two</p><p>New paragraph</p>",
		},

		// unordered lists
		{
			Name:     "single list item",
			Input:    "\t\u2022 apples",
			Expected: "<ul><li>apples</li></ul>",
		},
		{
			Name:     "multiple list items",
			Input:    "\t\u2022 apples\n\t\u2022 bananas\n\t\u2022 cherries",
			Expected: "<ul><li>apples</li><li>bananas</li><li>cherries</li></ul>",
		},
		{
			Name:     "list item with continuation line",
			Input:    "\t\u2022 first item\n\tcontinues here\n\t\u2022 second item",
			Expected: "<ul><li>first item<br />continues here</li><li>second item</li></ul>",
		},
		{
			Name:     "list continuation only strips one leading tab",
			Input:    "\t\u2022 item\n\t\textra indented",
			Expected: "<ul><li>item<br />\textra indented</li></ul>",
		},
		{
			Name:     "list with bold item",
			Input:    "\t\u2022 **bold item**\n\t\u2022 normal item",
			Expected: "<ul><li><strong>bold item</strong></li><li>normal item</li></ul>",
		},
		{
			Name:     "list followed by paragraph",
			Input:    "\t\u2022 apples\n\t\u2022 bananas\n\nSome text",
			Expected: "<ul><li>apples</li><li>bananas</li></ul><p>Some text</p>",
		},
		{
			Name:     "paragraph followed by list",
			Input:    "Some text\n\n\t\u2022 apples\n\t\u2022 bananas",
			Expected: "<p>Some text</p><ul><li>apples</li><li>bananas</li></ul>",
		},

		// links
		{
			Name:     "link",
			Input:    "[[https://example.com]]",
			Expected: `<p><a href="https://example.com">https://example.com</a></p>`,
		},
		{
			Name:     "link in sentence",
			Input:    "Visit [[https://example.com]] for more info",
			Expected: `<p>Visit <a href="https://example.com">https://example.com</a> for more info</p>`,
		},
		{
			Name:     "link with query parameters",
			Input:    "[[https://example.com/search?q=hello&lang=en]]",
			Expected: `<p><a href="https://example.com/search?q=hello&amp;lang=en">https://example.com/search?q=hello&amp;lang=en</a></p>`,
		},
		{
			Name:     "link with http and query parameters",
			Input:    "[[http://example.com/something.php?a=1&bb=22]]",
			Expected: `<p><a href="http://example.com/something.php?a=1&amp;bb=22">http://example.com/something.php?a=1&amp;bb=22</a></p>`,
		},
		{
			Name:     "bold and link",
			Input:    "**bold** and [[https://example.com]]",
			Expected: `<p><strong>bold</strong> and <a href="https://example.com">https://example.com</a></p>`,
		},

		// U+0085 next line
		{
			Name:     "U+0085 next line",
			Input:    "Line one\u0085Line two",
			Expected: "<p>Line one<br />Line two</p>",
		},

		// U+2028 line separator
		{
			Name:     "U+2028 line separator",
			Input:    "Line one\u2028Line two",
			Expected: "<p>Line one<br />Line two</p>",
		},

		// U+2029 paragraph separator
		{
			Name:     "U+2029 paragraph separator",
			Input:    "First paragraph\u2029Second paragraph",
			Expected: "<p>First paragraph</p><p>Second paragraph</p>",
		},

		// Persian numbers
		{
			Name:     "persian digits",
			Input:    "۰۱۲۳۴۵۶۷۸۹",
			Expected: "<p>۰۱۲۳۴۵۶۷۸۹</p>",
		},
		{
			Name:     "persian number in sentence",
			Input:    "قیمت ۱۲۳ تومان است",
			Expected: "<p>قیمت ۱۲۳ تومان است</p>",
		},
		{
			Name:     "bold persian number",
			Input:    "**۴۲**",
			Expected: "<p><strong>۴۲</strong></p>",
		},
		{
			Name:     "persian numbers in two paragraphs",
			Input:    "شماره ۱\n\nشماره ۲",
			Expected: "<p>شماره ۱</p><p>شماره ۲</p>",
		},

		// Persian words
		{
			Name:     "persian dorood",
			Input:    "درود",
			Expected: "<p>درود</p>",
		},
		{
			Name:     "persian dorood bold",
			Input:    "**درود**",
			Expected: "<p><strong>درود</strong></p>",
		},
		{
			Name:     "persian sentence",
			Input:    "درود بر شما",
			Expected: "<p>درود بر شما</p>",
		},
		{
			Name:     "persian sentence with formatting",
			Input:    "درود **دوست** //عزیز//",
			Expected: "<p>درود <strong>دوست</strong> <em>عزیز</em></p>",
		},
		{
			Name:     "persian two paragraphs",
			Input:    "درود بر شما\n\nخوش آمدید",
			Expected: "<p>درود بر شما</p><p>خوش آمدید</p>",
		},
		{
			Name:     "persian with line break",
			Input:    "درود\nخداحافظ",
			Expected: "<p>درود<br />خداحافظ</p>",
		},

		// Korean words
		{
			Name:     "korean plain text",
			Input:    "안녕하세요",
			Expected: "<p>안녕하세요</p>",
		},
		{
			Name:     "korean bold",
			Input:    "**안녕하세요**",
			Expected: "<p><strong>안녕하세요</strong></p>",
		},
		{
			Name:     "korean sentence with formatting",
			Input:    "이것은 **굵은** 글씨와 //기울임// 글씨입니다",
			Expected: "<p>이것은 <strong>굵은</strong> 글씨와 <em>기울임</em> 글씨입니다</p>",
		},
		{
			Name:     "korean two paragraphs",
			Input:    "첫 번째 단락\n\n두 번째 단락",
			Expected: "<p>첫 번째 단락</p><p>두 번째 단락</p>",
		},
		{
			Name:     "korean with line break",
			Input:    "안녕하세요\n감사합니다",
			Expected: "<p>안녕하세요<br />감사합니다</p>",
		},
		{
			Name:     "korean with highlight",
			Input:    "||중요한|| 내용",
			Expected: "<p><mark>중요한</mark> 내용</p>",
		},
	}

	for testNumber, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := ToHTML(test.Input)

			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual HTML from rtxt is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("INPUT:    %q", test.Input)
				return
			}
		})
	}
}

func TestToHTML_UnorderedList(t *testing.T) {
	tests := []struct {
		Name     string
		Input    string
		Expected string
	}{
		{
			Name: "grocery list",
			Input: "\t\u2022 apples\n" +
				"\t\u2022 bananas\n" +
				"\t\u2022 cherries",
			Expected: "<ul><li>apples</li><li>bananas</li><li>cherries</li></ul>",
		},
		{
			Name: "list with formatted items",
			Input: "\t\u2022 **apples**\n" +
				"\t\u2022 //bananas//\n" +
				"\t\u2022 __cherries__\n" +
				"\t\u2022 ||grapes||",
			Expected: "<ul><li><strong>apples</strong></li><li><em>bananas</em></li><li><u>cherries</u></li><li><mark>grapes</mark></li></ul>",
		},
		{
			Name: "list item with continuation",
			Input: "\t\u2022 first item\n" +
				"\tthis continues the first item\n" +
				"\t\u2022 second item",
			Expected: "<ul><li>first item<br />this continues the first item</li><li>second item</li></ul>",
		},
		{
			Name: "list with link",
			Input: "\t\u2022 visit [[https://example.com]]\n" +
				"\t\u2022 another item",
			Expected: `<ul><li>visit <a href="https://example.com">https://example.com</a></li><li>another item</li></ul>`,
		},
		{
			Name: "paragraph then list then paragraph",
			Input: "Here is a list:\n\n" +
				"\t\u2022 one\n" +
				"\t\u2022 two\n" +
				"\t\u2022 three\n\n" +
				"That was the list.",
			Expected: "<p>Here is a list:</p><ul><li>one</li><li>two</li><li>three</li></ul><p>That was the list.</p>",
		},
	}

	for testNumber, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := ToHTML(test.Input)

			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual HTML from rtxt is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("INPUT:    %q", test.Input)
				return
			}
		})
	}
}
