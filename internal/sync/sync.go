package sync

import (
	"fmt"
	"bytes"
	"time"
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
        
		var outBuf, errBuf bytes.Buffer
		cmd.Stdout = &outBuf
		cmd.Stderr = &errBuf

		start := time.Now()
		err := cmd.Run()
		duration := time.Since(start).Round(time.Millisecond)

		if err != nil {
            fmt.Printf("🚨 Installation failed for %s (after %s)\n", spec, duration)
			fmt.Printf("=== STDOUT ===\n%s\n", outBuf.String())
			fmt.Printf("=== STDERR ===\n%s\n", errBuf.String())
			return fmt.Errorf("pip install failed: %w", err)
        }
		
		fmt.Printf("✅ %s installed in %s\n", spec, duration)
	}

	return nil
}
