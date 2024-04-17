package main

import (
	"fmt"

	"github.com/alopes2/go-studies/price-calculator/filemanager"
	"github.com/alopes2/go-studies/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Print("Could not process job")
		// 	fmt.Print(err)
		// }
	}

	for index := range taxRates {
		select {
		case <-doneChans[index]:
			fmt.Println("Done")
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
