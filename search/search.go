package search

import (
	"fmt"

	ty "github.com/ajean-zd/codeclub-challenge/types"
)

//return a simple query
func ReturnQuery(search_term int, users []ty.User) (string, error) {

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
