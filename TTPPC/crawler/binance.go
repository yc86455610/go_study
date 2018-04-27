package crawler

// package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// FindForBinance .
func FindForBinance(url string) (title string, text string, err error) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	_, title = filepath.Split(url)
	fmt.Println("title: ", title)
	text, err = doc.Find("#readme").Find("article.markdown-body").Html()
	if err != nil {
		log.Println(err)
		return
	}
	if len(text) == 0 {
		log.Println("error: empty text. ", url)
		return title, text, errors.New("empty text")
	}
	fmt.Println("text: ", text)

	return title, text, nil
}

// CleanDataForBinance .
func CleanDataForBinance(title, text string) (string, string) {
	title, err := url.QueryUnescape(title)
	if err != nil {
		log.Println("QueryUnescape error")
		return title, text
	}
	n := strings.IndexByte(title, '.')
	if n != -1 {
		return title[:n], text
	}
	return title, text
}

// Binance API
// func Binance() error {
// 	s, err := LinkScrapeForBinance()
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	for i := 0; i < len(s); i++ {
// 		title, text, err := FindForBinance(s[i])
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		title, text = CleanDataForBinance(title, text)
// 		if err = SaveForMd(title, text, "./binance"); err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 	}
// 	return nil
// }
