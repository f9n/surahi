/*
	How to install this tool?
	go get github.com/pleycpl/surahi

	Task List:
		- Task 1 = You use terminal spinner in "FindAllPyPackagesPath" function, because it takes a lot of time.
		- Task 2 = Research that We often use bash script in functions, it is not good? Because we need bash.
		- Task 3 = Use color package for stdout.
		- Task 4 = In future, we use specific directory for project, like we created file in code directory.Example: allvenvpaths.txt
*/
package main

import (
	"fmt"

	"github.com/pleycpl/surahi/pythonpackages"
	"github.com/pleycpl/surahi/virtualenv"
)

// It stores filenames in "venv/lib/python*/site-packages" that when you create init venv directory with virtualenv.
//VirtualenvInitFilenames []string

func main() {
	fmt.Println("[+] Program Started!")
	pathlist, err := pythonpackages.GetAllPyPackagesPath()
	if err != nil {
		fmt.Println("[-] Error: ", err)
		return
	}
	fmt.Println(pathlist)
	initfiles, err2 := virtualenv.GetVirtualEnvInitFiles()
	if err2 != nil {
		fmt.Println("[-] Error: ", err2)
		return
	}
	fmt.Println(initfiles)
	fmt.Println("[+] Program Finished!")
}
