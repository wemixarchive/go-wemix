# Croissant Hard Fork

Croissant brings the EVM-level part of Ethereum's Shanghai upgrade to go-wemix.
This document describes what changes at the fork block, how the fork is
activated, and where go-wemix deliberately differs from go-ethereum.

Fork names are desserts in alphabetical order of activation: Applepie,
Brioche, Croissant, Dalgona. The WBFT consensus switch was planned as
Croissant. It now comes after this fork and is called Dalgona.

## Scope

Ethereum's Shanghai bundles five EIPs. Croissant adopts the three that change
EVM execution and leaves out the two that do not apply to a PoA chain.

| EIP | Title | Status in go-wemix |
|-----|-------|--------------------|
| [EIP-3855](https://eips.ethereum.org/EIPS/eip-3855) | PUSH0 instruction | Adopted |
| [EIP-3860](https://eips.ethereum.org/EIPS/eip-3860) | Limit and meter initcode | Adopted |
| [EIP-3651](https://eips.ethereum.org/EIPS/eip-3651) | Warm COINBASE | Adopted |
| [EIP-4895](https://eips.ethereum.org/EIPS/eip-4895) | Beacon chain push withdrawals | Not adopted |
| [EIP-6049](https://eips.ethereum.org/EIPS/eip-6049) | Deprecate SELFDESTRUCT | No change needed |

Why only part of Shanghai:

- **EIP-4895** pushes withdrawals from the beacon chain into the execution
  layer. go-wemix has no beacon chain, so there is nothing to withdraw, and the
  header and block body fields it adds would have no source.
- **EIP-6049** only announces that SELFDESTRUCT will change later. It does not
  change any behaviour, so there is nothing to implement.
- The three adopted EIPs are what Solidity 0.8.20 and later target by default
  (`evmVersion: shanghai`). Adopting them lets contracts compiled with current
  toolchains run on go-wemix without lowering the EVM version.

## Activation

Every go-wemix fork is gated by block number, and Croissant follows that
convention. Ethereum activates Shanghai by timestamp instead.

The fork block is `ChainConfig.CroissantBlock`, `croissantBlock` in a genesis
`config`:

```json
{
  "config": {
    "croissantBlock": 1000000
  }
}
```

- `nil` (absent): the fork never activates and the EVM is unchanged.
- `0`: the chain starts on Croissant.
- `N`: blocks `N` and later run the Croissant rules.

Mainnet and testnet leave `croissantBlock` and `dalgonaBlock` unset until the
activation blocks are scheduled, so neither fork can activate. A release that
activates Croissant sets the same block in go-wemix and go-wbft, because
go-wbft re-executes the blocks go-wemix produced. Once a node has passed a
fork block, moving it is rejected as an incompatible config change.

### Order with Dalgona

go-wemix does not enforce an order between Croissant and Dalgona, as with its
other forks. go-wemix stops at the Dalgona block, so a Croissant block at or
after Dalgona never runs in go-wemix. go-wbft, which takes over at Dalgona,
applies the three EIPs from Dalgona on whether or not Croissant is scheduled.

### Renamed genesis key

Before this fork existed, the WBFT switch used the key `croissantBlock`. It is
now `dalgonaBlock`, and `croissantBlock` schedules this fork. Unknown keys in
a genesis `config` are ignored, so a private network genesis written for the
old meaning still loads, but its `croissantBlock` now activates the EIPs below
instead of stopping go-wemix. Rename the key to `dalgonaBlock` in such files.

## Changes at the fork block

### EIP-3855: PUSH0

The opcode `0x5f` pushes the constant zero onto the stack for 2 gas. Before the
fork it is an invalid opcode, as it was in every earlier go-wemix release.

### EIP-3860: Limit and meter initcode

The initcode limit is defined as twice the maximum contract size, as in the EIP.
go-wemix raises `MaxCodeSize` to 253,952 bytes, so the limit is larger than on
Ethereum:

| | Ethereum | go-wemix |
|-|----------|----------|
| `MaxCodeSize` | 24,576 | 253,952 |
| `MaxInitCodeSize` (`2 * MaxCodeSize`) | 49,152 | 507,904 |

`MaxInitCodeSize` is derived from `MaxCodeSize` rather than written as a
literal, so it follows any future change to the contract size limit.

From the fork block on:

- A creation transaction pays 2 gas per 32-byte word of initcode on top of its
  intrinsic gas.
- A creation transaction whose initcode exceeds `MaxInitCodeSize` is invalid
  and is rejected with `max initcode size exceeded`. It is not included in a
  block.
- `CREATE` and `CREATE2` pay 2 gas per initcode word and fail when the initcode
  exceeds `MaxInitCodeSize`.

### EIP-3651: Warm COINBASE

The block's coinbase is added to the access list at the start of every
transaction. The first access to the coinbase in a transaction then costs the
warm price (100 gas) instead of the cold price (2,600 gas).

## Code paths before and after the fork

Every Croissant change is selected by `Rules.IsCroissant`, which is
`ChainConfig.IsCroissant(blockNumber)`. Blocks below `CroissantBlock` run the
pre-Croissant code path unchanged.

| EIP | Where | Before the fork | From the fork block on |
|-----|-------|-----------------|------------------------|
| EIP-3855 | `core/vm/interpreter.go` | london jump table, `0x5f` undefined | `croissantInstructionSet` with `PUSH0` |
| EIP-3860 | `core/vm/jump_table.go` | `gasCreate`, `gasCreate2` | `gasCreateEip3860`, `gasCreate2Eip3860` |
| EIP-3860 | `core.IntrinsicGas` | `isEIP3860` is false, no initcode word charge | initcode word charge for creation transactions |
| EIP-3860 | `StateTransition.TransitionDb` | no initcode size check | `ErrMaxInitCodeSizeExceeded` over `MaxInitCodeSize` |
| EIP-3860 | `TxPool.validateTx`, light `TxPool.validateTx` | no initcode size check | `ErrMaxInitCodeSizeExceeded` over `MaxInitCodeSize` |
| EIP-3651 | `StateDB.PrepareAccessList` | coinbase not in the access list | coinbase added to the access list |

The transaction pools compute the same flag for the next block, so a
transaction is priced with the rules of the block it can be included in.

## Differences from go-ethereum

**Instruction set base.** go-ethereum builds the Shanghai instruction set on
top of the merge set, where `0x44` is `PREVRANDAO` and reads the block's random
value. go-wemix never went through the merge. Its PoA headers carry difficulty
1, so the EVM has no random value, and executing `0x44` on the merge set would
dereference nil and crash the node. go-wemix builds the Croissant set on the
London set instead, so `0x44` stays `DIFFICULTY` and keeps returning 1.

**Access list preparation.** go-ethereum renames `StateDB.PrepareAccessList` to
`Prepare` when adding EIP-3651. go-wemix already has an unrelated
`Prepare(thash, ti)`, and the upstream rename of that method touches the miner,
tracers and state processor. go-wemix keeps the name `PrepareAccessList` and
only adds the `rules` and `coinbase` parameters. The behaviour is the same.

## Notes for contract developers

- Before the fork block, compile with `evmVersion: paris` or earlier. Solidity
  0.8.20 and later emit `PUSH0` by default, and such bytecode fails with
  `invalid opcode` until Croissant is active.
- After the fork block, `evmVersion: shanghai` is supported. Later EVM versions
  (`cancun` and up) are not: their opcodes such as `TLOAD`, `TSTORE` and
  `MCOPY` are not part of this fork.

## Testing

Unit tests cover each EIP on both sides of the fork:

| Package | What is checked |
|---------|-----------------|
| `params` | block-number gating, `Rules` propagation, compatibility check, `MaxInitCodeSize` derivation |
| `core/vm` | PUSH0 and the initcode gas functions in the Croissant jump table, `0x44` staying `DIFFICULTY` |
| `core/vm/runtime` | PUSH0 execution, `CREATE`/`CREATE2` size limit, `DIFFICULTY` execution without a random value |
| `core/state` | coinbase warmed only under Croissant |
| `core` | initcode word gas, rejection of oversized creation transactions, 2,500 gas saved on a coinbase `BALANCE` |

Run them with:

```sh
go test -run 'Croissant|EIP3855|EIP3860|EIP3651|MaxInitCode|PrepareAccessList' \
  ./params/ ./core/ ./core/state/ ./core/vm/...
```
