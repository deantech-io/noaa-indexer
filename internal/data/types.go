package data

import (
	"database/sql/driver"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Hash is stored as lowercase 0x-prefixed hex TEXT.
type Hash common.Hash

func (h Hash) Value() (driver.Value, error) {
	return common.Hash(h).Hex(), nil
}

func (h *Hash) Scan(src interface{}) error {
	b, err := decodeHex(src, common.HashLength)
	if err != nil {
		return err
	}
	*h = Hash(common.BytesToHash(b))
	return nil
}

func (h Hash) Hex() string {
	return common.Hash(h).Hex()
}

// Address returns the stored form: lowercase 0x-prefixed hex.
func Address(a common.Address) string {
	return strings.ToLower(a.Hex())
}

// Numeric maps a uint256 to a NUMERIC(78,0) column. A nil Int is NULL.
type Numeric struct {
	*big.Int
}

func NewNumeric(i *big.Int) Numeric {
	return Numeric{Int: i}
}

// NumericFromBytes reads a bytes32 value as an unsigned big-endian integer.
func NumericFromBytes(b [32]byte) Numeric {
	return Numeric{Int: new(big.Int).SetBytes(b[:])}
}

func (n Numeric) Value() (driver.Value, error) {
	if n.Int == nil {
		return nil, nil
	}
	return n.String(), nil
}

func (n *Numeric) Scan(src interface{}) error {
	if src == nil {
		n.Int = nil
		return nil
	}

	s, err := asString(src)
	if err != nil {
		return err
	}
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return fmt.Errorf("invalid numeric %q", s)
	}
	n.Int = i
	return nil
}

func (n Numeric) IsZero() bool {
	return n.Int == nil || n.Sign() == 0
}

// Equal treats nil as 0.
func (n Numeric) Equal(other Numeric) bool {
	if n.IsZero() || other.IsZero() {
		return n.IsZero() == other.IsZero()
	}
	return n.Cmp(other.Int) == 0
}

// Time converts unix seconds to UTC. 0 means none.
func Time(unix uint64) *time.Time {
	if unix == 0 {
		return nil
	}
	t := time.Unix(int64(unix), 0).UTC()
	return &t
}

func decodeHex(src interface{}, size int) ([]byte, error) {
	s, err := asString(src)
	if err != nil {
		return nil, err
	}
	b, err := hexutil.Decode(s)
	if err != nil || len(b) != size {
		return nil, fmt.Errorf("invalid %d-byte hex %q", size, s)
	}
	return b, nil
}

func asString(src interface{}) (string, error) {
	switch v := src.(type) {
	case []byte:
		return string(v), nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("unsupported source type %T", src)
	}
}
