package config

import (
	"fmt"
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

type PyProject struct {
	Project struct {
		Dependencies []string `toml:"dependencies"`
	} `toml:"project"`
}

func Load() (*PyProject, error) {
	data, err := os.ReadFile("pyproject.toml")
	if err != nil {
		return nil, fmt.Errorf("error: missing pyproject.toml")
	}

	var cfg PyProject
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error: invalid TOML format (%w)", err)
	}
	return &cfg, nil
}
