package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/mjehanno/dnd-spell-counter/components"
	playermanager "github.com/mjehanno/dnd-spell-counter/player-manager"
)

func main() {
	spellCounterApp := app.NewWithID("com.mjehanno.dnd-spellcounter")
	playermanager.CurrentPlayer = new(playermanager.Player)

	playermanager.PlayerBinding = binding.BindStruct(playermanager.CurrentPlayer)
	spellCounterApp.SetIcon(resourceIconPng)
	mainW := spellCounterApp.NewWindow("D&D Spell-counter")
	mainC := container.NewGridWithColumns(2, components.CreateTopLayout(), components.CreateBottomLayout())
	mainW.SetContent(mainC)
	mainW.Resize(fyne.NewSize(500, 800))
	mainW.ShowAndRun()
}
