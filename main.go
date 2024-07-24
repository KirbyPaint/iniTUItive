package main

import (
	"github.com/rivo/tview"
)

func main() {
	// The base application
	app := tview.NewApplication()

	// Data entry
	form := tview.NewForm().
		AddInputField("Name", "", 0, nil, nil).
		AddInputField("Initiative", "", 0, nil, nil).
		AddButton("Save", func() {
			// Save the data
		}).
		AddButton("Cancel", func() {
			// Cancel the data entry
		})

	listBox := tview.NewBox().SetBorder(true).SetTitle(" List ")
	headerBox := tview.NewBox().SetBorder(true).SetTitle(" Initiative Tracker ")
	addNewBox := tview.NewBox().SetBorder(true).SetTitle(" Add New ")
	currentBox := tview.NewBox().SetBorder(true).SetTitle(" Current ")

	// The layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(headerBox, 2, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(listBox, 0, 1, false).
			AddItem(currentBox, 15, 1, false), 0, 2, false).
		AddItem(addNewBox, 8, 1, false)

	// Run the application
	if err := app.SetRoot(flex, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}
