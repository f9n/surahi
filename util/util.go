package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

//	Get Lines in file with path arguman.
//	And return lines and error status.
func GetLinesInFile(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(path, err)
		return []string{}, errors.New("We can't read file")
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}
