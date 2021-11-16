package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	classmanager "github.com/mjehanno/dnd-spell-counter/class-manager"
	"github.com/mjehanno/dnd-spell-counter/components"
	playermanager "github.com/mjehanno/dnd-spell-counter/player-manager"
)

func main() {

	spellCounterApp := app.NewWithID("com.mjehanno.dnd-spellcounter")
	playermanager.CurrentPlayer = new(playermanager.Player)

	playermanager.PlayerBinding = binding.BindStruct(playermanager.CurrentPlayer)
	spellCounterApp.SetIcon(resourceIconPng)
	mainW := spellCounterApp.NewWindow("D&D Spell-counter")
	mainC := container.NewGridWithColumns(2, createTopLayout(), createBottomLayout())
	mainW.SetContent(mainC)
	mainW.Resize(fyne.NewSize(500, 800))
	mainW.ShowAndRun()
}

func createTopLayout() *fyne.Container {
	firstClass := components.CreateClassSelector(false)

	mainTitle := canvas.NewText("D&D Spell and Resource Counter", color.NRGBA{R: 130, G: 0, B: 18, A: 255})
	mainDesc := canvas.NewText("This application will help you keep track of how many spells or resources you still have in your D&D Game.", color.Black)
	mainDescNext := canvas.NewText("For exemple, it will keep track or your left Ki point, or Bardic Inspiration.", color.Black)

	titleC := container.NewVBox(container.NewCenter(mainTitle))
	descC := container.NewVBox(container.NewVBox(container.NewCenter(mainDesc)),
		container.NewVBox(container.NewCenter(mainDescNext)))

	topLayout := container.NewVBox(titleC, descC, widget.NewSeparator(), firstClass)

	multiClassContainer := container.NewVBox()
	multiClassFlagWidget := widget.NewCheck("MultiClass", func(b bool) {
		if b {
			playermanager.CurrentPlayer.MultiClass = true
			multiClassContainer.Add(components.CreateClassSelector(true))
		} else {
			playermanager.CurrentPlayer.MultiClass = false
			multiClassContainer.Objects = multiClassContainer.Objects[:0]
			multiClassContainer.Refresh()
		}
	})

	class, _ := playermanager.PlayerBinding.GetItem("Class")
	subClass, _ := playermanager.PlayerBinding.GetItem("SubClass")
	secClass, _ := playermanager.PlayerBinding.GetItem("SecondClass")
	secSubClass, _ := playermanager.PlayerBinding.GetItem("SecondSubClass")
	lvl, _ := playermanager.PlayerBinding.GetItem("Lvl")
	seclvl, _ := playermanager.PlayerBinding.GetItem("SecondLvl")

	caracContainer := container.NewHBox()

	onClassChanged := func() {
		class := playermanager.CurrentPlayer.Class.Name
		secClass := playermanager.CurrentPlayer.SecondClass.Name
		subClass := playermanager.CurrentPlayer.SubClass.Name
		secSubClass := playermanager.CurrentPlayer.SecondSubClass.Name
		if class == "Bard" || secClass == "Bard" {
			if len(caracContainer.Objects) == 0 {
				caracContainer.Add(widget.NewLabel("Charism : "))

				w := components.NewNumericalEntry()
				w.OnChanged = func(s string) {
					value, _ := strconv.Atoi(s)
					if playermanager.StatModificator[value] < 1 {
						value = 1
					} else {
						value = playermanager.StatModificator[value]
					}
					if class == "Bard" {

						playermanager.PlayerBinding.SetValue("FeatsValue", value)
					} else {
						playermanager.PlayerBinding.SetValue("SecondFeatsValue", value)
					}
				}
				caracContainer.Add(w)
				caracContainer.Refresh()
			}
		}

		if class != "Bard" && secClass != "Bard" {
			caracContainer.Objects = caracContainer.Objects[:0]
			caracContainer.Refresh()
		}

		if (class == "Fighter" && subClass == "Psi Warrior") ||
			(secClass == "Fighter" && secSubClass == "Psi Warrior") {
			if class == "Fighter" {
				value := playermanager.MasteryByLevel[playermanager.CurrentPlayer.Lvl]
				playermanager.PlayerBinding.SetValue("FeatsValue", value)
			} else {
				value := playermanager.MasteryByLevel[playermanager.CurrentPlayer.SecondLvl]
				playermanager.PlayerBinding.SetValue("SecondFeatsValue", value)
			}
		}

		if class == "Artificer" || class == "Barbarian" || class == "Monk" || class == "Sorcerer" {
			value := playermanager.CurrentPlayer.Class.FeatsAmountByLevel[playermanager.CurrentPlayer.Lvl]
			playermanager.PlayerBinding.SetValue("FeatsValue", value)
		}

		if secClass == "Artificer" || secClass == "Barbarian" || secClass == "Monk" || secClass == "Sorcerer" {
			value := playermanager.CurrentPlayer.SecondClass.FeatsAmountByLevel[playermanager.CurrentPlayer.SecondLvl]
			playermanager.PlayerBinding.SetValue("SecondFeatsValue", value)
		}

	}

	onLvlChanged := func() {
		class := playermanager.CurrentPlayer.Class.Name
		secClass := playermanager.CurrentPlayer.SecondClass.Name
		if class == "Artificer" || class == "Barbarian" || class == "Monk" || class == "Sorcerer" {
			value := playermanager.CurrentPlayer.Class.FeatsAmountByLevel[playermanager.CurrentPlayer.Lvl]
			playermanager.PlayerBinding.SetValue("FeatsValue", value)
		}

		if secClass == "Artificer" || secClass == "Barbarian" || secClass == "Monk" || secClass == "Sorcerer" {
			value := playermanager.CurrentPlayer.SecondClass.FeatsAmountByLevel[playermanager.CurrentPlayer.SecondLvl]
			playermanager.PlayerBinding.SetValue("SecondFeatsValue", value)
		}
	}

	onSubChanged := func() {}

	class.AddListener(binding.NewDataListener(onClassChanged))
	secClass.AddListener(binding.NewDataListener(onClassChanged))

	subClass.AddListener(binding.NewDataListener(onSubChanged))
	secSubClass.AddListener(binding.NewDataListener(onSubChanged))

	lvl.AddListener(binding.NewDataListener(onLvlChanged))
	seclvl.AddListener(binding.NewDataListener(onLvlChanged))

	/*shortRest := widget.NewButton("Short Rest", func() {

	})

	longRest := widget.NewButton("Long Rest", func() {})*/
	//rest := container.NewGridWithColumns(2, shortRest, longRest)
	topLayout.Add(multiClassFlagWidget)
	topLayout.Add(multiClassContainer)
	topLayout.Add(caracContainer)
	//topLayout.Add(rest)

	return topLayout
}

func createBottomLayout() *fyne.Container {
	lvl, _ := playermanager.PlayerBinding.GetItem("Lvl")
	class, _ := playermanager.PlayerBinding.GetItem("Class")
	subClass, _ := playermanager.PlayerBinding.GetItem("SubClass")
	secondSubClass, _ := playermanager.PlayerBinding.GetItem("SecondSubClass")
	secondClass, _ := playermanager.PlayerBinding.GetItem("SecondClass")
	secondLvl, _ := playermanager.PlayerBinding.GetItem("SecondLvl")
	feats, _ := playermanager.PlayerBinding.GetItem("FeatsValue")
	secondFeats, _ := playermanager.PlayerBinding.GetItem("SecondFeatsValue")

	bottomLayout := container.NewGridWithColumns(2)

	drawFirstPanel := func() {
		var temp fyne.CanvasObject
		if len(bottomLayout.Objects) > 0 {
			if len(bottomLayout.Objects) > 1 {
				temp = bottomLayout.Objects[1]
			}

			bottomLayout.Objects = bottomLayout.Objects[:0]
		}
		bottomLayout.Refresh()
		c := createSkillPanel(
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

	drawSecondPanel := func() {
		if len(bottomLayout.Objects) > 0 {
			bottomLayout.Objects = bottomLayout.Objects[:1]
		}
		bottomLayout.Refresh()
		c := createSkillPanel(
			playermanager.CurrentPlayer.SecondClass,
			playermanager.CurrentPlayer.SecondLvl,
			playermanager.CurrentPlayer.SecondSubClass,
			playermanager.CurrentPlayer.SecondFeatsValue,
		)
		bottomLayout.Add(c)
		bottomLayout.Refresh()
	}

	lvl.AddListener(binding.NewDataListener(drawFirstPanel))
	class.AddListener(binding.NewDataListener(drawFirstPanel))
	subClass.AddListener(binding.NewDataListener(drawFirstPanel))
	feats.AddListener(binding.NewDataListener(drawFirstPanel))

	secondClass.AddListener(binding.NewDataListener(drawSecondPanel))
	secondSubClass.AddListener(binding.NewDataListener(drawSecondPanel))
	secondLvl.AddListener(binding.NewDataListener(drawSecondPanel))
	secondFeats.AddListener(binding.NewDataListener(drawSecondPanel))

	return bottomLayout
}

func createSkillPanel(class classmanager.Class, lvl int, subClass classmanager.SubClass, featValue int) *fyne.Container {
	mainClass := container.NewGridWithRows(11)
	titleBinding := binding.NewString()
	titleString := class.Name + " - " + subClass.Name + " - lvl " + strconv.Itoa(lvl)
	title := widget.NewLabelWithData(titleBinding)
	titleBinding.Set(titleString)
	featBinding := binding.NewString()
	featName := widget.NewLabelWithData(featBinding)
	featCheckLine := container.NewHBox()
	for i := 0; i < featValue; i++ {
		featCheckLine.Add(widget.NewCheck("", func(b bool) {}))
	}
	featBinding.Set(class.Feats.Name)
	mainClass.Objects = mainClass.Objects[:0]
	mainClass.Refresh()
	mainClass.Add(title)

	featRow := container.NewGridWithRows(2, featName, featCheckLine)
	mainClass.Add(featRow)
	if playermanager.CurrentPlayer.Class.HasSpell {
		for i := 0; i < playermanager.CurrentPlayer.Class.SpellLevelLimit; i++ {

			spellLine := container.NewGridWithRows(2)
			spellLine.Add(widget.NewLabel("Spells Level " + strconv.Itoa(i+1)))
			check := playermanager.CurrentPlayer.Class.SpellByLevel[lvl][i+1]
			checkRow := container.NewHBox()
			for j := 0; j < check; j++ {
				checkRow.Add(widget.NewCheck("", func(b bool) {}))
			}
			spellLine.Add(checkRow)
			mainClass.Add(spellLine)
		}
	}

	mainClass.Refresh()
	return mainClass
}
