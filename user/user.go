package user

import (
	format "github.com/ajean-zd/codeclub-challenge/format"
)

func menuWorkflow() {

	//refactor this: needs to be a starting point menu and then ifs/fors/cases for each menu
	//implement this later
	//the logic however goes start with menu 1
	userSelectOne, _ := menu(MenuOne, 4)
	//exit if 4 selected]
	//go get correct data to display based on selection
	//it will return the the right format
	data, _ := format.ReturnFormattedData(userSelectOne)

	//load some data to display
	displayData(data)
	//then loop around to display the next menu, go back a step, or exit if 4 selected
	//do this logic wooooooo challenge
}

func displayData(data string) {
	//from whatever the user has selected display the correct menu
	//string for now. strings are great holding patterns. wooot.

}

//accepts user input from terminal
func menu(userMenuType string, boundaries int) (string, error) {

	//needs to be a loopy loop that goes through menus and back
	//this is a challenge
	menu := Menu{Display: userMenuType,
		Boundaries: boundaries,
	}

	//Display prompt to terminal
	displayMenuTerminal(menu.Display)
	userSelection, _ := getUserInput() //however the fk this works
	userSelection, _ = validateInput(userSelection, menu.Boundaries)
	// handle error
	// Re-prompt back to menu do it again -  or exit menu
	// specific validation:
	//loop menu on selection to get a final firstSelection
	//probably update Menu struct with specifics for errors here

	return userSelection, nil
}

func displayMenuTerminal(menu string) {
	//display the menu to terminal
	//somehow
}

func getUserInput() (string, error) {
	//user input in
	// - Enter selection
	//  Read result from terminal (STDIN)
	//  io.Readline (or something)
	//  Save to a variable
	input := "this will be the std.in or whatevs"
	return input, nil
}

func validateInput(input string, boundaries int) (string, error) {

	//do the validation
	// Ensure not blank
	// Ensure is within expected boundaries
	validated := "this is the validated user input"
	return validated, nil
}
