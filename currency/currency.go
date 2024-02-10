package currency

import (
	"fmt"
	"os"
	"sync"

	"github.com/olekukonko/tablewriter"
)

const (
	CURRENCY     = "CURRENCY"
	BUYING_RATE  = "BUYING RATE"
	SELLING_RATE = "SELLING RATE"
)

func CurrencyRun() {
	var currencyList [][]string
	var wg sync.WaitGroup

	wg.Add(2)
	go goldHandler(&currencyList, &wg)
	go currencyHandler(&currencyList, &wg)
	wg.Wait()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{CURRENCY, BUYING_RATE, SELLING_RATE})

	for _, v := range currencyList {
		table.Append(v)
	}
	table.Render()
	var input string
	fmt.Scanln(&input)
}
