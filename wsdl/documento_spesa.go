package wsdl

import "encoding/xml"

type InserimentoDocumentoSpesaRequest struct {
	XMLName                       xml.Name                      `xml:"inserimentoDocumentoSpesaRequest"`
	XMLNS                         string                        `xml:"xmlns,attr"`
	Pincode                       *string                       `xml:"pincode"`
	Proprietario                  Proprietario                  `xml:"Proprietario"`
	IdInserimentoDocumentoFiscale IdInserimentoDocumentoFiscale `xml:"idInserimentoDocumentoFiscale"`
}

type Proprietario struct {
	CfProprietario *string `xml:"cfProprietario"`
}

type IdInserimentoDocumentoFiscale struct {
	IdSpesa            IdSpesa     `xml:"idSpesa"`
	DataPagamento      *string     `xml:"dataPagamento"`
	CfCittadino        *string     `xml:"cfCittadino"`
	VoceSpesa          []VoceSpesa `xml:"voceSpesa"`
	PagamentoTracciato *string     `xml:"pagamentoTracciato"`
	TipoDocumento      *string     `xml:"tipoDocumento"`
	FlagOpposizione    int8        `xml:"flagOpposizione"`
}

type IdSpesa struct {
	PIva                *string             `xml:"pIva"`
	DataEmissione       *string             `xml:"dataEmissione"`
	NumDocumentoFiscale NumDocumentoFiscale `xml:"numDocumentoFiscale"`
}

type NumDocumentoFiscale struct {
	Dispositivo  *string `xml:"dispositivo"`
	NumDocumento *string `xml:"numDocumento"`
}

type VoceSpesa struct {
	TipoSpesa *string `xml:"tipoSpesa"`
	Importo   float32 `xml:"importo"`
	NaturaIVA *string `xml:"naturaIVA"`
}
