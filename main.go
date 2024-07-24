package main

import (
	"github.com/rivo/tview"
)

func AddNewCharacter(name, init, hp, team string) {
	// Add a new character to the list
}

func main() {
	// The base application
	app := tview.NewApplication()

	headerBox := tview.NewBox().SetBorder(true).SetTitle(" Initiative Tracker ")
	// listBox := tview.NewBox().SetBorder(true).SetTitle(" List ")
	currentBox := tview.NewBox().SetBorder(true).SetTitle(" Current ")

	addNewForm := tview.NewForm()
	addNewForm.AddInputField("Name", "", 20, nil, nil).
		AddInputField("Init", "", 3, nil, nil).
		AddInputField("HP", "", 4, nil, nil).
		AddDropDown("Team", []string{"Player", "Ally", "Enemy", "Unknown"}, 0, nil).
		AddButton("Save", func() {
			// AddNewCharacter(addNewForm.GetFormItemByLabel("Name").(*tview.InputField).GetText(),
			// 	addNewForm.GetFormItemByLabel("Init").(*tview.InputField).GetText(),
			// 	addNewForm.GetFormItemByLabel("HP").(*tview.InputField).GetText(),
			// 	addNewForm.GetFormItemByLabel("Team").(*tview.DropDown).GetCurrentOption())
			app.SetFocus(currentBox)
		}).
		AddButton("Clear", func() {
			addNewForm.Clear(false)
		})

	// addNewForm := tview.NewForm().
	// 	AddInputField("Name", "", 20, nil, nil).
	// 	AddInputField("Init", "", 3, nil, nil).
	// 	AddInputField("HP", "", 4, nil, nil).
	// 	AddDropDown("Team", []string{"Player", "Ally", "Enemy", "Unknown"}, 0, nil).
	// 	AddButton("Save", func() {
	// 		app.SetFocus(currentBox)
	// 	}).
	// 	AddButton("Clear", func() {
	// 		tview.NewForm().Clear(false)
	// 	})

	list := tview.NewList().
		AddItem("Add New", "", 'n', func() {
			app.SetFocus(addNewForm)
		})

	addNewForm.SetHorizontal(true).SetBorder(true).SetTitle(" Add New ")
	// .SetTitleAlign(tview.AlignLeft) // I like it but come back to this later
	list.SetBorder(true).SetTitle(" List ")

	// Set the input field to accept only numbers
	addNewForm.GetFormItemByLabel("Init").(*tview.InputField).SetAcceptanceFunc(tview.InputFieldInteger)

	// The layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(headerBox, 2, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(list, 0, 1, false).
			AddItem(currentBox, 15, 1, false), 0, 2, false).
		AddItem(addNewForm, 5, 1, false)

	// Run the application
	if err := app.SetRoot(flex, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
