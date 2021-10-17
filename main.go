package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
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
	fmt.Println(blockNumber)

	gasFee, _ := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasFeeReadable := new(big.Float)
	gasFeeReadable.SetString(gasFee.String())
	readable := new(big.Float).Quo(gasFeeReadable, big.NewFloat(math.Pow10(int(9))))
	fmt.Println(gasFee)
	fmt.Printf("gasFee: %f", readable)

}
