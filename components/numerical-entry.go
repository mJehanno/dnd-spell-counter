package components

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

// NumericalEntry is a numerical input
type NumericalEntry struct {
	widget.Entry
}

// NewNumericalEntry generate a new numerical input
func NewNumericalEntry() *NumericalEntry {
	entry := &NumericalEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

//TypedRune enforce the available rune in the input to numeral value
func (e *NumericalEntry) TypedRune(r rune) {
	switch r {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		e.Entry.TypedRune(r)
	}
}

//TypedShortcut enforce the copy/paste to be a number
func (e *NumericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

//Keyboard enforce the number keyboard to be displayed when clikcing on the input on mobile device.
func (e *NumericalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}
