package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "venv":
		createVenv()
	case "help":
		printHelp()
	}

}

func createVenv() {
	cmd := exec.Command("python3", "-m", "venv", ".venv")
	if _, err_check := os.Stat(".venv"); os.IsNotExist(err_check) {
		output, err_exec := cmd.CombinedOutput()
		if err_exec != nil {
			panic(err_exec)
		}
		fmt.Println(string(output))
	} else {
		fmt.Println(".venv already exists")
	}
}


func printHelp() {
	fmt.Println("VIPERs - Virtual Python Environment Resolver and Simplified Python Package Manager")
	fmt.Println("Commands:")
	fmt.Println("  venv     Create virtual environment")
	fmt.Println("  help     Show help")
	// fmt.Println("  install  Install packages")
	// fmt.Println("  lock     Generate lock file")
}
