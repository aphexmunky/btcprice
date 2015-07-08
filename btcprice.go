package main

import (
	"fmt"
	"github.com/Toorop/go-bitpay"
)

func printRate(currency string, currencySymbol string) {
    bp := bitpay.New("")
    rate,_ := bp.GetBitcoinBestRate(currency)
    bitsPerUnit := (1.0 / rate.Rate) * 1000000

    fmt.Printf("%v%.2f (%v1 = %.2f bits)\n", currencySymbol, rate.Rate, currencySymbol, bitsPerUnit)
}

func main() {
    printRate("USD", "$")
    printRate("EUR", "€")
    printRate("GBP", "£")
}
