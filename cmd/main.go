package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/omani/go-sideshiftai-client"
)

func main() {
	client := sideshiftai.New(sideshiftai.Config{})

	pair, err := client.GetFetchPairs("btc/xmr")
	if err != nil {
		panic(err)
	}

	log.Printf("Min: %s - Max: %s - Rate: %s\n", pair.Min, pair.Max, pair.Rate)

	// Now create a variable (variable rate) order with that pair
	order, err := client.PostCreateVariableOrders(&sideshiftai.RequestVariableOrders{
		Type:            "variable",
		DepositMethodId: "btc",
		SettleMethodId:  "xmr",
		SettleAddress:   "89woiq9b5byQ89SsUL4Bd66MNfReBrTwNEDk9GoacgESjfiGnLSZjTD5x7CcUZba4PBbE3gUJRQyLWD4Akz8554DR4Lcyoj", // random monero address for testing purposes.
		RefundAddress:   "1F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX",                                                              // random BTC address from blockexplorer for testing purposes.
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(order)

	// check our order if filled (after we paid to the DepositAddress in the response or order)
	spew.Dump(client.GetFetchOrders(order.ID))
}
