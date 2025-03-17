package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swarleynunez/NxGenT/core/utils"
	"math/big"
	"strconv"
	"sync"
)

var (
	// Unexported global variables
	_nonce  uint64
	_nmutex *sync.Mutex
)

func Connect(url string) (ethc *ethclient.Client) {

	ethc, err := ethclient.Dial(url)
	utils.CheckError(err, utils.ErrorMode)

	return
}

func LoadKeystore(keydir string) (ks *keystore.KeyStore) {

	ks = keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)

	return
}

func LoadAccount(ks *keystore.KeyStore, passphrase string) (from accounts.Account) {

	// Unlock the loaded account
	from = ks.Accounts()[0]
	err := ks.Unlock(from, passphrase)
	utils.CheckError(err, utils.ErrorMode)

	return
}

func Transactor(ctx context.Context, ethc *ethclient.Client, ks *keystore.KeyStore, from accounts.Account) *bind.TransactOpts {

	// Get and parse chain ID (transaction replay protection)
	chainID, err := strconv.ParseUint(utils.GetEnv("ETH_CHAIN_ID"), 10, 64)
	utils.CheckError(err, utils.ErrorMode)

	// Get transactor from keystore and account
	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, from, big.NewInt(int64(chainID)))
	utils.CheckError(err, utils.ErrorMode)

	// Get account's next nonce
	_nonce, err = ethc.PendingNonceAt(ctx, auth.From)
	utils.CheckError(err, utils.ErrorMode)
	_nmutex = &sync.Mutex{} // Mutex to synchronize access to the next nonce

	return auth
}

func GetNextNonce() (r *big.Int) {

	_nmutex.Lock()
	r = new(big.Int).SetUint64(_nonce)
	_nonce++
	_nmutex.Unlock()

	return
}
