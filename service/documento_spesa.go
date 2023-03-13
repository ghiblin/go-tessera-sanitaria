package service

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ghiblin/go-tessera-sanitaria/wsdl"
)

const url = "https://invioSS730pTest.sanita.finanze.it/DocumentoSpesa730pWeb/DocumentoSpesa730pPort"

type DocumentoSpesa struct{}

func (s *DocumentoSpesa) Inserimento(request *wsdl.InserimentoDocumentoSpesaRequest) string {
	var root = wsdl.Envelope{}
	root.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	// root.Doc = "http://documentospesap730.sanita.finanze.it"
	root.Header = &wsdl.Header{}
	root.Body = &wsdl.Body{}
	root.Body.Request = request

	out, _ := xml.MarshalIndent(&root, " ", "  ")
	body := string(out)
	log.Println(body)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "text/xml")
	req.SetBasicAuth("MTOMRA66A41G224M", "Salve123")
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	result := strings.TrimSpace(string(content))
	return result
}
