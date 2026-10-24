package recovery

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/noaa-indexer/internal/contracts"
	"github.com/rarimo/noaa-indexer/internal/data"
)

func TestDecode(t *testing.T) {
	abi, err := contracts.NoAARecoveryMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}

	passport := common.HexToHash("0xaa")
	event := abi.Events["RecoveryRequested"]
	body, err := event.Inputs.NonIndexed().Pack(big.NewInt(42), uint64(1000))
	if err != nil {
		t.Fatal(err)
	}

	got, err := Decode(types.Log{
		Topics:      []common.Hash{event.ID, passport},
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
