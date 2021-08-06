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
		fmt.Println("OH NO", err)
		os.Exit(1)
	}

	//fmt.Printf("%#v\n", file)
	//fmt.Printf(file.Name())

	thingy, err := ioutil.ReadAll(file)
	fmt.Println(thingy)

	defer file.Close()

}

// start
// take users
// search by id
// ask a question of each record (DO YOU MATCH?????)
// if they match output the record as it is
