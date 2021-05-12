GO SIDESHIFT.AI Client
================

[![GoDoc](https://godoc.org/github.com/omani/go-sideshiftai-client?status.svg)](https://godoc.org/github.com/omani/go-sideshiftai-client)


<p align="center">
<img src="https://raw.githubusercontent.com/omani/go-sideshiftai-client/main/assets/img/icon.png" alt="Logo" width="1250" />
</p>

A client implementation for the [sideshift.ai](https://sideshift.ai) service written in go.

### Installation

```sh
go get -u github.com/omani/go-sideshiftai-client
```

#### Example code:

```Go
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

```

# Run the example code
```Go
cd cmd
go run main.go
```

# Contribution
* You can fork this, extend it and contribute back.
* You can contribute with pull requests.

# Donations
I love Monero (XMR) and building applications for and on top of Monero.

You can make me happy by donating Monero to the following address:

```
89woiq9b5byQ89SsUL4Bd66MNfReBrTwNEDk9GoacgESjfiGnLSZjTD5x7CcUZba4PBbE3gUJRQyLWD4Akz8554DR4Lcyoj
```

# LICENSE
MIT License
