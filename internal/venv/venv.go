package venv


import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	python "github.com/LeonSideln1kov/viper/internal/python"
)


const (
    venvDir    = ".venv"
    minPython  = "3.10"
)


func CreateVenv() {
	pythonPath, err := python.FindPython()
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


func PipPath() (string, error) {
	var path string

    if runtime.GOOS == "windows" {
        path = filepath.Join(venvDir, "Scripts", "pip.exe") 
    } else {
		path = filepath.Join(venvDir, "bin", "pip")
	}
    
	if _, err := os.Stat(path); os.IsNotExist(err) {
        return "", fmt.Errorf("pip not found in virtual environment")
    }

	return path, nil
}
