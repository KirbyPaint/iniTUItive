package main

import (
	"sort"
	"strconv"

	"github.com/google/uuid"
	"github.com/rivo/tview"
)

// Team struct
type Team struct {
	Id   int
	Text string
}

// Character struct
type Character struct {
	ID       uuid.UUID
	Name     string
	Init     int
	HP       int
	Priority int
	Team     Team
}

var characters []Character

// Add a new character to the list
func addNewCharacter(class Character) {
	characters = append(characters, class)
}

// Sort the characters by initiative
func getCharactersSorted() []Character {
	sort.Slice(characters, func(i, j int) bool {
		// If there's a tie, return enemies before anyone else
		if characters[i].Init == characters[j].Init {
			if characters[i].Team.Id == 2 && characters[j].Team.Id != 2 {
				return true
			}
			if characters[i].Team.Id != 2 && characters[j].Team.Id == 2 {
				return false
			}
		}
		// Then if there's a tie between anyone else, sort by Priority
		if characters[i].Init == characters[j].Init {
			return characters[i].Priority > characters[j].Priority
		}
		return characters[i].Init > characters[j].Init
	})
	return characters
}

func getCharacterByID(id uuid.UUID) Character {
	for _, character := range characters {
		if character.ID == id {
			return character
		}
	}
	return Character{}
}

func removeCharacterByID(id uuid.UUID) {
	for i, character := range characters {
		if character.ID == id {
			characters = append(characters[:i], characters[i+1:]...)
			break
		}
	}
}

func generateID() uuid.UUID {
	id := uuid.New()
	return id
}

func main() {
	// The base application
	app := tview.NewApplication()

	// The main components
	addNewForm := tview.NewForm()
	displayList := tview.NewList()
	headerBox := tview.NewBox()
	commandList := tview.NewList()

	// Always focus Add New command for speed
	focusCommandsList := func() {
		commandList.SetCurrentItem(0)
		app.SetFocus(commandList)
	}

	// Add the return command to the display list first always
	displayList.AddItem("Return", "", 'r', func() {
		focusCommandsList()
	}).SetWrapAround(true)

	// Function to clear the form fields of data
	refreshAddNewForm := func() {
		addNewForm.GetFormItemByLabel("Name").(*tview.InputField).SetText("")
		addNewForm.GetFormItemByLabel("Init").(*tview.InputField).SetText("")
		addNewForm.GetFormItemByLabel("HP").(*tview.InputField).SetText("")
		addNewForm.GetFormItemByLabel("Team").(*tview.DropDown).SetCurrentOption(0)
		addNewForm.GetFormItemByLabel("Prio").(*tview.InputField).SetText("")
	}

	// Function to re-sort the display list
	refreshDisplayList := func() {
		displayList.Clear()
		displayList.AddItem("Return", "", 'r', func() {
			focusCommandsList()
		})
		for _, character := range getCharactersSorted() {
			var characterNameColored string
			switch character.Team.Id {
			case 0:
				characterNameColored = "[blue]" + character.Name + "[-]"
			case 1:
				characterNameColored = "[green]" + character.Name + "[-]"
			case 2:
				characterNameColored = "[red]" + character.Name + "[-]"
			case 3:
				characterNameColored = "[yellow]" + character.Name + "[-]"
			}
			hpValue := func() string {
				if character.HP <= 0 {
					return ""
				} else {
					return " (" + strconv.Itoa(character.HP) + ")"
				}
			}
			lineItem := strconv.Itoa(character.Init) + "   " + characterNameColored + hpValue()
			displayList.AddItem(lineItem, "", 0, nil)
		}
	}

	// Add styling
	addNewForm.SetHorizontal(true).SetBorder(true).SetTitle(" Add New Line ")
	displayList.SetBorder(true).SetTitle(" Initiative ")
	headerBox.SetBorder(true).SetTitle(" Initiative Tracker ")
	commandList.SetBorder(true).SetTitle(" Commands ")

	// Create dropdown and add options
	teamDropDown := tview.NewDropDown()
	teamDropDown.SetLabel("Team").SetOptions([]string{"PLYR", "ALLY", "ENMY", "UNKN"}, nil).SetCurrentOption(0)

	// Input form for adding a new character
	addNewForm.AddInputField("Name", "", 10, nil, nil).
		AddInputField("Init", "", 3, nil, nil).
		AddInputField("HP", "", 4, nil, nil).
		AddFormItem(teamDropDown).
		AddInputField("Prio", "", 2, nil, nil).
		AddButton("S", func() {
			teamId, teamText := addNewForm.GetFormItemByLabel("Team").(*tview.DropDown).GetCurrentOption()
			character := Character{
				ID:   generateID(),
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
				Priority: func() int {
					i, _ := strconv.Atoi(addNewForm.GetFormItemByLabel("Prio").(*tview.InputField).GetText())
					return i
				}(),
			}
			addNewCharacter(character)
			refreshAddNewForm()
			refreshDisplayList()
			focusCommandsList()
		}).
		AddButton("C", func() {
			refreshAddNewForm()
		})

	// Set the numeric input fields to accept only numbers
	addNewForm.GetFormItemByLabel("Init").(*tview.InputField).SetAcceptanceFunc(tview.InputFieldInteger)
	addNewForm.GetFormItemByLabel("HP").(*tview.InputField).SetAcceptanceFunc(tview.InputFieldInteger)
	addNewForm.GetFormItemByLabel("Prio").(*tview.InputField).SetAcceptanceFunc(tview.InputFieldInteger)

	// Command list buttons
	commandList.AddItem("Add New", "", 'n', func() {
		addNewForm.SetFocus(0)
		app.SetFocus(addNewForm)
	}).AddItem("List", "", 'l', func() {
		app.SetFocus(displayList)
	}).AddItem("Exit", "", 'q', func() {
		app.Stop()
	}).SetSelectedFocusOnly(true)

	// Edit a selected character with Enter
	displayList.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if index == 0 {
			focusCommandsList()
		} else {
			character := getCharacterByID(characters[index-1].ID)
			removeCharacterByID(characters[index-1].ID)
			addNewForm.GetFormItemByLabel("Name").(*tview.InputField).SetText(character.Name)
			addNewForm.GetFormItemByLabel("Init").(*tview.InputField).SetText(strconv.Itoa(character.Init))
			addNewForm.GetFormItemByLabel("HP").(*tview.InputField).SetText(strconv.Itoa(character.HP))
			addNewForm.GetFormItemByLabel("Team").(*tview.DropDown).SetCurrentOption(character.Team.Id)
			addNewForm.GetFormItemByLabel("Prio").(*tview.InputField).SetText(strconv.Itoa(character.Priority))
			addNewForm.SetFocus(2) // focuses HP since that is most likely to be edited
			app.SetFocus(addNewForm)
		}
	}).SetSelectedFocusOnly(true)

	// If d is pressed on a list item remove it
	// displayList.RemoveItem(displayList.GetCurrentItem()).SetSelectedFocusOnly(true)

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
