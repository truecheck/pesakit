package airtel

type TokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type PushRequest struct {
	Reference  string `json:"reference"`
	Subscriber struct {
		Country  string `json:"country"`
		Currency string `json:"currency"`
		Msisdn   string `json:"msisdn"`
	} `json:"subscriber"`
	Transaction struct {
		Amount   int64  `json:"amount"`
		Country  string `json:"country"`
		Currency string `json:"currency"`
		ID       string `json:"id"`
	} `json:"transaction"`
}

type RefundRequest struct {
	CountryOfTransaction string `json:"-"`
	Transaction          struct {
		AirtelMoneyID string `json:"airtel_money_id"`
	} `json:"transaction"`
}

type PushEnquiryRequest struct {
	ID                   string `json:"id"`
	CountryOfTransaction string `json:"country"`
}

type CallbackRequest struct {
	Transaction struct {
		ID            string `json:"id"`
		Message       string `json:"message"`
		StatusCode    string `json:"status_code"`
		AirtelMoneyID string `json:"airtel_money_id"`
	} `json:"transaction"`
	Hash string `json:"hash"`
}

type UserEnquiryRequest struct {
	MSISDN               string
	CountryOfTransaction string
}

type BalanceEnquiryRequest struct {
	MSISDN               string
	CountryOfTransaction string
}

type DisburseEnquiryRequest struct {
	CountryOfTransaction string
	ID                   string `json:"id"`
}

type iDisburseRequest struct {
	CountryOfTransaction string `json:"-"`
	Payee                struct {
		Msisdn string `json:"msisdn"`
	} `json:"payee"`
	Reference   string `json:"reference"`
	Pin         string `json:"pin"`
	Transaction struct {
		Amount int64  `json:"amount"`
		ID     string `json:"id"`
	} `json:"transaction"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	Message     string `json:"message,omitempty"`
}

type iPushResponse struct {
	Data struct {
		Transaction struct {
			ID     string `json:"id,omitempty"`
			Status string `json:"status,omitempty"`
		} `json:"transaction,omitempty"`
	} `json:"data,omitempty"`
	Status struct {
		Code       string `json:"code,omitempty"`
		Message    string `json:"message,omitempty"`
		ResultCode string `json:"result_code,omitempty"`
		Success    bool   `json:"success,omitempty"`
	} `json:"status,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	Error            string `json:"error,omitempty"`
	StatusMessage    string `json:"status_message,omitempty"`
	StatusCode       string `json:"status_code,omitempty"`
}

type RefundResponse struct {
	Data struct {
		Transaction struct {
			AirtelMoneyID string `json:"airtel_money_id,omitempty"`
			Status        string `json:"status,omitempty"`
		} `json:"transaction,omitempty"`
	} `json:"data,omitempty"`
	Status struct {
		Code       string `json:"code,omitempty"`
		Message    string `json:"message,omitempty"`
		ResultCode string `json:"result_code,omitempty"`
		Success    bool   `json:"success,omitempty"`
	} `json:"status,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	Error            string `json:"error,omitempty"`
	StatusMessage    string `json:"status_message,omitempty"`
	StatusCode       string `json:"status_code,omitempty"`
}

type PushEnquiryResponse struct {
	Data struct {
		Transaction struct {
			AirtelMoneyID string `json:"airtel_money_id,omitempty"`
			ID            string `json:"id,omitempty"`
			Message       string `json:"message,omitempty"`
			Status        string `json:"status,omitempty"`
		} `json:"transaction,omitempty"`
	} `json:"data,omitempty"`
	Status struct {
		Code       string `json:"code,omitempty"`
		Message    string `json:"message,omitempty"`
		ResultCode string `json:"result_code,omitempty"`
		Success    bool   `json:"success,omitempty"`
	} `json:"status,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	Error            string `json:"error,omitempty"`
	StatusMessage    string `json:"status_message,omitempty"`
	StatusCode       string `json:"status_code,omitempty"`
}

type UserEnquiryResponse struct {
	Data struct {
		FirstName    string `json:"first_name,omitempty"`
		Grade        string `json:"grade,omitempty"`
		IsBarred     bool   `json:"is_barred,omitempty"`
		IsPinSet     bool   `json:"is_pin_set,omitempty"`
		LastName     string `json:"last_name,omitempty"`
		Msisdn       int    `json:"msisdn,omitempty"`
		RegStatus    string `json:"reg_status,omitempty"`
		RegisteredID string `json:"registered_id,omitempty"`
		Registration struct {
			ID     string `json:"id,omitempty"`
			Status string `json:"status,omitempty"`
		} `json:"registration,omitempty"`
	} `json:"data,omitempty"`
	Status struct {
		Code       string `json:"code,omitempty"`
		Message    string `json:"message,omitempty"`
		ResultCode string `json:"result_code,omitempty"`
		Success    bool   `json:"success,omitempty"`
	} `json:"status,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	Error            string `json:"error,omitempty"`
	StatusMessage    string `json:"status_message,omitempty"`
	StatusCode       string `json:"status_code,omitempty"`
}

type BalanceEnquiryResponse struct {
	Data struct {
		Balance       string `json:"balance"`
		Currency      string `json:"currency"`
		AccountStatus string `json:"account_status"`
	} `json:"data"`
	Status struct {
		Code       string `json:"code"`
		Message    string `json:"message"`
		ResultCode string `json:"result_code"`
		Success    bool   `json:"success"`
	} `json:"status"`
	ErrorDescription string `json:"error_description,omitempty"`
	Error            string `json:"error,omitempty"`
	StatusMessage    string `json:"status_message,omitempty"`
	StatusCode       string `json:"status_code,omitempty"`
}

type DisburseEnquiryResponse struct {
	Data struct {
		Transaction struct {
			ID      string `json:"id"`
			Message string `json:"message"`
			Status  string `json:"status"`
		} `json:"transaction"`
	} `json:"data"`
	Status struct {
		Code       string `json:"code"`
		Message    string `json:"message"`
		ResultCode string `json:"result_code"`
		Success    bool   `json:"success"`
	} `json:"status"`
	ErrorDescription string `json:"error_description,omitempty"`
	Error            string `json:"error,omitempty"`
	StatusMessage    string `json:"status_message,omitempty"`
	StatusCode       string `json:"status_code,omitempty"`
}

type ListTransactionsResponse struct {
	Data struct {
		ErrorDescription string `json:"error_description,omitempty"`
		Error            string `json:"error,omitempty"`
		StatusMessage    string `json:"status_message,omitempty"`
		StatusCode       string `json:"status_code,omitempty"`
		Count            int    `json:"count"`
		Transactions     []struct {
			Charges struct {
				Service int `json:"service"`
			} `json:"charges"`
			Payee struct {
				Currency string `json:"currency"`
				Msisdn   int    `json:"msisdn"`
				Name     string `json:"name"`
			} `json:"payee"`
			Payer struct {
				Currency string `json:"currency"`
				Msisdn   int    `json:"msisdn"`
				Name     string `json:"name"`
			} `json:"payer"`
			Service struct {
				Type string `json:"type"`
			} `json:"service"`
			Transaction struct {
				AirtelMoneyID   string `json:"airtel_money_id"`
				Amount          int    `json:"amount"`
				CreatedAt       int    `json:"created_at"`
				ID              int64  `json:"id"`
				ReferenceNumber string `json:"reference_number"`
				Status          string `json:"status"`
			} `json:"transaction"`
		} `json:"transactions"`
	} `json:"data"`
	Status struct {
		Code       int    `json:"code"`
		Message    string `json:"message"`
		ResultCode string `json:"result_code"`
		Success    bool   `json:"success"`
	} `json:"status"`
}

type iDisburseResponse struct {
	Data struct {
		Transaction struct {
			ReferenceID   string `json:"reference_id,omitempty"`
			AirtelMoneyID string `json:"airtel_money_id,omitempty"`
			ID            string `json:"id,omitempty"`
		} `json:"transaction,omitempty"`
	} `json:"data,omitempty"`
	Status struct {
		Code       string `json:"code,omitempty"`
		Message    string `json:"message,omitempty"`
		ResultCode string `json:"result_code,omitempty"`
		Success    bool   `json:"success,omitempty"`
	} `json:"status,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	Error            string `json:"error,omitempty"`
	StatusMessage    string `json:"status_message,omitempty"`
	StatusCode       string `json:"status_code,omitempty"`
}
