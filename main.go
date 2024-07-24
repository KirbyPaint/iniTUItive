package main

import (
	"sort"
	"strconv"

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
	Init int
	HP   int
	Team Team
}

var characters []Character

// Add a new character to the list
func AddNewCharacter(class Character) {
	characters = append(characters, class)
}

func GetCharactersSorted() []Character {
	// Sort the characters by initiative
	sort.Slice(characters, func(i, j int) bool {
		return characters[i].Init < characters[j].Init
	})
	return characters
}

func main() {
	// The base application
	app := tview.NewApplication()

	// The components
	addNewForm := tview.NewForm()
	displayList := tview.NewList()
	headerBox := tview.NewBox()
	commandList := tview.NewList()

	// Add styling
	addNewForm.SetHorizontal(true).SetBorder(true).SetTitle(" Add New Line ")
	displayList.SetBorder(true).SetTitle(" Initiative ")
	headerBox.SetBorder(true).SetTitle(" Initiative Tracker ")
	commandList.SetBorder(true).SetTitle(" Commands ")

	addNewForm.AddInputField("Name", "", 20, nil, nil).
		AddInputField("Init", "", 3, nil, nil).
		AddInputField("HP", "", 4, nil, nil).
		AddDropDown("Team", []string{"Player", "Ally", "Enemy", "Unknown"}, 0, nil).
		AddButton("Save", func() {
			teamId, teamText := addNewForm.GetFormItemByLabel("Team").(*tview.DropDown).GetCurrentOption()
			character := Character{
				Name: addNewForm.GetFormItemByLabel("Name").(*tview.InputField).GetText(),
				Init: func() int {
					i, _ := strconv.Atoi(addNewForm.GetFormItemByLabel("Init").(*tview.InputField).GetText())
					return i
				}(),
				HP: func() int {
					i, _ := strconv.Atoi(addNewForm.GetFormItemByLabel("HP").(*tview.InputField).GetText())
					return i
				}(),
				Team: Team{
					Id:   teamId,
					Text: teamText,
				},
			}
			AddNewCharacter(character)
			app.SetFocus(commandList)
			displayList.Clear()
			for _, character := range GetCharactersSorted() {
				lineItem := strconv.Itoa(character.Init) + ": " + character.Name + " (" + strconv.Itoa(character.HP) + ")"
				displayList.AddItem(lineItem, "", 0, nil)
			}
		}).
		AddButton("Clear", func() {
			addNewForm.Clear(false)
		})

	// List of options
	commandList.AddItem("Add New", "", 'n', func() {
		app.SetFocus(addNewForm)
	})

	// Set the input field to accept only numbers
	addNewForm.GetFormItemByLabel("Init").(*tview.InputField).SetAcceptanceFunc(tview.InputFieldInteger)

	// Set the layout of components
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(headerBox, 2, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(commandList, 0, 1, false).
			AddItem(displayList, 0, 2, false), 0, 2, false).
		AddItem(addNewForm, 5, 1, false)

	// Run the application
	if err := app.SetRoot(flex, true).SetFocus(commandList).Run(); err != nil {
		panic(err)
	}
}
