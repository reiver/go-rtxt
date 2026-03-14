package rtxt

import (
	"testing"
)

func TestEOLIndex(t *testing.T) {
	tests := []struct{
		String string
		Expected int
	}{
		{
			String: "",
			Expected: 0,
		},



		{
			String: "e",
			Expected: 1,
		},
		{
			String: "ex",
			Expected: 2,
		},
		{
			String: "exa",
			Expected: 3,
		},
		{
			String: "exam",
			Expected: 4,
		},
		{
			String: "examp",
			Expected: 5,
		},
		{
			String: "exampl",
			Expected: 6,
		},
		{
			String: "example",
			Expected: 7,
		},
		{
			String: "example ",
			Expected: 8,
		},
		{
			String: "example l",
			Expected: 9,
		},
		{
			String: "example li",
			Expected: 10,
		},
		{
			String: "example lin",
			Expected: 11,
		},
		{
			String: "example line",
			Expected: 12,
		},
		{
			String: "example line ",
			Expected: 13,
		},
		{
			String: "example line w",
			Expected: 14,
		},
		{
			String: "example line wi",
			Expected: 15,
		},
		{
			String: "example line wit",
			Expected: 16,
		},
		{
			String: "example line with",
			Expected: 17,
		},
		{
			String: "example line with ",
			Expected: 18,
		},
		{
			String: "example line with n",
			Expected: 19,
		},
		{
			String: "example line with no",
			Expected: 20,
		},
		{
			String: "example line with no ",
			Expected: 21,
		},
		{
			String: "example line with no e",
			Expected: 22,
		},
		{
			String: "example line with no eo",
			Expected: 23,
		},
		{
			String: "example line with no eol",
			Expected: 24,
		},



		{
			String: "example line\r\n",
			Expected: 12,
		},
		{
			String: "example line\n",
			Expected: 12,
		},
		{
			String: "example line\u0085",
			Expected: 12,
		},
		{
			String: "example line\u2028",
			Expected: 12,
		},
		{
			String: "example line\u2029",
			Expected: 12,
		},



		{
			String: "example line\r\nline two",
			Expected: 12,
		},
		{
			String: "example line\nline two",
			Expected: 12,
		},
		{
			String: "example line\u0085line two",
			Expected: 12,
		},
		{
			String: "example line\u2028line two",
			Expected: 12,
		},
		{
			String: "example line\u2029line two",
			Expected: 12,
		},
	}

	for testNumber, test := range tests {

		actual := eolIndex(test.String)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual EOL-index is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("STRING: %q", test.String)
			continue
		}
	}
}
