package main

import (
	"fmt"

	ty "github.com/ajean-zd/codeclub-challenge/types"
)

func main() {
	fmt.Println("why hello and welcome to code club")

	users := ty.PopulateUserData("./files/users.json")
	fmt.Println(users[0])
	tickets := ty.PopulateTicketData("./files/tickets.json")
	fmt.Println(tickets[0])
	organizations := ty.PopulateOrgData("./files/organizations.json")
	fmt.Println(organizations[0])
	// fileContents, err := ty.GetFileContent("users.json")
	// if err != nil {
	// 	fmt.Println("check error and start again", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(fileContents))
	// json.Unmarshal([]byte(fileContents), &ty.User{})

}

// start
// take users
// search by id
// ask a question of each record (DO YOU MATCH?????)
// if they match output the record as it is
