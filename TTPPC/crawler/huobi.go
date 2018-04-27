package crawler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
)

// LinkScrapeForHuobi .
func LinkScrapeForHuobi() ([]string, error) {
	s := make([]string, 0)
	doc, err := goquery.NewDocument("https://github.com/huobiapi/API_Docs/wiki")
	if err != nil {
		log.Println(err)
		return s, err
	}

	t := doc.Find(".wiki-pages a")
	for i := 0; i < t.Length(); i++ {
		d, _ := t.Eq(i).Attr("href")
		// fmt.Println(d)
		d = "https://github.com" + d
		s = append(s, d)
	}
	// fmt.Println(s)
	return s, nil
}

// FindForHuobi .
func FindForHuobi(url string) (title string, text string, err error) {
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
	text, err = doc.Find("#wiki-body").Find(".markdown-body").Eq(0).Html()
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

// CleanDataForHuobi .
func CleanDataForHuobi(title, text string) (string, string) {
	return title, text
}

// Huobi  API.
func Huobi() error {
	s, err := LinkScrapeForHuobi()
	if err != nil {
		log.Println(err)
		return err
	}
	for i := 0; i < len(s); i++ {
		title, text, err := FindForHuobi(s[i])
		if err != nil {
			log.Println(err)
			continue
		}
		title, text = CleanDataForHuobi(title, text)
		if err = SaveForMd(title, text, "./huobi"); err != nil {
			log.Println(err)
		}
	}
	return nil
}
