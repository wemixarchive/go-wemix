// deploy-governance.js

// uses offline wallet
var GovernanceDeployer = new function() {
    this.wallet = null
    this.from = null
    this.gas = 21000 * 1400
    this.gasPrice = eth.gasPrice
    this._nonce = 0
    this.receiptCheckParams = { "interval": 100, "count": 300 }

    // bool log(var args...)
    this.log = function() {
        var msg = ""
        for (var i in arguments) {
            if (msg.length > 0)
                msg += " "
            msg += arguments[i]
        }
        console.log(msg)
        return true
    }

    // void verifyCfg(json data)
    // verifies config data, and normalize addresses
    // throws exception on failure
    this.verifyCfg = function(data) {
        if (data.accounts.length == 0 || data.members.length == 0)
            throw "At least one account and node are required"
        var bootnodeExists = false
        for (var i in data.members) {
            var m = data.members[i]
            if (!web3.isAddress(m.addr))
                throw "Invalid address 1 " + m.addr
            data.members[i].addr = web3.toChecksumAddress(m.addr)
            if (m.bootnode)
                bootnodeExists = true
        }
        if (!bootnodeExists)
            throw "Bootnode is not designated"
        for (var i in data.accounts) {
            var a = data.accounts[i]
            if (!web3.isAddress(a.addr))
                throw "Invalid address " + a.addr
            data.accounts[i].addr = web3.toChecksumAddress(a.addr)
        }
        if (data.staker) {
            if (!web3.isAddress(data.staker))
                throw "Invalid staker address " + data.staker
            data.staker = web3.toChecksumAddress(data.staker)
        }
        if (data.ecosystem) {
            if (!web3.isAddress(data.ecosystem))
                throw "Invalid ecosystem address " + data.ecosystem
            data.ecosystem = web3.toChecksumAddress(data.ecosystem)
        }
        if (data.maintenance) {
            if (!web3.isAddress(data.maintenance))
                throw "Invalid maintenance address " + data.maintenance
            data.maintenance = web3.toChecksumAddress(data.maintenance)
        }
    }

    // bytes packNum(int num)
    // pack a number into 256 bit bytes
    this.packNum = function(num) {
        return web3.padLeft(web3.toHex(num).substr(2), 64, "0")
    }

    // { "nodes": string, "stakes": string, "staker": address, "ecosystem": address, "maintenance": address } getInitialGovernanceMembersAndNodes(json data)
    this.getInitialGovernanceMembersAndNodes = function(data) {
        var nodes = "0x", stakes = "0x"

        for (var i = 0, l = data.members.length; i < l; i++) {
            var m = data.members[i], id, addr
            if (m.id.length != 128 && m.id.length != 130)
                throw "Invalid enode id " + m.id
            id = m.id.length == 128 ? m.id : m.id.substr(2)
            addr = m.addr.indexOf("0x") != 0 ? m.addr : m.addr.substr(2)

            nodes += web3.padLeft(addr, 64, "0") +
                this.packNum(m.name.length) + web3.fromAscii(m.name).substr(2) +
                this.packNum(id.length/2) + id +
                this.packNum(m.ip.length) + web3.fromAscii(m.ip).substr(2) +
                this.packNum(m.port)

            stakes += web3.padLeft(addr, 64, "0") +
                this.packNum(m.stake)
        }
        return {
            "nodes": nodes,
            "stakes": stakes,
            "staker": data.staker,
            "ecosystem": data.ecosystem,
            "maintenance": data.maintenance
        }
    }

    this.nonce = function() {
        return this._nonce++
    }

    // returns transaction hash, or throws error
    this.deployContract = function(data) {
        var tx = {
            from: this.from,
            data: data,
            gas: this.gas,
            gasPrice: this.gasPrice,
            nonce: this.nonce()
        }
        var stx = offlineWalletSignTx(this.wallet.id, tx, eth.chainId())
        return eth.sendRawTransaction(stx)
    }

    // wait for transaction receipt for contract address, then
    // load a contract
    // Contract resolveContract(ABI abi, hash txh)
    this.resolveContract = function(abi, txh) {
        for (var i = 0; i < this.receiptCheckParams.count; i++ ) {
            var r = eth.getTransactionReceipt(txh)
            if (r != null && r.contractAddress != null) {
                var ctr = web3.eth.contract(abi).at(r.contractAddress)
                ctr.transactionHash = txh
                return ctr
            }
            msleep(this.receiptCheckParams.interval)
        }
        throw "Cannot get contract address for " + txh
    }

    // sends a simple or method transaction, returns transaction hash
    this.sendTx = function(to, value, data) {
        var tx = {from:this.from, to:to, gas:this.gas,
                  gasPrice:this.gasPrice, nonce:this.nonce()}
        if (value)
            tx.value = value
        if (data)
            tx.data = data
        var stx = offlineWalletSignTx(this.wallet.id, tx, eth.chainId())
        return eth.sendRawTransaction(stx)
    }

    this.checkReceipt = function(tx) {
        for (var i = 0; i < this.receiptCheckParams.count; i++ ) {
            var r = eth.getTransactionReceipt(tx)
            if (r != null)
                return web3.toBigNumber(r.status) == 1
            msleep(this.receiptCheckParams.interval)
        }
        throw "Cannot get a transaction receipt for " + tx
    }

    this.sendStakingDeposit = function (to, data) {
        var tx = { from: this.from, to: to, gas: this.gas, gasPrice: this.gasPrice, nonce: this.nonce(), value: "0" }
        tx.value = "1500000" + "0".repeat(18)
        if (data) tx.data = data
        var stx = offlineWalletSignTx(this.wallet.id, tx, eth.chainId())

        return eth.sendRawTransaction(stx)
    }

    // bool deploy(string walletUrl, string password, string cfg)
    this.deploy = function(walletUrl, password, cfg, doInitOnce) {
        w = offlineWalletOpen(walletUrl, password)
        if (!w || !w.id || !w.address) {
            throw "Offline wallet is not loaded"
        }
        this.wallet = w
        this.from = this.wallet.address
        this._nonce = eth.getTransactionCount(this.from, 'pending')

        var data
        if (!(data = loadFile(cfg)))
            throw "cannot load governance contract .js or config .json file"

        // check if contracts exist
        var contractNames = [ "Registry", "EnvStorageImp", "Staking",
                              "BallotStorage", "EnvStorage", "GovImp", "Gov" ]
        for (var i in contractNames) {
            var cn = contractNames[i]
            if (eval("typeof " + cn + "_data") == "undefined" ||
                eval("typeof " + cn + "_contract") == "undefined")
                throw cn + " not found"
        }

        // check config.js
        eval("var data = " + data)
        this.verifyCfg(data)

        // initial members and nodes data
        var initData = this.getInitialGovernanceMembersAndNodes(data)

        // bootnode
        var bootNode = {
            "name": web3.fromAscii(data.members[0].name),
            "id": data.members[0].id,
            "ip": web3.fromAscii(data.members[0].ip),
            "port": data.members[0].port,
            "lockAmount": data.members[0].lockAmount
        }

        // contacts, transactions to be deployed
        var registry, envStorageImp, staking, ballotStorage, envStorage, govImp, gov
        var txs = new Array()

        // 1. deploy Registry and EnvStorageImp contracts
        this.log("Deploying Registry and EnvStorageImp...")
        registry = this.deployContract(Registry_data)
        envStorageImp = this.deployContract(EnvStorageImp_data)

        this.log("Waiting for receipts...")
        envStorageImp = this.resolveContract(EnvStorageImp_contract.abi, envStorageImp)
        registry = this.resolveContract(Registry_contract.abi, registry)

        // 2. deploy Staking, BallotStorage, EnvStorage, GovImp, Gov
        this.log("Deploying Staking, BallotStorage, EnvStorage, GovImp & Gov...")
        var code = Staking_contract.getData(registry.address,
            doInitOnce ? initData.stakes : "", {data: Staking_data})
        staking = this.deployContract(code)
        code = BallotStorage_contract.getData(registry.address, {data: BallotStorage_data})
        ballotStorage = this.deployContract(code)
        code = EnvStorage_contract.getData(envStorageImp.address, {data: EnvStorage_data})
        envStorage = this.deployContract(code)
        govImp = this.deployContract(GovImp_data)

        this.log("Waiting for receipts...")
        govImp = this.resolveContract(GovImp_contract.abi, govImp)
        code = Gov_contract.getData(govImp.address, { data: Gov_data })
        gov = this.deployContract(code)

        this.log("Waiting for gov contract...")
        gov = this.resolveContract(Gov_contract.abi, gov)
        envStorage = this.resolveContract(EnvStorage_contract.abi, envStorage)
        ballotStorage = this.resolveContract(BallotStorage_contract.abi, ballotStorage)
        staking = this.resolveContract(Staking_contract.abi, staking)

        // 3. setup registry
        this.log("Setting registry...")
        txs.length = 0
        txs[txs.length] = this.sendTx(registry.address, null,
            registry.setContractDomain.getData(
                "Staking", staking.address))
        txs[txs.length] = this.sendTx(registry.address, null,
            registry.setContractDomain.getData(
                "BallotStorage", ballotStorage.address))
        txs[txs.length] = this.sendTx(registry.address, null,
            registry.setContractDomain.getData(
                "EnvStorage", envStorage.address))
        txs[txs.length] = this.sendTx(registry.address, null,
            registry.setContractDomain.getData(
                "GovernanceContract", gov.address))
        if (initData.staker)
            txs[txs.length] = this.sendTx(registry.address, null,
                registry.setContractDomain.getData(
                    "StakingReward", initData.staker))
        if (initData.ecosystem)
            txs[txs.length] = this.sendTx(registry.address, null,
                registry.setContractDomain.getData(
                    "Ecosystem", initData.ecosystem))
        if (initData.maintenance)
            txs[txs.length] = this.sendTx(registry.address, null,
                registry.setContractDomain.getData(
                    "Maintenance", initData.maintenance))

        // no need to wait for the receipts for the above

        // 4. initialize environment storage data:
        // blocksPer, ballotDurationMin, ballotDurationMax,
        // stakingMin, stakingMax, gasPrice
        this.log("Initializing environment storage...")
        // Just changing address doesn't work here. Address is embedded in
        // the methods. Have to re-construct temporary EnvStorageImp here.
        var tmpEnvStorageImp = web3.eth.contract(envStorageImp.abi).at(envStorage.address)
        var envNames = [
            web3.sha3("blocksPer"),
            web3.sha3("ballotDurationMin"), web3.sha3("ballotDurationMax"),
            web3.sha3("stakingMin"), web3.sha3("stakingMax"),
            web3.sha3("MaxIdleBlockInterval"),
            web3.sha3("blockCreationTime"),
            web3.sha3("blockRewardAmount"),
            web3.sha3("maxPriorityFeePerGas"),
            web3.sha3("blockRewardDistributionBlockProducer"),
            web3.sha3("blockRewardDistributionStakingReward"),
            web3.sha3("blockRewardDistributionEcosystem"),
            web3.sha3("blockRewardDistributionMaintenance"),
            web3.sha3("maxBaseFee"),
            web3.sha3("blockGasLimit"),
            web3.sha3("baseFeeMaxChangeRate"),
            web3.sha3("gasTargetPercentage") ],
            envValues = [
                1,
                86400, 604800,
                1500000000000000000000000, 1500000000000000000000000,
                5,
                1000,
                web3.toWei(1, 'ether'),    // mint amount: 1 wemix
                web3.toWei(100, 'gwei'),   // tip: 100 gwei
                4000, 1000, 2500, 2500,    // NCPs, WEMIX Staker, Eco System, Maintenance
                web3.toWei(50000, 'gwei'), // maxBaseFee * 21000 -> 1.05 wemix
                5000 * 21000, 55, 30 ]
        txs[txs.length] = this.sendTx(envStorage.address, null,
            tmpEnvStorageImp.initialize.getData(registry.address, envNames, envValues))

        // 5. deposit staking
        txs[txs.length] = this.sendStakingDeposit(staking.address, staking.deposit.getData());

        if (!this.checkReceipt(txs[0]))
            throw "Failed to initialize environment storage data. Tx is " + txs[0]
        if (!this.checkReceipt(txs[1]))
            throw "Failed to initialize staking data. Tx is " + txs[1]

        // 5. Gov initialize
        this.log("Initializing governance members and nodes...")
        var tmpGovImp = web3.eth.contract(govImp.abi).at(gov.address);
        if (!doInitOnce) {
            txs.length = 0
            txs[txs.length] = this.sendTx(gov.address, null,
                tmpGovImp.init.getData(registry.address,
                    bootNode.lockAmount, bootNode.name, bootNode.id,
                    bootNode.ip, bootNode.port))
        } else {
            txs.length = 0
            txs[txs.length] = this.sendTx(gov.address, null,
                tmpGovImp.initOnce.getData(registry.address, initData.nodes))
        }
        if (!this.checkReceipt(txs[0]))
            throw "Failed to initialize with gov.init. Tx is " + txs[0]

        // 7. print the addresses
        this.log('{\n' +
                 '  "REGISTRY_ADDRESS": "' + registry.address + '",\n' +
                 '  "STAKING_ADDRESS": "' + staking.address + '",\n' +
                 '  "ENV_STORAGE_ADDRESS": "' + envStorage.address + '",\n' +
                 '  "BALLOT_STORAGE_ADDRESS": "' + ballotStorage.address + '",\n' +
                 '  "GOV_ADDRESS": "' + gov.address + '",\n' +
                 '  "GOV_IMP_ADDRESS": "' + govImp.address + '"\n' +
                 '}')

        return true
    }
}()

// EOF
