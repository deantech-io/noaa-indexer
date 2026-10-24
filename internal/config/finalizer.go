package config

import (
	"crypto/ecdsa"
	"slices"
	"time"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const finalizerYamlKey = "finalizer"

const (
	ModeRelayer = "relayer"
	ModeWallet  = "wallet"
)

type Finalizer struct {
	Period time.Duration `fig:"period"`
	Mode   string        `fig:"mode,required"`
	// UnknownRetryAfter allows a new attempt after an unknown outcome or a dropped transaction.
	// 0 never retries.
	UnknownRetryAfter time.Duration `fig:"unknown_retry_after"`

	RelayerURL     string        `fig:"relayer_url"`
	RelayerTimeout time.Duration `fig:"relayer_timeout"`

	// Wallet mode is for local anvil only.
	WalletPrivateKey      *ecdsa.PrivateKey `fig:"wallet_private_key"`
	WalletAllowedChainIDs []int64           `fig:"wallet_allowed_chain_ids"`
}

func (c *config) Finalizer() Finalizer {
	return c.finalizer.Do(func() interface{} {
		result := Finalizer{
			Period:         30 * time.Second,
			RelayerTimeout: 30 * time.Second,
		}

		err := figure.
			Out(&result).
			With(figure.BaseHooks, figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, finalizerYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out config", logan.F{"key": finalizerYamlKey}))
		}

		if result.Period <= 0 || result.RelayerTimeout <= 0 || result.UnknownRetryAfter < 0 {
			panic(errors.From(errors.New("period and relayer_timeout must be positive"), logan.F{"key": finalizerYamlKey}))
		}

		switch result.Mode {
		case ModeRelayer:
			if result.RelayerURL == "" {
				panic(errors.New("finalizer.relayer_url is required in relayer mode"))
			}
		case ModeWallet:
			if result.WalletPrivateKey == nil {
				panic(errors.New("finalizer.wallet_private_key is required in wallet mode"))
			}
			chainID := int64(c.Network().ChainID)
			if !slices.Contains(result.WalletAllowedChainIDs, chainID) {
				panic(errors.From(errors.New("wallet mode is not allowed on this chain"), logan.F{"chain_id": chainID}))
			}
		default:
			panic(errors.From(errors.New("unknown finalizer mode"), logan.F{"mode": result.Mode}))
		}

		return result
	}).(Finalizer)
}
