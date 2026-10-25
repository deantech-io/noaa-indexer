package config

import (
	"time"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const indexerYamlKey = "indexer"

type Indexer struct {
	PollPeriod time.Duration `fig:"poll_period"`
	// ChunkSize is the max eth_getLogs range.
	ChunkSize uint64 `fig:"chunk_size"`
}

func (c *config) Indexer() Indexer {
	return c.indexer.Do(func() any {
		result := Indexer{
			PollPeriod: 5 * time.Second,
			ChunkSize:  5000,
		}

		err := figure.
			Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, indexerYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out config", logan.F{"key": indexerYamlKey}))
		}
		if result.ChunkSize == 0 || result.PollPeriod <= 0 {
			panic(errors.From(errors.New("chunk_size and poll_period must be positive"), logan.F{"key": indexerYamlKey}))
		}

		return result
	}).(Indexer)
}
