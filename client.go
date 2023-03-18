package main

import (
	"log"

	"github.com/ghiblin/go-tessera-sanitaria/service"
	"github.com/ghiblin/go-tessera-sanitaria/util"
)

func main() {
	invoices, err := util.LoadInvoices("./fatture.csv")
	if err != nil {
		log.Fatalf("Failed to load invoices: %s", err)
	}

	srv, err := service.NewDocumentoSpesaService()
	if err != nil {
		log.Fatalf("Failed to initialize service: %s", err)
	}

	for _, invoice := range invoices {
		err = srv.Inserimento(invoice)
		if err != nil {
			log.Printf("Failed to send %s: %s", invoice, err)
		} else {
			log.Printf("Invoice %s sent correctly", invoice)
		}
	}
}
