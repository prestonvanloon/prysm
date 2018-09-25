package main

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prysmaticlabs/prysm/contracts/validator-registration-contract/testutils"
)

func TestSendDepositTransaction(t *testing.T) {
	acc, err := testutils.Setup()
	if err != nil {
		t.Fatalf("Failed to setup test %v", err)
	}

	p := &powchainclient{
		contractAddress: acc.ContractAddr,
		priv:            acc.PrivKey,
	}

	_, err = p.sendDepositTransaction(context.Background(), acc.Backend, acc.PubKey[:])
	acc.Backend.Commit()
	if err != nil {
		t.Fatalf("Validator registration failed: %v", err)
	}
	log, err := acc.Contract.FilterValidatorRegistered(&bind.FilterOpts{}, [][32]byte{}, []common.Address{}, [][32]byte{})
	if err != nil {
		t.Fatal(err)
	}
	log.Next()
	if log.Event.PubKey != acc.PubKey {
		t.Errorf("validatorRegistered event public key mismatched. Want: %v, Got: %v", acc.PubKey, log.Event.PubKey)
	}
}
