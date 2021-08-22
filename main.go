package main

import (
	"fmt"
	"os"

	cd "github.com/ajean-zd/codeclub-challenge/configuredata"
)

func main() {
	fmt.Println("why hello and welcome to code club")

	fileContents, err := cd.GetFileContent("users.json")
	if err != nil {
		fmt.Println("check error and start again", err)
		os.Exit(1)
	}
	//fmt.Println(string(fileContents))
	userlist, err := cd.GetUsers(fileContents)

	fmt.Println(userlist)

}

// start
// take users
// search by id
// ask a question of each record (DO YOU MATCH?????)
// if they match output the record as it is
