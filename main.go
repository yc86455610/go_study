package main

import (
	"github.com/no1charleswu/go_study/TTPPC/crawler"
)

func main() {
	testGate()
	// testOkex()
	// testHuobi()
	// crawler.FindForHuobi("https://github.com#")
	// crawler.FindForOkex("https://github.com/okcoin-okex/API-docs-OKEx.com/blob/master/%E5%B8%81%E5%B8%81%E6%9C%80%E5%B0%8F%E4%BA%A4%E6%98%93%E9%87%8F(min_trade_size%20for%20spot).md")
	// testDecode()
}

func testHuobi() {
	crawler.Huobi()
}

func testOkex() {
	crawler.Okex()
}

func testGate() {
	crawler.Gate()
}

// func testBinance() {
// crawler.Binance()
// }
