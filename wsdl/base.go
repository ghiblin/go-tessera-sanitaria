package wsdl

import "encoding/xml"

type Header struct {
	XMLName xml.Name `xml:"soapenv:Header"`
}

type Body struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Request interface{}
}

type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Header  *Header
	Body    *Body
}
