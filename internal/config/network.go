package config

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rarimo/noaa-indexer/internal/chain"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const networkYamlKey = "network"

type Network struct {
	ChainID uint64
	Client  *chain.Client
}

func (c *config) Network() Network {
	return c.network.Do(func() interface{} {
		cfg := struct {
			ChainID uint64        `fig:"chain_id,required"`
			RPC     []string      `fig:"rpc,required"`
			Timeout time.Duration `fig:"timeout"`
			Retries int           `fig:"retries"`
			Backoff time.Duration `fig:"backoff"`
		}{
			Timeout: 10 * time.Second,
			Retries: 3,
			Backoff: time.Second,
		}

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, networkYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out config", logan.F{"key": networkYamlKey}))
		}
		if cfg.Timeout <= 0 || cfg.Retries < 0 || cfg.Backoff < 0 {
			panic(errors.From(errors.New("timeout must be positive, retries and backoff not negative"), logan.F{"key": networkYamlKey}))
		}

		client, err := chain.Dial(cfg.RPC, cfg.Timeout, cfg.Retries, cfg.Backoff)
		if err != nil {
			panic(errors.Wrap(err, "failed to dial rpc"))
		}

		var chainID uint64
		err = client.Do(context.Background(), func(ctx context.Context, ec *ethclient.Client) error {
			id, err := ec.ChainID(ctx)
			if err == nil {
				chainID = id.Uint64()
			}
			return err
		})
		if err != nil {
			panic(errors.Wrap(err, "failed to get chain ID"))
		}
		if chainID != cfg.ChainID {
			panic(errors.From(errors.New("rpc chain ID mismatch"), logan.F{"expected": cfg.ChainID, "actual": chainID}))
		}

		return Network{ChainID: cfg.ChainID, Client: client}
	}).(Network)
}
