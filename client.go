package main

import (
	"log"

	"github.com/ghiblin/go-tessera-sanitaria/service"
	"github.com/ghiblin/go-tessera-sanitaria/util"
	"github.com/ghiblin/go-tessera-sanitaria/wsdl"
)

var (
	data          = "2023-02-20"
	dispositivo   = "1"
	numDocumento  = "101"
	cfCittadino   = "VTLLSN79H29A794O"
	tipoSpesa     = "SP"
	naturaIVA     = "N2.2"
	si            = "SI"
	tipoDocumento = "F"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatalf("Failed to get config: %s", err)
	}

	c, err := util.NewCipher()
	if err != nil {
		log.Fatalf("Failed to initiate cipher: %s", err)
	}

	cfProprietario, err := c.Encrypt(cfg.User.Username)
	if err != nil {
		log.Fatalf("Failed to encrypt username: %s", err)
	}

	encodedPincode, err := c.Encrypt(cfg.User.Pincode)
	if err != nil {
		log.Fatalf("Failed to encrypt pincode: %s", err)
	}
	encodedCF, err := c.Encrypt(cfCittadino)
	if err != nil {
		log.Fatalf("Failed to encrypt cfCittadino: %s", err)
	}

	s := &service.DocumentoSpesa{}
	resp := s.Inserimento(&wsdl.InserimentoDocumentoSpesaRequest{
		XMLNS:   "http://documentospesap730.sanita.finanze.it",
		Pincode: &encodedPincode,
		Proprietario: wsdl.Proprietario{
			CfProprietario: &cfProprietario,
		},
		IdInserimentoDocumentoFiscale: wsdl.IdInserimentoDocumentoFiscale{
			IdSpesa: wsdl.IdSpesa{
				PIva:          &cfg.User.PIva,
				DataEmissione: &data,
				NumDocumentoFiscale: wsdl.NumDocumentoFiscale{
					Dispositivo:  &dispositivo,
					NumDocumento: &numDocumento,
				},
			},
			DataPagamento: &data,
			CfCittadino:   &encodedCF,
			VoceSpesa: []wsdl.VoceSpesa{
				{
					TipoSpesa: &tipoSpesa,
					Importo:   100.80,
					NaturaIVA: &naturaIVA,
				}, {
					TipoSpesa: &tipoSpesa,
					Importo:   2,
					NaturaIVA: &naturaIVA,
				}},
			PagamentoTracciato: &si,
			TipoDocumento:      &tipoDocumento,
			FlagOpposizione:    0,
		},
	})
	log.Printf("Resp: %s\n", resp)
}
