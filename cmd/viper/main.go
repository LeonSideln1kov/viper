package main


import (
	"fmt"
	"os"
	"os/exec"
	"github.com/LeonSideln1kov/viper/internal/venv"
	"github.com/LeonSideln1kov/viper/internal/config"
	"github.com/LeonSideln1kov/viper/internal/resolver"
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
	case "lock", "--lock", "-l":
		generateLock()
	case "sync", "--sync", "-s":
		syncWithLock()
	}
}


func printHelp() {
	fmt.Println("VIPER - Virtual Python Environment Resolver and Simplified Python Package Manager")
	fmt.Println("Commands:")
	fmt.Println("  venv     Create virtual environment")
	fmt.Println("  help     Show help")
	fmt.Println("  install  Install packages")
	fmt.Println("  lock     Generate/update lock file")
	fmt.Println("  sync     Install from lock file")
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


func generateLock() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}

	for _, pkg := range cfg.Project.Dependencies {
		fmt.Println(resolver.ResolveVersion(pkg))
	}
}


func syncWithLock() {
	return
}