package bitcoin

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

const transferReceipt = `{
	"txid":"df63ddab7d4eed2fb6cb40d4d0519e7e5ac7cf5ad556b2edbd45963ea1a2931c",
	"version":1,
	"vin":[
		{
			"txid":"bf19be44d7dc3e1e6771801a1d250c7207fa9b09d8df9b0ee1b028b6c153475e",
			"sequence":4294967295,
			"n":0,
			"addresses":[
				"3QJmV3qfvL9SuYo34YihAf3sRCW3qSinyC"
			],
			"value":"777200",
			"hex":"00483045022100e9d0db3bb20a5828ab9dae7cd8373064ce087cc9c8e3def87034d5c2f6f3abb9022047d7c27b355c6487cff40bfbd45742d26d727f3135b2396d8f1abc371c51870c01473044022016280108af73079a69f378218ad4259f02c4e4b6f52c573729650cb3645bc9180220785973cb4e5c4ec6340dc77dc56cec3fb74ebd7296cf1d14344d4f3e157658bb014cc952410491bba2510912a5bd37da1fb5b1673010e43d2c6d812c514e91bfa9f2eb129e1c183329db55bd868e209aac2fbc02cb33d98fe74bf23f0c235d6126b1d8334f864104865c40293a680cb9c020e7b1e106d8c1916d3cef99aa431a56d253e69256dac09ef122b1a986818a7cb624532f062c1d1f8722084861c5c3291ccffef4ec687441048d2455d2403e08708fc1f556002f1b6cd83f992d085097f9974ab08a28838f07896fbab08f39495e15fa6fad6edbfb1e754e35fa1c7844c41f322a1863d4621353ae"}
	],
	"vout":[
		{
			"value":"677012",
			"n":0,
			"hex":"a91499fa965ad13a9580ed7a64ac24b2ecad30f1209a87",
			"addresses":["3FjBW1KL9L8aYtdKzJ8FhCNxmXB7dXDRw4"]
		}
	],
	"blockHash":"00000000000000000011b58c01ede5a602eec61ebaf097aaa6e682ef2819536e",
	"blockHeight":585094,
	"confirmations":1997,
	"blockTime":1562945790,
	"value":"677012",
	"valueIn":"777200",
	"fees":"100188",
	"hex":"01000000015e4753c1b628b0e10e9bdfd8099bfa07720c251d1a8071671e3edcd744be19bf00000000fd5d0100483045022100e9d0db3bb20a5828ab9dae7cd8373064ce087cc9c8e3def87034d5c2f6f3abb9022047d7c27b355c6487cff40bfbd45742d26d727f3135b2396d8f1abc371c51870c01473044022016280108af73079a69f378218ad4259f02c4e4b6f52c573729650cb3645bc9180220785973cb4e5c4ec6340dc77dc56cec3fb74ebd7296cf1d14344d4f3e157658bb014cc952410491bba2510912a5bd37da1fb5b1673010e43d2c6d812c514e91bfa9f2eb129e1c183329db55bd868e209aac2fbc02cb33d98fe74bf23f0c235d6126b1d8334f864104865c40293a680cb9c020e7b1e106d8c1916d3cef99aa431a56d253e69256dac09ef122b1a986818a7cb624532f062c1d1f8722084861c5c3291ccffef4ec687441048d2455d2403e08708fc1f556002f1b6cd83f992d085097f9974ab08a28838f07896fbab08f39495e15fa6fad6edbfb1e754e35fa1c7844c41f322a1863d4621353aeffffffff0194540a000000000017a91499fa965ad13a9580ed7a64ac24b2ecad30f1209a8700000000"
}`

var expectedTransferTrx = blockatlas.Tx{
	ID:       "df63ddab7d4eed2fb6cb40d4d0519e7e5ac7cf5ad556b2edbd45963ea1a2931c",
	Coin:     coin.BTC,
	From:     "",
	To:       "",
	Inputs:   []string{"3QJmV3qfvL9SuYo34YihAf3sRCW3qSinyC"},
	Outputs:  []string{"3FjBW1KL9L8aYtdKzJ8FhCNxmXB7dXDRw4"},
	Fee:      "100188",
	Date:     1562945790,
	Type:     "transfer",
	Status:   "completed",
	Block:    585094,
	Sequence: 585094,
	Meta: blockatlas.Transfer{
		Value: "677012",
	},
}

func TestNormalizeTransfer(t *testing.T) {
	var tests = []struct {
		Receipt  string
		Expected blockatlas.Tx
	}{
		{transferReceipt, expectedTransferTrx},
	}

	for _, test := range tests {
		var receipt TransferReceipt

		rErr := json.Unmarshal([]byte(test.Receipt), &receipt)
		if rErr != nil {
			t.Fatal(rErr)
		}

		var readyTx blockatlas.Tx
		normTx, ok := NormalizeTransfer(&receipt, 0)
		if !ok {
			t.Fatal("Bitcoin: Can't normalize transaction", readyTx)
		}
		readyTx = normTx

		actual, err := json.Marshal(&readyTx)
		if err != nil {
			t.Fatal(err)
		}

		expected, err := json.Marshal(&test.Expected)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(actual, expected) {
			println(string(actual))
			println(string(expected))
			t.Error("Transactions not equal")
		}
	}
}
