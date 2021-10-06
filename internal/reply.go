package internal

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func Reply(r *Response, writer http.ResponseWriter) {
	if r.Payload == nil {

		for key, value := range r.Headers {
			writer.Header().Add(key, value)
		}
		writer.WriteHeader(r.StatusCode)

		return
	}

	pType := categorizeContentType(r.Headers["Content-Type"])

	switch pType {
	case XmlPayload:
		payload, err := xml.MarshalIndent(r.Payload, "", "  ")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		for key, value := range r.Headers {
			writer.Header().Set(key, value)
		}

		writer.WriteHeader(r.StatusCode)
		writer.Header().Set("Content-Type", cTypeAppXml)
		_, err = writer.Write(payload)
		return

	case JsonPayload:
		payload, err := json.Marshal(r.Payload)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		for key, value := range r.Headers {
			writer.Header().Set(key, value)
		}
		writer.Header().Set("Content-Type", cTypeJson)
		writer.WriteHeader(r.StatusCode)
		_, err = writer.Write(payload)
		return
	}
}
