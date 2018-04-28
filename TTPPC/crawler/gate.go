package crawler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
)

// FindForGate .
func FindForGate(url string) (title string, text string, err error) {
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
	text, err = doc.Find(".itemContent").Html()
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

// CleanDataForGate .
func CleanDataForGate(title, text string) (string, string) {
	return title, text
}

// Gate  API.
func Gate() error {
	s := "https://gate.io/api2"
	title, text, err := FindForGate(s)
	if err != nil {
		log.Println(err)
	}
	title, text = CleanDataForGate(title, text)
	if err = SaveForMd(title, text, "./gate"); err != nil {
		log.Println(err)
	}
	return nil
}
