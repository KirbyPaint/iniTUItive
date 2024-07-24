package main

import (
	"github.com/rivo/tview"
)

func main() {
	// The base application
	app := tview.NewApplication()

	headerBox := tview.NewBox().SetBorder(true).SetTitle(" Initiative Tracker ")
	listBox := tview.NewBox().SetBorder(true).SetTitle(" List ")
	currentBox := tview.NewBox().SetBorder(true).SetTitle(" Current ")

	addNewForm := tview.NewForm().
		AddInputField("Name", "", 0, nil, nil).
		AddInputField("Init", "", 0, nil, nil).
		AddInputField("HP", "", 0, nil, nil).
		AddDropDown("Title", []string{"Player", "Ally", "Enemy", "Unknown"}, 0, nil).
		AddButton("Save", func() {
		}).
		AddButton("Clear", func() {
		})

	addNewForm.SetHorizontal(true).SetBorder(true).SetTitle(" Add New ")
	// .SetTitleAlign(tview.AlignLeft) // I like it but come back to this later

	// Set the input field to accept only numbers
	addNewForm.GetFormItemByLabel("Init").(*tview.InputField).SetAcceptanceFunc(tview.InputFieldInteger)

	// The layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(headerBox, 2, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(listBox, 0, 1, false).
			AddItem(currentBox, 15, 1, false), 0, 2, false).
		AddItem(addNewForm, 8, 1, false)

	// Run the application
	if err := app.SetRoot(flex, true).SetFocus(addNewForm).Run(); err != nil {
		panic(err)
	}
}
