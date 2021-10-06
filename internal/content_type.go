package internal

import "strings"

const (
	JsonPayload PayloadType = iota
	XmlPayload
	FormPayload
	TextXml
	UnsupportedPayload
)

const (
	cTypeJson        = "application/json"
	cTypeTextXml     = "text/xml"
	cTypeAppXml      = "application/xml"
	cTypeForm        = "application/x-www-form-urlencoded"
	cTypeUnsupported = "unsupported"
)

type (
	PayloadType int
)

func (p PayloadType) String() string {
	types := []string{
		cTypeJson,
		cTypeAppXml,
		cTypeForm,
		cTypeTextXml,
		cTypeUnsupported,
	}

	return types[p]
}

func categorizeContentType(headerStr string) PayloadType {
	j := JsonPayload.String()
	x := XmlPayload.String()
	xml2 := TextXml.String()
	form := FormPayload.String()
	if strings.Contains(headerStr, j) {
		return JsonPayload
	} else if strings.Contains(headerStr, xml2) || strings.Contains(headerStr, x) {
		return XmlPayload
	} else if strings.Contains(headerStr, form) {
		return FormPayload
	} else {
		//todo: figure out proper way to return this
		return UnsupportedPayload
	}
}
