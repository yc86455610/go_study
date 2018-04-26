package main

import (
	"fmt"
	"log"

	"github.com/no1charleswu/go_study/TTPPC/crawler"
)

func main() {
	// title, text := crawler.FindForBinance("https://github.com/binance-exchange/binance-official-api-docs/blob/master/errors.md")
	title, text := crawler.FindForOkex("https://github.com/okcoin-okex/API-docs-OKEx.com/blob/master/API-For-Futures-CN/%E5%90%88%E7%BA%A6%E4%BA%A4%E6%98%93REST%20API.md")
	fmt.Println(title, text)
	if err := crawler.SaveForMd(title, text, "."); err != nil {
		log.Fatal(err)
	}
}
