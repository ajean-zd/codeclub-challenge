package configuredata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetFileContent_ReturnsContent(t *testing.T) {

	fileName := "users.json"
	result, err := GetFileContent(fileName)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func Test_NoFileErrors(t *testing.T) {
	fileName := "losers.json"

	result, err := GetFileContent(fileName)

	assert.Error(t, err)
	assert.Nil(t, result)
}
