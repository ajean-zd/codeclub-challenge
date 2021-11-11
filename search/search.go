package search

import (
	"fmt"

	"github.com/ajean-zd/codeclub-challenge/load"
)

//TODO plan this package

// Given various options on the screen you can have a few options
//     Store the data type, the key, and the value and choose what search to do. A search may have all or none of these

//     If search is a key/value pair
// Mapping from menu item to key in selected data type
// On selection of key, search value given
// Value can be blank and will search for blank
// Remove all cases for search purposes

//     If search is a value only
// Search the datatype for the value. If found, return key and value
// If no match, search the next datatype for the value. And do this for all data types. Return the data type with the found key/value

//     If search is a key only
// Return all the values for the key

// Inputs:
// The search type
// The data to search

// Outputs:
// Some sort of structure containing search results:
// Needs to cater for results that are
// Nil
// One
// Many
// All records

//return a simple query
func ReturnQuery(search_term int, users []load.User) (string, error) {

	result := ""

	fmt.Println(users)
	for _, user := range users {
		fmt.Println(user.ID)

	}

	// }
	// us := users
	// fmt.Printf("%#v\n", us)

	return result, nil

}
