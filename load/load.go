package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// Given files: tickets, organisations, users

var (
	Orgpath    = "./../files/organizations.json"
	Userpath   = "./../files/users.json"
	Ticketpath = "./../files/organizations.json"
)

func load_data() {
	// parse json from files to data structures and unmarshall
	//needs a cleanup and refactor

	users := PopulateUserData(Userpath)
	fmt.Println(users[0])
	tickets := PopulateTicketData(Ticketpath)
	fmt.Println(tickets[0])
	organizations := PopulateOrgData(Orgpath)
	fmt.Println(organizations[0])

	//
	// figure schema?
	// do this dynamically from data (at runtime)
	//
	//

	// Outputs:
	// loaded relational data in some form of data type , with a searchable schema
	// List of fields and their data types
	// Identify any relationships & foreign keys  - how?

	// some rando info we did
	//fmt.Println(organizations[0])
	// fileContents, err := ty.GetFileContent("users.json")
	// if err != nil {
	// 	fmt.Println("check error and start again", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(fileContents))
	// json.Unma

}
func PopulateUserData(fileName string) []User {
	fileContents, err := GetFileContent(fileName)
	if err != nil {
		fmt.Println("check error and start again", err)
		os.Exit(1)
	}
	var users []User

	json.Unmarshal([]byte(fileContents), &users)
	return users
}

func PopulateTicketData(fileName string) []Ticket {
	fileContents, err := GetFileContent(fileName)
	if err != nil {
		fmt.Println("check error and start again", err)
		os.Exit(1)
	}
	var tickets []Ticket
	json.Unmarshal([]byte(fileContents), &tickets)
	return tickets
}

func PopulateOrgData(fileName string) []Org {
	fileContents, err := GetFileContent(fileName)
	if err != nil {
		fmt.Println("check error and start again", err)
		os.Exit(1)
	}
	var orgs []Org
	json.Unmarshal([]byte(fileContents), &orgs)
	return orgs
}

func GetFileContent(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "File could not open")
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "File unable to be read")
	}
	//to get the binary output
	//fmt.Printf("%#v\n", fleContents)
	//to confirm what is in a string
	fmt.Println(string(fileContents))

	return fileContents, nil
}
