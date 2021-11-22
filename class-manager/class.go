package classmanager

import (
	"encoding/json"
)

//Class is a dnd character class
type Class struct {
	Name                                     string
	HasSpell, HasSpellWithSubClass, HasFeats bool
	Feats                                    Feats
	FeatsAmountByLevel                       map[int]int
	SubClasses                               []SubClass
	SpellLevelLimit                          int
	SpellByLevel                             map[int]map[int]int
}

//String method return a string representation of class struct
func (c Class) String() string {
	b, _ := json.Marshal(c)

	return string(b)
}

const (
	Artificer   = "Artificer"
	Barbarian   = "Barbarian"
	Bard        = "Bard"
	BloodHunter = "BloodHunter"
	Cleric      = "Cleric"
	Druid       = "Druid"
	Fighter     = "Fighter"
	Monk        = "Monk"
	Paladin     = "Paladin"
	Ranger      = "Ranger"
	Rogue       = "Rogue"
	Sorcerer    = "Sorcerer"
	Warlock     = "Warlock"
	Wizard      = "Wizard"
)

//SubClass is a dnd character subclass (usually a choice made at level 3)
type SubClass struct {
	Name               string
	HasSpell, HasFeats bool
	SpellLevelLimit    int
}

//Abilities are character's basic stat
type Abilities string

const (
	//None is the default ability
	None Abilities = "none"
	//Constitution stat
	Constitution Abilities = "cons"
	//Charism stat
	Charism Abilities = "char"
	//Strength stat
	Strength Abilities = "str"
	//Agility stat
	Agility Abilities = "agi"
	//Intelligence stat
	Intelligence Abilities = "int"
	//Wisdom stat
	Wisdom Abilities = "wis"
	//Mastery stat
	Mastery Abilities = "mas"
)

//Feats define a class special feature like Ki for monk or Rage for Barbarian
type Feats struct {
	Name         string
	CharModifier Abilities
}

var (
	//NoFeats default value
	NoFeats = Feats{"", None}
	//BardicInspiration default value
	BardicInspiration = Feats{"Bardic Inspiration", Charism}
	//HemocraftDice default value
	HemocraftDice = Feats{"EmocraftDice", None}
	//Infusion default value
	Infusion = Feats{"Infusion", None}
	//Ki default value
	Ki = Feats{"Ki", None}
	//PsionicEnergyDice default value
	PsionicEnergyDice = Feats{"Psionic Energy Dice", Mastery}
	//Rage default value
	Rage = Feats{"Rage", None}
	//SorceryPoint default value
	SorceryPoint = Feats{"Sorcery Point", None}
)

//ClassList is a list of dnd classes and subclasses
var ClassList []Class = []Class{}

//FindClassByName will look for the given name in the classlist
func FindClassByName(name string) Class {
	var class Class
	for _, v := range ParseClassList() {
		if v.Name == name {
			class = v
		}
	}
	return class
}

//ParseClassList read the classlist from bundled json file.
func ParseClassList() []Class {

	err := json.Unmarshal(resourceClasslistJson.StaticContent, &ClassList)
	if err != nil {
		panic("error while parsing classlist")
	}
	return ClassList
}

//FindSubClassByName will look for subclasses with the given name within the given class
func FindSubClassByName(name string, className string) SubClass {
	var sub SubClass

	for _, v := range FindClassByName(className).SubClasses {
		if v.Name == name {
			sub = v
		}
	}
	return sub
}
