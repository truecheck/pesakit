/*
 * MIT License
 *
 * Copyright (c) 2021 PESAKIT - MOBILE MONEY TOOLBOX
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

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

// IsFileExist checks if a file exists. It returns true if the file exists,
// and it is not a directory as well as the file is regular file.
func IsFileExist(path string) bool {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	// return true if file is regular and not directory
	return !fi.IsDir() && fi.Mode().IsRegular()

}

// At creates home directory named .pesakit at the specified root path.
// If the root path is ".", the home directory will be created at the current
// working directory.
// If the root path is "", the home directory will be created at the home
// directory of the current user.
// At returns the path of the created home directory.
func At(rootPath string) (string, error) {
	pesakit := ".pesakit"
	switch {
	case rootPath == ".":
		cwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("could not get the current working directory, %w", err)
		}
		pesakitHome := filepath.Join(cwd, pesakit)

		return At(pesakitHome)
	case strings.TrimSpace(rootPath) == "":

		homePath, err := Get()
		if err != nil {
			return "", fmt.Errorf("error: %w", err)
		}

		return At(homePath)
	default:
		pesakitHome := filepath.Join(rootPath, pesakit)
		err := os.MkdirAll(pesakitHome, fs.ModePerm)
		if err != nil {
			return "", err
		}

		return pesakitHome, nil

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
