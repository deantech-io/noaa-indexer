package chain

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestTransportErrorHidesEndpoint(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := l.Addr().String()
	_ = l.Close()

	c, err := Dial([]string{"http://" + addr + "/secret-key"}, time.Second, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	err = c.Do(context.Background(), func(ctx context.Context, ec *ethclient.Client) error {
		_, err := ec.BlockNumber(ctx)
		return err
	})
	if err == nil || err.Error() != "rpc[0]: request failed" {
		t.Fatalf("error = %v, want rpc[0]: request failed", err)
	}
}
