package main

import (
	"log"

	"github.com/ghiblin/go-tessera-sanitaria/service"
	"github.com/ghiblin/go-tessera-sanitaria/util"
	"github.com/ghiblin/go-tessera-sanitaria/wsdl"
)

var (
	cfProprietario = "SsFrZY1plknIYKxk2MxIsgCyH2X3cfnrbg7B1aywMzw4SYwfzCa797Bb40vZMlS1pRjBki3SYZT/dao7W7SCwarTTLQqFmfXu7SGBStGzfAyVWcXAZapnW3d8QWfY7EgbktdHPfcoslCqY+K4JJrHQA9H2bk2ngSA7n+xOjuLVw="
	// pincode        = "W+cy4Lz7rOOgldsbK98TEAwR6OIikKMkQJ1nWS09LgnJ6up+4e2LfIHe9z6aOQ9NocHOqepHuUcqmNNXpOq2JDCZQms65cX2oif8VhSUsvOk/9mc/8A9A7tpLnHcoGNrrjrg0z3fat7JmENXo5LF5uQV2IqvT4z5BDJbNa5XZpQ="
	pincode      = "3489543096"
	pIva         = "65498732105"
	data         = "2023-02-20"
	dispositivo  = "1"
	numDocumento = "100"
	// cfCittadino   = "iKvd9JQntqxPBT2UA/OFfztSNLidocP8Op+NfODzfTdxFWzkcdZrJz5gvCuqv7Dh/r3Cin1ZQMmg/BofIqYCyq2PcC+PJzbvQCocDdl6FrXVXs3W5JhnX7VpWFGCLPYYY2WL+RWKxhfkGqeY8+NCVfQ1lEA15g3W5AabJ15Tthk="
	cfCittadino   = "VTLLSN79H29A794O"
	tipoSpesa     = "SP"
	naturaIVA     = "N2.2"
	si            = "SI"
	tipoDocumento = "F"
)

func main() {
	c, err := util.NewCipher()
	if err != nil {
		log.Fatalf("Failed to initiate cipher: %s", err)
	}

	encodedPincode, err := c.Encrypt(pincode)
	if err != nil {
		log.Fatalf("Failed to encrypt pincode: %s", err)
	}
	encodedCF, err := c.Encrypt(cfCittadino)
	if err != nil {
		log.Fatalf("Failed to encrypt pincode: %s", err)
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
				PIva:          &pIva,
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
				}},
			PagamentoTracciato: &si,
			TipoDocumento:      &tipoDocumento,
			FlagOpposizione:    0,
		},
	})
	log.Printf("Resp: %s\n", resp)
}
