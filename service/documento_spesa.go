package service

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ghiblin/go-tessera-sanitaria/util"
	"github.com/ghiblin/go-tessera-sanitaria/wsdl"
)

const url = "https://invioSS730pTest.sanita.finanze.it/DocumentoSpesa730pWeb/DocumentoSpesa730pPort"

type DocumentoSpesa struct {
	config *util.Config
	cipher *util.Cipher
}

func NewDocumentoSpesaService() (*DocumentoSpesa, error) {
	cfg, err := util.GetConfig()
	if err != nil {
		return nil, err
	}
	c, err := util.NewCipher()
	if err != nil {
		return nil, err
	}
	return &DocumentoSpesa{
		config: cfg,
		cipher: c,
	}, nil
}

func (s *DocumentoSpesa) Inserimento(invoice *util.Invoice) error {
	request := s.toEnvelop(invoice)
	soap, err := s.marshalRequest(request)
	if err != nil {
		return err
	}

	response, err := s.sendRequest(soap)
	if err != nil {
		return err
	}

	result, err := s.unmarshalResponse(response)
	if err != nil {
		return err
	}

	if result.Body.Content.Esito > 0 {
		return errors.New(result.Body.Content.Messaggi[0].Descrizione)
	}
	return nil
}

func (s *DocumentoSpesa) toEnvelop(invoice *util.Invoice) *wsdl.Envelope[*wsdl.InserimentoDocumentoSpesaRequest] {
	return &wsdl.Envelope[*wsdl.InserimentoDocumentoSpesaRequest]{
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		Header:  &wsdl.Header{},
		Body: &wsdl.Body[*wsdl.InserimentoDocumentoSpesaRequest]{
			Content: &wsdl.InserimentoDocumentoSpesaRequest{
				XMLNS:        "http://documentospesap730.sanita.finanze.it",
				Pincode:      s.cipher.Encrypt(s.config.User.Pincode),
				Proprietario: s.cipher.Encrypt(s.config.User.Username),
				IdInserimentoDocumentoFiscale: &wsdl.InsDocumentoFiscale{
					IdSpesa: &wsdl.IdSpesa{
						PIva:          s.config.User.PIva,
						DataEmissione: invoice.Date.String(),
						NumDocumentoFiscale: &wsdl.NumDocumentoFiscale{
							Dispositivo:  "1",
							NumDocumento: invoice.Id,
						},
					},
					DataPagamento: invoice.Date.String(),
					CfCittadino:   s.cipher.Encrypt(invoice.TaxCode),
					VoceSpesa: []*wsdl.VoceSpesa{
						{
							TipoSpesa: "SP",
							Importo:   float32(invoice.Amount) - float32(invoice.TaxStamp),
							NaturaIVA: "N2.2",
						},
						{
							TipoSpesa: "SP",
							Importo:   float32(invoice.TaxStamp),
							NaturaIVA: "N2.2",
						},
					},
					PagamentoTracciato: "SI",
					TipoDocumento:      "F",
				},
			},
		},
	}
}

func (s *DocumentoSpesa) marshalRequest(env *wsdl.Envelope[*wsdl.InserimentoDocumentoSpesaRequest]) (string, error) {
	soap, err := xml.MarshalIndent(env, "  ", "  ")
	if err != nil {
		return "", err
	}
	return string(soap), nil
}

func (s *DocumentoSpesa) unmarshalResponse(response string) (*wsdl.Envelope[*wsdl.InserimentoDocumentoSpesaResponse], error) {
	env := &wsdl.Envelope[*wsdl.InserimentoDocumentoSpesaResponse]{
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: &wsdl.Body[*wsdl.InserimentoDocumentoSpesaResponse]{
			Content: &wsdl.InserimentoDocumentoSpesaResponse{
				XMLNS: "http://documentospesap730.sanita.finanze.it",
			},
		},
	}
	err := xml.Unmarshal([]byte(response), env)
	if err != nil {
		return nil, err
	}
	return env, nil
}

func (s *DocumentoSpesa) sendRequest(soap string) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBufferString(soap))
	req.Header.Add("Content-Type", "text/xml")
	req.SetBasicAuth("MTOMRA66A41G224M", "Salve123")
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	result := strings.TrimSpace(string(content))
	return result, nil
}
