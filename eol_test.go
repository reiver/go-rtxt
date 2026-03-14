package rtxt

import (
	"testing"
)

func TestEOLIndex(t *testing.T) {
	tests := []struct{
		String      string
		ExpectedIndex int
		ExpectedEOL   string
	}{
		{
			String: "",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},



		{
			String: "e",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "ex",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "exa",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "exam",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "examp",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "exampl",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example ",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example l",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example li",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example lin",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line ",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line w",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line wi",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line wit",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line with",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line with ",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line with n",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line with no",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line with no ",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line with no e",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line with no eo",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},
		{
			String: "example line with no eol",
			ExpectedIndex: -1,
			ExpectedEOL: "",
		},



		{
			String: "example line\r\n",
			ExpectedIndex: 12,
			ExpectedEOL: "\r\n",
		},
		{
			String: "example line\n",
			ExpectedIndex: 12,
			ExpectedEOL: "\n",
		},
		{
			String: "example line\u0085",
			ExpectedIndex: 12,
			ExpectedEOL: "\u0085",
		},
		{
			String: "example line\u2028",
			ExpectedIndex: 12,
			ExpectedEOL: "\u2028",
		},
		{
			String: "example line\u2029",
			ExpectedIndex: 12,
			ExpectedEOL: "\u2029",
		},



		{
			String: "example line\r\nline two",
			ExpectedIndex: 12,
			ExpectedEOL: "\r\n",
		},
		{
			String: "example line\nline two",
			ExpectedIndex: 12,
			ExpectedEOL: "\n",
		},
		{
			String: "example line\u0085line two",
			ExpectedIndex: 12,
			ExpectedEOL: "\u0085",
		},
		{
			String: "example line\u2028line two",
			ExpectedIndex: 12,
			ExpectedEOL: "\u2028",
		},
		{
			String: "example line\u2029line two",
			ExpectedIndex: 12,
			ExpectedEOL: "\u2029",
		},
	}

	for testNumber, test := range tests {

		actualIndex, actualEOL := eolIndex(test.String)

		if test.ExpectedIndex != actualIndex || test.ExpectedEOL != actualEOL {
			t.Errorf("For test #%d, the actual EOL-index is not what was expected.", testNumber)
			t.Logf("EXPECTED INDEX: %d", test.ExpectedIndex)
			t.Logf("ACTUAL   INDEX: %d", actualIndex)
			t.Logf("EXPECTED EOL:   %q", test.ExpectedEOL)
			t.Logf("ACTUAL   EOL:   %q", actualEOL)
			t.Logf("STRING: %q", test.String)
			continue
		}
	}
}
