package crawler

// package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// FindForBinance .
func FindForBinance(url string) (title, text string) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	title = doc.Find("#readme").Find("h1").Text()
	fmt.Println("title: ", title)
	text, err = doc.Find("#readme").Find("article.markdown-body").Html()
	fmt.Println("text: ", text)

	return title, text
}

// CleanDataForBinance .
func CleanDataForBinance(title, text string) (string, string) {
	return title, text
}
