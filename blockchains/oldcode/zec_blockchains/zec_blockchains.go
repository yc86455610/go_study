package main

import (
	"encoding/json"
	"fmt"

	"github.com/levigross/grequests"
)

func main() {
	netReceiveAmount, confirmations, err := ZecBlocksChainCheck(0.02588889, "t1aZ2DGuiokCxHVfb4cGQqXghxy9hUpE6xQ", "t1J7StFmqay9sWHMtUhXseLRMzvm1vLE7i7")
	if err != nil {
		fmt.Println("request failed...")
		return
	}
	fmt.Println(fmt.Sprintf("net_withdraw_amount: %f, confirmations: %d", netReceiveAmount, confirmations))
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

type script struct {
	Addresses []string `json:"addresses"`
}

type vout struct {
	Value        float64 `json:"value"`
	ScriptPubkey script  `json:"scriptPubkey"`
}
type address struct {
	Vout []vout `json:"vout"`
}

// ZecBlocksChainCheck .
func ZecBlocksChainCheck(withdrawAmount float64, originalAddress string, targetAddress string) (netWithdrawAmount float64, confirmations int64, err error) {
	url := fmt.Sprintf("https://api.zcha.in/v2/mainnet/accounts/%s/recv?limit=20&offset=0", targetAddress)
	bData, err := HTTPGet(url, nil)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// fmt.Println(string(bData[:]))
	var add []address
	err = json.Unmarshal(bData, &add)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// fmt.Printf("%+v\n", add)
	for i := 0; i < len(add); i++ {
		out := add[i].Vout
		// fmt.Printf("%+v\n", out)
		for j := 0; j < len(out); j++ {
			// fmt.Printf("%+v\n", out[j])
			if out[j].Value == withdrawAmount && out[j].ScriptPubkey.Addresses[0] == targetAddress {
				netWithdrawAmount = out[j].Value
				return
			}
		}
	}
	return
}
