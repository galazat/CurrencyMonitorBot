package currency

import "encoding/xml"

type Currency struct {
	XMLName  xml.Name `xml:"Valute"`
	NumCode  int      `xml:"NumCode"`
	CharCode string   `xml:"CharCode"`
	Nominal  int      `xml:"Nominal"`
	Name     string   `xml:"Name"`
	Value    string   `xml:"Value"`
}

type Currencies struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Carrencyes []Currency `xml:"Valute"`
}

func NewCurrencies() *Currencies {
	return &Currencies{}
}
