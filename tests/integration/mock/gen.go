// Package mock holds test bindings for MockNoAARecovery.
package mock

//go:generate sh -c "forge inspect --root ../../../contracts MockNoAARecovery abi --json > MockNoAARecovery.abi.json && forge inspect --root ../../../contracts MockNoAARecovery bytecode > MockNoAARecovery.bin"
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.16.9 --abi MockNoAARecovery.abi.json --bin MockNoAARecovery.bin --pkg mock --type MockNoAARecovery --out mock.go
