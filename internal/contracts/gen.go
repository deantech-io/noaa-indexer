// Package contracts holds abigen bindings for NoAARecovery.
// NoAARecovery.abi.json is the ABI from the compiled NoAARecovery contract.
package contracts

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.16.9 --abi NoAARecovery.abi.json --pkg contracts --type NoAARecovery --out noaa_recovery.go
