package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/LeonSideln1kov/viper/internal/config"
	"github.com/LeonSideln1kov/viper/internal/resolver"
	"github.com/LeonSideln1kov/viper/internal/sync"
	"github.com/LeonSideln1kov/viper/internal/venv"
	"github.com/pelletier/go-toml/v2"
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
	
    if err != nil {
        return fmt.Errorf("venv pip missing: %w", err)
    }
	
	for _, pkg := range cfg.Project.Dependencies {
		venv.InstallPackage(pkg)
    }
	return nil
}


func generateLock() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	lockData := struct {
		Metadata struct {
			ViperVersion string `toml:"viper-version"`
			PythonVersion string `toml:"python-version"`
			GeneratedAt string `toml:"generated-ad"`
		}`toml:"metadata"`
		Packages map[string]string `toml:"packages"`
	}{}
	
	// TODO get from cfg
	lockData.Metadata.ViperVersion = "0.1.0"
	lockData.Metadata.GeneratedAt = time.Now().UTC().Format(time.RFC3339)

	// TODO raise error if python or .venv havn't been found  
	pyPath, _ := venv.PythonPath()
	cmd := exec.Command(pyPath, "--version")
    out, _ := cmd.CombinedOutput()
    lockData.Metadata.PythonVersion = strings.TrimSpace(string(out))

	lockData.Packages = make(map[string]string)

	for _, pkg := range cfg.Project.Dependencies {
		name, ver, err := resolver.ResolveVersion(pkg)
		if err != nil {
			panic(err)
		}
		lockData.Packages[name] = ver
	}

	data, err := toml.Marshal(lockData)
	if err != nil {
		return fmt.Errorf("TOML encoding error: %w", err)
	}

	if err := os.WriteFile("viper.lock", data, 0644); err != nil {
		return fmt.Errorf("file write error: %w", err)
	}

	fmt.Println("Lock file generated: viper.lock")
    return nil
}


func syncWithLock() {
	sync.SyncFromLock("viper.lock")
}