package playermanager

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    int
	expected int
}

func TestGetStatModificator(t *testing.T) {

	cases := []TestCase{
		{1, -5},
		{3, -4},
		{4, -3},
		{10, 0},
		{14, 2},
		{18, 4},
		{21, 5},
	}

	for i, v := range cases {
		r := GetStatModificator(v.input)
		if r != v.expected {
			t.Error(fmt.Sprintf("Error in test case number %v, expected %v, got %v", i, v.expected, r))
			t.Fail()
		}
	}
}

func TestGetMasteryByLevel(t *testing.T) {
	cases := []TestCase{
		{1, 2},
		{4, 2},
		{5, 3},
		{7, 3},
		{10, 4},
		{12, 4},
		{18, 6},
		{20, 6},
	}

	for i, v := range cases {
		r := GetMasteryByLevel(v.input)
		if r != v.expected {
			t.Error(fmt.Sprintf("Error in test case number %v, expected %v, got %v", i, v.expected, r))
			t.Fail()
		}
	}
}
