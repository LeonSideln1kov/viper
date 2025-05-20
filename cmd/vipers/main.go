package main


import (
	"fmt"
	"os"
	"os/exec"
)


const (
    venvDir    = ".venv"
    minPython  = "3.10"
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
	pythonPath, err := findPython()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Printf("VIPERs requires Python %s+ to create virtual environments\n", minPython)
		os.Exit(1)
	}

	if _, err_check := os.Stat(venvDir); os.IsNotExist(err_check) {
		cmd := exec.Command(pythonPath, "-m", "venv", venvDir)
	
		output, err_exec := cmd.CombinedOutput()
		if err_exec != nil {
			fmt.Printf("Failed to create virtual environment:\n")
            fmt.Printf("Command: %s\n", cmd.String())
            fmt.Printf("Output:\n%s\n", output)
            fmt.Printf("Error Details: %v\n", err)
            os.Exit(1)
		}
		fmt.Println("Virtual environment created successfully")
	} else {
		fmt.Printf("Directory %s already exists\n", venvDir)
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


func findPython() (string, error) {
    versions := []string{
        "python3.13", "python3.12", "python3.11", "python3.10", 
        "python3", "python", "py",
    }
    
    for _, ver := range versions {
        path, err := exec.LookPath(ver)
        if err == nil {
            return path, nil
        }
    }
    return "", fmt.Errorf("no Python installation found")
}
