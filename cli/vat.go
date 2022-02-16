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
	"fmt"
	"github.com/pesakit/pesalib/vat"
	"github.com/urfave/cli/v2"
	"io"
	"text/tabwriter"
	"text/template"
)

const (
	flagAmountName  = "amount"
	flagAmountUsage = "amount involved in transaction"
	flagFromName    = "from"
	flagFromUsage   = "transaction initiator medium"
	flagToName      = "to"
	flagToUsage     = "transaction recipient medium"
	vatCommandName  = "vat"
	vatCommandUsage = "calculate VAT for different transactions"
	vatCommandDesc  = `pesakit vat - calculate VAT for a given transaction. It can calculate VAT for
same MNO transactions,cross-MNO transactions, and VAT for sending money to the
banks as well as VAT for cash withdraw.

Below command calculates VAT for a given transaction. if amount is not specified,
it will print the breakdown for all ranges.

pesakit vat --from=<from> --to=<to> --amount=<amount>
e.g pesakit vat --from=vodacom --to=bank --amount=10000
`
)

func VatCommand(writer io.Writer, calculator vat.Calculator) *cli.Command {
	// Flags are from to and amount
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     flagFromName,
			Usage:    flagFromUsage,
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagToName,
			Usage:    flagToUsage,
			Required: true,
		},
		&cli.Float64Flag{
			Name:  flagAmountName,
			Usage: flagAmountUsage,
			Value: 0,
		},
	}

	return &cli.Command{
		Name:        vatCommandName,
		Usage:       vatCommandUsage,
		Description: vatCommandDesc,
		Flags:       flags,
		Action: func(cliCtx *cli.Context) error {
			// Get the values from the flags
			from := cliCtx.String(flagFromName)
			receiver := cliCtx.String(flagToName)
			amount := cliCtx.Float64(flagAmountName)

			// create request
			request := vat.Request{
				From:   from,
				To:     receiver,
				Amount: amount,
			}

			return vatAction(cliCtx.Context, writer, calculator, request)
		},
	}
}

func vatAction(context context.Context, writer io.Writer, calculator vat.Calculator, request vat.Request) error {
	responses, err := calculator.Calculate(context, request)
	if err != nil {
		return fmt.Errorf("failed to calc VAT: %w", err)
	}

	// Print the response  to the writer using text templates and a tab writer
	return printVatResponse(writer, responses...)
}

func printVatResponse(writer io.Writer, responses ...vat.Response) error {
	if len(responses) == 0 {
		return fmt.Errorf("no response to print")
	}
	if len(responses) == 1 {
		return printOneResponse(writer, responses[0])
	}

	return printManyResponses(writer, responses...)
}

const oneResponseTemplate = `{{"From:\t"}}{{.From}}
{{"To:\t"}}{{.To}}
{{"Amount\t"}}{{.Amount}}
{{"Vat\t"}}{{.ServiceCharge}}
{{"Total\t"}}{{.Total}}
`

func printOneResponse(writer io.Writer, response vat.Response) error {
	t := template.Must(template.New("vat").Parse(oneResponseTemplate))

	newWriter := tabwriter.NewWriter(writer, 4, 4, 4, ' ', tabwriter.TabIndent)
	defer newWriter.Flush()
	return t.Execute(newWriter, response)
}

type longResponse struct {
	Title string
	Items []responseBody
}

type responseBody struct {
	MinAmount      float64
	MaxAmount      float64
	ServiceCharge  float64
	GovernmentLevy float64
	Total          float64
}

func createParsableResponse(responses []vat.Response) longResponse {
	if len(responses) == 0 {
		return longResponse{}
	}
	title := fmt.Sprintf("VAT Breakdown for %s to %s transactions", responses[0].From, responses[0].To)

	items := make([]responseBody, len(responses))
	for i, response := range responses {
		items[i] = responseBody{
			MinAmount:      response.MinAmount,
			MaxAmount:      response.MaxAmount,
			ServiceCharge:  response.ServiceCharge,
			GovernmentLevy: response.GovernmentLevy,
			Total:          response.Total,
		}
	}

	return longResponse{
		Title: title,
		Items: items,
	}
}

const logResponseParsableTemplate = `{{.Title}}
{{"Min Amount\t"}}{{"Max Amount\t"}}{{"Charge\t"}}{{"Levy\t"}}{{"Total\t"}}{{range .Items}}
{{.MinAmount}}{{"\t"}}{{.MaxAmount}}{{"\t"}}{{.ServiceCharge}}{{"\t"}}{{.GovernmentLevy}}{{"\t"}}{{.Total}}{{end}}
`

func printManyResponses(writer io.Writer, responses ...vat.Response) error {
	parsableResponse := createParsableResponse(responses)
	t := template.Must(template.New("logResponse").Parse(logResponseParsableTemplate))
	newWriter := tabwriter.NewWriter(writer, 8, 8, 8, '\t', 0)
	defer newWriter.Flush()
	return t.Execute(newWriter, parsableResponse)
}
