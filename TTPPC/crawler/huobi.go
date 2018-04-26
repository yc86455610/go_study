package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

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
	text = doc.Find("#wiki-body").Find(".markdown-body").Eq(0).Text()
	fmt.Println("text: ", text)

	return title, text
}

// CleanDataForHuobi .
func CleanDataForHuobi(title, text string) (string, string) {
	return title, text
}

// SaveForMd .
func SaveForMd(title, text string, loc string) (err error) {
	name := loc + "/" + title + ".md"
	fmt.Println("file name: ", name)
	f, err := os.Create(name) //创建文件
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n, err := w.WriteString(text)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("写入 %d 个字节", n)
	w.Flush()
	return nil
}
func main() {
	title, text := FindHuobi("https://github.com/huobiapi/API_Docs/wiki/FIX-v2-API-instances")
	title, text = CleanDataForHuobi(title, text)
	if err := SaveForMd(title, text, "."); err != nil {
		log.Fatal(err)
	}
}
