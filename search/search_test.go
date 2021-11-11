package search

import (
	"fmt"
	"testing"

	load "github.com/ajean-zd/codeclub-challenge/load"
	"github.com/stretchr/testify/assert"
)

func Test_Returns_Single_UserId(t *testing.T) {

	search_id := 1
	users := load.PopulateUserData("../files/users.json")
	fmt.Println(users)

	result, err := ReturnQuery(search_id, users)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
