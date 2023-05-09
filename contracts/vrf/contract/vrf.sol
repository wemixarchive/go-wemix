// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

/**
 * @author @lukepark327
 * @title VRF
 */
abstract contract VRF {
    function _callVrfVerify(
        bytes32 publicKey,
        bytes memory pi,
        bytes memory message
    ) internal view returns (bool result) {
        return _vrf(abi.encodePacked(publicKey, pi, message));
    }

    function _callVrfVerifyRaw(
        bytes memory input
    ) internal view returns (bool result) {
        return _vrf(input);
    }

    function _vrf(bytes memory input) internal view returns (bool result) {
        assembly {
            let memPtr := mload(0x40)
            let success := staticcall(
                gas(),
                0x13,
                add(input, 0x20),
                mload(input),
                memPtr,
                0x20
            )
            switch success
            case 0 {
                revert(0, 0)
            }
            default {
                result := mload(memPtr)
            }
        }
    }
}
