package wsdl

import "encoding/xml"

type Header struct {
	XMLName xml.Name `xml:"Header"`
}

type Body[T any] struct {
	Content T
	Fault   *Fault `xml:"Fault"`
}

type Fault struct {
	Code    string `xml:"faultcode"`
	Message string `xml:"faultstring"`
}

type Envelope[T any] struct {
	XMLName xml.Name `xml:"Envelope"`
	Soapenv string   `xml:"xmlns,attr"`
	Header  *Header
	Body    *Body[T] `xml:"Body"`
}

type Bool struct {
	bool
}

func (b *Bool) MarshalCSV(e *xml.Encoder, start xml.StartElement) error {
	var v int8
	if b.bool {
		v = 1
	} else {
		v = 0
	}
	return e.EncodeElement(v, start)
}
