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

package cli

import (
	"context"
	"github.com/pesakit/pesalib/mna"
	"github.com/urfave/cli/v2"
	"io"
	"strings"
	"text/tabwriter"
	"text/template"
)

const (
	flagPhoneNumber       = "phone-number"
	flagPhoneNumberDesc   = "phone number to inspect"
	mnaCommandName        = "mna"
	mnaCommandUsage       = "mobile number assignment (mna) command"
	mnaCommandDescription = `mna command deals with mobile number assignment operations.
like checking the MNO of a given mobile number`
)

func MnaCommand(writer io.Writer) *cli.Command {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  flagPhoneNumber,
			Usage: flagPhoneNumberDesc,
			Value: "",
		},
	}
	return &cli.Command{
		Name:        mnaCommandName,
		Usage:       mnaCommandUsage,
		Description: mnaCommandDescription,
		Flags:       flags,
		Action: func(c *cli.Context) error {
			// get phone number
			phoneNumber := c.String(flagPhoneNumber)
			return checkPhoneNumberAction(c.Context, writer, phoneNumber)
		},
	}
}

type mnaCommandResponse struct {
	PhoneNumber    string
	CommonName     string
	RegisteredName string
	Prefixes       string
	Status         string
}

const (
	mnaCommandResponseTemplate = `{{"Phone Number:\t"}}{{.PhoneNumber}}
{{"Common Name:\t"}}{{.CommonName}}
{{"Registered Name:\t"}}{{.RegisteredName}}
{{"Prefixes:\t"}}{{.Prefixes}}
{{"Status:\t"}}{{.Status}}
`
)

func checkPhoneNumberAction(ctx context.Context, writer io.Writer, number string) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()
	operator, phoneNumber, err := mna.OperatorFromPhoneNumber(number)
	if err != nil {
		return err
	}

	response := mnaCommandResponse{
		PhoneNumber:    phoneNumber,
		CommonName:     operator.CommonName(),
		RegisteredName: operator.RegisteredName(),
		Prefixes:       strings.Join(operator.Prefixes(), ","),
		Status:         operator.Status(),
	}

	t := template.Must(template.New("mna").Parse(mnaCommandResponseTemplate))

	newWriter := tabwriter.NewWriter(writer, 4, 4, 4, ' ', tabwriter.TabIndent)
	defer newWriter.Flush()

	return t.Execute(newWriter, response)
}
