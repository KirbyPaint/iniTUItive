package main

import (
	"sort"
	"strconv"

	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type Team struct {
	Id   int
	Text string
}

type Character struct {
	ID       uuid.UUID
	Name     string
	Init     int
	HP       int
	Priority int
	Team     Team
}

var characters []Character

func addNewCharacter(class Character) {
	characters = append(characters, class)
}

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
	app := tview.NewApplication()

	characterInputForm := tview.NewForm()
	initiativeList := tview.NewList()
	headerBox := tview.NewBox()
	commandsList := tview.NewList()

	focusCommandsList := func() {
		commandsList.SetCurrentItem(0)
		app.SetFocus(commandsList)
	}

	refreshCharacterInputForm := func() {
		deleteButtonIndex := characterInputForm.GetButtonIndex("D")
		if deleteButtonIndex != -1 {
			characterInputForm.RemoveButton(deleteButtonIndex)
		}
		characterInputForm.GetFormItemByLabel("Name").(*tview.InputField).SetText("")
		characterInputForm.GetFormItemByLabel("Init").(*tview.InputField).SetText("")
		characterInputForm.GetFormItemByLabel("HP").(*tview.InputField).SetText("")
		characterInputForm.GetFormItemByLabel("Team").(*tview.DropDown).SetCurrentOption(0)
		characterInputForm.GetFormItemByLabel("Prio").(*tview.InputField).SetText("")
	}

	renderInitiativeList := func() {
		initiativeList.Clear()
		initiativeList.AddItem("Return", "", 'r', func() {
			focusCommandsList()
		}).SetWrapAround(true)
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
			initiativeList.AddItem(lineItem, "", 0, nil)
		}
	}

	refreshAndFocusInitiativeList := func() {
		refreshCharacterInputForm()
		renderInitiativeList()
		focusCommandsList()
	}

	characterInputForm.SetHorizontal(true).SetBorder(true).SetTitle(" Add New Line ")
	initiativeList.SetBorder(true).SetTitle(" Initiative ")
	headerBox.SetBorder(true).SetTitle(" Initiative Tracker ")
	commandsList.SetBorder(true).SetTitle(" Commands ")

	teamDropDown := tview.NewDropDown()
	teamDropDown.SetLabel("Team").SetOptions([]string{"PLYR", "ALLY", "ENMY", "UNKN"}, nil).SetCurrentOption(0)

	characterInputForm.AddInputField("Name", "", 10, nil, nil).
		AddInputField("Init", "", 3, nil, nil).
		AddInputField("HP", "", 4, nil, nil).
		AddFormItem(teamDropDown).
		AddInputField("Prio", "", 2, nil, nil).
		AddButton("S", func() {
			teamId, teamText := characterInputForm.GetFormItemByLabel("Team").(*tview.DropDown).GetCurrentOption()
			character := Character{
				ID:   generateID(),
				Name: characterInputForm.GetFormItemByLabel("Name").(*tview.InputField).GetText(),
				Init: func() int {
					i, _ := strconv.Atoi(characterInputForm.GetFormItemByLabel("Init").(*tview.InputField).GetText())
					return i
				}(),
				HP: func() int {
					i, _ := strconv.Atoi(characterInputForm.GetFormItemByLabel("HP").(*tview.InputField).GetText())
					return i
				}(),
				Team: Team{
					Id:   teamId,
					Text: teamText,
				},
				Priority: func() int {
					i, _ := strconv.Atoi(characterInputForm.GetFormItemByLabel("Prio").(*tview.InputField).GetText())
					return i
				}(),
			}
			addNewCharacter(character)
			refreshAndFocusInitiativeList()
		}).
		AddButton("C", func() {
			refreshCharacterInputForm()
			characterInputForm.SetFocus(0)
			app.SetFocus(characterInputForm)
		}).SetCancelFunc(func() {
		refreshCharacterInputForm()
		focusCommandsList()
	})

	characterInputForm.GetFormItemByLabel("Init").(*tview.InputField).SetAcceptanceFunc(tview.InputFieldInteger)
	characterInputForm.GetFormItemByLabel("HP").(*tview.InputField).SetAcceptanceFunc(tview.InputFieldInteger)
	characterInputForm.GetFormItemByLabel("Prio").(*tview.InputField).SetAcceptanceFunc(tview.InputFieldInteger)

	commandsList.AddItem("Add New", "", 'n', func() {
		characterInputForm.SetFocus(0)
		app.SetFocus(characterInputForm)
	}).AddItem("List", "", 'l', func() {
		app.SetFocus(initiativeList)
	}).AddItem("Exit", "", 'q', func() {
		app.Stop()
	}).SetSelectedFocusOnly(true)

	initiativeList.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if index == 0 {
			focusCommandsList()
		} else {
			character := getCharacterByID(characters[index-1].ID)
			removeCharacterByID(characters[index-1].ID)
			characterInputForm.GetFormItemByLabel("Name").(*tview.InputField).SetText(character.Name)
			characterInputForm.GetFormItemByLabel("Init").(*tview.InputField).SetText(strconv.Itoa(character.Init))
			characterInputForm.GetFormItemByLabel("HP").(*tview.InputField).SetText(strconv.Itoa(character.HP))
			characterInputForm.GetFormItemByLabel("Team").(*tview.DropDown).SetCurrentOption(character.Team.Id)
			characterInputForm.GetFormItemByLabel("Prio").(*tview.InputField).SetText(strconv.Itoa(character.Priority))
			characterInputForm.AddButton("[white:red]D[-]", func() {
				removeCharacterByID(character.ID)
				refreshAndFocusInitiativeList()
			})
			characterInputForm.SetFocus(2) // focuses HP since that is most likely to be edited
			app.SetFocus(characterInputForm)
		}
	}).SetSelectedFocusOnly(true)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(headerBox, 2, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(commandsList, 16, 1, false).
			AddItem(initiativeList, 0, 2, false), 0, 2, false).
		AddItem(characterInputForm, 5, 1, false)

	renderInitiativeList()

	if err := app.SetRoot(flex, true).SetFocus(commandsList).Run(); err != nil {
		panic(err)
	}
}
