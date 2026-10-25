package config

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const contractYamlKey = "noaa_recovery"

type Contract struct {
	Address common.Address `fig:"address,required"`
	// Block is the deployment block.
	Block uint64 `fig:"block"`
}

func (c *config) Contract() Contract {
	return c.contract.Do(func() any {
		var result Contract

		err := figure.
			Out(&result).
			With(figure.BaseHooks, figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, contractYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out config", logan.F{"key": contractYamlKey}))
		}

		return result
	}).(Contract)
}
