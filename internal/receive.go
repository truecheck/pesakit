package internal

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

//
//type Receipt struct {
//	Method       string
//	OriginalIP   string
//	ForwardingIP string
//	Headers      map[string]string
//	Params       map[string]string
//	Payload      interface{}
//}

// ReceivePayload takes *http.Request from clients like during then unmarshal the provided
// request into given interface v
// The expected Content-Type should also be declared. If its cTypeJson or
// application/xml
func ReceivePayload(r *http.Request, v interface{}) error {

	contentType := r.Header.Get("Content-Type")
	payloadType := categorizeContentType(contentType)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if v == nil {
		return fmt.Errorf("v can not be nil")
	}
	// restore request body
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	switch payloadType {
	case JsonPayload:
		err := json.NewDecoder(r.Body).Decode(v)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		return err

	case XmlPayload:
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		return xml.Unmarshal(body, v)
	}

	return err
}
