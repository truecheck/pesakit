package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/pesakit/pesakit/internal/io"
	stdio "io"
	"net/http"
	"net/http/httputil"
	"strings"
)

var (
	_ Receiver = (*receiver)(nil)
)

type receiver struct {
	Logger stdio.Writer
	DebugMode bool
}

func NewReceiver(writer stdio.Writer, debug bool)Receiver {
	return &receiver{
		Logger:    writer,
		DebugMode: debug,
	}
}

func (rc *receiver) Receive(ctx context.Context,rn string, r *http.Request, v interface{}) (*Receipt, error) {
	receipt := new(Receipt)
	rClone := r.Clone(ctx)
	receipt.Request = rClone
	contentType := r.Header.Get("Content-Type")
	payloadType := categorizeContentType(contentType)
	body, err := stdio.ReadAll(r.Body)
	if err != nil {
		return nil,err
	}
	if v == nil {
		return nil,fmt.Errorf("v can not be nil")
	}
	// restore request body
	r.Body = stdio.NopCloser(bytes.NewBuffer(body))

	defer func(debug bool) {
		if debug{
			rc.logRequest(rn,r)
		}
	}(rc.DebugMode)

	switch payloadType {
	case JsonPayload:
		err := json.NewDecoder(r.Body).Decode(v)
		defer func(Body stdio.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(r.Body)
		r.Body = stdio.NopCloser(bytes.NewBuffer(body))
		return receipt,err

	case XmlPayload:
		r.Body = stdio.NopCloser(bytes.NewBuffer(body))
		return receipt,xml.Unmarshal(body, v)
	}

	return receipt,err
}

type Receiver interface {
	Receive(ctx context.Context, rn string,r *http.Request, v interface{})(*Receipt,error)
	LogPayload(prefix string,response *Response)
}

type Receipt struct {
	Request *http.Request
}

type ReceiveParams struct {
	DebugMode bool
	Logger stdio.Writer
}

type ReceiveOption func(params *ReceiveParams)

func ReceiveDebugMode(mode bool) ReceiveOption{
	return func(params *ReceiveParams) {
		params.DebugMode = mode
	}
}

func ReceiveLogger(writer stdio.Writer) ReceiveOption{
	return func(params *ReceiveParams) {
		params.Logger = writer
	}
}

// ReceivePayloadWithParams takes *http.Request from clients like during then unmarshal the provided
// request into given interface v
// The expected Content-Type should also be declared. If its cTypeJson or
// application/xml
func ReceivePayloadWithParams(r *http.Request, v interface{},opts...ReceiveOption) error {

	rp := &ReceiveParams{
		DebugMode: true,
		Logger:    io.Stderr,
	}

	for _, opt := range opts {
		opt(rp)
	}
	contentType := r.Header.Get("Content-Type")
	payloadType := categorizeContentType(contentType)
	body, err := stdio.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if v == nil {
		return fmt.Errorf("v can not be nil")
	}
	// restore request body
	r.Body = stdio.NopCloser(bytes.NewBuffer(body))

	switch payloadType {
	case JsonPayload:
		err := json.NewDecoder(r.Body).Decode(v)
		defer func(Body stdio.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(r.Body)
		r.Body = stdio.NopCloser(bytes.NewBuffer(body))
		return err

	case XmlPayload:
		r.Body = stdio.NopCloser(bytes.NewBuffer(body))
		return xml.Unmarshal(body, v)
	}

	return err
}

// ReceivePayload takes *http.Request from clients like during then unmarshal the provided
// request into given interface v
// The expected Content-Type should also be declared. If its cTypeJson or
// application/xml
func ReceivePayload(r *http.Request, v interface{}) error {

	contentType := r.Header.Get("Content-Type")
	payloadType := categorizeContentType(contentType)
	body, err := stdio.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if v == nil {
		return fmt.Errorf("v can not be nil")
	}
	// restore request body
	r.Body = stdio.NopCloser(bytes.NewBuffer(body))

	switch payloadType {
	case JsonPayload:
		err := json.NewDecoder(r.Body).Decode(v)
		defer func(Body stdio.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(r.Body)
		r.Body = stdio.NopCloser(bytes.NewBuffer(body))
		return err

	case XmlPayload:
		r.Body = stdio.NopCloser(bytes.NewBuffer(body))
		return xml.Unmarshal(body, v)
	}

	return err
}

// logRequest is called to print the details of http.Request received
func (rc *receiver) logRequest(name string, request *http.Request) {

	if request != nil && rc.DebugMode {
		reqDump, _ := httputil.DumpRequest(request, true)
		_, err := fmt.Fprintf(rc.Logger, "%s REQUEST: %s\n", name, reqDump)
		if err != nil {
			fmt.Printf("error while logging %s request: %v\n",
				strings.ToLower(name), err)
			return
		}
		return
	}
	return
}

func (rc *receiver) LogPayload(prefix string,response *Response) {
	contentType := response.Headers["Content-Type"]
	payloadType := categorizeContentType(contentType)
	buffer, _ := MarshalPayload(payloadType,response.Payload)
	_, _ = rc.Logger.Write([]byte(fmt.Sprintf("%s response: %s\n\n", prefix, buffer.String())))
}
