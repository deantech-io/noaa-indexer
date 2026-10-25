package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer

	Network() Network
	Contract() Contract
	Indexer() Indexer
	Finalizer() Finalizer
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer
	getter kv.Getter

	network   comfig.Once
	contract  comfig.Once
	indexer   comfig.Once
	finalizer comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Databaser:  pgdb.NewDatabaser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
