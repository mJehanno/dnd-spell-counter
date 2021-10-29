package components

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	classmanager "github.com/mjehanno/dnd-spell-counter/class-manager"
	playermanager "github.com/mjehanno/dnd-spell-counter/player-manager"
)

var classBind binding.ExternalString
var lvlBind binding.ExternalInt
var subClassBinding binding.StringList = binding.NewStringList()

func CreateClassSelector(multi bool) *fyne.Container {
	selectedLvl := 0
	selectedClass := ""
	classBind = binding.BindString(&selectedClass)
	lvlBind = binding.BindInt(&selectedLvl)
	optionalInputC := container.NewGridWithColumns(2)
	levelOptions := createLevelArray()
	selectLvl := widget.NewSelect(
		levelOptions, func(s string) {
			lvl, _ := strconv.Atoi(s)
			lvlBind.Set(lvl)
			if multi {
				playermanager.PlayerBinding.SetValue("SecondLvl", lvl)
			} else {
				playermanager.PlayerBinding.SetValue("Lvl", lvl)
			}
			addSubClassesSelect(optionalInputC, multi)
		})

	options := []string{}

	for _, c := range classmanager.ClassList {
		options = append(options, c.Name)
	}

	selectClass := widget.NewSelect(options, func(s string) {
		classBind.Set(s)
		options := []string{}
		if multi {
			playermanager.CurrentPlayer.SecondClass = classmanager.FindClassByName(s)
			for _, v := range playermanager.CurrentPlayer.SecondClass.SubClasses {
				options = append(options, v.Name)
			}
			playermanager.PlayerBinding.SetValue("SecondClass", playermanager.CurrentPlayer.SecondClass)
		} else {
			playermanager.CurrentPlayer.Class = classmanager.FindClassByName(s)
			for _, v := range playermanager.CurrentPlayer.Class.SubClasses {
				options = append(options, v.Name)
			}
			playermanager.PlayerBinding.SetValue("Class", playermanager.CurrentPlayer.Class)
		}
		subClassBinding.Set(options)
		addSubClassesSelect(optionalInputC, multi)
	})

	mainInputC := container.NewGridWithColumns(2,
		selectClass,
		selectLvl,
	)

	return container.NewVBox(mainInputC, optionalInputC)
}

func createLevelArray() []string {
	levels := make([]int, 20)
	for i := range levels {
		levels[i] = i + 1
	}
	levelOptions := make([]string, 20)

	for i, v := range levels {
		levelOptions[i] = strconv.Itoa(v)
	}
	return levelOptions
}

func addSubClassesSelect(newC *fyne.Container, multi bool) {
	class, _ := classBind.Get()
	lvl, _ := lvlBind.Get()

	subclasses, _ := subClassBinding.Get()

	OnSubClassChange := func(s string) {
		if multi {
			playermanager.PlayerBinding.SetValue("SecondSubClass", classmanager.FindSubClassByName(s, playermanager.CurrentPlayer.SecondClass.Name))
		} else {
			playermanager.PlayerBinding.SetValue("SubClass", classmanager.FindSubClassByName(s, playermanager.CurrentPlayer.Class.Name))
		}
	}

	if class == "Cleric" || (lvl == 2 && class == "Druid" || class == "Wizard") || lvl >= 3 {
		w := widget.NewSelect(subclasses, OnSubClassChange)
		if len(newC.Objects) == 0 {
			newC.Add(w)
		} else {
			newC.Remove(newC.Objects[0])
			newC.Add(w)
			newC.Refresh()
		}
	} else {
		if len(newC.Objects) > 0 {
			newC.Remove(newC.Objects[0])
		}
		newC.Refresh()
	}
}
