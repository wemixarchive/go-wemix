// ooo.js

// web3.padLeft("123", 5, "0") -> "00123"
// web3.padRight("123", 5, ".") -> "123.."

function log() {
    var msg = ""
    for (var i in arguments) {
        if (msg.length > 0)
            msg += " "
        msg += arguments[i]
    }
    console.log(msg)
    return true
}

// bool deployGovernance(string datadir, string cfg)
// will throw exception if something fails
// won't roll back when error occurs
function deployGovernance(datadir, cfg) {
    var from = eth.accounts[0]
    var gas = "0xE000000", gasPrice = eth.gasPrice
    var receiptCheckParams = { "interval": 100, "count": 300 }

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
    verifyCfg()

    // initial members and nodes data
    var initData = getInitialGovernanceMembersAndNodes(data)

    // contacts, transactions to be deployed
    var registry, envStorageImp, staking, ballotStorage, envStorage, govImp, gov
    var txs = new Array()

    // 1. deploy Registry and EnvStorageImp contracts
    log("Deploying Registry and EnvStorageImp...")
    registry = deployContract(Registry_contract, Registry_data)
    envStorageImp = deployContract(EnvStorageImp_contract, EnvStorageImp_data)

    log("Waiting for receipts...")
    envStorageImp = getContractAddress(envStorageImp)
    registry = getContractAddress(registry)

    // 2. deploy Staking, BallotStorage, EnvStorage, GovImp, Gov
    log("Deploying Staking, BallotStorage, EnvStorage, GovImp & Gov...")
    var code = Staking_contract.getData(registry.address, initData.stakes,
                                        {data: Staking_data})
    staking = deployContract(Staking_contract, code)
    ballotStorage = deployContract(BallotStorage_contract, BallotStorage_data)
    code = EnvStorage_contract.getData(registry.address, envStorageImp.address,
                                       {data: EnvStorage_data})
    envStorage = deployContract(EnvStorage_contract, code)
    govImp = deployContract(GovImp_contract, GovImp_data)
    gov = deployContract(Gov_contract, Gov_data)

    log("Waiting for receipts...")
    gov = getContractAddress(gov)
    govImp = getContractAddress(govImp)
    envStorage = getContractAddress(envStorage)
    ballotStorage = getContractAddress(ballotStorage)
    staking = getContractAddress(staking)

    // 3. setup registry
    log("Setting registry...")
    txs.length = 0
    txs[txs.length] = registry.setContractDomain(
        "Staking", staking.address, {from:from, gas:gas})
    txs[txs.length] = registry.setContractDomain(
        "BallotStorage", ballotStorage.address, {from:from, gas:gas})
    txs[txs.length] = registry.setContractDomain(
        "EnvStorage", envStorage.address, {from:from, gas:gas})
    txs[txs.length] = registry.setContractDomain(
        "GovernanceContract", gov.address, {from:from, gas:gas})
    if (initData.pool)
        txs[txs.length] = registry.setContractDomain(
            "RewardPool", initData.pool, {from:from, gas:gas})
    if (initData.maintenance)
        txs[txs.length] = registry.setContractDomain(
            "Maintenance", initData.pool, {from:from, gas:gas})

    // no need to wait for the receipts for the above

    // 4. deposit staking - not needed

    // 5. Gov.initOnce()
    log("Initializing governance members and nodes...")
    txs.length = 0
    txs[txs.length] = gov.initOnce(registry.address, govImp.address,
                                   initData.nodes, {from:from, gas:gas})

    // 6. initialize environment storage data:
    // blocksPer, ballotDurationMin, ballotDurationMax, stakingMin, stakingMax,
    // gasPrice
    log("Initializing environment storage...")
    txs[txs.length] = envStorageImp.initialize(
        100, // blocksPer
        86400, 604800, // ballotDurationMin, ...Max
        // stakingMin, ...Max
        4980000000000000000000000, 39840000000000000000000000,
        80000000000, // gasPrice
        5, // maxIdleBlockInterval
        {from:from, gas:gas})

    if (!checkReceipt(txs[0]))
        throw "Failed to initialize with initOnce. Tx is " + txs[0]
    if (!checkReceipt(txs[1]))
        throw "Failed to initialize environment storage data. Tx is " + txs[1]

    // 7. print the addresses
    log('{\n' +
'  "REGISTRY_ADDRESS": "' + registry.address + '",\n' +
'  "STAKING_ADDRESS": "' + staking.address + '",\n' +
'  "ENV_STORAGE_ADDRESS": "' + envStorage.address + '",\n' +
'  "BALLOT_STORAGE_ADDRESS": "' + ballotStorage.address + '",\n' +
'  "GOV_ADDRESS": "' + gov.address + '",\n' +
'  "GOV_IMP_ADDRESS": "' + govImp.address + '"\n' +
'}')

    return true


    // internal functions follow
    function verifyCfg() {
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

    function packNum(num) {
        return web3.padLeft(web3.toHex(num).substr(2), 64, "0")
    }

    // { string nodes, string stakes, address pool, address maintenance }  getInitialGovernanceMembersAndNodes(json data)
    function getInitialGovernanceMembersAndNodes(data) {
        var nodes = "0x", stakes = "0x", pool, maintenance

        for (var i = 0, l = data.members.length; i < l; i++) {
            var m = data.members[i], id, addr
            if (m.id.length != 128 && m.id.length != 130)
                throw "Invalid enode id " + m.id
            id = m.id.length == 128 ? m.id : m.id.substr(2)
            addr = m.addr.indexOf("0x") != 0 ? m.addr : m.addr.substr(2)

            nodes += web3.padLeft(addr, 64, "0") +
                packNum(m.name.length) + web3.fromAscii(m.name).substr(2) +
                packNum(id.length/2) + id +
                packNum(m.ip.length) + web3.fromAscii(m.ip).substr(2) +
                packNum(m.port)

            stakes += web3.padLeft(addr, 64, "0") +
                packNum(m.stake)
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
    function deployContract(abi, data) {
        return abi.new({
            from: from,
            data: data,
            gas: gas,
            gasPrice: gasPrice
        })
    }

    // TODO: needs to explain why a contract gets reloaded here
    function getContractAddress(ctr) {
        var tx = ctr.transactionHash
        for (var i = 0; i < receiptCheckParams.count; i++ ) {
            var r = eth.getTransactionReceipt(tx)
            if (r != null && r.contractAddress != null) {
                ctr = web3.eth.contract(ctr.abi).at(r.contractAddress)
                ctr.transactionHash = tx
                return ctr
            }
            msleep(receiptCheckParams.interval)
        }
        throw "Cannot get contract address for " + ctr.transactionHash
    }

    function checkReceipt(tx) {
        for (var i = 0; i < receiptCheckParams.count; i++ ) {
            var r = eth.getTransactionReceipt(tx)
            if (r != null)
                return web3.toBigNumber(r.status ) == 1
            msleep(receiptCheckParams.interval)
        }
        throw "Cannot get a transaction receipt for " + tx
    }
}

// EOF
