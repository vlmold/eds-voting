package main

import (
	"encoding/json"
	"fmt"
	"github.com/eds-voting/configs"
	"github.com/eds-voting/internal/storage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	configBytes, err := ioutil.ReadFile("./configs/config.json")
	if err != nil {
		panic(err)
	}
	var config configs.Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		panic(err)
	}

	keyjson, err := ioutil.ReadFile(config.Keypath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(keyjson))
	url := config.Endpoint
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(string(keyjson)), config.Keypassword)

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	address, _, _, _ := contracts.DeployVotingStorage(
		auth,
		blockchain,
	)

	fmt.Printf("Contract pending deploy: 0x%x\n", address)
}
