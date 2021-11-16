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

//StatModificator is the dnd stat modificator scale.
var StatModificator map[int]int = map[int]int{
	1:  -5,
	2:  -4,
	3:  -4,
	4:  -3,
	5:  -3,
	6:  -2,
	7:  -2,
	8:  -1,
	9:  -1,
	10: 0,
	11: 0,
	12: 1,
	13: 1,
	14: 2,
	15: 2,
	16: 3,
	17: 3,
	18: 4,
	19: 4,
	20: 5,
	21: 5,
}

//MasteryByLevel is the mastery scale system for every class in dnd
var MasteryByLevel map[int]int = map[int]int{
	1:  2,
	2:  2,
	3:  2,
	4:  2,
	5:  3,
	6:  3,
	7:  3,
	8:  3,
	9:  4,
	10: 4,
	11: 4,
	12: 4,
	13: 5,
	14: 5,
	15: 5,
	16: 5,
	17: 6,
	18: 6,
	19: 6,
	20: 6,
}
