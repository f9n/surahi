package virtualenv

import (
	"errors"
	"os"
	"os/exec"

	"github.com/pleycpl/surahi/util"
)

//	Create new venv with virtualenv command, and then read filenames in new venv.
// 	After that Update VirtualEnvInitFiles global variable.
// 	Maybe VirtualEnv file list change next time.
func GetVirtualEnvInitFiles() ([]string, error) {
	initfiles := []string{}
	err := CreateVenv()
	if err != nil {
		return []string{}, err
	}
	err = SaveVirtualenvInitFileList()
	if err != nil {
		return []string{}, err
	}
	HOME := os.Getenv("HOME")
	path := HOME + "/go/src/github.com/pleycpl/surahi/virtualenvinitfilelist.txt"
	initfiles, err = util.GetLinesInFile(path)
	if err != nil {
		return []string{}, errors.New("We can't read virtualenvinitfilelist.txt")
	}
	return initfiles, nil
}

// Save virtualenv init filenames to virtualenvinitfilelist.txt
func SaveVirtualenvInitFileList() error {
	cmd := exec.Command(
		"/bin/bash",
		"-c",
		//"ls $HOME/go/src/github.com/pleycpl/surahi/venv/lib/python*/site-packages > $HOME/go/src/github.com/pleycpl/surahi/virtualenvinitfilelist.txt",
		"./shellscripts/save_virtualenv_init_filelist.sh",
	)
	err := cmd.Run()
	if err != nil {
		return errors.New("We can't list file in venv and pipe virtualenvinitfilelist.txt file.")
	}
	return nil
}

// Create new venv with virtualenv command.
func CreateVenv() error {
	cmd := exec.Command(
		"/bin/bash",
		"-c",
		//"virtualenv $HOME/go/src/github.com/pleycpl/surahi/venv",
		"./shellscripts/create_venv.sh",
	)
	err := cmd.Run()
	if err != nil {
		return errors.New("We can't create venv file. ")
	}
	return nil
}
