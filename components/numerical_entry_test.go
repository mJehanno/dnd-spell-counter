package components

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	value    string
	expected string
}

func TestNumericalEntry_WithWrittenValue(t *testing.T) {
	cases := []TestCase{
		{"125", "125"},
		{"72.5", "725"},
		{"abcd", ""},
	}

	for _, v := range cases {
		entry := NewNumericalEntry()
		test.Type(entry, v.value)
		assert.Equal(t, v.expected, entry.Text)
	}
}
func TestNumericalEntry_WithPastValue(t *testing.T) {
	cases := []TestCase{
		{"125", "125"},
		{"72.5", "725"},
		{"abcd", ""},
	}

	for _, v := range cases {
		entry := NewNumericalEntry()
		clip := test.NewClipboard()
		clip.SetContent(v.value)
		test.Type(entry, clip.Content())
		assert.Equal(t, v.expected, entry.Text)
	}

}
