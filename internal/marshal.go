package internal

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"
)

var (
	//ErrInvalidFormPayload is returned when the PayloadType passed is
	// FormPayload but its not of type url.Values
	ErrInvalidFormPayload = errors.New("invalid form submitted: type url.Values is expected")
)

// MarshalPayload returns the JSON/XML encoding of payload.
func MarshalPayload(payloadType PayloadType, payload interface{}) (buffer *bytes.Buffer, err error) {

	switch payloadType {
	case JsonPayload:
		buf, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		buffer = bytes.NewBuffer(buf)

		return buffer, nil
	case XmlPayload, TextXml:
		buf, err := xml.MarshalIndent(payload, "", "  ")
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(buf)
		return buffer, nil

	case FormPayload:

		form, ok := payload.(url.Values)
		if !ok {
			return nil, ErrInvalidFormPayload
		}

		buffer = bytes.NewBufferString(form.Encode())

		return buffer, nil

	default:
		err := fmt.Errorf("can not marshal the payload: invalid payload type")
		return nil, err
	}

}
