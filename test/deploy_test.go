package test

import (
	"github.com/eds-voting/internal/storage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

// Test inbox contract gets deployed correctly
func TestDeployInbox(t *testing.T) {

	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 10000000)

	//Deploy contract
	address, _, _, err := contracts.DeployVotingStorage(
		auth,
		blockchain,
	)
	// commit all pending transactions
	blockchain.Commit()

	if err != nil {
		t.Fatalf("Failed to deploy the Inbox contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}

}

//Test initial message gets set up correctly
func TestOwner(t *testing.T) {

	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 10000000)

	//Deploy contract
	_, _, contract, err := contracts.DeployVotingStorage(
		auth,
		blockchain,
	)
	if err != nil {
		t.Errorf("Contract wasn't deployed %s", err)

	}
	// commit all pending transactions
	blockchain.Commit()

	if got, _ := contract.Owner(nil); got != auth.From {
		t.Errorf("Expected message to be: Hello World. Go: %s", got)
	}

}
