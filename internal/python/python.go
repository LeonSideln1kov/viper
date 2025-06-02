package python

import (
	"fmt"
	"os/exec"
)

func FindPython() (string, error) {
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