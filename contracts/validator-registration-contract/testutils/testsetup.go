package testutils

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	vrc "github.com/prysmaticlabs/prysm/contracts/validator-registration-contract"
)

// TestAccount container
type TestAccount struct {
	Addr              common.Address
	WithdrawalAddress common.Address
	RandaoCommitment  [32]byte
	PubKey            [32]byte
	Contract          *vrc.ValidatorRegistration
	Backend           *backends.SimulatedBackend
	TxOpts            *bind.TransactOpts
	PrivKey           *ecdsa.PrivateKey
	ContractAddr      common.Address
}

// Setup test account with a deployed validator registration contract.
func Setup() (*TestAccount, error) {
	genesis := make(core.GenesisAlloc)
	privKey, _ := crypto.GenerateKey()
	pubKeyECDSA, ok := privKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	// strip off the 0x and the first 2 characters 04 which is always the EC prefix and is not required.
	publicKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)[4:]
	var pubKey [32]byte
	copy(pubKey[:], []byte(publicKeyBytes))

	addr := crypto.PubkeyToAddress(privKey.PublicKey)
	txOpts := bind.NewKeyedTransactor(privKey)
	startingBalance, _ := new(big.Int).SetString("100000000000000000000", 10)
	genesis[addr] = core.GenesisAccount{Balance: startingBalance}
	backend := backends.NewSimulatedBackend(genesis, 2100000)

	contractAddr, _, contract, err := vrc.DeployValidatorRegistration(txOpts, backend)
	if err != nil {
		return nil, err
	}

	backend.Commit()
	return &TestAccount{addr, common.Address{}, [32]byte{}, pubKey, contract, backend, txOpts, privKey, contractAddr}, nil
}
