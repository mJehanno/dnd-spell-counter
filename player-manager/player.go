package playermanager

import (
	"fyne.io/fyne/v2/data/binding"
	classmanager "github.com/mjehanno/dnd-spell-counter/class-manager"
)

//PlayerBinding is a Fyne binding on the current player
var PlayerBinding binding.Struct

//CurrentPlayer is the current player (aka user)
var CurrentPlayer *Player

//Player is a struct defining a user
type Player struct {
	MultiClass       bool
	Class            classmanager.Class
	Lvl              int
	SubClass         classmanager.SubClass
	SecondClass      classmanager.Class
	SecondLvl        int
	SecondSubClass   classmanager.SubClass
	FeatsValue       int
	SecondFeatsValue int
}

const (
	MultiClass       = "MultiClass"
	Class            = "Class"
	Lvl              = "Lvl"
	SubClass         = "SubClass"
	SecondClass      = "SecondClass"
	SecondLvl        = "SecondLvl"
	SecondSubClass   = "SecondSubClass"
	FeatsValue       = "FeatsValue"
	SecondFeatsValue = "SecondFeatsValue"
)

// GetStatModificator return the modificator to apply for the given stat value.
func GetStatModificator(stat int) int {
	if stat < 10 {
		return (stat - 11) / 2
	}
	return (stat - 10) / 2
}

// GetMasteryByLevel return the mastery bonus for the given lvl
func GetMasteryByLevel(lvl int) int {
	if lvl%4 == 0 {
		return lvl/4 + 1
	}
	return lvl/4 + 2
}
