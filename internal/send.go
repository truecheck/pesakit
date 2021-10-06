package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	stdio "io"
	"net/http"
	"strings"
)

const errStatusCodeMargin = 400

var DoErr = errors.New("result code is above or equal to 400")

func (client *BaseClient) Do(ctx context.Context, rn string, request *Request, body interface{}) (*Response, error) {

	var (
		errDecodingBody  = errors.New("error while decoding response body")
		errUnknownHeader = errors.New("unknown content-type header")
	)

	var (
		_, cancel    = context.WithTimeout(ctx, defaultTimeout)
		req          *http.Request
		res          *http.Response
		reqBodyBytes []byte
		resBodyBytes []byte
	)
	defer cancel()
	defer func(debug bool) {
		if debug {
			req.Body = stdio.NopCloser(bytes.NewBuffer(reqBodyBytes))
			if res == nil {
				client.logOut(strings.ToUpper(rn), req, nil)

				return
			}
			res.Body = stdio.NopCloser(bytes.NewBuffer(resBodyBytes))
			client.logOut(strings.ToUpper(rn), req, res)
		}
	}(client.DebugMode)
	req, err := NewRequestWithContext(ctx, request)

	if err != nil {
		return nil, err
	}

	if req.Body != nil {
		reqBodyBytes, _ = stdio.ReadAll(req.Body)
	}

	req.Body = stdio.NopCloser(bytes.NewBuffer(reqBodyBytes))
	res, doErr := client.Http.Do(req)

	if doErr != nil {
		return nil, doErr
	}

	if res.Body != nil {
		resBodyBytes, _ = stdio.ReadAll(res.Body)
	}

	response := new(Response)
	statusCode := res.StatusCode
	response.StatusCode = statusCode

	contentType := res.Header.Get("Content-Type")
	headers := make(map[string]string)
	for k, v := range res.Header {
		headers[strings.ToLower(k)] = v[0]
	}

	response.Headers = headers
	cType := categorizeContentType(contentType)

	isJSON := cType == JsonPayload
	isXML := cType == XmlPayload || cType == TextXml
	isOK := statusCode < errStatusCodeMargin

	if body != nil {
		if isJSON {
			dErr := json.NewDecoder(bytes.NewBuffer(resBodyBytes)).Decode(body)
			isDecodeErr := dErr != nil && !errors.Is(dErr, stdio.EOF)

			if isDecodeErr {
				return nil, fmt.Errorf("%w: %v", dErr, errDecodingBody)
			}

			response.Payload = body

			if !isOK {
				response.Error = DoErr
				return response, nil
			}

			return response, nil

		} else if isXML {

			dErr := xml.NewDecoder(bytes.NewBuffer(resBodyBytes)).Decode(body)
			isDecodeErr := dErr != nil && !errors.Is(dErr, stdio.EOF)
			if isDecodeErr {
				return nil, fmt.Errorf("%w: %v", dErr, errDecodingBody)
			}

			response.Payload = body
			if !isOK {
				response.Error = DoErr
				return response, nil
			}
			return response, nil

		} else {
			return nil, errUnknownHeader
		}
	}

	if !isOK {
		response.Error = DoErr
	}
	return response, nil

}
