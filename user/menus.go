package user

const MenuOne = "format this // Select data type// 1: User// 2: Ticket/ 3: Organisation// 4:Exit"

const MenuTwo = "format of menu2"

// 	On selection of data type, list fields available to query on for the selected record
//     Option to go back to main/last menu
//     Option to search on a value without selecting a key
// Option to exit
// Fields can be Selected by number or value.
// Ie, 1 OR user_id
// On entry of field selection (ie, user types 1 or ‘user-id’)  display chosen key, and also the data type (alphanumerical, number, etc):
//         “Please enter the value (alphanumeric) for blah”
// What happens if user has selected org but they are looking for a userid and it’s in another data type? -  Consider returning a prompt that states “you want to search for x but this is in y menu. Remain here or search in blah instead?”
// Option to display all values of the key selected
// Option to search on only blank values
// Options to return only blank search data
// How does user choose value if displayed? Number, text?

const MenuThree = "format of menu 3"

// Menu contains all the information on the menus and will expand to fit the needs
type Menu struct {
	Display    string
	Boundaries int
}
