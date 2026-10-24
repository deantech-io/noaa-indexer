// Command mockrelayer is a local stand-in for the registration relayer /register endpoint.
//
// Env: MOCK_RELAYER_ADDR, MOCK_RELAYER_RPC, MOCK_RELAYER_PRIVATE_KEY, MOCK_RELAYER_WHITELIST.
package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type relayer struct {
	client    *ethclient.Client
	signer    types.Signer
	key       *ecdsa.PrivateKey
	from      common.Address
	whitelist map[common.Address]bool
	mu        sync.Mutex // serializes nonces
}

type request struct {
	Data struct {
		TxData      string  `json:"tx_data"`
		Destination *string `json:"destination"`
		NoSend      bool    `json:"no_send"`
	} `json:"data"`
}

func main() {
	key, err := crypto.HexToECDSA(os.Getenv("MOCK_RELAYER_PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("invalid private key: %v", err)
	}
	client, err := ethclient.Dial(os.Getenv("MOCK_RELAYER_RPC"))
	if err != nil {
		log.Fatalf("failed to dial rpc: %v", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("failed to get chain id: %v", err)
	}

	whitelist := make(map[common.Address]bool)
	for _, a := range strings.Split(os.Getenv("MOCK_RELAYER_WHITELIST"), ",") {
		if a = strings.TrimSpace(a); a != "" {
			whitelist[common.HexToAddress(a)] = true
		}
	}

	r := &relayer{
		client:    client,
		signer:    types.LatestSignerForChainID(chainID),
		key:       key,
		from:      crypto.PubkeyToAddress(key.PublicKey),
		whitelist: whitelist,
	}

	http.HandleFunc("POST /integrations/registration-relayer/v1/register", r.register)

	addr := os.Getenv("MOCK_RELAYER_ADDR")
	log.Printf("mock relayer listening on %s, sender %s", addr, r.from.Hex())
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (r *relayer) register(w http.ResponseWriter, req *http.Request) {
	var body request
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil || body.Data.Destination == nil {
		writeError(w, http.StatusBadRequest, "invalid request")
		return
	}
	to := common.HexToAddress(*body.Data.Destination)
	if !r.whitelist[to] {
		writeError(w, http.StatusBadRequest, "specified contract address not allowed")
		return
	}
	calldata, err := hexutil.Decode(body.Data.TxData)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid tx_data")
		return
	}

	tx, status, err := r.send(req.Context(), to, calldata, body.Data.NoSend)
	if err != nil {
		log.Printf("register failed: %v", err)
		writeError(w, status, err.Error())
		return
	}
	log.Printf("sent tx %s to %s", tx.Hash().Hex(), to.Hex())

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"data": map[string]any{
		"id":         tx.Hash().Hex(),
		"type":       "txs",
		"attributes": map[string]any{"tx_hash": tx.Hash().Hex()},
	}})
}

// Like the real relayer: 400 if gas estimation fails, 500 after that.
func (r *relayer) send(ctx context.Context, to common.Address, calldata []byte, noSend bool) (*types.Transaction, int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	gas, err := r.client.EstimateGas(ctx, ethereum.CallMsg{From: r.from, To: &to, Data: calldata})
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	gasPrice, err := r.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	nonce, err := r.client.PendingNonceAt(ctx, r.from)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	tx, err := types.SignNewTx(r.key, r.signer, &types.LegacyTx{
		Nonce:    nonce,
		Gas:      gas,
		GasPrice: gasPrice,
		To:       &to,
		Data:     calldata,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if noSend {
		return tx, 0, nil
	}
	if err = r.client.SendTransaction(ctx, tx); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return tx, 0, nil
}

func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"errors": []map[string]any{{"title": http.StatusText(status), "detail": msg}}})
}
