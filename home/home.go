package home

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	home "github.com/mitchellh/go-homedir"
)

// IsDirExist checks if a directory exists
func IsDirExist(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// At creates home directory named .pesakit at the specified root path.
// If the root path is ".", the home directory will be created at the current
// working directory.
// If the root path is "", the home directory will be created at the home
// directory of the current user.
func At(rootPath string) error {
	pesakit := ".pesakit"
	switch {
	case rootPath == ".":
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get the current working directory, %w", err)
		}
		pesakitHome := filepath.Join(cwd, pesakit)

		return At(pesakitHome)
	case strings.TrimSpace(rootPath) == "":

		homePath, err := Get()
		if err != nil {
			return fmt.Errorf("error: %w", err)
		}

		return At(homePath)
	default:
		pesakitHome := filepath.Join(rootPath, pesakit)
		return os.MkdirAll(pesakitHome, fs.ModePerm)
	}

}

// Get returns the home directory of the current user.
func Get() (string, error) {
	dir, err := home.Dir()
	if err != nil {
		return "", fmt.Errorf("could not get the home directory, %w", err)
	}

	return dir, nil
}
