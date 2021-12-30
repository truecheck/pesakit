package pesakit

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
