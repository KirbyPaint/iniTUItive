package main

import (
	"github.com/rivo/tview"
)

// Team struct
type Team struct {
	Id   int
	Text string
}

// Character struct
type Character struct {
	Name string
	Init string
	HP   string
	Team Team
}

var characters []Character

func AddNewCharacter(class Character) {
	// Add a new character to the list
	characters = append(characters, class)
}

func main() {
	// The base application
	app := tview.NewApplication()

	headerBox := tview.NewBox()
	displayList := tview.NewList()
	list := tview.NewList()
	addNewForm := tview.NewForm()

	headerBox.SetBorder(true).SetTitle(" Initiative Tracker ")
	displayList.SetBorder(true).SetTitle(" Current ")

	addNewForm.AddInputField("Name", "", 20, nil, nil).
		AddInputField("Init", "", 3, nil, nil).
		AddInputField("HP", "", 4, nil, nil).
		AddDropDown("Team", []string{"Player", "Ally", "Enemy", "Unknown"}, 0, nil).
		AddButton("Save", func() {
			teamId, teamText := addNewForm.GetFormItemByLabel("Team").(*tview.DropDown).GetCurrentOption()
			character := Character{
				Name: addNewForm.GetFormItemByLabel("Name").(*tview.InputField).GetText(),
				Init: addNewForm.GetFormItemByLabel("Init").(*tview.InputField).GetText(),
				HP:   addNewForm.GetFormItemByLabel("HP").(*tview.InputField).GetText(),
				Team: Team{
					Id:   teamId,
					Text: teamText,
				},
			}
			AddNewCharacter(character)
			// displayList.SetText("Current: " + character.Name + " (" + character.Init + ")")
			displayList.AddItem(character.Name+" ("+character.Init+")", character.HP, 0, nil)
			app.SetFocus(list)
		}).
		AddButton("Clear", func() {
			addNewForm.Clear(false)
		})

	list.AddItem("Add New", "", 'n', func() {
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
			AddItem(displayList, 0, 2, false), 0, 2, false).
		AddItem(addNewForm, 5, 1, false)

	// Run the application
	if err := app.SetRoot(flex, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
