package pythonpackages

import (
	"errors"
	"os"
	"os/exec"

	"github.com/pleycpl/surahi/util"
)

// Get All python package paths
func GetAllPyPackagesPath() ([]string, error) {
	venvpaths := []string{}
	err := SaveAllPyPackagesPath()
	if err != nil {
		return []string{}, err
	}
	HOME := os.Getenv("HOME")
	path := HOME + "/go/src/github.com/pleycpl/surahi/allvenvpaths.txt"
	venvpaths, err = util.GetLinesInFile(path)
	if err != nil {
		return []string{}, errors.New("We can't read lines.")
	}
	return venvpaths, nil
}

//	Get all venv path and save to allpyvenvpaths.txt
//	Most python developer are using venv directory name for virtualenv.
func SaveAllPyPackagesPath() error {
	cmd := exec.Command(
		"/bin/bash",
		"-c",
		// "find $HOME -name site-packages -type d | grep 'python3' 2>/dev/null > $HOME/go/src/github.com/pleycpl/surahi/allpyvenvpaths.txt",
		"find $HOME -name venv -type d > $HOME/go/src/github.com/pleycpl/surahi/allvenvpaths.txt",
	)
	err := cmd.Run()
	if err != nil {
		return errors.New("We cant't find venv directorys")
	}
	return nil
}
