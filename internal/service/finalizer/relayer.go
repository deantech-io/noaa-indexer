package finalizer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
		// Transport errors carry the relayer URL and IPs.
		if uerr, ok := err.(*url.Error); ok && uerr.Timeout() {
			return common.Hash{}, errors.New("relayer request timed out")
		}
		return common.Hash{}, errors.New("relayer request failed")
	}
	defer func() { _ = resp.Body.Close() }()

	// The body carries the revert reason, so it is kept in the error.
	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
	detail := fmt.Sprintf("status %d: %s", resp.StatusCode, respBody)

	// The relayer returns 4xx before signing.
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return common.Hash{}, errors.Wrap(errRejected, "relayer rejected request: "+detail)
	}
	if resp.StatusCode != http.StatusOK {
		return common.Hash{}, errors.New("unexpected relayer response: " + detail)
	}

	var parsed relayerResponse
	if err = json.Unmarshal(respBody, &parsed); err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to decode relayer response: "+detail)
	}
	hash, err := hexutil.Decode(parsed.Data.Attributes.TxHash)
	if err != nil || len(hash) != common.HashLength {
		return common.Hash{}, errors.New("invalid tx hash in relayer response: " + detail)
	}
	return common.BytesToHash(hash), nil
}
