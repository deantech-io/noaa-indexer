package recovery

import (
	"math/big"
	"testing"

	"github.com/deantech-io/noaa-indexer/internal/contracts"
	"github.com/deantech-io/noaa-indexer/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestDecode(t *testing.T) {
	abi, err := contracts.NoAARecoveryMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}

	passport := common.HexToHash("0xaa")
	event := abi.Events["RecoveryRequested"]
	body, err := event.Inputs.Pack(passport, common.BigToHash(big.NewInt(42)), big.NewInt(1000))
	if err != nil {
		t.Fatal(err)
	}

	got, err := Decode(types.Log{
		Topics:      []common.Hash{event.ID},
		Data:        body,
		BlockNumber: 5,
		Index:       3,
	})
	if err != nil {
		t.Fatal(err)
	}

	if got.Type != data.EventRequested || got.PassportHash != data.Hash(passport) || got.NewIdentityKey.Int64() != 42 ||
		got.ReadyAt == nil || got.ReadyAt.Unix() != 1000 || got.BlockNumber != 5 || got.LogIndex != 3 {
		t.Errorf("unexpected event %+v", got)
	}

	if _, err := Decode(types.Log{Topics: []common.Hash{common.HexToHash("0x01")}}); err == nil {
		t.Error("expected error for unknown topic")
	}

	if len(Topics()) != len(eventTypes) {
		t.Errorf("topics = %d, want %d", len(Topics()), len(eventTypes))
	}
}
