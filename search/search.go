package search

import (
	"fmt"

	ty "github.com/ajean-zd/codeclub-challenge/types"
)

//return a simple query
func ReturnQuery(id int) string {

	users := ty.PopulateTicketData("../files/users.json")

	searchID := 1

	fmt.Println(users, searchID)

}
