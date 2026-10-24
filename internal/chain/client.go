package chain

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client calls a primary RPC endpoint with fallbacks.
type Client struct {
	endpoints []*ethclient.Client
	timeout   time.Duration
	retries   int
	backoff   time.Duration
}

func Dial(urls []string, timeout time.Duration, retries int, backoff time.Duration) (*Client, error) {
	c := &Client{timeout: timeout, retries: retries, backoff: backoff}
	for _, url := range urls {
		ec, err := ethclient.Dial(url)
		if err != nil {
			return nil, err
		}
		c.endpoints = append(c.endpoints, ec)
	}
	if len(c.endpoints) == 0 {
		return nil, errors.New("no rpc endpoints")
	}
	return c, nil
}

// Do runs f on each endpoint until one succeeds, retrying with exponential backoff.
// Errors returned by the node itself are not retried. f must be safe to repeat.
func (c *Client) Do(ctx context.Context, f func(ctx context.Context, ec *ethclient.Client) error) error {
	var err error
	backoff := c.backoff
	for attempt := 0; attempt <= c.retries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(backoff):
			}
			backoff *= 2
		}

		for _, ec := range c.endpoints {
			if err = c.call(ctx, ec, f); err == nil || isFinal(err) || ctx.Err() != nil {
				return err
			}
		}
	}
	return err
}

func (c *Client) call(ctx context.Context, ec *ethclient.Client, f func(context.Context, *ethclient.Client) error) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	return f(ctx, ec)
}

func isFinal(err error) bool {
	var rpcErr rpc.Error
	return errors.As(err, &rpcErr) || errors.Is(err, ethereum.NotFound)
}
