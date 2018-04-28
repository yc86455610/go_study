package crawler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
)

// LinkScrapeForZb .
func LinkScrapeForZb() ([]string, error) {
	s := make([]string, 0)
	doc, err := goquery.NewDocument("https://www.zb.com/i/developer")
	if err != nil {
		log.Println(err)
		return s, err
	}

	doc.Find("ul.nav").Each(func(i int, sel *goquery.Selection) {
		d, _ := sel.Attr("href")
		d = "https://www.zb.com" + d
		s = append(s, d)
	})
	return s, nil
}

// FindForZb .
func FindForZb(url string) (title string, text string, err error) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s\n", res.StatusCode, res.Status)
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
	text, err = doc.Find("article.table-responsive").Html()
	if err != nil {
		log.Println(err)
		return
	}
	if len(text) == 0 {
		log.Println("error: empty text. ", url)
		return title, text, errors.New("empty text")
	}
	// fmt.Println("text: ", text)

	return title, text, nil
}

// CleanDataForZb .
func CleanDataForZb(title, text string) (string, string) {
	return title, text
}

// Zb  API.
func Zb() error {
	s, err := LinkScrapeForZb()
	if err != nil {
		log.Println(err)
		return err
	}
	for i := 0; i < len(s); i++ {
		title, text, err := FindForZb(s[i])
		if err != nil {
			log.Println(err)
			continue
		}
		title, text = CleanDataForZb(title, text)
		if err = SaveForMd(title, text, "./Zb"); err != nil {
			log.Println(err)
		}
	}
	return nil
}
