// deploy-governance.js

// web3.padLeft("123", 5, "0") -> "00123"
// web3.padRight("123", 5, ".") -> "123.."

var GovernanceDeployer = new function() {
    this.from = eth.accounts[0]
    this.gas = "0xE000000";
    this.gasPrice = eth.gasPrice
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
        if (data.pool) {
            if (!web3.isAddress(data.pool))
                throw "Invalid pool address " + data.pool
            data.pool = web3.toChecksumAddress(data.pool)
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

    // { "nodes": string, "stakes": string, "pool": address, "maitenance": address } getInitialGovernanceMembersAndNodes(json data)
    this.getInitialGovernanceMembersAndNodes = function(data) {
        var nodes = "0x", stakes = "0x", pool, maintenance

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
            "pool": data.pool,
            "maintenance": data.maintenance
        }
    }

    // returns contract object, or throws error
    // 1. <object>.transactionHash is what it says it is
    // 2. <object>.address needs to be set by <receipt>.contractAddress later
    this.deployContract = function(abi, data) {
        return abi.new({
            from: this.from,
            data: data,
            gas: this.gas,
            gasPrice: this.gasPrice
        })
    }

    // TODO: needs to explain why a contract gets reloaded here
    this.getContractAddress = function(ctr) {
        var tx = ctr.transactionHash
        for (var i = 0; i < this.receiptCheckParams.count; i++ ) {
            var r = eth.getTransactionReceipt(tx)
            if (r != null && r.contractAddress != null) {
                ctr = web3.eth.contract(ctr.abi).at(r.contractAddress)
                ctr.transactionHash = tx
                return ctr
            }
            msleep(this.receiptCheckParams.interval)
        }
        throw "Cannot get contract address for " + ctr.transactionHash
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

    // bool deploy(string datadir, string cfg)
    this.deploy = function(datadir, cfg) {
        var data
        if (!loadScript(datadir + "/conf/MetadiumGovernance.js") ||
            !(data = loadFile(cfg)))
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

        // contacts, transactions to be deployed
        var registry, envStorageImp, staking, ballotStorage, envStorage, govImp, gov
        var txs = new Array()

        // 1. deploy Registry and EnvStorageImp contracts
        this.log("Deploying Registry and EnvStorageImp...")
        registry = this.deployContract(Registry_contract, Registry_data)
        envStorageImp = this.deployContract(EnvStorageImp_contract, EnvStorageImp_data)

        this.log("Waiting for receipts...")
        envStorageImp = this.getContractAddress(envStorageImp)
        registry = this.getContractAddress(registry)

        // 2. deploy Staking, BallotStorage, EnvStorage, GovImp, Gov
        this.log("Deploying Staking, BallotStorage, EnvStorage, GovImp & Gov...")
        var code = Staking_contract.getData(registry.address, initData.stakes,
                                            {data: Staking_data})
        staking = this.deployContract(Staking_contract, code)
        ballotStorage = this.deployContract(BallotStorage_contract, BallotStorage_data)
        code = EnvStorage_contract.getData(registry.address, envStorageImp.address,
                                           {data: EnvStorage_data})
        envStorage = this.deployContract(EnvStorage_contract, code)
        govImp = this.deployContract(GovImp_contract, GovImp_data)
        gov = this.deployContract(Gov_contract, Gov_data)

        this.log("Waiting for receipts...")
        gov = this.getContractAddress(gov)
        govImp = this.getContractAddress(govImp)
        envStorage = this.getContractAddress(envStorage)
        ballotStorage = this.getContractAddress(ballotStorage)
        staking = this.getContractAddress(staking)

        // 3. setup registry
        this.log("Setting registry...")
        txs.length = 0
        txs[txs.length] = registry.setContractDomain(
            "Staking", staking.address, {from:this.from, gas:this.gas})
        txs[txs.length] = registry.setContractDomain(
            "BallotStorage", ballotStorage.address, {from:this.from, gas:this.gas})
        txs[txs.length] = registry.setContractDomain(
            "EnvStorage", envStorage.address, {from:this.from, gas:this.gas})
        txs[txs.length] = registry.setContractDomain(
            "GovernanceContract", gov.address, {from:this.from, gas:this.gas})
        if (initData.pool)
            txs[txs.length] = registry.setContractDomain(
                "RewardPool", initData.pool, {from:this.from, gas:this.gas})
        if (initData.maintenance)
            txs[txs.length] = registry.setContractDomain(
                "Maintenance", initData.pool, {from:this.from, gas:this.gas})

        // no need to wait for the receipts for the above

        // 4. deposit staking - not needed

        // 5. Gov.initOnce()
        this.log("Initializing governance members and nodes...")
        txs.length = 0
        txs[txs.length] = gov.initOnce(registry.address, govImp.address,
                                       initData.nodes, {from:this.from, gas:this.gas})

        // 6. initialize environment storage data:
        // blocksPer, ballotDurationMin, ballotDurationMax, stakingMin, stakingMax,
        // gasPrice
        this.log("Initializing environment storage...")
        // Just changing address doesn't work here. Address is embeded in
        // the methods. Have to re-construct here.
        var tmpEnvStorageImp = web3.eth.contract(envStorageImp.abi).at(envStorage.address)
        txs[txs.length] = tmpEnvStorageImp.initialize(
            100, // blocksPer
            86400, 604800, // ballotDurationMin, ...Max
            // stakingMin, ...Max
            4980000000000000000000000, 39840000000000000000000000,
            80000000000, // gasPrice
            5, // maxIdleBlockInterval
            {from:this.from, gas:this.gas})

        if (!this.checkReceipt(txs[0]))
            throw "Failed to initialize with initOnce. Tx is " + txs[0]
        if (!this.checkReceipt(txs[1]))
            throw "Failed to initialize environment storage data. Tx is " + txs[1]

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
