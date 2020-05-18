package prose

import (
	"testing"
)

type testData struct {
	list     []string
	expected string
}

// table driven test is good approach as adding new test is easy
func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{[]string{"apple"}, "apple"},
		{[]string{"apple", "orange"}, "apple and orange"},
		{[]string{"apple", "orange", "mango"}, "apple ,orange and mango"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.expected {
			t.Error(errorString(test.list, got, test.expected))
		}
	}
}
