package finalizer

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/deantech-io/noaa-indexer/internal/chain"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// wallet sends transactions with a hot key. Local anvil only.
type wallet struct {
	client  *chain.Client
	key     *ecdsa.PrivateKey
	from    common.Address
	to      common.Address
	chainID *big.Int
}

func newWallet(client *chain.Client, key *ecdsa.PrivateKey, chainID uint64, to common.Address) wallet {
	return wallet{
		client:  client,
		key:     key,
		from:    crypto.PubkeyToAddress(key.PublicKey),
		to:      to,
		chainID: new(big.Int).SetUint64(chainID),
	}
}

func (w wallet) submit(ctx context.Context, calldata []byte) (common.Hash, error) {
	var (
		gas      uint64
		gasPrice *big.Int
		nonce    uint64
	)
	err := w.client.Do(ctx, func(ctx context.Context, ec *ethclient.Client) (err error) {
		if gas, err = ec.EstimateGas(ctx, ethereum.CallMsg{From: w.from, To: &w.to, Data: calldata}); err != nil {
			return err
		}
		if gasPrice, err = ec.SuggestGasPrice(ctx); err != nil {
			return err
		}
		nonce, err = ec.PendingNonceAt(ctx, w.from)
		return err
	})
	if err != nil {
		return common.Hash{}, errors.Wrap(errRejected, err.Error())
	}

	tx, err := types.SignNewTx(w.key, types.LatestSignerForChainID(w.chainID), &types.LegacyTx{
		Nonce:    nonce,
		Gas:      gas,
		GasPrice: gasPrice,
		To:       &w.to,
		Data:     calldata,
	})
	if err != nil {
		return common.Hash{}, errors.Wrap(errRejected, err.Error())
	}

	// Resending the same signed transaction is safe.
	err = w.client.Do(ctx, func(ctx context.Context, ec *ethclient.Client) error {
		return ec.SendTransaction(ctx, tx)
	})
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to send transaction")
	}
	return tx.Hash(), nil
}
