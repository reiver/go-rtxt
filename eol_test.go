package rtxt

import (
	"testing"
)

func TestEOLIndex(t *testing.T) {
	tests := []struct{
		String         string
		ExpectedIndex  int
		ExpectedLength int
	}{
		{
			String: "",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},



		{
			String: "e",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "ex",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "exa",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "exam",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "examp",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "exampl",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example ",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example l",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example li",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example lin",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line ",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line w",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line wi",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line wit",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line with",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line with ",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line with n",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line with no",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line with no ",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line with no e",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line with no eo",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},
		{
			String: "example line with no eol",
			ExpectedIndex: -1,
			ExpectedLength: 0,
		},



		{
			String: "example line\r\n",
			ExpectedIndex: 12,
			ExpectedLength: 2,
		},
		{
			String: "example line\n",
			ExpectedIndex: 12,
			ExpectedLength: 1,
		},
		{
			String: "example line\u0085",
			ExpectedIndex: 12,
			ExpectedLength: 2,
		},
		{
			String: "example line\u2028",
			ExpectedIndex: 12,
			ExpectedLength: 3,
		},
		{
			String: "example line\u2029",
			ExpectedIndex: 12,
			ExpectedLength: 3,
		},



		{
			String: "example line\r\nline two",
			ExpectedIndex: 12,
			ExpectedLength: 2,
		},
		{
			String: "example line\nline two",
			ExpectedIndex: 12,
			ExpectedLength: 1,
		},
		{
			String: "example line\u0085line two",
			ExpectedIndex: 12,
			ExpectedLength: 2,
		},
		{
			String: "example line\u2028line two",
			ExpectedIndex: 12,
			ExpectedLength: 3,
		},
		{
			String: "example line\u2029line two",
			ExpectedIndex: 12,
			ExpectedLength: 3,
		},
	}

	for testNumber, test := range tests {

		actualIndex, actualLength := eolIndex(test.String)

		if test.ExpectedIndex != actualIndex || test.ExpectedLength != actualLength {
			t.Errorf("For test #%d, the actual EOL-index is not what was expected.", testNumber)
			t.Logf("EXPECTED INDEX:  %d", test.ExpectedIndex)
			t.Logf("ACTUAL   INDEX:  %d", actualIndex)
			t.Logf("EXPECTED LENGTH: %d", test.ExpectedLength)
			t.Logf("ACTUAL   LENGTH: %d", actualLength)
			t.Logf("STRING: %q", test.String)
			continue
		}
	}
}
