package configuredata

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

//get the data from the files and configure it for use

func getFileContent(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, errors.Wrap(err, "File could not open")
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "File unable to be read")
	}

	fmt.Println(string(fileContents))
	defer file.Close()
	return fileContents, nil

}
