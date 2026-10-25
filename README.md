# noaa-indexer

Indexer and finalizer for the `NoAARecovery` contract.

- **indexer** reads contract events up to the `finalized` block, stores them in `recovery_event`
  and keeps `recovery_state` (one row per passport hash) in sync.
- **finalizer** periodically finalizes pending recoveries whose `readyAt` has passed, via the
  registration relayer (or, on local anvil only, a hot wallet).

> **PROVISIONAL ABI.** The contract does not exist yet. The service is built against the assumed
> interface in [`contracts/src/INoAARecovery.sol`](contracts/src/INoAARecovery.sol). To switch to
> the real contract:
> 1. replace `internal/contracts/NoAARecovery.abi.json` with the real ABI;
> 2. `go generate ./internal/contracts/`;
> 3. adjust `internal/recovery` (decoder and reducer; see the `TODO(real ABI)` there).

## Commands

```
noaa-indexer run indexer       # index events, serve /health and /metrics
noaa-indexer run finalizer     # finalize ready recoveries, serve /health and /metrics
noaa-indexer migrate up|down
noaa-indexer rebuild           # recompute recovery_state from recovery_event
```

The config file is passed via `KV_VIPER_FILE`.

## Config

```yaml
log:
  level: info
  disable_sentry: true

db:
  url: postgres://...

listener:
  addr: :8000            # /health and /metrics

network:
  chain_id: 1            # must match the RPC, checked at start
  rpc:                   # primary first, then fallbacks
    - https://...
    - https://...
  timeout: 10s           # per call
  retries: 3             # rounds over all endpoints, exponential backoff
  backoff: 1s

noaa_recovery:
  address: 0x...
  block: 0               # deployment block, indexing starts here

indexer:
  poll_period: 5s
  chunk_size: 5000       # max eth_getLogs range; halved on RPC errors, grows back

finalizer:
  period: 30s
  mode: relayer          # relayer | wallet
  unknown_retry_after: 0 # retry after an unknown outcome or dropped tx; 0 = never
  retry_base: 1m         # backoff after a failed, rejected or expired attempt, doubled each time
  retry_max: 1h
  max_attempts: 5        # attempts per request, then it is no longer retried
  relayer_url: https://.../integrations/registration-relayer/v1/register
  relayer_timeout: 30s
  wallet_private_key:    # wallet mode only
  wallet_allowed_chain_ids: []  # wallet mode refuses to start on other chains
```

In relayer mode the NoAARecovery address must be in the relayer's `whitelist`: it is sent as
`destination`.

## How it works

**Indexer.** Each step reads the `finalized` header and fetches logs for the next range after the
checkpoint with chunked `eth_getLogs`. In one transaction it stores the new events, applies them to
`recovery_state` and moves the checkpoint. Events that are already stored are skipped. Before each
range it checks that the next block's parent hash equals the checkpoint hash. On a mismatch the
indexer stops: `noaa_indexer_halted` becomes 1, `/health` returns 503, and there is no automatic
recovery. Events the contract should not have emitted in the stored state are logged as warnings
and counted in `noaa_indexer_unexpected_transitions_total{type}`; indexing continues.

**State.** `recovery_state.status` is one of `pending`, `contested`, `locked`, `active`. A contest
lock expires with time, so readers should use `effective_status` from the
`recovery_state_effective` view. A lock keeps a pending request and blocks its finalization until
unlock (provisional).

**Finalizer.** Every period it takes `pending` rows with `ready_at` at or before the latest block
time, re-checks `requests()` and `lockedUntil()` on chain, and submits `finalizeRecovery`. Each
attempt is a `finalizer_submission` row for the request event. A unique index allows one
blocking attempt per request, and a passport has at most one open (`submitting`, `submitted`,
`unknown`) attempt. After a `failed`, `rejected` or `expired` attempt the request is retried with
exponential backoff from `retry_base` to `retry_max`. After `max_attempts` it is logged at error
level, counted in `noaa_finalizer_exhausted_requests_total` and no longer retried.

| status | meaning | blocks another attempt |
|---|---|---|
| submitted | tx hash known, waiting for the receipt | until the request is gone on chain, or `unknown_retry_after` passes (then `expired`) |
| confirmed | finalized | yes |
| failed | tx reverted, nothing applied | no |
| rejected | nothing was sent (relayer 4xx, gas estimation failed) | no |
| unknown | timeout or 5xx, may have been sent; logged at error level | until the request is gone on chain, or `unknown_retry_after` passes (then `expired`) |
| resolved | request finalized or cancelled by someone else | no |

Run a single finalizer instance: on start-up it marks leftover `submitting` rows as `unknown`.

## Database

| table | contents |
|---|---|
| `block` | finalized blocks referenced by events or the checkpoint |
| `transaction` | transactions of stored events |
| `contract` | indexed deployments (`chain_id`, `address`); `block_id` is the checkpoint |
| `recovery_event` | decoded events |
| `recovery_state` | one row per passport hash, rebuildable from `recovery_event` |
| `finalizer_submission` | finalization attempts |

Hashes and addresses are lowercase hex `TEXT`, times are UTC `TIMESTAMP(3)`, and statuses are
enums. Rows are never deleted; every table has `is_active`.

## Metrics

`noaa_indexer_indexed_block`, `noaa_indexer_head_block`, `noaa_indexer_lag_blocks`,
`noaa_indexer_events_total{type}`, `noaa_indexer_unexpected_transitions_total{type}`,
`noaa_indexer_halted`, `noaa_finalizer_submissions_total{outcome}`,
`noaa_finalizer_exhausted_requests_total`, `noaa_finalizer_unknown_submissions`.

Hourly rates: `increase(noaa_indexer_events_total{type="requested"}[1h])`, same with
`type="contested"`. Note that the counters also count events found during a backfill.

## Local development

Requires Docker. Everything runs on a local anvil. Nothing is persisted: anvil and Postgres start
empty on every restart.

```sh
docker compose up --build
```

| service | role |
|---|---|
| postgres | database (not published to the host; `docker compose exec postgres psql -U noaa`) |
| anvil | local chain, `localhost:8545`, a block every 2 s |
| deployer | one-shot: deploys `MockNoAARecovery` and writes `/shared/config.yaml` from `docker/config.template.yaml` |
| migrate | one-shot: `migrate up` |
| indexer | `localhost:8000/health`, `/metrics` |
| finalizer | `localhost:8001/health`, `/metrics`; relayer mode via mock-relayer |
| mock-relayer | stand-in for the registration relayer's `/register` endpoint, sends to anvil |

The local recovery delay is 60 s and the contest lock 30 days.

### Emitting test events

The mock has no proofs: anyone can trigger any transition.

```sh
RPC=http://localhost:8545
KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80   # anvil account #0
NOAA=$(docker compose exec indexer /bin/sh -c '. /shared/deploy.env && echo $NOAA_ADDRESS')
P=$(cast keccak passport-1)

send() { cast send -q --rpc-url $RPC --private-key $KEY $NOAA "$@"; }

# Finalization: request, skip past readyAt, then move anvil's finalized block (latest - 64) forward.
send "requestRecovery(bytes32,uint256)" $P 111
cast rpc --rpc-url $RPC evm_increaseTime 61
cast rpc --rpc-url $RPC anvil_mine 70

# Contest: the second request must come while the first is still pending,
# i.e. before readyAt, or the finalizer will have finalized it already.
P2=$(cast keccak passport-2)
send "requestRecovery(bytes32,uint256)" $P2 111
send "requestRecovery(bytes32,uint256)" $P2 222

# Lock, unlock, cancel
P3=$(cast keccak passport-3)
send "requestRecovery(bytes32,uint256)" $P3 111
send "lockRecovery(bytes32)" $P3
send "unlockRecovery(bytes32)" $P3
send "cancelRecovery(bytes32)" $P3

cast rpc --rpc-url $RPC anvil_mine 70
docker compose exec postgres psql -U noaa -c "select * from recovery_state_effective"
docker compose exec postgres psql -U noaa -c "select * from finalizer_submission"
```

To try wallet mode, set `finalizer.mode: wallet` and `wallet_private_key` in
`docker/config.template.yaml`.

### Tests

```sh
go test ./...                                    # unit tests (reducer, decoder)
TEST_DB_URL=postgres://user@host:port/db?sslmode=disable \
  go test -tags integration -p 1 ./tests/integration/
```

Integration tests start their own `anvil` (must be in `PATH`) and reset the schema of the given
database. They cover backfill, restart from checkpoint, idempotent re-indexing, rebuild, the
parent-hash halt, the header-hash check for Prague headers, and the finalizer (wallet and relayer
mode, state changed before submit, unknown outcome not retried, relayer rejection, retry backoff
and attempt cap, one open attempt per passport), config checks (chain ID mismatch, wallet mode on
an unlisted chain) and RPC errors without endpoint URLs.

Regenerating bindings: `go generate ./internal/contracts/` (provisional ABI) and
`go generate ./tests/integration/mock/` (mock with bytecode, needs `forge`).
