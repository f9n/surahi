package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	VirtualEnvInitFiles []string
	HOME                = os.Getenv("HOME")
)

func main() {
	fmt.Println("[+] Program Started!")
	if FindAllPyPackagesPath() != nil {
		fmt.Println("Error: we can't find all py packages path")
		return
	}
	path := HOME + "/go/src/github.com/pleycpl/surahi/allvenvpaths.txt"
	lines, err := GetLinesInFile(path)
	if err != nil {
		fmt.Println("Error: we can't get lines in file\n", err)
		return
	}
	fmt.Println(lines)
}

//	Get all venv path and write to allpyvenvpaths.txt
//	Most python developer are using venv directory name for virtualenv.
func FindAllPyPackagesPath() error {
	// Task 1 = You use terminal spinner here, because it takes a lot of time.
	cmd := exec.Command(
		"/bin/bash",
		"-c",
		// "find $HOME -name site-packages -type d | grep 'python3' 2>/dev/null > $HOME/go/src/github.com/pleycpl/surahi/allpyvenvpaths.txt",
		"find $HOME -name venv -type d > $HOME/go/src/github.com/pleycpl/surahi/allvenvpaths.txt",
	)
	err := cmd.Run()
	return err
}

//	Get Lines in file with path arguman.
//	And return lines and error status.
func GetLinesInFile(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}

//	Create new venv with virtualenv command, and then read filenames in new venv.
// 	After that Update VirtualEnvInitFiles global variable.
// 	Maybe VirtualEnv file list change next time.
func UpdateVirtualEnvInitFiles() {}
