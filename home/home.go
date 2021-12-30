package home

import (
	"fmt"
	"os"
	"path/filepath"

	home "github.com/mitchellh/go-homedir"
)

// At creates home directory named .pesakit at the specified root path.
// If the root path is not specified, the current working directory will be used.
func At(rootPath string) error {
	pesakit := ".pesakit"
	if rootPath == "" || rootPath == "." {
		// use the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get the current working directory, %w", err)
		}
		pesakitHome := filepath.Join(cwd, pesakit)
		return At(pesakitHome)
	}
	pesakitHome := filepath.Join(rootPath, pesakit)
	return os.MkdirAll(pesakitHome, 0755)
}

// Init creates a home directory named at $HOME/.pesakit if not available
func Init() error {
	homePath, err := home.Dir()
	if err != nil {
		return fmt.Errorf("could not get the home directory, %w", err)
	}
	return At(homePath)
}
