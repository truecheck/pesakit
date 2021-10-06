package mpesa

type (
	Request struct {
		ThirdPartyID string  `json:"id,omitempty"`
		Reference    string  `json:"reference,omitempty"`
		Amount       float64 `json:"amount,omitempty"`
		MSISDN       string  `json:"msisdn,omitempty"`
		Desc         string  `json:"description,omitempty"`
	}

	SessionResponse struct {
		Code      string `json:"output_ResponseCode,omitempty"`
		Desc      string `json:"output_ResponseDesc,omitempty"`
		ID        string `json:"output_SessionID,omitempty"`
		OutputErr string `json:"output_error,omitempty"`
	}

	//pushPayRequest
	//  Amount	The transaction amount. This amount will be moved from the organization's account to the customer's account.	True	^\d*\.?\d+$	10.00
	//  CustomerMSISDN	The MSISDN of the customer where funds will be debitted from.	True	^[0-9]{12,14}$	254707161122
	//  Country	The country of the mobile money platform where the transaction needs happen on.	True	N/A	GHA
	//  Currency	The currency in which the transaction should take place.	True	^[a-zA-Z]{1,3}$	GHS
	//  ServiceProviderCode	The shortcode of the organization where funds will be creditted to.	True	^([0-9A-Za-z]{4,12})$	ORG001
	//  TransactionReference	The customer's transaction reference	True	^[0-9a-zA-Z \w+]{1,20}$	T12344C
	//  ThirdPartyConversationID	The third party's transaction reference on their system.	True	^[0-9a-zA-Z \w+]{1,40}$	1e9b774d1da34af78412a498cbc28f5e
	//  PurchasedItemsDesc	Description of purchased items	True	^[0-9a-zA-Z \w+]{1,256}$	Handbag, Black shoes
	pushPayRequest struct {
		Amount                   string `json:"input_Amount"`
		Country                  string `json:"input_Country"`
		Currency                 string `json:"input_Currency"`
		CustomerMSISDN           string `json:"input_CustomerMSISDN"`
		ServiceProviderCode      string `json:"input_ServiceProviderCode"`
		ThirdPartyConversationID string `json:"input_ThirdPartyConversationID"`
		TransactionReference     string `json:"input_TransactionReference"`
		PurchasedItemsDesc       string `json:"input_PurchasedItemsDesc"`
	}

	PushAsyncResponse struct {
		ResponseCode             string `json:"output_ResponseCode"`
		ResponseDesc             string `json:"output_ResponseDesc"`
		ConversationID           string `json:"output_ConversationID"`
		ThirdPartyConversationID string `json:"output_ThirdPartyConversationID"`
		OutputErr                string `json:"output_error,omitempty"`
	}

	PushCallbackRequest struct {
		OriginalConversationID   string `json:"input_OriginalConversationID"`
		TransactionID            string `json:"input_TransactionID"`
		ResultCode               string `json:"input_ResultCode"`
		ResultDesc               string `json:"input_ResultDesc"`
		ThirdPartyConversationID string `json:"input_ThirdPartyConversationID"`
	}

	PushCallbackResponse struct {
		OriginalConversationID   string `json:"output_OriginalConversationID"`
		ResponseCode             string `json:"output_ResponseCode"`
		ResponseDesc             string `json:"output_ResponseDesc"`
		ThirdPartyConversationID string `json:"output_ThirdPartyConversationID"`
	}

	// disburseRequest
	//  Amount	The transaction amount. This amount will be moved from the organization's account to the customer's account.	True	^\d*\.?\d+$	10.00
	//  CustomerMSISDN	The MSISDN of the customer where funds will be credited to.	True	^[0-9]{12,14}$	254707161122
	//  Country	The country of the mobile money platform where the transaction needs happen on.	True	N/A	GHA
	//  Currency	The currency in which the transaction should take place.	True	^[a-zA-Z]{1,3}$	GHS
	//  ServiceProviderCode	The shortcode of the organization where funds will be creditted to.	True	^([0-9A-Za-z]{4,12})$	ORG001
	//  TransactionReference	The customer's transaction reference	True	^[0-9a-zA-Z \w+]{1,20}$	T12344C
	//  ThirdPartyConversationID	The third party's transaction reference on their system.	True	^[0-9a-zA-Z \w+]{1,40}$	1e9b774d1da34af78412a498cbc28f5e
	//  PaymentItemsDesc	Description of payment items	True	^[0-9a-zA-Z \w+]{1,256}$	Salary payment
	disburseRequest struct {
		Amount                   string `json:"input_Amount"`
		Country                  string `json:"input_Country"`
		Currency                 string `json:"input_Currency"`
		CustomerMSISDN           string `json:"input_CustomerMSISDN"`
		ServiceProviderCode      string `json:"input_ServiceProviderCode"`
		ThirdPartyConversationID string `json:"input_ThirdPartyConversationID"`
		TransactionReference     string `json:"input_TransactionReference"`
		PaymentItemsDesc         string `json:"input_PaymentItemsDesc"`
	}

	// DisburseResponse ...
	// ResponseCode	The result code for the transaction.	INS-0
	// ResponseDesc	The result description for the transaction.	Request processed successfully
	// TransactionID	The transaction identifier that gets generated on the Mobile Money platform. This is used to query transactions on the Mobile Money Platform.	hv9ahxcg4ccv
	// ConversationID	The OpenAPI platform generates this as a reference to the transaction.	fd1e9143d22544459f7c66e1860ef276
	// ThirdPartyConversationID	The incoming reference from the third party system. When there are queries about transactions, this will usually be used to track a transaction.	1e9b774d1da34af78412a498cbc28f5e
	DisburseResponse struct {
		ConversationID           string `json:"output_ConversationID"`
		ResponseCode             string `json:"output_ResponseCode"`
		ResponseDesc             string `json:"output_ResponseDesc"`
		TransactionID            string `json:"output_TransactionID"`
		ThirdPartyConversationID string `json:"output_ThirdPartyConversationID"`
		OutputErr                string `json:"output_error,omitempty"`
	}
)
