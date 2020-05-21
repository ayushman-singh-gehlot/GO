package components

import (
	"fmt"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	tests := []struct {
		input     int
		expected  int
		expected1 int
		expected2 string
	}{
		{2, 2, 4, " "},
		{3, 3, 9, " "},
		{4, 4, 16, " "},
	}
	for _, test := range tests {
		var b NewBoard
		b.CreateBoard(test.input)
		if b.size != test.expected {
			t.Error(errorString(test, b.size, test.expected))
		}
		if len(b.values) != test.expected1 {
			t.Error(errorString(test, len(b.values), test.expected1))
		}
		for _, value := range b.values {
			if value != test.expected2 {
				t.Error(errorString(test, value, test.expected2))
			}
		}
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		input1   int
		input2   int
		input3   string
		expected string
	}{
		{2, 0, "O", "O"},
		{3, 3, "O", "O"},
		{4, 11, "O", "O"},
		{3, 8, "X", "X"},
		{4, 15, "X", "X"},
		{2, 3, "X", "X"},
		{3, 8, "X", "X"},
		{5, 16, "X", "X"}, // for any general case
	}
	for _, test := range tests {
		var b NewBoard
		b.CreateBoard(test.input1) // to test update i m using CreateBoard .....can we have this dependency?
		err := b.Update(test.input2, test.input3)
		if err != nil {
			t.Error(err)
		} else if b.values[test.input2] != test.expected {
			t.Error(errorString(test, test.input3, test.expected))
		}
	}
}
func TestCheckRows(t *testing.T) {
	tests := []struct {
		size     int
		values   []string
		input    string
		expected bool
	}{
		{2, []string{"X", "X", " ", " "}, "X", true},
		{3, []string{" ", " ", " ", "O", "O", "O", " ", " ", " "}, "X", false},
		{4, []string{"X", "X", " X", "O", "O", "O", " X", " O", "X", "X", "X", "X", " ", " ", " ", " "}, "X", true},
	}
	for _, test := range tests {
		var b NewBoard
		b.size = test.size
		b.values = append(b.values, test.values...)
		if checkRows(&b, test.input) != test.expected {
			t.Error(errorString(test, test.input, test.expected))
		}
	}
}
func TestCheckColumns(t *testing.T) {
	tests := []struct {
		size     int
		values   []string
		input    string
		expected bool
	}{
		{2, []string{" ", "O", " ", "O"}, "O", true},
		{3, []string{" ", "X", " ", "X", "O", "O", " ", "X", " "}, "X", false},
		{4, []string{"X", "X", "X", "O", "O", "O", "X", "O", "X", "X", "X", "X", " ", " ", "X", " "}, "X", true},
	}
	for _, test := range tests {
		var b NewBoard
		b.size = test.size
		b.values = append(b.values, test.values...)
		if checkColumns(&b, test.input) != test.expected {
			t.Error(errorString(test, test.input, test.expected))
		}
	}
}

func TestCheckDiagonal(t *testing.T) {
	tests := []struct {
		size     int
		values   []string
		input    string
		expected bool
	}{
		{2, []string{" ", "O", "O", " "}, "O", true},
		{3, []string{"O", "X", "X", "O", "X", "O", "X", "O", " "}, "X", true},
		{4, []string{"X", "X", "X", "O", "O", "O", "X", "O", "X", "X", "X", "X", " ", " ", "X", " "}, "X", false},
	}
	for _, test := range tests {
		var b NewBoard
		b.size = test.size
		b.values = append(b.values, test.values...)
		if checkDiagonal(&b, test.input) != test.expected {
			t.Error(errorString(test, test.input, test.expected))
		}
	}
}
func TestCheckFull(t *testing.T) {
	tests := []struct {
		size     int
		values   []string
		expected bool
	}{
		{2, []string{"X", "O", "O", " "}, false},
		{3, []string{"O", "X", "X", "O", "X", "O", "X", "O", "X"}, true},
		{4, []string{"X", "X", "X", "O", "O", "O", "X", "O", "X", "X", "X", "X", " ", " ", "X", " "}, false},
	}
	for _, test := range tests {
		var b NewBoard
		b.size = test.size
		b.values = append(b.values, test.values...)
		if checkFull(&b) != test.expected {
			t.Error(errorString(test, test.values, test.expected))
		}
	}
}

func errorString(test, actual, expected interface{}) string {
	return fmt.Sprintf("fail :\ntest : %v\tactual : %v\texpected : %v", test, actual, expected)
}
