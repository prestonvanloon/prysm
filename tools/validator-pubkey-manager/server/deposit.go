package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/opentracing/opentracing-go"
	vrc "github.com/prysmaticlabs/prysm/contracts/validator-registration-contract"
)

var (
	amount32Eth, _ = new(big.Int).SetString("32000000000000000000", 10)
)

type powchainclient struct {
	httpPath        string
	priv            *ecdsa.PrivateKey
	contractAddress common.Address
}

func newPowchainclient(httpPath, address, privKey string) *powchainclient {
	priv, err := crypto.HexToECDSA(privKey)
	if err != nil {
		panic(err)
	}

	return &powchainclient{
		contractAddress: common.HexToAddress(address),
		httpPath:        httpPath,
		priv:            priv,
	}
}

func (p *powchainclient) Deposit(ctx context.Context, pubkey []byte) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "deposit_validator")
	defer span.Finish()

	fmt.Println("dialing RPC")
	client, err := p.dialRPC(ctx)
	if err != nil {
		return err
	}
	fmt.Println("depositing transaction")
	tx, err := p.sendDepositTransaction(ctx, client, pubkey)
	if err != nil {
		return err
	}
	fmt.Println("waiting for completion")
	if err := p.waitForTransaction(ctx, client, tx); err != nil {
		return err
	}

	return nil
}

func (p *powchainclient) dialRPC(ctx context.Context) (*ethclient.Client, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "dial_rpc")
	defer span.Finish()

	rpcClient, err := rpc.Dial(p.httpPath)
	if err != nil {
		return nil, err
	}

	return ethclient.NewClient(rpcClient), nil
}

func (p *powchainclient) sendDepositTransaction(ctx context.Context, backend bind.ContractBackend, pubkey []byte) (*types.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "send_deposit_transaction")
	defer span.Finish()

	contract, err := vrc.NewValidatorRegistration(p.contractAddress, backend)
	if err != nil {
		return nil, err
	}

	txOps := bind.NewKeyedTransactor(p.priv)
	txOps.Value = amount32Eth

	var pkey [32]byte
	copy(pkey[:], pubkey)

	t, err := contract.UsedPubkey(&bind.CallOpts{}, pkey)
	if err != nil {
		return nil, err
	}
	if t {
		return nil, fmt.Errorf("pubkey %x is already deposited", pubkey)
	}

	withdrawalShardID := big.NewInt(99)
	withdrawalAddress := common.Address{'A', 'D', 'D', 'R', 'E', 'S', 'S'}
	randaoCommitment := [32]byte{'S', 'H', 'H', 'H', 'H', 'I', 'T', 'S', 'A', 'S', 'E', 'C', 'R', 'E', 'T'}

	return contract.Deposit(txOps, pkey, withdrawalShardID, withdrawalAddress, randaoCommitment)
}

func (p *powchainclient) waitForTransaction(ctx context.Context, client *ethclient.Client, tx *types.Transaction) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wait_for_transaction")
	defer span.Finish()

	var err error
	for pending := true; pending; _, pending, err = client.TransactionByHash(ctx, tx.Hash()) {
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	r, err := client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return err
	}

	if r.Status != types.ReceiptStatusSuccessful {
		rJSON, _ := json.Marshal(r)
		tJSON, _ := tx.MarshalJSON()
		return fmt.Errorf("Transaction failed. Transaction: %s Receipt: %s", tJSON, rJSON)
	}

	return nil
}
