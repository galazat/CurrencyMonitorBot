package currency

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
)

var TodayCurrensies *Currencies

func GetCurrencies() *Currencies {
	xmlBytes, err := getXML(getUrl())
	if err != nil {
		log.Println(err)
	}

	currencies := NewCurrencies()

	r := bytes.NewReader([]byte(xmlBytes))
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel

	err = d.Decode(currencies)
	if err != nil {
		log.Println(err)
		return currencies
	}

	return currencies
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func UpdateCurrencyData(hour, min, sec int) error {
	TodayCurrensies = GetCurrencies()

	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	// Вычисляем время первого запуска.
	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc)
	if firstCallTime.Before(now) {
		// Если получилось время раньше текущего, прибавляем сутки.
		firstCallTime = firstCallTime.Add(time.Hour * 24)
	}

	// Вычисляем временной промежуток до запуска.
	duration := firstCallTime.Sub(time.Now().Local())

	go func() {
		time.Sleep(duration)
		for {
			TodayCurrensies = GetCurrencies()
			// Следующий запуск через сутки.
			time.Sleep(time.Hour * 24)
		}
	}()

	return nil
}

func (c *Currencies) Get(charCode string) *Currency {
	for i := 0; i < len(c.Carrencyes); i++ {
		if c.Carrencyes[i].CharCode == charCode {
			return &c.Carrencyes[i]
		}
	}
	return nil
}
