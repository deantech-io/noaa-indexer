// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

// The part of NoAARecovery the service and the mock rely on.
interface INoAARecovery {
    struct RecoveryRequest {
        bytes32 oldIdentityKey;
        bytes32 newIdentityKey;
        uint256 dgCommit;
        uint64 readyAt;
    }

    event RecoveryRequested(bytes32 passportHash, bytes32 newIdentityKey, uint256 readyAt);
    event RecoveryContested(bytes32 passportHash, bytes32 pendingIdentityKey, bytes32 newIdentityKey);
    event RecoveryCancelled(bytes32 passportHash, bytes32 pendingIdentityKey);
    event RecoveryFinalized(bytes32 passportHash, bytes32 oldIdentityKey, bytes32 newIdentityKey);
    event RecoveryLocked(bytes32 passportHash);
    event RecoveryUnlocked(bytes32 passportHash);

    function getRecoveryRequest(bytes32 passportHash_) external view returns (RecoveryRequest memory);

    function lockedUntil(bytes32 passportHash_) external view returns (uint64);

    function finalizeRecovery(bytes32 passportHash_) external;
}
