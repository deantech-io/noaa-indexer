// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {INoAARecovery} from "./INoAARecovery.sol";

// NoAARecovery without proofs and StateKeeper: anyone can trigger any transition.
// Passports start with the active identity 1.
contract MockNoAARecovery is INoAARecovery {
    uint64 public immutable recoveryDelay;
    uint64 public immutable lockDuration;

    mapping(bytes32 => RecoveryRequest) internal _requests;
    mapping(bytes32 => uint64) public lockedUntil;
    mapping(bytes32 => bool) public isRecoveryLocked;
    mapping(bytes32 => bytes32) internal _identities;

    constructor(uint64 recoveryDelay_, uint64 lockDuration_) {
        recoveryDelay = recoveryDelay_;
        lockDuration = lockDuration_;
    }

    function requestRecovery(
        bytes32,
        uint256 identityKey_,
        uint256 dgCommit_,
        bytes32 passportHash_,
        bytes32,
        bytes calldata
    ) external {
        require(identityKey_ > 0, "NoAARecovery: identity can not be zero");
        require(!isRecoveryLocked[passportHash_], "NoAARecovery: recovery is locked");
        require(lockedUntil[passportHash_] <= block.timestamp, "NoAARecovery: passport is locked");

        bytes32 activeIdentity_ = activeIdentity(passportHash_);

        require(activeIdentity_ != bytes32(identityKey_), "NoAARecovery: identity already active");

        RecoveryRequest storage request_ = _requests[passportHash_];

        if (request_.readyAt == 0) {
            uint64 readyAt_ = uint64(block.timestamp) + recoveryDelay;

            request_.oldIdentityKey = activeIdentity_;
            request_.newIdentityKey = bytes32(identityKey_);
            request_.dgCommit = dgCommit_;
            request_.readyAt = readyAt_;

            emit RecoveryRequested(passportHash_, bytes32(identityKey_), readyAt_);

            return;
        }

        bytes32 pendingIdentityKey_ = request_.newIdentityKey;

        require(pendingIdentityKey_ != bytes32(identityKey_), "NoAARecovery: recovery already requested");

        delete _requests[passportHash_];
        lockedUntil[passportHash_] = uint64(block.timestamp) + lockDuration;

        emit RecoveryContested(passportHash_, pendingIdentityKey_, bytes32(identityKey_));
    }

    function finalizeRecovery(bytes32 passportHash_) external {
        RecoveryRequest memory request_ = _requests[passportHash_];

        require(request_.readyAt != 0, "NoAARecovery: recovery is not requested");
        require(request_.readyAt <= block.timestamp, "NoAARecovery: recovery is not ready");
        require(activeIdentity(passportHash_) == request_.oldIdentityKey, "NoAARecovery: active identity changed");

        delete _requests[passportHash_];
        _identities[passportHash_] = request_.newIdentityKey;

        emit RecoveryFinalized(passportHash_, request_.oldIdentityKey, request_.newIdentityKey);
    }

    function cancelRecovery(bytes32 passportHash_, bytes calldata) external {
        bytes32 pendingIdentityKey_ = _requests[passportHash_].newIdentityKey;

        require(pendingIdentityKey_ != bytes32(0), "NoAARecovery: recovery is not requested");

        delete _requests[passportHash_];
        lockedUntil[passportHash_] = uint64(block.timestamp) + lockDuration;

        emit RecoveryCancelled(passportHash_, pendingIdentityKey_);
    }

    function lockRecovery(bytes32 passportHash_, bytes calldata) external {
        require(!isRecoveryLocked[passportHash_], "NoAARecovery: recovery is locked");

        isRecoveryLocked[passportHash_] = true;

        bytes32 pendingIdentityKey_ = _requests[passportHash_].newIdentityKey;

        if (pendingIdentityKey_ != bytes32(0)) {
            delete _requests[passportHash_];

            emit RecoveryCancelled(passportHash_, pendingIdentityKey_);
        }

        emit RecoveryLocked(passportHash_);
    }

    function unlockRecovery(bytes32 passportHash_, bytes calldata) external {
        require(isRecoveryLocked[passportHash_], "NoAARecovery: recovery is not locked");

        isRecoveryLocked[passportHash_] = false;

        emit RecoveryUnlocked(passportHash_);
    }

    function getRecoveryRequest(bytes32 passportHash_) external view returns (RecoveryRequest memory) {
        return _requests[passportHash_];
    }

    function activeIdentity(bytes32 passportHash_) public view returns (bytes32) {
        bytes32 identity_ = _identities[passportHash_];
        return identity_ == bytes32(0) ? bytes32(uint256(1)) : identity_;
    }
}
