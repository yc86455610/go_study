package crawler

import (
	"fmt"
	"log"
	"testing"
)

func Test(t *testing.T) {
	title, text := FindHuobi("https://github.com/huobiapi/API_Docs/wiki/FIX-v2-API-instances")
	fmt.Println(title, text)
	title, text = CleanDataForHuobi(title, text)
	fmt.Println(title, text)
	if err := SaveForMd(title, text, "."); err != nil {
		log.Fatal(err)
	}
}
