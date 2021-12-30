package home

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	home "github.com/mitchellh/go-homedir"
)

// At creates home directory named .pesakit at the specified root path.
// If the root path is ".", the home directory will be created at the current
// working directory.
// If the root path is "", the home directory will be created at the home
// directory of the current user.
func At(rootPath string) error {
	pesakit := ".pesakit"
	if rootPath == "." {
		// use the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get the current working directory, %w", err)
		}
		pesakitHome := filepath.Join(cwd, pesakit)
		return At(pesakitHome)
	} else if strings.TrimSpace(rootPath) == "" {
		homePath, err := home.Dir()
		if err != nil {
			return fmt.Errorf("could not get the home directory, %w", err)
		}
		return At(homePath)
	} else {
		pesakitHome := filepath.Join(rootPath, pesakit)
		return os.MkdirAll(pesakitHome, 0755)
	}

}
