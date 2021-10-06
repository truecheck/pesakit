//

package airtel

import (
	"context"
	"fmt"
	"github.com/techcraftlabs/pesakit/internal"
	"github.com/techcraftlabs/pesakit/pkg/countries"
)

type DisbursementService interface {
	disburse(ctx context.Context, request iDisburseRequest) (iDisburseResponse, error)
	Disburse(ctx context.Context, request DisburseRequest) (DisburseResponse, error)
	DisburseEnquiry(ctx context.Context, enquiryRequest DisburseEnquiryRequest) (DisburseEnquiryResponse, error)
}

func (c *Client) Disburse(ctx context.Context, request DisburseRequest) (DisburseResponse, error) {
	disburseRequest, err := c.reqAdapter.ToDisburseRequest(request)
	if err != nil {
		return DisburseResponse{}, err
	}

	disburseResponse, err := c.disburse(ctx, disburseRequest)
	if err != nil {
		return DisburseResponse{}, err
	}
	response := c.resAdapter.ToDisburseResponse(disburseResponse)

	return response, nil
}

func (c *Client) disburse(ctx context.Context, request iDisburseRequest) (iDisburseResponse, error) {
	token, err := c.checkToken(ctx)
	if err != nil {
		return iDisburseResponse{}, err
	}

	countryName := request.CountryOfTransaction
	country, err := countries.GetByName(countryName)
	if err != nil {
		return iDisburseResponse{}, err
	}
	var opts []internal.RequestOption

	hs := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "*/*",
		"X-Country":     country.CodeName,
		"X-Currency":    country.CurrencyCode,
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	headersOpt := internal.WithRequestHeaders(hs)
	opts = append(opts, headersOpt)

	req := c.makeInternalRequest(Disbursement, request, opts...)
	res := new(iDisburseResponse)
	rn := Disbursement.String()
	_, err = c.base.Do(ctx, rn, req, res)
	if err != nil {
		return iDisburseResponse{}, err
	}
	return *res, nil
}

func (c *Client) DisburseEnquiry(ctx context.Context, request DisburseEnquiryRequest) (DisburseEnquiryResponse, error) {
	token, err := c.checkToken(ctx)
	if err != nil {
		return DisburseEnquiryResponse{}, err
	}

	countryName := request.CountryOfTransaction
	country, err := countries.GetByName(countryName)
	if err != nil {
		return DisburseEnquiryResponse{}, err
	}
	var opts []internal.RequestOption

	hs := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "*/*",
		"X-Country":     country.CodeName,
		"X-Currency":    country.CurrencyCode,
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	headersOpt := internal.WithRequestHeaders(hs)
	endpointOption := internal.WithEndpoint(request.ID)
	opts = append(opts, headersOpt, endpointOption)
	req := c.makeInternalRequest(DisbursementEnquiry, request, opts...)
	res := new(DisburseEnquiryResponse)
	_, err = c.base.Do(ctx, "disbursement enquiry", req, res)
	if err != nil {
		return DisburseEnquiryResponse{}, err
	}
	return *res, nil
}
