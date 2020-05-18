package prose

import (
	"fmt"
	"testing"
)

// file name should end with "_test.go"
func TestThreeElements(t *testing.T) { // func name should start with Test followed by anything
	list := []string{"apple", "oranges", "mango"}
	actual := JoinWithCommas(list)
	expected := "apple ,oranges and mango"
	if actual != expected {
		t.Error(errorString(list, actual, expected))
		// just like printf there is Errorf method also for formatted error display
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "mango"}
	actual := JoinWithCommas(list)
	expected := "apple and mango"
	if actual != expected {
		t.Error(errorString(list, actual, expected))
	}
}

func errorString(list []string, actual, expected string) string {
	return fmt.Sprintf("joinWithCommas(%#v) =\"%s\",want \"%s\"", list, actual, expected)
	// Sprintf returns a formatted string
}
