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
	r, err := HsrBlocksChainCheck("hsr", "cd5f32fe90f42ea30e5e81fe3186e0de2b32345e3b961d9efba34820ab64e4b0", "HHPcyzQ2rVdkH47XbMNsFLPmjH9etw46sY")
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

// Txs .
type Txs struct {
	Type string `json:"type"`
	Txid string `json:"txid"`
}

type address struct {
	Address string `json:"address"`
	LastTxs []Txs  `json:"last_txs"`
}

type script struct {
	Addresses []string `json:"addresses"`
}
type vout struct {
	Value        float64 `json:"value"`
	ScriptPubkey script  `json:"scriptPubkey"`
}

type transaction struct {
	Time          int64  `json:"time"`
	Vout          []vout `json:"vout"`
	Confirmations int64  `json:"confirmations"`
}

// HsrBlocksChainCheck .
func HsrBlocksChainCheck(categorization, txid, address string) (result Result, err error) {
	if strings.ToLower(categorization) != "hsr" {
		fmt.Println("type error: The first parameter is not HSR.")
		return result, errors.New("type")
	}
	url := fmt.Sprintf("http://explorer.h.cash/api/getaddress/%s", address)
	bData, err := HTTPGet(url, nil)
	if err != nil {
		fmt.Println("address error: ", err)
		return
	}
	url = fmt.Sprintf("http://explorer.h.cash/api/getrawtransaction/%s/1", txid)
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
	// fmt.Printf("%+v\n", transaction)

	for j := 0; j < len(tx.Vout); j++ {
		if tx.Vout[j].ScriptPubkey.Addresses[0] == address {
			result.value = tx.Vout[j].Value
			if tx.Confirmations >= 6 {
				result.res = true
			}
		}
	}
	result.categorization = categorization
	result.timestamp = tx.Time
	result.address = address
	result.txid = txid
	return result, nil
}
