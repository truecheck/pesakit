//

package airtel

import (
	"context"
	"fmt"
	"github.com/pesakit/pesakit/internal"
)

type (
	// Params
	//From	query	integer(int64)	true	Date from which transactions are to be fetched.
	//To	query	integer(int64)	true	Date until transactions are to be fetched.
	//Limit	query	integer(int64)	true	The number of transactions to be fetched on a page.
	//Offset	query	integer(int64)	true	Page number from which transactions are to be fetched.
	Params struct {
		From   int64 `json:"from"`
		To     int64 `json:"to"`
		Limit  int64 `json:"limit"`
		Offset int64 `json:"offset"`
	}
	TransactionService interface {
		Summary(ctx context.Context, params Params) (ListTransactionsResponse, error)
	}
)

func queryParamsOptions(params Params, m map[string]string) internal.RequestOption {
	from, to, limit, offset := params.From, params.To, params.Limit, params.Offset
	if from > 0 {
		m["from"] = fmt.Sprintf("%d", from)
	}
	if to > 0 {
		m["to"] = fmt.Sprintf("%d", to)
	}
	if limit > 0 {
		m["limit"] = fmt.Sprintf("%d", limit)
	}
	if offset > 0 {
		m["offset"] = fmt.Sprintf("%d", offset)
	}

	return internal.WithQueryParams(m)
}

func (c *Client) Summary(ctx context.Context, params Params) (ListTransactionsResponse, error) {

	token, err := c.checkToken(ctx)
	if err != nil {
		return ListTransactionsResponse{}, err
	}

	var opts []internal.RequestOption

	hs := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "*/*",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	queryMap := make(map[string]string, 4)
	queryMapOpt := queryParamsOptions(params, queryMap)
	headersOpt := internal.WithRequestHeaders(hs)
	opts = append(opts, headersOpt, queryMapOpt)
	req := c.makeInternalRequest(TransactionSummary, nil, opts...)

	res := new(ListTransactionsResponse)

	_, err = c.base.Do(ctx, "transaction summary", req, res)
	if err != nil {
		return ListTransactionsResponse{}, err
	}
	return *res, nil
}
