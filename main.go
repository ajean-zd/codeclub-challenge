package main

import (
	"fmt"
	"os"
	// cd "github.com/ajean-zd/codeclub-challenge/configuredata"
)

func main() {
	fmt.Println("why hello and welcome to code club")

	err, fileContents := cd.getFileContent("users.json")
	if err != nil {
		fmt.Println("nope nope nope", err)
		os.Exit(1)
	}
	fmt.Println(string(fileContents))

}

// start
// take users
// search by id
// ask a question of each record (DO YOU MATCH?????)
// if they match output the record as it is
