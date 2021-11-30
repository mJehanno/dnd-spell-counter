package components

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	classmanager "github.com/mjehanno/dnd-spell-counter/class-manager"
	playermanager "github.com/mjehanno/dnd-spell-counter/player-manager"
)

func CreateTopLayout() *fyne.Container {
	topLayout := CreateTitleTopLayout()

	multiClassContainer, multiClassFlagWidget := CreateMultiClassFlagLayout()

	class, _ := playermanager.PlayerBinding.GetItem(playermanager.Class)
	subClass, _ := playermanager.PlayerBinding.GetItem(playermanager.SubClass)
	secClass, _ := playermanager.PlayerBinding.GetItem(playermanager.SecondClass)
	secSubClass, _ := playermanager.PlayerBinding.GetItem(playermanager.SecondSubClass)
	lvl, _ := playermanager.PlayerBinding.GetItem(playermanager.Lvl)
	seclvl, _ := playermanager.PlayerBinding.GetItem(playermanager.SecondLvl)

	class.AddListener(binding.NewDataListener(OnClassChanged))
	secClass.AddListener(binding.NewDataListener(OnClassChanged))

	subClass.AddListener(binding.NewDataListener(OnSubChanged))
	secSubClass.AddListener(binding.NewDataListener(OnSubChanged))

	lvl.AddListener(binding.NewDataListener(OnLvlChanged))
	seclvl.AddListener(binding.NewDataListener(OnLvlChanged))

	topLayout.Add(multiClassFlagWidget)
	topLayout.Add(multiClassContainer)
	topLayout.Add(caracContainer)

	return topLayout
}

func CreateTitleTopLayout() *fyne.Container {
	firstClass := CreateClassSelector(false)

	mainTitle := canvas.NewText("D&D Spell and Resource Counter", color.NRGBA{R: 130, G: 0, B: 18, A: 255})
	mainDesc := canvas.NewText("This application will help you keep track of how many spells or resources you still have in your D&D Game.", color.Black)
	mainDescNext := canvas.NewText("For example, it will keep track or your left Ki point, or Bardic Inspiration.", color.Black)
	titleC := container.NewVBox(container.NewCenter(mainTitle))
	descC := container.NewVBox(container.NewVBox(container.NewCenter(mainDesc)),
		container.NewVBox(container.NewCenter(mainDescNext)))

	topLayout := container.NewVBox(titleC, descC, widget.NewSeparator(), firstClass)
	return topLayout
}

func CreateMultiClassFlagLayout() (*fyne.Container, *widget.Check) {
	multiClassContainer := container.NewVBox()
	multiClassFlagWidget := widget.NewCheck("MultiClass", func(b bool) {
		if b {
			playermanager.CurrentPlayer.MultiClass = true
			multiClassContainer.Add(CreateClassSelector(true))
		} else {
			playermanager.CurrentPlayer.MultiClass = false
			multiClassContainer.Objects = multiClassContainer.Objects[:0]
			multiClassContainer.Refresh()
		}
	})
	return multiClassContainer, multiClassFlagWidget
}

func CreateBottomLayout() *fyne.Container {
	lvl, _ := playermanager.PlayerBinding.GetItem(playermanager.Lvl)
	class, _ := playermanager.PlayerBinding.GetItem(playermanager.Class)
	subClass, _ := playermanager.PlayerBinding.GetItem(playermanager.SubClass)
	secondSubClass, _ := playermanager.PlayerBinding.GetItem(playermanager.SecondSubClass)
	secondClass, _ := playermanager.PlayerBinding.GetItem(playermanager.SecondClass)
	secondLvl, _ := playermanager.PlayerBinding.GetItem(playermanager.SecondLvl)
	feats, _ := playermanager.PlayerBinding.GetItem(playermanager.FeatsValue)
	secondFeats, _ := playermanager.PlayerBinding.GetItem(playermanager.SecondFeatsValue)

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

func CreateSkillPanel(class classmanager.Class, lvl int, subClass classmanager.SubClass, featValue int) *fyne.Container {
	mainClass := container.NewGridWithRows(11)
	titleBinding := binding.NewString()
	titleString := class.Name + " - " + subClass.Name + " - lvl " + strconv.Itoa(lvl)
	title := widget.NewLabelWithData(titleBinding)
	titleBinding.Set(titleString)
	featBinding := binding.NewString()
	featName := widget.NewLabelWithData(featBinding)
	featCheckLine := container.NewHBox()
	for i := 0; i < featValue; i++ {
		featCheckLine.Add(widget.NewCheck("", nil))
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
				checkRow.Add(widget.NewCheck("", nil))
			}
			spellLine.Add(checkRow)
			mainClass.Add(spellLine)
		}
	}

	mainClass.Refresh()
	return mainClass
}
