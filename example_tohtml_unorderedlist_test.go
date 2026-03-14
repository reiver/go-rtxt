package rtxt_test

import (
	"fmt"

	"github.com/reiver/go-rtxt"
)

func ExampleToHTML_unorderedList() {
	input :=
		"\t"+ "• apples"   +"\n" +
		"\t"+ "• bananas"  +"\n" +
		"\t"+ "• cherries" +"\n"

	result := rtxt.ToHTML(input)

	fmt.Println(result)

	// Output: <ul><li>apples</li><li>bananas</li><li>cherries</li></ul>
}
