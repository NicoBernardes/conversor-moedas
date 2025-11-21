package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type exchange struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Utilize: ./convert [valor] [moeda_destino]")
		os.Exit(1)
	}

	real := os.Args[1]
	moeda_destino := os.Args[2]

	realBrl, err := strconv.ParseFloat(real, 64)
	if err != nil {
		fmt.Println("Valor invalido")
		os.Exit(1)
	}

	moeda_destino = strings.ToUpper(strings.TrimSpace(moeda_destino))

	str := `
{
  "base": "BRL",
  "date": "2025-04-14",
  "rates": {
    "USD": 0.151,
    "EUR": 0.137,
    "JPY": 16.29,
    "GBP": 0.13,
    "CHF": 0.1402,
    "AUD": 0.2712,
    "CAD": 0.2374,
    "CNY": 1.251,
    "HKD": 1.326,
    "NZD": 0.2922,
    "SEK": 1.655,
    "NOK": 1.806,
    "DKK": 1.122,
    "SGD": 0.2249,
    "KRW": 242.97,
    "ZAR": 3.239,
    "MXN": 3.454,
    "INR": 14.71,
    "ILS": 0.63,
    "THB": 5.74,
    "IDR": 2875.0,
    "MYR": 0.754,
    "PHP": 9.74,
    "PLN": 0.644,
    "CZK": 3.77,
    "HUF": 61.59,
    "TRY": 6.49,
    "BGN": 0.293,
    "RON": 0.746
  }
}
`
	var data exchange

	if err := json.Unmarshal([]byte(str), &data); err != nil {
		fmt.Println("erro ao carregar dados", err)
		os.Exit(1)
	}

	cotacao, ok := data.Rates[moeda_destino]
	if !ok {
		fmt.Println("moeda não encontrada")
		os.Exit(1)
	}

	vlrConvertido := realBrl * cotacao

	fmt.Printf("%.2f\n", vlrConvertido)
}
