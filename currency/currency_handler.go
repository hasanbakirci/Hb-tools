package currency

import (
	"log"
	"sync"

	"github.com/gocolly/colly"
)

const (
	GOLD_URL     = "https://bigpara.hurriyet.com.tr/altin/gram-altin-fiyati/"
	CURRENCY_URL = "https://www.bloomberght.com/doviz"
	GOLD         = "GRAM ALTIN"
)

func goldHandler(currencyList *[][]string, wg *sync.WaitGroup) {
	defer wg.Done()
	c := colly.NewCollector()
	var goldList []string

	goldList = append(goldList, GOLD)

	c.OnHTML("div.kurDetail", func(e *colly.HTMLElement) {
		buyingRate := e.ChildText("div.kurBox:nth-child(2) span.value")
		goldList = append(goldList, buyingRate)
		sellingRate := e.ChildText("div.kurBox:nth-child(3) span.value")
		goldList = append(goldList, sellingRate)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	err := c.Visit(GOLD_URL)
	if err != nil {
		log.Fatal(err)
	}
	*currencyList = append(*currencyList, goldList)
}

func currencyHandler(currencyList *[][]string, wg *sync.WaitGroup) {
	defer wg.Done()

	c := colly.NewCollector()
	c.OnHTML("div.widget-table-data.type3.goldPriceWidget", func(e *colly.HTMLElement) {
		e.ForEach(".box-4", func(_ int, box *colly.HTMLElement) {
			var currencies []string
			currencyName := box.ChildText("th a")
			currencies = append(currencies, currencyName)
			buyingRate := box.ChildText("tbody tr:nth-child(1) td:nth-child(2)")
			currencies = append(currencies, buyingRate)
			sellingRate := box.ChildText("tbody tr:nth-child(2) td:nth-child(2)")
			currencies = append(currencies, sellingRate)

			*currencyList = append(*currencyList, currencies)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	err := c.Visit(CURRENCY_URL)
	if err != nil {
		log.Fatal(err)
	}
}
