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
	"github.com/pesakit/pesakit/config"
	"github.com/pesakit/pesakit/mno"
	"github.com/spf13/cobra"
	"strings"
)

const (
	mnoName    = "mno"
	mnoUsage   = "mobile network operator"
	mnoDefault = ""
)

func SetMnoFlag(cmd *cobra.Command, p Type) {

	defaultMnoConf := config.DefaultMnoConf()
	switch p {
	case LOCAL:
		cmd.Flags().StringVar(&defaultMnoConf.Value, mnoName, mnoDefault, mnoUsage)
	case PERSISTENT:
		cmd.PersistentFlags().StringVar(&defaultMnoConf.Value, mnoName, mnoDefault, mnoUsage)
	}
}

func LoadMnoConfig(cmd *cobra.Command, p Type) (mno.Mno, error) {
	var (
		mnoValue mno.Mno
		err      error
		value    string
	)

	switch p {
	case LOCAL:
		value, err = cmd.Flags().GetString(mnoName)
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(value) == "" {
			return "", fmt.Errorf("mno is required [tigo,airtel,vodacom]")
		}

	case PERSISTENT:
		value, err = cmd.PersistentFlags().GetString(mnoName)
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(value) == "" {
			return "", fmt.Errorf("mno is required [tigo,airtel,vodacom]")
		}
	}

	mnoValue = mno.Mno(value)
	if mnoValue == mno.Unknown {
		return "", fmt.Errorf("mno is required [tigo,airtel,vodacom]")
	}

	return mnoValue, nil
}
