package pesakit

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io"
	"os"
	"path/filepath"
)

func markHiddenExcept(flags *pflag.FlagSet, unhidden ...string) {
	flags.VisitAll(func(flag *pflag.Flag) {
		name := flag.Name
		if !contains(unhidden, name) {
			flag.Hidden = true
		}
	})
}

// contains returns true if the string is in the slice
func contains(b []string, i string) bool {
	for _, s := range b {
		if s == i {
			return true
		}
	}

	return false
}

func markFlagsRequired(command *cobra.Command, flagType flagType, required ...string) error {
	switch flagType {
	case globalFlagType:
		for _, s := range required {
			err := command.MarkPersistentFlagRequired(s)
			if err != nil {
				return err
			}
			continue
		}

		return nil

	case localFlagType:
		for _, s := range required {
			err := command.MarkFlagRequired(s)
			if err != nil {
				return err
			}
			continue
		}

		return nil

	default:
		return fmt.Errorf("unknown flag type: %v", flagType)
	}
}

// copyFile copies a file from src to dst. If src or dst files do not exist
// it returns an error. src is expected to be a regular file and dst is expected
// to be a directory or else error is returned.
// copyFile if successful returns the new path of the file in dst.
func copyToDir(src, dst string) (string, error) {
	srcFileStat, err := os.Stat(src)
	if err != nil {
		return "", err
	}
	if !srcFileStat.Mode().IsRegular() {
		return "", fmt.Errorf("%s is not a regular file", src)
	}

	dstDirStat, err := os.Stat(dst)
	if err != nil {
		return "", err
	}
	if !dstDirStat.Mode().IsDir() {
		return "", fmt.Errorf("%s is not a directory", dst)
	}

	dstFile := filepath.Join(dst, srcFileStat.Name())
	srcFile, err := os.Open(src)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()
	destination, err := os.Create(dstFile)
	if err != nil {
		return "", err
	}
	defer destination.Close()
	_, err = io.Copy(destination, srcFile)
	if err != nil {
		return "", err
	}
	return dstFile, nil
}
