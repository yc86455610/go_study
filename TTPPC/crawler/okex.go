package crawler

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

// LinkScrapeForOkex .
func LinkScrapeForOkex() ([]string, error) {
	s := make([]string, 0)
	doc, err := goquery.NewDocument("https://github.com/okcoin-okex/API-docs-OKEx.com")
	if err != nil {
		log.Println(err)
		return s, err
	}

	t := doc.Find(".file-wrap a")
	for i := 0; i < t.Length(); i++ {
		d, _ := t.Eq(i).Attr("href")
		fmt.Println(d)
		d = "http://github.com" + d
		s = append(s, d)
	}
	return s, nil
}

// FindForOkex .
func FindForOkex(url string) (title string, text string, err error) {
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

// CleanDataForOkex .
func CleanDataForOkex(title, text string) (string, string) {
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

// Okex API
func Okex() error {
	s, err := LinkScrapeForOkex()
	if err != nil {
		log.Println(err)
		return err
	}
	for i := 0; i < len(s); i++ {
		title, text, err := FindForOkex(s[i])
		if err != nil {
			log.Println(err)
			continue
		}
		title, text = CleanDataForOkex(title, text)
		if err = SaveForMd(title, text, "./okex"); err != nil {
			log.Println(err)
			continue
		}
	}
	return nil
}
