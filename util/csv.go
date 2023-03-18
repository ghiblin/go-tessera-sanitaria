package util

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

type DateTime time.Time

// Convert the internal date as CSV string
func (date DateTime) MarshalCSV() (string, error) {
	// yyyy-mm-dd
	return time.Time(date).Format("2006-01-02"), nil
}

func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	dt, err := time.Parse("1/2/06", csv)
	if err != nil {
		return err
	}
	*date = DateTime(dt)
	return nil
}

func (date DateTime) String() string {
	return time.Time(date).Format("2006-01-02")
}

type Bool bool

func (b Bool) MarshalCSV() (string, error) {
	if b {
		return "Si", nil
	}
	return "No", nil
}

func (b *Bool) UnmarshalCSV(csv string) (err error) {
	*b = csv == "Si"
	return err
}

type Currency float64

func (c Currency) MarshalCSV() (string, error) {
	return fmt.Sprintf("%.2f €", c), nil
}

func (c *Currency) UnmarshalCSV(csv string) (err error) {
	s := strings.Replace(csv, "€", "", -1)
	s = strings.Replace(s, ",", ".", -1)
	s = strings.TrimSpace(s)
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*c = Currency(v)
	return nil
}

func (c Currency) String() string {
	return fmt.Sprintf("%.2f €", c)
}

type Invoice struct {
	Id           string   `csv:"N° Fattura"`
	Date         DateTime `csv:"Data"`
	Name         string   `csv:"Paziente"`
	Price        Currency `csv:"Visita"`
	Amount       Currency `csv:"Fattura"`
	TaxStamp     Currency `csv:"Marca da bollo"`
	Payment      string   `csv:"Tipo di Pagamento"`
	TaxCode      string   `csv:"CF"`
	Traceability Bool     `csv:"Tracciabilità"`
}

func LoadInvoices(filename string) ([]*Invoice, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	invoices := []*Invoice{}
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = ';'
		return r
	})
	if err := gocsv.UnmarshalFile(file, &invoices); err != nil {
		return nil, err
	}

	return invoices, nil
}

func (i *Invoice) String() string {
	return fmt.Sprintf("%s: [%s] %s %s", i.Id, i.Date, i.Name, i.Amount)
}
