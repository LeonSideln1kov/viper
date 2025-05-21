package main


import (
	"fmt"
	"os"
	"os/exec"
	venv "github.com/LeonSideln1kov/vipers/internal/venv"
	// python "github.com/LeonSideln1kov/vipers/internal/python"
	config "github.com/LeonSideln1kov/vipers/internal/config"
)


func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "help", "--help", "-h":
		printHelp()
	case "venv", "--venv", "-v":
		venv.CreateVenv()
	case "install", "--install", "-i":
		installPackages()
	}

}


func printHelp() {
	fmt.Println("VIPERs - Virtual Python Environment Resolver and Simplified Python Package Manager")
	fmt.Println("Commands:")
	fmt.Println("  venv     Create virtual environment")
	fmt.Println("  help     Show help")
	fmt.Println("  install  Install packages")
	// fmt.Println("  lock     Generate lock file")
}

func installPackages() error{
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}
	
	pipPath, err := venv.PipPath()
    if err != nil {
        return fmt.Errorf("venv pip missing: %w", err)
    }
	
	for _, pkg := range cfg.Project.Dependencies {
        cmd := exec.Command(pipPath, "install", pkg)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        
        if err := cmd.Run(); err != nil {
            return fmt.Errorf("failed to install %s: %w", pkg, err)
        }
    }
	return nil
}
