package crawler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// FindHuobi .
func FindHuobi(url string) (title string, text string) {
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
	title = doc.Find("h1.gh-header-title").Text()
	fmt.Println("title: ", title)
	text, err = doc.Find("#wiki-body").Find(".markdown-body").Eq(0).Html()
	fmt.Println("text: ", text)

	return title, text
}

// CleanDataForHuobi .
func CleanDataForHuobi(title, text string) (string, string) {
	return title, text
}
