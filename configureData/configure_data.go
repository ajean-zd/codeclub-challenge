package configuredata

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

//get the data from the files and configure it for use

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
	//get the binary output
	//fmt.Printf("%#v\n", fleContents)
	//to confirm what is in a string
	fmt.Println(string(fileContents))
	return fileContents, nil
}

type User struct {
	ID int `json:"page"`
}

func GetUsers(fileContents []byte) (User, error) {

	users := User{}

	return users, nil //return a struct of users

}
