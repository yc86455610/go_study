package main

import (
	"fmt"
	"net/url"

	"github.com/axgle/mahonia"
	"github.com/no1charleswu/go_study/TTPPC/crawler"
)

func main() {
	testOkex()
	// testHuobi()
	// crawler.FindForHuobi("https://github.com#")
	// crawler.FindForOkex("https://github.com/okcoin-okex/API-docs-OKEx.com/blob/master/%E5%B8%81%E5%B8%81%E6%9C%80%E5%B0%8F%E4%BA%A4%E6%98%93%E9%87%8F(min_trade_size%20for%20spot).md")
	testDecode()
}

func testHuobi() {
	crawler.Huobi()
}

func testOkex() {
	crawler.Okex()
}

// func testBinance() {
// crawler.Binance()
// }

func testDecode() {
	s := "%E6%9B%B4%E6%96%B0%E6%97%A5%E5%BF%97.md"
	fmt.Println(s)
	enc := mahonia.NewEncoder("UTF-8")
	//converts a  string from UTF-8 to gbk encoding.
	fmt.Println(enc.ConvertString("hello,世界"))
	fmt.Println(enc.ConvertString(s))
	dec := mahonia.NewDecoder("UTF-8")
	fmt.Println(dec.ConvertString(s))
	a := "123"
	fmt.Println(url.QueryEscape(a))
	fmt.Println(url.QueryUnescape(s))
}
