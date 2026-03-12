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
