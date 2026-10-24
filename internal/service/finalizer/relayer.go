package finalizer

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type relayer struct {
	url         string
	destination common.Address
	client      *http.Client
}

func newRelayer(url string, timeout time.Duration, destination common.Address) relayer {
	return relayer{url: url, destination: destination, client: &http.Client{Timeout: timeout}}
}

type relayerRequest struct {
	Data struct {
		TxData      string `json:"tx_data"`
		Destination string `json:"destination"`
		NoSend      bool   `json:"no_send"`
	} `json:"data"`
}

type relayerResponse struct {
	Data struct {
		Attributes struct {
			TxHash string `json:"tx_hash"`
		} `json:"attributes"`
	} `json:"data"`
}

// submit does not retry: the relayer signs a new transaction per call.
func (r relayer) submit(ctx context.Context, calldata []byte) (common.Hash, error) {
	var body relayerRequest
	body.Data.TxData = hexutil.Encode(calldata)
	body.Data.Destination = r.destination.Hex()

	raw, err := json.Marshal(body)
	if err != nil {
		return common.Hash{}, errors.Wrap(errRejected, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.url, bytes.NewReader(raw))
	if err != nil {
		return common.Hash{}, errors.Wrap(errRejected, err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		// Keep the relayer URL out of logs and the database.
		if uerr, ok := err.(*url.Error); ok {
			err = uerr.Err
		}
		return common.Hash{}, errors.Wrap(err, "relayer request failed")
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
	fields := logan.F{"status": resp.StatusCode, "body": string(respBody)}

	// The relayer returns 4xx before signing.
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return common.Hash{}, errors.Wrap(errRejected, "relayer rejected request", fields)
	}
	if resp.StatusCode != http.StatusOK {
		return common.Hash{}, errors.From(errors.New("unexpected relayer response"), fields)
	}

	var parsed relayerResponse
	if err = json.Unmarshal(respBody, &parsed); err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to decode relayer response", fields)
	}
	hash, err := hexutil.Decode(parsed.Data.Attributes.TxHash)
	if err != nil || len(hash) != common.HashLength {
		return common.Hash{}, errors.From(errors.New("invalid tx hash in relayer response"), fields)
	}
	return common.BytesToHash(hash), nil
}
