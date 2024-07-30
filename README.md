# IniTUItive Tracker

![application image](https://github.com/KirbyPaint/iniTUItive/blob/main/img/example_many.png?raw=true)

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Contributing](#contributing)
5. [License](#license)
6. [Contact](#contact)
7. [Acknowledgements](#acknowledgements)

## Introduction

IniTUItive Tracker is a simple RPG turn-order tracker for tabletop role-playing games. It allows you to track players and monsters in combat, and easily manage their initiative order with only a terminal interface. IniTUItive Tracker is designed with Pathfinder 2nd Edition in mind.

## Installation

To install IniTUItive Tracker, simply clone the [GitHub repository](https://github.com/KirbyPaint/iniTUItive). You can run the application by executing `go run main.go` or by running `run.sh`. This means you must have `go` installed to use this application.

## Usage

IniTUItive Tracker is meant to be easy to use and understand. Each command has an assigned key as a shortcut, and the commands are listed in each panel. From the Commands panel, you can see all the available commands and their keys:

```shell
Commands:
 n: Add a new combatant
 l: Navigate to the Initiative List panel
 q: Quit
```

From the Initiative List panel, you can see all the combatants in the current initiative order, and you can edit or delete a combatant by selecting their entry and pressing Enter. You can also see the available commands:

```shell
Commands:
 r: Return to the Commands panel
```

When adding or editing a character, you can enter their name, initiative bonus, and any other relevant information. The Initiative and HP fields accept integers only. The Team field is optional and can be used to group combatants together. The four Teams are Player (blue), Enemy (red), Ally (green), and Neutral (yellow), and the selected Team will cause that character's name to render in the terminal with that color. In Pathfinder 2e, in the case of a tie, Enemies attack first, so the program will automatically sort the initiative order to reflect this. Should two characters have the same initiative, you may set their order manually by entering a number in the Priority field. The higher the number, the earlier the character will act in the case of a tie. The buttons following the input fields are Save and Clear, respectively. Save will save the character to the initiative order, and Clear will clear the input fields for re-entry.  
![image showing input fields and team labels](https://github.com/KirbyPaint/iniTUItive/blob/main/img/example_1.png?raw=true)
When selecting a character for editing, a new button will appear: Delete. This button will remove the selected character from the initiative order.  

![image showing example with priority input and visible delete button](https://github.com/KirbyPaint/iniTUItive/blob/main/img/example_with_delete.png?raw=true)  

There are a few other sample images in the [img](/img) directory.

## Contributing

If you would like to contribute to IniTUItive Tracker, please submit a pull request. I welcome any contributions, including bug fixes, feature enhancements, and translations. If you have any questions or feedback, please feel free to [contact us](#contact).

## License

IniTUItive Tracker is licensed under the [MIT License](LICENSE). You are free to use, modify, and distribute this software as you see fit. I only ask that you include the original license and attribution in your distribution.

## Contact

If you have any questions or feedback about IniTUItive Tracker, please feel free to submit an issue. You can also contact the author, KirbyPaint, on Discord (@kirbypaint)

## Acknowledgements

I would like to thank the following individuals and organizations for their contributions to IniTUItive Tracker:

- [KirbyPaint](https://github.com/KirbyPaint)
- [rivo (tview)](https://github.com/rivo/tview)
- [GoLang](https://golang.org/)
- [Pathfinder 2nd Edition](https://paizo.com/pathfinder)
