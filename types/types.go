package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type User struct {
	ID              int    `json:"_id"`
	URL             string `json:"url"`
	EXTERNAL_ID     string `json:"external_id"`
	NAME            string `json:"name"`
	ALIAS           string `json:"alias"`
	CREATED_AS      string `json:"created_at"`
	ACTIVE          bool   `json:"active"`
	VERIFIED        bool   `json:"verified"`
	SHARED          bool   `json:"shared"`
	LOCALE          string `json:"locale"`
	TIMEZONE        string `json:"timezone"`
	LAST_LOGIN_AT   string `json:"last_login_at"`
	EMAIL           string `json:"email"`
	PHONE           string `json:"phone"`
	SIGNATURE       string `json:"signature"`
	ORGANIZATION_ID int    `json:"organization_id"`
	SUSPENDED       bool   `json:"suspended"`
	ROLE            string `json:"role"`
	TAGS            Tags
}

type Tags struct {
	TAG string
}

type Ticket struct {
	ID              string `json:"_id"`
	URL             string `json:"url"`
	EXTERNAL_ID     string `json:"external_id"`
	CREATED_AT      string `json:"created_at"`
	TYPE            string `json:"type"`
	SUBJECT         string `json:"subject"`
	DESCRIPTION     string `json:"description"`
	PRIORITY        string `json:"priority"`
	STATUS          string `json:"status"`
	SUBMITTER_ID    int    `json:"submitter_id"`
	ASSIGNEE_ID     int    `json:"assignee_id"`
	ORGANIZATION_ID int    `json:"organization_id"`
	TAGS            Tags
	HAS_INCIDENTS   bool   `json:"has_incidents"`
	DUE_AT          string `json:"due_at"`
	VIA             string `json:"via"`
}

type Org struct {
	ID             string `json:"_id"`
	URL            string `json:"url"`
	EXTERNAL_ID    string `json:"external_id"`
	NAME           string `json:"name"`
	DOMAIN_NAME    string `json:"domain_name"`
	CREATED_AT     string `json:"created_at"`
	DETAILS        string `json:"details"`
	SHARED_TICKETS bool   `json:"shared_tickets"`
	TAGS           Tags
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
