package sync

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/LeonSideln1kov/viper/internal/venv"
	"github.com/pelletier/go-toml/v2"
)


type LockFile struct {
	Packages map[string]string `toml:"packages"`
}


func SyncFromLock(lockPath string) error {
	data, err := os.ReadFile(lockPath)
	if err != nil {
		return fmt.Errorf("lock file read error: %w", err)
	}

	var lock LockFile
	if err := toml.Unmarshal(data, &lock); err != nil {
		return fmt.Errorf("lock file parse error: %w", err)
	}

	pipPath, err := venv.PipPath()
	if err != nil {
		return fmt.Errorf("no .venv created: %w", err)
	}

	for pkg, version := range lock.Packages {
		spec := fmt.Sprintf("%s==%s", pkg, version)
		cmd := exec.Command(pipPath, "install", "--force-reinstall", "-v", spec)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
            return fmt.Errorf("failed to install %s: %w", spec, err)
        }

		fmt.Printf("✅ Installed %s@%s\n", pkg, version)
	}

	return nil
}
