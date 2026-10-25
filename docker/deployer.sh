#!/bin/sh
set -eu

# forge writes build output next to the sources, which are mounted read-only.
cp -r /contracts /tmp/contracts
cd /tmp/contracts

forge script script/Deploy.s.sol --rpc-url "$RPC_URL" --broadcast

. "$DEPLOY_OUT"
sed -e "s|__NOAA_ADDRESS__|$NOAA_ADDRESS|" -e "s|__NOAA_BLOCK__|$NOAA_BLOCK|" /config.template.yaml > /shared/config.yaml
echo "deployed MockNoAARecovery at $NOAA_ADDRESS (block $NOAA_BLOCK)"
