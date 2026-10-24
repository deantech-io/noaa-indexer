// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {INoAARecovery} from "./INoAARecovery.sol";

// Local mock without proofs: anyone can trigger any transition.
contract MockNoAARecovery is INoAARecovery {
    struct Request {
        uint256 newIdentityKey;
        uint64 readyAt;
    }

    uint64 public immutable recoveryDelay;
    uint64 public immutable contestLock;

    mapping(bytes32 => Request) public requests;
    mapping(bytes32 => uint64) public lockedUntil;
    mapping(bytes32 => bool) public locked;
    mapping(bytes32 => uint256) public identityKey;

    constructor(uint64 recoveryDelay_, uint64 contestLock_) {
        recoveryDelay = recoveryDelay_;
        contestLock = contestLock_;
    }

    // A different key while a request is pending contests it.
    function requestRecovery(bytes32 passportHash, uint256 newIdentityKey) external {
        require(newIdentityKey != 0, "zero key");
        require(!locked[passportHash], "locked");
        require(lockedUntil[passportHash] <= block.timestamp, "contest lock");

        Request storage r = requests[passportHash];
        if (r.newIdentityKey != 0) {
            require(r.newIdentityKey != newIdentityKey, "already requested");
            delete requests[passportHash];
            uint64 until = uint64(block.timestamp) + contestLock;
            lockedUntil[passportHash] = until;
            emit RecoveryContested(passportHash, until);
            return;
        }

        uint64 readyAt = uint64(block.timestamp) + recoveryDelay;
        requests[passportHash] = Request(newIdentityKey, readyAt);
        emit RecoveryRequested(passportHash, newIdentityKey, readyAt);
    }

    function cancelRecovery(bytes32 passportHash) external {
        require(requests[passportHash].newIdentityKey != 0, "no request");
        delete requests[passportHash];
        emit RecoveryCancelled(passportHash);
    }

    function finalizeRecovery(bytes32 passportHash) external {
        Request memory r = requests[passportHash];
        require(r.newIdentityKey != 0, "no request");
        require(!locked[passportHash], "locked");
        require(block.timestamp >= r.readyAt, "not ready");

        identityKey[passportHash] = r.newIdentityKey;
        delete requests[passportHash];
        emit RecoveryFinalized(passportHash, r.newIdentityKey);
    }

    // PROVISIONAL: a lock keeps the pending request and blocks its finalization.
    function lockRecovery(bytes32 passportHash) external {
        require(!locked[passportHash], "already locked");
        locked[passportHash] = true;
        emit RecoveryLocked(passportHash);
    }

    function unlockRecovery(bytes32 passportHash) external {
        require(locked[passportHash], "not locked");
        locked[passportHash] = false;
        emit RecoveryUnlocked(passportHash);
    }
}
