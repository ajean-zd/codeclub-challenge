package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("why hello and welcome to code club")

	file, err := os.Open("users.json")

	if err != nil {
		fmt.Println("There was some sort of problemo with opening the file", err)
		os.Exit(1)
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("the file read failed", err)
		os.Exit(1)
	}

	fmt.Println(string(fileContents))

	defer file.Close()

}

// start
// take users
// search by id
// ask a question of each record (DO YOU MATCH?????)
// if they match output the record as it is
