package unittesting

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) { // name convention for func TextXxxxx
	actual := add(20, 30)
	var expected int64 = 50

	if actual != expected {
		t.Error("Actual is : ", actual, "but expected is : ", expected)
	}

}

func TestAddBulk(t *testing.T) {
	var list = []struct {
		first    int
		second   int
		expected int64
	}{
		{
			20, 30, 50,
		},
		{
			-20, 30, 10,
		},
		{
			-20, -30, -50,
		},
		{
			20, -30, -10,
		},
	}
	fmt.Println(list)
	for _, str := range list {
		actual := add(str.first, str.second)
		if actual != str.expected {
			t.Error("Actual is : ", actual, "but expected is : ", str.expected)
		}
	}
}
