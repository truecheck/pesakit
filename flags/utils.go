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

package flags

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func MarkHiddenExcept(flags *pflag.FlagSet, unhidden ...string) {
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

func MarkRequired(command *cobra.Command, flagType Type, required ...string) error {
	switch flagType {
	case PERSISTENT:
		for _, s := range required {
			err := command.MarkPersistentFlagRequired(s)
			if err != nil {
				return err
			}
			continue
		}

		return nil

	case LOCAL:
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
