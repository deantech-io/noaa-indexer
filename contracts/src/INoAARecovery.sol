// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

// PROVISIONAL: assumed interface, the real contract does not exist yet.
interface INoAARecovery {
    event RecoveryRequested(bytes32 indexed passportHash, uint256 newIdentityKey, uint64 readyAt);
    event RecoveryCancelled(bytes32 indexed passportHash);
    event RecoveryContested(bytes32 indexed passportHash, uint64 lockedUntil);
    event RecoveryFinalized(bytes32 indexed passportHash, uint256 newIdentityKey);
    event RecoveryLocked(bytes32 indexed passportHash);
    event RecoveryUnlocked(bytes32 indexed passportHash);

    function requests(bytes32 passportHash) external view returns (uint256 newIdentityKey, uint64 readyAt);

    function lockedUntil(bytes32 passportHash) external view returns (uint64);

    function finalizeRecovery(bytes32 passportHash) external;
}
