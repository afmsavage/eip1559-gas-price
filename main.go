package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

type gas struct {
	Block uint64     `json:"block"`
	Price *big.Float `json:"price"`
}

func gasPrice(w http.ResponseWriter, req *http.Request) {
	endpoint, present := os.LookupEnv("ALCHEMY")
	if present == false {
		endpoint = "https://cloudflare-eth.com"
	}
	fmt.Printf("Using %v for the endpoint\n", endpoint)

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully made connection with endpoint")

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(blockNumber)

	gasFee, _ := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasFeeReadable := new(big.Float)
	gasFeeReadable.SetString(gasFee.String())
	readable := new(big.Float).Quo(gasFeeReadable, big.NewFloat(math.Pow10(int(9))))
	// fmt.Println(gasFee)
	// fmt.Printf("gasFee: %f", readable)

	response := &gas{
		Block: blockNumber,
		Price: readable,
	}
	newResponse, _ := json.Marshal(response)
	fmt.Fprintf(w, string(newResponse))
}
func main() {

	http.HandleFunc("/gasPrice", gasPrice)
	http.ListenAndServe(":8090", nil)
}
