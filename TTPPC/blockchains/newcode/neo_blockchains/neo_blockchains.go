package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/levigross/grequests"
)

// Result .
type Result struct {
	res            bool
	categorization string
	value          float64
	timestamp      int64
	address        string
	txid           string
}

func main() {
	r, err := NeoBlocksChainCheck("neo", "c9f3c922f8205937173e8ab7403fabe85aeffe3ec8c304ae393a13ba3fc0124b", "ALTTfbqUam5T2gRaJDF2B6RTmhVTj3QosX")
	if err != nil {
		fmt.Println("request failed...")
		return
	}
	displayResult(r)
}

func displayResult(result Result) {
	if result.res {
		fmt.Printf("成功 ")
	} else {
		fmt.Printf("失败 ")
	}
	fmt.Println(strings.ToUpper(result.categorization))
	fmt.Println(result.value)
	fmt.Println(time.Unix(result.timestamp, 0).Format("2006-01-02 15:04:05"))
	fmt.Println("地址：\t", result.address)
	fmt.Println("Txid:\t", result.txid)
}

// HTTPGet asfdadsf.
func HTTPGet(url string, requestOptions *grequests.RequestOptions) (response []byte, err error) {
	httpResponse, err := grequests.Get(url, requestOptions)
	if err == nil {
		if httpResponse.StatusCode == 200 {
			response = httpResponse.Bytes()
		}
	}
	return
}

type vout struct {
	Value   float64 `json:"value"`
	Address string  `json:"address"`
}

type transaction struct {
	Time int64  `json:"time"`
	Vout []vout `json:"vouts"`
}

// NeoBlocksChainCheck .
func NeoBlocksChainCheck(categorization, txid, address string) (result Result, err error) {
	if strings.ToLower(categorization) != "neo" {
		fmt.Println("type error: The first parameter is not HSR.")
		return result, errors.New("type")
	}
	url := fmt.Sprintf("https://neoscan.io/api/main_net/v1/get_address/%s", address)
	bData, err := HTTPGet(url, nil)
	if err != nil {
		fmt.Println("address error: ", err)
		return
	}
	url = fmt.Sprintf("https://neoscan.io/api/main_net/v1/get_transaction/%s", txid)
	bData, err = HTTPGet(url, nil)
	if err != nil {
		fmt.Println("txid error: ", err)
		return
	}
	// fmt.Println(string(bData[:]))

	var tx transaction
	err = json.Unmarshal(bData, &tx)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// fmt.Printf("%+v\n", tx)

	for j := 0; j < len(tx.Vout); j++ {
		if tx.Vout[j].Address == address {
			result.value = tx.Vout[j].Value
		}
	}
	result.res = true
	result.categorization = categorization
	result.timestamp = tx.Time
	result.address = address
	result.txid = txid
	return result, nil
}
