# noaa-indexer

Indexer and finalizer for the `NoAARecovery` contract.

- **indexer** reads contract events up to the `finalized` block, stores them in `recovery_event`
  and keeps `recovery_state` (one row per passport hash) equal to the contract storage.
- **finalizer** periodically finalizes pending recoveries whose `readyAt` has passed, via the
  registration relayer (or, on local anvil only, a hot wallet).

The bindings in `internal/contracts` are generated from the `NoAARecovery` ABI.

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
  addr: :8000 # /health and /metrics

network:
  chain_id: 1 # must match the RPC, checked at start
  rpc: # primary first, then fallbacks
    - https://...
    - https://...
  timeout: 10s # per call
  retries: 3 # rounds over all endpoints, exponential backoff
  backoff: 1s

noaa_recovery:
  address: 0x...
  block: 0 # deployment block, indexing starts here

indexer:
  poll_period: 5s
  chunk_size: 5000 # max eth_getLogs range; halved on RPC errors, grows back

finalizer:
  period: 30s
  mode: relayer # relayer | wallet
  unknown_retry_after: 0 # retry after an unknown outcome or dropped tx; 0 = never
  retry_base: 1m # backoff after a failed, rejected or expired attempt, doubled each time
  retry_max: 1h
  max_attempts: 5 # attempts per request, then it is no longer retried
  relayer_url: https://.../integrations/registration-relayer/v1/register
  relayer_timeout: 30s
  wallet_private_key: # wallet mode only
  wallet_allowed_chain_ids: [] # wallet mode refuses to start on other chains
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

Events do not carry everything the contract stores, so the indexer reads it at the event's block:
`getRecoveryRequest` for `RecoveryRequested` (old identity key and DG1 commitment) and
`lockedUntil` for `RecoveryContested` and `RecoveryCancelled`. Catching up only reads recent
blocks, but a backfill from an old deployment block needs an RPC that serves historical state
(archive). If a request is dropped later in the same block, its old identity key and commitment
stay NULL.

**State.** `recovery_state` mirrors the contract storage of each passport: the pending request
(`old_identity_key`, `new_identity_key`, `dg_commit`, `ready_at`, all NULL without one),
`locked_until` (NULL for 0) and `is_recovery_locked`. On chain:

- a request is possible when the passport is not locked and `locked_until` has passed;
- a second request with another key contests the first: the request is dropped and
  `locked_until` is set to now + `lockDuration`;
- a cancel drops the request and sets `locked_until` the same way;
- a lock drops the request (emitting `RecoveryCancelled`) without touching `locked_until`, and
  lasts until unlock;
- anyone can finalize a request once `ready_at` has passed.

The `recovery_state_detail` view adds the chain, the contract address and the block of the last
applied event.

**Finalizer.** Every period it takes pending requests with `ready_at` at or before the latest
block time, re-checks `getRecoveryRequest()` on chain, and submits `finalizeRecovery`. The
contract also requires the passport's active identity to be unchanged; otherwise the attempt is
`rejected` and stops after `max_attempts`. Each
attempt is a `finalizer_submission` row for the request event. A unique index allows one
blocking attempt per request, and a passport has at most one open (`submitting`, `submitted`,
`unknown`) attempt. After a `failed`, `rejected` or `expired` attempt the request is retried with
exponential backoff from `retry_base` to `retry_max`. After `max_attempts` it is logged at error
level, counted in `noaa_finalizer_exhausted_requests_total` and no longer retried.

| status    | meaning                                                   | blocks another attempt                                                               |
| --------- | --------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| submitted | tx hash known, waiting for the receipt                    | until the request is gone on chain, or `unknown_retry_after` passes (then `expired`) |
| confirmed | finalized                                                 | yes                                                                                  |
| failed    | tx reverted, nothing applied                              | no                                                                                   |
| rejected  | nothing was sent (relayer 4xx, gas estimation failed)     | no                                                                                   |
| unknown   | timeout or 5xx, may have been sent; logged at error level | until the request is gone on chain, or `unknown_retry_after` passes (then `expired`) |
| resolved  | request finalized or cancelled by someone else            | no                                                                                   |

Run a single finalizer instance: on start-up it marks leftover `submitting` rows as `unknown`.

## Database

| table                  | contents                                                                  |
| ---------------------- | ------------------------------------------------------------------------- |
| `block`                | finalized blocks referenced by events or the checkpoint                   |
| `transaction`          | transactions of stored events                                             |
| `contract`             | indexed deployments (`chain_id`, `address`); `block_id` is the checkpoint |
| `recovery_event`       | decoded events                                                            |
| `recovery_state`       | one row per passport hash, rebuildable from `recovery_event`              |
| `finalizer_submission` | finalization attempts                                                     |

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

| service      | role                                                                                                     |
| ------------ | -------------------------------------------------------------------------------------------------------- |
| postgres     | database (not published to the host; `docker compose exec postgres psql -U noaa`)                        |
| anvil        | local chain, `localhost:8545`, a block every 2 s                                                         |
| deployer     | one-shot: deploys `MockNoAARecovery` and writes `/shared/config.yaml` from `docker/config.template.yaml` |
| migrate      | one-shot: `migrate up`                                                                                   |
| indexer      | `localhost:8000/health`, `/metrics`                                                                      |
| finalizer    | `localhost:8001/health`, `/metrics`; relayer mode via mock-relayer                                       |
| mock-relayer | stand-in for the registration relayer's `/register` endpoint, sends to anvil                             |

The local recovery delay is 60 s and the lock duration 30 days. Mock passports start with the
active identity `1`.

### Emitting test events

The mock has no proofs: anyone can trigger any transition.

```sh
RPC=http://localhost:8545
KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80   # anvil account #0
NOAA=$(docker compose exec indexer /bin/sh -c '. /shared/deploy.env && echo $NOAA_ADDRESS')
Z=0x0000000000000000000000000000000000000000000000000000000000000000

send() { cast send -q --rpc-url $RPC --private-key $KEY $NOAA "$@"; }
# request <passport> <identity key>: certificates root, zk type and proof are ignored by the mock.
request() { send "requestRecovery(bytes32,uint256,uint256,bytes32,bytes32,bytes)" $Z $2 0 $1 $Z 0x; }

# Finalization: request, skip past readyAt, then move anvil's finalized block (latest - 64) forward.
P1=$(cast keccak passport-1)
request $P1 111
cast rpc --rpc-url $RPC evm_increaseTime 61
cast rpc --rpc-url $RPC anvil_mine 70

# Contest: the second request must come while the first is still pending,
# i.e. before readyAt, or the finalizer will have finalized it already.
P2=$(cast keccak passport-2)
request $P2 111
request $P2 222

# Cancel
P3=$(cast keccak passport-3)
request $P3 111
send "cancelRecovery(bytes32,bytes)" $P3 0x

# Lock drops the request, unlock
P4=$(cast keccak passport-4)
request $P4 111
send "lockRecovery(bytes32,bytes)" $P4 0x
send "unlockRecovery(bytes32,bytes)" $P4 0x

cast rpc --rpc-url $RPC anvil_mine 70
docker compose exec postgres psql -U noaa -c "select * from recovery_state_detail"
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
database. They cover backfill with the state compared to the contract storage, a request dropped
in its own block, restart from checkpoint, idempotent re-indexing, rebuild, the
parent-hash halt, the header-hash check for Prague headers, and the finalizer (wallet and relayer
mode, state changed before submit, unknown outcome not retried, relayer rejection, retry backoff
and attempt cap, one open attempt per passport), config checks (chain ID mismatch, wallet mode on
an unlisted chain) and RPC errors without endpoint URLs.

Regenerating bindings: `go generate ./internal/contracts/` (provisional ABI) and
`go generate ./tests/integration/mock/` (mock with bytecode, needs `forge`).
