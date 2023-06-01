# VRF

# Install Dependencies

```bash
$ go get github.com/ethereum/go-ethereum/crypto/vrf
```

# Get Results of `Prove()`

## curl

```bash
$ curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0", "method":"personal_edPubKey","params":["0x14e98077090336e914b8c76ef1a7c999c9e4d76e", <PASSWORD>], "id":11}' <IP>:<PORT>

{"jsonrpc":"2.0","id":11,"result":"0xf1b59f7bcb2ff6f33e8cbb01fcab71e98d80ac4d94a3b1197d2f433c3c5ced9b"}
```

```bash
$ curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0", "method":"personal_prove","params":["0x14e98077090336e914b8c76ef1a7c999c9e4d76e", "", "0x48656c6c6f2c20576f726c6421"], "id":11}' <IP>:<PORT>

{"jsonrpc":"2.0","id":11,"result":"0x035d91cb46308d31126f4308b3f69872a4f835cfa162e06a8588d1bb3a50ee608101bba5caec2da6f1c2733a60fb055c470a4e00d50ae07fa57f6d4e637fb71f45359ad6bb514bdb5e37a10fe44e99a93b"}
```

```bash
$ curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0", "method":"personal_verify","params":["0xf1b59f7bcb2ff6f33e8cbb01fcab71e98d80ac4d94a3b1197d2f433c3c5ced9b", "0x035d91cb46308d31126f4308b3f69872a4f835cfa162e06a8588d1bb3a50ee608101bba5caec2da6f1c2733a60fb055c470a4e00d50ae07fa57f6d4e637fb71f45359ad6bb514bdb5e37a10fe44e99a93b", "0x48656c6c6f2c20576f726c6421"], "id":11}' <IP>:<PORT>

{"jsonrpc":"2.0","id":11,"result":true}
```

## In geth console:

```bash
> accounts = personal.listAccounts

["0x14e98077090336e914b8c76ef1a7c999c9e4d76e", "0x5257cee5ce8381c89b1535939fc37b357e33fe48", "0x86a1551f03fa47e85afa1360637766adcf274449"]
```

```bash
> personal.edPubKey(accounts[0], <PASSWORD>)

"0xf1b59f7bcb2ff6f33e8cbb01fcab71e98d80ac4d94a3b1197d2f433c3c5ced9b"
```

```bash
> personal.prove(accounts[0], <PASSWORD>, "0x48656c6c6f2c20576f726c6421")

"0x035d91cb46308d31126f4308b3f69872a4f835cfa162e06a8588d1bb3a50ee60819dbd1abfb575f8dc792ddc78c718769f0508287cb0852497518731773a2c66bd9efb27b73338aff59e00a04aabdbb791"
```

```bash
> personal.verify("0xf1b59f7bcb2ff6f33e8cbb01fcab71e98d80ac4d94a3b1197d2f433c3c5ced9b", "0x035d91cb46308d31126f4308b3f69872a4f835cfa162e06a8588d1bb3a50ee60819dbd1abfb575f8dc792ddc78c718769f0508287cb0852497518731773a2c66bd9efb27b73338aff59e00a04aabdbb791", "0x48656c6c6f2c20576f726c6421")

true
```

# Test

## Test `Verify()`

```bash
$ /usr/local/go/bin/go test -timeout 30s -run ^TestPrecompiledVRF$ github.com/ethereum/go-ethereum/core/vm -v
```

## Benchmark

```bash
$ /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkPrecompiledVRF$ github.com/ethereum/go-ethereum/core/vm
```

# Node

```bash
$ ./build/bin/geth --datadir data account new

$ ./build/bin/geth --datadir data --networkid 327 --allow-insecure-unlock --http --http.api 'web3,eth,net,debug,personal' --http.corsdomain '*' --http.addr 0.0.0.0 --http.port 8545 --port 30301
```

```bash
$ ./build/bin/geth attach http://localhost:8545
```
