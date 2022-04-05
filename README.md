# API for currency exchange API https://currencylayer.com

## Install

> git clone https://github.com/bananoamo/currencylayer

## Example
```
package main

import (
  "https://github.com/bananoamo/currencylayer"
  "log"
)
  

func main() {
  // init `options`
  options := currencylayer.New()
  
  // to get currencyLayerToken you have to register at https://currencylayer.com
  options.AddAccessKey(currencyLayerToken)
  
  // you can get API url at after sing up at https://currencylayer.com
  // if you use free version you can use parameter empty string, because package has default not secure API url 
  quotes, err := options.GetQuotes(apiURL)
  if err != nil {
    log.Fatal(err)
  }
  
  // you can get all quotes as map[string]float64, where keys are like `USDKES`
  quoteList := quotes.QuotesList()
  
  // or you can get current quote
  quote, found := quotes.GetQuote()
  if !found {
    log.Fatal(`quote is empty`)
  }
  
  // customization
  
  // also you can customize you option before do options.GetQuotes()
  // if you specifies currency slice, you will get quotes only with list's of currencies as parameter
  quotes.AddCurrencies([]string{`usd`})
  
  // also you can change source currency, for example by default all currency quotes are `USD...any other currency`
  // if you set source currecy `KES`, all availiable quotes will be `KESUSD`, `KESGBP` etc
  // this option only availiable at paid API version
  options.AddSource(`KES`)
  
  // also you can specify output API format, by default it is `1`
  quotes.EditFormat(`2`)
  
}
```
