package rtxt

import (
	"testing"
)

func TestToHTML_paragraph(t *testing.T) {
	tests := []struct {
		Name     string
		Input    string
		Expected string
	}{
		{
			Name: "(no final newlie) one paragraph with one sentence.",
			Input:
				`Hello world!` +""+
				"",
			Expected:
				`<p>`+
					`Hello world!`+
				`</p>`,
		},
		{
			Name: "(yes final newlie) one paragraph with one sentence.",
			Input:
				`Hello world!` +"\n"+
				"",
			Expected:
				`<p>`+
					`Hello world!`+
				`</p>`,
		},



		{
			Name: "(no final newlie) one paragraph with two sentences.",
			Input:
				`Hello world!` +"\n"+
				`How do //you// do?` +""+
				"",
			Expected:
				`<p>`+
					`Hello world!`+
					`<br />`+
					`How do <em>you</em> do?`+
				`</p>`,
		},
		{
			Name: "(yes final newlie) one paragraph with two sentences.",
			Input:
				`Hello world!` +"\n"+
				`How do //you// do?` +""+
				"",
			Expected:
				`<p>`+
					`Hello world!`+
					`<br />`+
					`How do <em>you</em> do?`+
				`</p>`,
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
