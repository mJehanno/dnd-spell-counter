package components

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	classmanager "github.com/mjehanno/dnd-spell-counter/class-manager"
	playermanager "github.com/mjehanno/dnd-spell-counter/player-manager"
)

var caracContainer *fyne.Container = container.NewHBox()
var bottomLayout *fyne.Container = container.NewGridWithColumns(2)

func HandleBardSelected(container *fyne.Container, className string) {
	if len(container.Objects) == 0 {
		container.Add(widget.NewLabel("Charism : "))

		w := NewNumericalEntry()
		w.OnChanged = func(s string) {
			value, _ := strconv.Atoi(s)
			if playermanager.GetStatModificator(value) < 1 {
				value = 1
			} else {
				value = playermanager.GetStatModificator(value)
			}
			if className == classmanager.Bard {

				playermanager.PlayerBinding.SetValue(playermanager.FeatsValue, value)
			} else {
				playermanager.PlayerBinding.SetValue(playermanager.SecondFeatsValue, value)
			}
		}
		container.Add(w)
		container.Refresh()
	}
}

func HandleFeatClass(className, secClassName string) {
	if className == classmanager.Artificer || className == classmanager.Barbarian || className == classmanager.Monk || className == classmanager.Sorcerer {
		value := playermanager.CurrentPlayer.Class.FeatsAmountByLevel[playermanager.CurrentPlayer.Lvl]
		playermanager.PlayerBinding.SetValue(playermanager.FeatsValue, value)
	}
	if secClassName == classmanager.Artificer || secClassName == classmanager.Barbarian || secClassName == classmanager.Monk || secClassName == classmanager.Sorcerer {
		value := playermanager.CurrentPlayer.SecondClass.FeatsAmountByLevel[playermanager.CurrentPlayer.SecondLvl]
		playermanager.PlayerBinding.SetValue(playermanager.SecondFeatsValue, value)
	}
}
func HandleHasNoFeat() {
	if !playermanager.CurrentPlayer.Class.HasFeats {
		playermanager.PlayerBinding.SetValue(playermanager.FeatsValue, 0)
	}

	if !playermanager.CurrentPlayer.SecondClass.HasFeats {
		playermanager.PlayerBinding.SetValue(playermanager.SecondFeatsValue, 0)
	}
}

func OnClassChanged() {
	class := playermanager.CurrentPlayer.Class.Name
	secClass := playermanager.CurrentPlayer.SecondClass.Name
	subClass := playermanager.CurrentPlayer.SubClass.Name
	secSubClass := playermanager.CurrentPlayer.SecondSubClass.Name

	if class != classmanager.Bard && secClass != classmanager.Bard {
		caracContainer.Objects = caracContainer.Objects[:0]
		caracContainer.Refresh()
	}

	if class == classmanager.Bard || secClass == classmanager.Bard {
		HandleBardSelected(caracContainer, class)
	}
	if (class == classmanager.Fighter && subClass == "Psi Warrior") ||
		(secClass == classmanager.Fighter && secSubClass == "Psi Warrior") {
		if class == classmanager.Fighter {
			value := playermanager.GetMasteryByLevel(playermanager.CurrentPlayer.Lvl)
			playermanager.PlayerBinding.SetValue(playermanager.FeatsValue, value)
		} else {
			value := playermanager.GetMasteryByLevel(playermanager.CurrentPlayer.SecondLvl)
			playermanager.PlayerBinding.SetValue(playermanager.SecondFeatsValue, value)
		}
	}

	HandleFeatClass(class, secClass)
	HandleHasNoFeat()
}

func OnSubChanged() {}

func OnLvlChanged() {
	class := playermanager.CurrentPlayer.Class.Name
	secClass := playermanager.CurrentPlayer.SecondClass.Name
	if class == classmanager.Artificer || class == classmanager.Barbarian || class == classmanager.Monk || class == classmanager.Sorcerer {
		value := playermanager.CurrentPlayer.Class.FeatsAmountByLevel[playermanager.CurrentPlayer.Lvl]
		playermanager.PlayerBinding.SetValue(playermanager.FeatsValue, value)
	}

	if secClass == classmanager.Artificer || secClass == classmanager.Barbarian || secClass == classmanager.Monk || secClass == classmanager.Sorcerer {
		value := playermanager.CurrentPlayer.SecondClass.FeatsAmountByLevel[playermanager.CurrentPlayer.SecondLvl]
		playermanager.PlayerBinding.SetValue(playermanager.SecondFeatsValue, value)
	}
}

func drawFirstPanel() {
	var temp fyne.CanvasObject
	if len(bottomLayout.Objects) > 0 {
		if len(bottomLayout.Objects) > 1 {
			temp = bottomLayout.Objects[1]
		}

		bottomLayout.Objects = bottomLayout.Objects[:0]
	}
	bottomLayout.Refresh()
	c := CreateSkillPanel(
		playermanager.CurrentPlayer.Class,
		playermanager.CurrentPlayer.Lvl,
		playermanager.CurrentPlayer.SubClass,
		playermanager.CurrentPlayer.FeatsValue,
	)
	bottomLayout.Add(c)
	if temp != nil {
		bottomLayout.Add(temp)
	}

	bottomLayout.Refresh()
}

func drawSecondPanel() {
	if len(bottomLayout.Objects) > 0 {
		bottomLayout.Objects = bottomLayout.Objects[:1]
	}
	bottomLayout.Refresh()
	c := CreateSkillPanel(
		playermanager.CurrentPlayer.SecondClass,
		playermanager.CurrentPlayer.SecondLvl,
		playermanager.CurrentPlayer.SecondSubClass,
		playermanager.CurrentPlayer.SecondFeatsValue,
	)
	bottomLayout.Add(c)
	bottomLayout.Refresh()
}
