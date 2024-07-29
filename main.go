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
func AddNewCharacter(class Character) {
	characters = append(characters, class)
}

// Sort the characters by initiative
func GetCharactersSorted() []Character {
	sort.Slice(characters, func(i, j int) bool {
		// If there's a tie, return enemies, then allies, then players
		if characters[i].Init == characters[j].Init {
			if characters[i].Team.Id == 2 {
				return true
			} else if characters[j].Team.Id == 2 {
				return false
			} else if characters[i].Team.Id == 1 {
				return true
			} else if characters[j].Team.Id == 1 {
				return false
			} else {
				return true
			}
		}
		return characters[i].Init > characters[j].Init
	})
	return characters
}

func GetCharacterByID(id uuid.UUID) Character {
	for _, character := range characters {
		if character.ID == id {
			return character
		}
	}
	return Character{}
}

func RemoveCharacterByID(id uuid.UUID) {
	for i, character := range characters {
		if character.ID == id {
			characters = append(characters[:i], characters[i+1:]...)
			break
		}
	}
}

func GenerateID() uuid.UUID {
	id := uuid.New()
	return id
}

func main() {
	// The base application
	app := tview.NewApplication()

	// The components
	addNewForm := tview.NewForm()
	displayList := tview.NewList()
	headerBox := tview.NewBox()
	commandList := tview.NewList()

	displayList.AddItem("Return", "", 'r', func() {
		app.SetFocus(commandList)
	}).SetWrapAround(true)

	// Clear the form fields of data
	RefreshAddNewForm := func() {
		addNewForm.GetFormItemByLabel("Name").(*tview.InputField).SetText("")
		addNewForm.GetFormItemByLabel("Init").(*tview.InputField).SetText("")
		addNewForm.GetFormItemByLabel("HP").(*tview.InputField).SetText("")
		addNewForm.GetFormItemByLabel("Team").(*tview.DropDown).SetCurrentOption(0)
	}

	// Re-sort the display list
	RefreshDisplayList := func() {
		displayList.Clear()
		displayList.AddItem("Return", "", 'r', func() {
			app.SetFocus(commandList)
		})
		for _, character := range GetCharactersSorted() {
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
			lineItem := strconv.Itoa(character.Init) + ": " + characterNameColored + " (" + strconv.Itoa(character.HP) + ")"
			displayList.AddItem(lineItem, "", 0, nil)
		}
	}

	// Add styling
	addNewForm.SetHorizontal(true).SetBorder(true).SetTitle(" Add New Line ")
	displayList.SetBorder(true).SetTitle(" Initiative ")
	headerBox.SetBorder(true).SetTitle(" Initiative Tracker ")
	commandList.SetBorder(true).SetTitle(" Commands ")

	// Input form for adding a new character
	addNewForm.AddInputField("Name", "", 20, nil, nil).
		AddInputField("Init", "", 3, nil, nil).
		AddInputField("HP", "", 4, nil, nil).
		AddDropDown("Team", []string{"Player", "Ally", "Enemy", "Unknown"}, 0, nil).
		AddButton("Save", func() {
			teamId, teamText := addNewForm.GetFormItemByLabel("Team").(*tview.DropDown).GetCurrentOption()
			character := Character{
				ID:   GenerateID(),
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
			RefreshAddNewForm()
			RefreshDisplayList()
			app.SetFocus(commandList)
		}).
		AddButton("Clear", func() {
			RefreshAddNewForm()
		})

	// Set the numeric input fields to accept only numbers
	addNewForm.GetFormItemByLabel("Init").(*tview.InputField).SetAcceptanceFunc(tview.InputFieldInteger)
	addNewForm.GetFormItemByLabel("HP").(*tview.InputField).SetAcceptanceFunc(tview.InputFieldInteger)

	// Command list options
	commandList.AddItem("Add New", "", 'n', func() {
		addNewForm.SetFocus(0)
		app.SetFocus(addNewForm)
	})
	commandList.AddItem("List", "", 'l', func() {
		app.SetFocus(displayList)
	})
	commandList.AddItem("Exit", "", 'q', func() {
		app.Stop()
	})

	// Edit a selected character
	displayList.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if index == 0 {
			app.SetFocus(commandList)
		} else {
			character := GetCharacterByID(characters[index-1].ID)
			RemoveCharacterByID(characters[index-1].ID)
			addNewForm.GetFormItemByLabel("Name").(*tview.InputField).SetText(character.Name)
			addNewForm.GetFormItemByLabel("Init").(*tview.InputField).SetText(strconv.Itoa(character.Init))
			addNewForm.GetFormItemByLabel("HP").(*tview.InputField).SetText(strconv.Itoa(character.HP))
			addNewForm.GetFormItemByLabel("Team").(*tview.DropDown).SetCurrentOption(character.Team.Id)
			addNewForm.SetFocus(2)
			app.SetFocus(addNewForm)
		}
	})

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
