package wsdl

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type InserimentoDocumentoSpesaRequest struct {
	XMLName                       xml.Name             `xml:"doc:inserimentoDocumentoSpesaRequest"`
	XMLNS                         string               `xml:"xmlns:doc,attr"`
	Pincode                       string               `xml:"doc:pincode"`
	Proprietario                  string               `xml:"doc:Proprietario>doc:cfProprietario"`
	IdInserimentoDocumentoFiscale *InsDocumentoFiscale `xml:"doc:idInserimentoDocumentoFiscale"`
}
type InserimentoDocumentoSpesaResponse struct {
	XMLName  xml.Name `xml:"inserimentoDocumentoSpesaResponse"`
	XMLNS    string   `xml:"xmlns,attr"`
	Esito    int16    `xml:"esitoChiamata"`
	Messaggi []struct {
		Codice      string `xml:"messaggio>codice"`
		Descrizione string `xml:"messaggio>descrizione"`
		Tipo        string `xml:"messaggio>tipo"`
	} `xml:"listaMessaggi"`
	Fault struct {
		Code    string `xml:"Fault>faultcode"`
		Message string `xml:"Fault>faultstring"`
	}
}

type InsDocumentoFiscale struct {
	IdSpesa            *IdSpesa     `xml:"doc:idSpesa"`
	DataPagamento      string       `xml:"doc:dataPagamento"`
	CfCittadino        string       `xml:"doc:cfCittadino"`
	VoceSpesa          []*VoceSpesa `xml:"doc:voceSpesa"`
	PagamentoTracciato string       `xml:"doc:pagamentoTracciato"`
	TipoDocumento      string       `xml:"doc:tipoDocumento"`
	FlagOpposizione    int8         `xml:"doc:flagOpposizione"`
}

type IdSpesa struct {
	PIva                string               `xml:"doc:pIva"`
	DataEmissione       string               `xml:"doc:dataEmissione"`
	NumDocumentoFiscale *NumDocumentoFiscale `xml:"doc:numDocumentoFiscale"`
}

type VoceSpesa struct {
	TipoSpesa string  `xml:"doc:tipoSpesa"`
	Importo   float32 `xml:"doc:importo"`
	NaturaIVA string  `xml:"doc:naturaIVA"`
}

type NumDocumentoFiscale struct {
	Dispositivo  string `xml:"doc:dispositivo"`
	NumDocumento string `xml:"doc:numDocumento"`
}

func (r *InserimentoDocumentoSpesaResponse) String() string {
	var messages []string
	for _, m := range r.Messaggi {
		messages = append(messages, fmt.Sprintf("(%s) %s", m.Tipo, m.Descrizione))
	}
	return fmt.Sprintf("%d: %s", r.Esito, strings.Join(messages, ", "))
}
