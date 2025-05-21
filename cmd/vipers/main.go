package main


import (
	"fmt"
	"os"
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

func installPackages() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}
	
	// pipPath := venv.PipPath()
	
	for _, pkg := range cfg.Project.Dependencies {
		fmt.Println(pkg)
	}
}
