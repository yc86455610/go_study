package main

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/levigross/grequests"
)

func main() {
	netReceiveAmount, confirmations, err := HsrBlocksChainCheck(0.94999999, "0x05ee546c1a62f90d7acbffd6d846c9c54c7cf94c", "HHPcyzQ2rVdkH47XbMNsFLPmjH9etw46sY")
	if err != nil {
		fmt.Println("request failed...")
		return
	}
	fmt.Println(fmt.Sprintf("net_withdraw_amount: %f, confirmations: %d", netReceiveAmount, confirmations))
}

const min = 0.000000000001

func isEqual(f1, f2 float64) bool {
	if f1 < f2 {
		return isEqual(f2, f1)
	}
	return math.Dim(f1, f2) < min
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

// Transaction .
type Transaction struct {
	Vout          []vout `json:"vout"`
	Confirmations int64  `json:"confirmations"`
}

// HsrBlocksChainCheck .
func HsrBlocksChainCheck(withdrawAmount float64, originalAddress string, targetAddress string) (netWithdrawAmount float64, confirmations int64, err error) {
	url := fmt.Sprintf("http://explorer.h.cash/api/getaddress/%s", targetAddress)
	bData, err := HTTPGet(url, nil)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	var add address
	err = json.Unmarshal(bData, &add)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// fmt.Printf("%+v", add)
	for i := 0; i < len(add.LastTxs); i++ {
		txid := add.LastTxs[i].Txid
		url = fmt.Sprintf("http://explorer.h.cash/api/getrawtransaction/%s/1", txid)
		bData, err = HTTPGet(url, nil)
		if err != nil {
			return
		}
		// fmt.Println(string(bData[:]))
		var transaction Transaction
		err = json.Unmarshal(bData, &transaction)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		// fmt.Printf("%+v", transaction)
		for j := 0; j < len(transaction.Vout); j++ {
			if transaction.Vout[j].Value == withdrawAmount && transaction.Vout[j].ScriptPubkey.Addresses[0] == targetAddress {
				confirmations = transaction.Confirmations
				netWithdrawAmount = transaction.Vout[j].Value
			}
		}
	}
	return
}
