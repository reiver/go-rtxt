package rtxt

import (
	"testing"
)

func TestToHTML_unorderedList(t *testing.T) {
	tests := []struct {
		Name     string
		Input    string
		Expected string
	}{
		{
			Name: "(no final newline) I like to eat, eat, eat (1)....",
			Input:
				"\t"+ "\u2022 "+ "apples"   +""+
				"",
			Expected:
				`<ul>`+
					`<li>`+
						`apples`+
					`</li>`+
				`</ul>`,
		},
		{
			Name: "(yes final newline) I like to eat, eat, eat (1)....",
			Input:
				"\t"+ "\u2022 "+ "apples"   +"\n"+
				"",
			Expected:
				`<ul>`+
					`<li>`+
						`apples`+
					`</li>`+
				`</ul>`,
		},

		{
			Name: "(no final newline) I like to eat, eat, eat (2)....",
			Input:
				"\t"+ "\u2022 "+ "apples"   +"\n"+
				"\t"+ "\u2022 "+ "bananas"  +""+
				"",
			Expected:
				`<ul>`+
					`<li>`+
						`apples`+
					`</li>`+
					`<li>`+
						`bananas`+
					`</li>`+
				`</ul>`,
		},
		{
			Name: "(yes final newline) I like to eat, eat, eat (2)....",
			Input:
				"\t"+ "\u2022 "+ "apples"   +"\n"+
				"\t"+ "\u2022 "+ "bananas"  +"\n"+
				"",
			Expected:
				`<ul>`+
					`<li>`+
						`apples`+
					`</li>`+
					`<li>`+
						`bananas`+
					`</li>`+
				`</ul>`,
		},

		{
			Name: "(no final newline) I like to eat, eat, eat (3)....",
			Input:
				"\t"+ "\u2022 "+ "apples"   +"\n"+
				"\t"+ "\u2022 "+ "bananas"  +"\n"+
				"\t"+ "\u2022 "+ "cherries" +""+
				"",
			Expected:
				`<ul>`+
					`<li>`+
						`apples`+
					`</li>`+
					`<li>`+
						`bananas`+
					`</li>`+
					`<li>`+
						`cherries`+
					`</li>`+
				`</ul>`,
		},
		{
			Name: "(yes final newline) I like to eat, eat, eat (3)....",
			Input:
				"\t"+ "\u2022 "+ "apples"   +"\n"+
				"\t"+ "\u2022 "+ "bananas"  +"\n"+
				"\t"+ "\u2022 "+ "cherries" +"\n"+
				"",
			Expected:
				`<ul>`+
					`<li>`+
						`apples`+
					`</li>`+
					`<li>`+
						`bananas`+
					`</li>`+
					`<li>`+
						`cherries`+
					`</li>`+
				`</ul>`,
		},



		{
			Name: "paragraph then unordered list",
			Input:
				"Look at these:"            +"\n"+
				""                          +"\n"+
				"\t"+ "\u2022 "+ "apples"   +"\n"+
				"\t"+ "\u2022 "+ "bananas"  +"\n"+
				"\t"+ "\u2022 "+ "cherries" +"\n"+
				"",
			Expected:
				`<p>`+
					`Look at these:`+
				`</p>`+
				`<ul>`+
					`<li>`+
						`apples`+
					`</li>`+
					`<li>`+
						`bananas`+
					`</li>`+
					`<li>`+
						`cherries`+
					`</li>`+
				`</ul>`,
		},
		{
			Name: "paragraph then unordered list then paragraph",
			Input:
				"Look at these:"            +"\n"+
				""                          +"\n"+
				"\t"+ "\u2022 "+ "apples"   +"\n"+
				"\t"+ "\u2022 "+ "bananas"  +"\n"+
				"\t"+ "\u2022 "+ "cherries" +"\n"+
				""                          +"\n"+
				"They are delicious!"       +"\n"+
				"",
			Expected:
				`<p>`+
					`Look at these:`+
				`</p>`+
				`<ul>`+
					`<li>`+
						`apples`+
					`</li>`+
					`<li>`+
						`bananas`+
					`</li>`+
					`<li>`+
						`cherries`+
					`</li>`+
				`</ul>`+
				`<p>`+
					`They are delicious!`+
				`</p>`+
				"",
		},



		{
			Name: "(no final newline) unordered list with formatted items",
			Input:
				"\t"+ "\u2022 "+ "**apples**"   +"\n"+
				"\t"+ "\u2022 "+ "//bananas//"  +"\n"+
				"\t"+ "\u2022 "+ "__cherries__" +"\n"+
				"\t"+ "\u2022 "+ "||dates||"    +""+
				"",
			Expected:
				`<ul>`+
					`<li>`+
						`<strong>apples</strong>`+
					`</li>`+
					`<li>`+
						`<em>bananas</em>`+
					`</li>`+
					`<li>`+
						`<u>cherries</u>`+
					`</li>`+
					`<li>`+
						`<mark>dates</mark>`+
					`</li>`+
				`</ul>`+
				"",
		},
		{
			Name: "(yes final newline) unordered list with formatted items",
			Input:
				"\t"+ "\u2022 "+ "**apples**"   +"\n"+
				"\t"+ "\u2022 "+ "//bananas//"  +"\n"+
				"\t"+ "\u2022 "+ "__cherries__" +"\n"+
				"\t"+ "\u2022 "+ "||dates||"    +"\n"+
				"",
			Expected:
				`<ul>`+
					`<li>`+
						`<strong>apples</strong>`+
					`</li>`+
					`<li>`+
						`<em>bananas</em>`+
					`</li>`+
					`<li>`+
						`<u>cherries</u>`+
					`</li>`+
					`<li>`+
						`<mark>dates</mark>`+
					`</li>`+
				`</ul>`+
				"",
		},



		{
			Name: "(no final newline) unordered list item with continuation",
			Input:
				"\t"+ "\u2022 "+ "first item"                    +"\n"+
				"\t"+            "this continues the first item" +"\n"+
				"\t"+ "\u2022 "+ "second item"                   +"\n"+
				"\t"+ "\u2022 "+ "third item"                    +"\n"+
				"\t"+            "to be"                         +"\n"+
				"\t"+            "continued"                     +""+
				"",
			Expected:
				`<ul>`+
					`<li>`+
						`first item`+
						`<br />`+
						`this continues the first item`+
					`</li>`+
					`<li>`+
						`second item`+
					`</li>`+
					`<li>`+
						`third item`+
						`<br />`+
						`to be`+
						`<br />`+
						`continued`+
					`</li>`+
				`</ul>`+
				"",
		},
		{
			Name: "(yes final newline) unordered list item with continuation",
			Input:
				"\t"+ "\u2022 "+ "first item"                    +"\n"+
				"\t"+            "this continues the first item" +"\n"+
				"\t"+ "\u2022 "+ "second item"                   +"\n"+
				"\t"+ "\u2022 "+ "third item"                    +"\n"+
				"\t"+            "to be"                         +"\n"+
				"\t"+            "continued"                     +"\n"+
				"",
			Expected:
				`<ul>`+
					`<li>`+
						`first item`+
						`<br />`+
						`this continues the first item`+
					`</li>`+
					`<li>`+
						`second item`+
					`</li>`+
					`<li>`+
						`third item`+
						`<br />`+
						`to be`+
						`<br />`+
						`continued`+
					`</li>`+
				`</ul>`+
				"",
		},



		{
			Name: "unordered list with link",
			Input:
				"\t"+ "\u2022 " +"visit [[https://example.com]]" +"\n"+
				"\t"+ "\u2022 " +"another item"                  +""+
				"",
			Expected:
				`<ul>`+
					`<li>`+
						`visit <a href="https://example.com">https://example.com</a>`+
					`</li>`+
					`<li>`+
						`another item`+
					`</li>`+
				`</ul>`+
				"",
		},



		{
			Name: "(no final newline) paragraph then list then paragraph",
			Input:
				"Here is a list:"         +"\n"+
				""                        +"\n"+
				"\t"+ "\u2022 "+ "once"   +"\n"+
				"\t"+ "\u2022 "+ "twice"  +"\n"+
				"\t"+ "\u2022 "+ "thrice" +"\n"+
				"\t"+ "\u2022 "+ "fource" +"\n"+
				""                        +"\n"+
				"That was the list."      +""+
				"",
			Expected:
				`<p>`+
					`Here is a list:`+
				`</p>`+
				`<ul>`+
					`<li>`+
						`once`+
					`</li>`+
					`<li>`+
						`twice`+
					`</li>`+
					`<li>`+
						`thrice`+
					`</li>`+
					`<li>`+
						`fource`+
					`</li>`+
				`</ul>`+
				`<p>`+
					`That was the list.`+
				`</p>`+
				"",
		},
		{
			Name: "(yes final newline) paragraph then list then paragraph",
			Input:
				"Here is a list:"         +"\n"+
				""                        +"\n"+
				"\t"+ "\u2022 "+ "once"   +"\n"+
				"\t"+ "\u2022 "+ "twice"  +"\n"+
				"\t"+ "\u2022 "+ "thrice" +"\n"+
				"\t"+ "\u2022 "+ "fource" +"\n"+
				""                        +"\n"+
				"That was the list."      +"\n"+
				"",
			Expected:
				`<p>`+
					`Here is a list:`+
				`</p>`+
				`<ul>`+
					`<li>`+
						`once`+
					`</li>`+
					`<li>`+
						`twice`+
					`</li>`+
					`<li>`+
						`thrice`+
					`</li>`+
					`<li>`+
						`fource`+
					`</li>`+
				`</ul>`+
				`<p>`+
					`That was the list.`+
				`</p>`+
				"",
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
