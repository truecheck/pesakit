//

package airtel

import (
	"fmt"
	"github.com/pesakit/pesakit/pkg/countries"
	"github.com/pesakit/pesakit/pkg/crypto"
	"math"
)

var (
	_ ResponseAdapter = (*adapter)(nil)
	_ RequestAdapter  = (*adapter)(nil)
)

type (
	ResponseAdapter interface {
		ToDisburseResponse(response iDisburseResponse) DisburseResponse
		ToPushPayResponse(response iPushResponse) PushPayResponse
	}

	// adapter acts as a default RequestAdapter and ResponseAdapter. This can be replaced
	// by other user defined adapters. and then injected to the client using the
	// Client.SetRequestAdapter or Client.SetResponseAdapter, it is not recommended as of now
	// 26th September 2021
	adapter struct {
		Conf *Config
	}
	RequestAdapter interface {
		ToPushPayRequest(request PushPayRequest) PushRequest
		ToDisburseRequest(request DisburseRequest) (iDisburseRequest, error)
	}
)

func (r *adapter) ToPushPayRequest(request PushPayRequest) PushRequest {

	amount := math.Floor(request.TransactionAmount*100/100)

	subCountry, _ := countries.GetByName(request.SubscriberCountry)
	transCountry, _ := countries.GetByName(request.TransactionCountry)
	return PushRequest{
		Reference: request.Reference,
		Subscriber: struct {
			Country  string `json:"country"`
			Currency string `json:"currency"`
			Msisdn   string `json:"msisdn"`
		}{
			Country:  subCountry.CodeName,
			Currency: subCountry.CurrencyCode,
			Msisdn:   request.SubscriberMsisdn,
		},
		Transaction: struct {
			Amount   float64 `json:"amount"`
			Country  string `json:"country"`
			Currency string `json:"currency"`
			ID       string `json:"id"`
		}{
			Amount:   amount,
			Country:  transCountry.CodeName,
			Currency: transCountry.CurrencyCode,
			ID:       request.TransactionID,
		},
	}
}

func (r *adapter) ToDisburseRequest(request DisburseRequest) (iDisburseRequest, error) {
	amount := math.Floor(request.Amount*100/100)
	pin := r.Conf.DisbursePIN
	pubKey := r.Conf.PublicKey
	encryptedPin, err := crypto.EncryptKey(pin, pubKey)
	if err != nil {
		return iDisburseRequest{}, fmt.Errorf("could not encrypt key: %w", err)
	}
	req := iDisburseRequest{
		CountryOfTransaction: request.CountryOfTransaction,
		Payee: struct {
			Msisdn string `json:"msisdn"`
		}{
			Msisdn: request.MSISDN,
		},
		Reference: request.Reference,
		Pin:       encryptedPin,
		Transaction: struct {
			Amount float64  `json:"amount"`
			ID     string `json:"id"`
		}{
			Amount: amount,
			ID:     request.ID,
		},
	}
	return req, nil
}

func (r *adapter) ToPushPayResponse(response iPushResponse) PushPayResponse {
	transaction := response.Data.Transaction
	status := response.Status

	if !status.Success {
		return PushPayResponse{
			ResultCode:    status.ResultCode,
			Success:       status.Success,
			StatusMessage: status.Message,
			StatusCode:    status.Code,
		}
	}
	isErr := response.Error != "" && response.ErrorDescription != ""
	if isErr {
		resp := PushPayResponse{
			ErrorDescription: response.ErrorDescription,
			Error:            response.Error,
			StatusMessage:    response.StatusMessage,
			StatusCode:       response.StatusCode,
		}
		return resp
	}

	return PushPayResponse{
		ID:               transaction.ID,
		Status:           transaction.Status,
		ResultCode:       status.ResultCode,
		Success:          status.Success,
		ErrorDescription: response.ErrorDescription,
		Error:            response.Error,
		StatusMessage:    response.StatusMessage,
		StatusCode:       response.StatusCode,
	}
}

func (r *adapter) ToDisburseResponse(response iDisburseResponse) DisburseResponse {

	isErr := response.Error != "" && response.ErrorDescription != ""
	if isErr {
		resp := DisburseResponse{
			ErrorDescription: response.ErrorDescription,
			Error:            response.Error,
			StatusMessage:    response.StatusMessage,
			StatusCode:       response.StatusCode,
		}

		return resp
	}
	transaction := response.Data.Transaction
	status := response.Status

	return DisburseResponse{
		ID:            transaction.ID,
		Reference:     transaction.ReferenceID,
		AirtelMoneyID: transaction.AirtelMoneyID,
		ResultCode:    status.ResultCode,
		Success:       status.Success,
		StatusMessage: status.Message,
		StatusCode:    status.Code,
	}

}
