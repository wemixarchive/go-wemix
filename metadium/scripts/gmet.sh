#!/bin/bash

[ "$DIR" = "" ] && DIR=/opt

CHAIN_ID=101
CONSENSUS_METHOD=2
FIXED_GAS_LIMIT=0x40000000
GAS_PRICE=1000000000
MAX_IDLE_BLOCK_INTERVAL=600
BLOCKS_PER_TURN=10

function get_data_dir ()
{
    if [ ! "$1" = "" ]; then
        d=${DIR}/$1
        if [ -x "$d/bin/gmet" ]; then
            echo $d
        fi
    else
	for i in $(/bin/ls -1 ${DIR}); do
	    if [ -x "${DIR}/$i/bin/gmet" ]; then
		echo ${DIR}/$i
		return
	    fi
	done
        echo "$(cd "$(dirname "$0")" && pwd)"
    fi
}

# void init(String node, String config_json)
function init ()
{
    NODE="$1"
    CONFIG="$2"

    if [ ! -f "$CONFIG" ]; then
	echo "Cannot find config file: $2"
	return 1
    fi

    d=$(get_data_dir "${NODE}")
    if [ -x "$d/bin/gmet" ]; then
	GMET="$d/bin/gmet"
    else
	echo "Cannot find gmet"
	return 1
    fi

    if [ ! -f "${d}/conf/genesis-template.json" ]; then
	echo "Cannot find template files."
	return 1
    fi

    echo "wiping out data..."
    wipe $NODE

    [ -d "$d/logs" ] || mkdir -p "$d/logs"

    ${GMET} metadium genesis --data "$CONFIG" --genesis "$d/conf/genesis-template.json" --out "$d/genesis.json"
    [ $? = 0 ] || return $?

    echo "PORT=8588" > $d/.rc
    ${GMET} --datadir $d init $d/genesis.json
    echo "Generating dags for epoch 0 and 1..."
    ${GMET} makedag 0     $d/.ethash &
    ${GMET} makedag 1     $d/.ethash &
    wait
}

# void init_gov(String node, String config_json, String account_file)
function init_gov ()
{
    NODE="$1"
    CONFIG="$2"
    ACCT="$3"

    if [ ! -f "$CONFIG" ]; then
	echo "Cannot find config file: $2"
	return 1
    fi

    d=$(get_data_dir "${NODE}")
    if [ -x "$d/bin/gmet" ]; then
	GMET="$d/bin/gmet"
    else
	echo "Cannot find gmet"
	return 1
    fi

    if [ ! -f "${d}/conf/MetadiumGovernance.js" ]; then
	echo "Cannot find ${d}/conf/MetadiumGovernance.js"
	return 1
    fi

    ${GMET} metadium deploy-governance --url http://localhost:8588 --gasprice 1 --gas 0xF000000 "$d/conf/MetadiumGovernance.js" "$CONFIG" "${ACCT}"
}

function wipe ()
{
    d=$(get_data_dir "$1")
    if [ ! -x "$d/bin/gmet" ]; then
	echo "Is '$1' metadium data directory?"
	return
    fi

    cd $d
    /bin/rm -rf geth/LOCK geth/chaindata geth/ethash geth/lightchaindata \
	geth/transactions.rlp geth/nodes geth.ipc logs/* etcd
}

function wipe_all ()
{
    for i in `/bin/ls -1 ${DIR}/`; do
        if [ ! -d "${DIR}/$i" -o ! -x "${DIR}/$i/bin/gmet" ]; then
            continue
        fi
        wipe $i
    done
}

function clean ()
{
    d=$(get_data_dir "$1")
    if [ -x "$d/bin/gmet" ]; then
        GMET="$d/bin/gmet"
    else
	echo "Cannot find gmet"
	return
    fi

    cd $d
    $GMET --datadir ${PWD} removedb
}

function clean_all ()
{
    for i in `/bin/ls -1 ${DIR}/`; do
        if [ ! -d "${DIR}/$i" -o ! -d "${DIR}/$i/geth" ]; then
            continue
        fi
        clean $i
    done
}

function start ()
{
    d=$(get_data_dir "$1")
    if [ -x "$d/bin/gmet" ]; then
        GMET="$d/bin/gmet"
    else
	echo "Cannot find gmet"
	return
    fi

    [ -f "$d/.rc" ] && source "$d/.rc"
    [ "$COINBASE" = "" ] && COINBASE="" || COINBASE="--miner.etherbase $COINBASE"

    RPCOPT="--rpc --rpcaddr 0.0.0.0"
    [ "$NONCE_LIMIT" = "" ] || NONCE_LIMIT="--noncelimit $NONCE_LIMIT"
    [ "$BOOT_NODES" = "" ] || BOOT_NODES="--bootnodes $BOOT_NODES"
    if [ ! "$DISCOVER" = "1" ]; then
        DISCOVER=--nodiscover
    else
        DISCOVER=
    fi  

    [ -d "$d/logs" ] || mkdir -p $d/logs

    cd $d
    if [ ! "$2" = "inner" ]; then
	$GMET --datadir ${PWD} --syncmode full --gcmode archive --metrics \
            $COINBASE $DISCOVER $RPCOPT $BOOT_NODES $NONCE_LIMIT | \
            ${d}/bin/logrot ${d}/logs/log 10M 5 &
    else
	exec $GMET --datadir ${PWD} --syncmode full --gcmode archive \
            --metrics $COINBASE $DISCOVER $RPCOPT $NONCE_LIMIT
    fi
}

function start_all ()
{
    for i in `/bin/ls -1 ${DIR}/`; do
        if [ ! -d "${DIR}/$i" -o ! -d "${DIR}/$i/geth" -o ! -f "${DIR}/$i/bin/gmet" ]; then
            continue
        fi
        start $i
	echo "started $i."
        return
    done

    echo "Cannot find gmet directory. Check if 'geth' directory and 'bin/gmet' are present in the data directory";
}

function get_gmet_pids ()
{
    ps axww | grep -v grep | grep "gmet.*datadir.*${NODE}" | awk '{print $1}'
}

function do_nodes ()
{
    LHN=$(hostname)
    CMD=${1/-nodes/}
    shift
    while [ ! "$1" = "" -a ! "$2" = "" ]; do
	if [ "$1" = "$LHN" -o "$1" = "${LHN/.*/}" ]; then
	    $0 ${CMD} $2
	else
	    ssh -f $1 ${DIR}/$2/bin/gmet.sh ${CMD} $2
	fi
	shift
	shift
    done
}

function usage ()
{
    echo "Usage: `basename $0` [init <node> <config.json> |
	init-gov <node> <config.json> <account-file> |
	clean [<node>] | wipe [<node>] | console [<node>] |
	[re]start [<node>] | stop [<node>] | [re]start-nodes | stop-nodes]

*-nodes uses NODES environment variable: [<host> <dir>]+
"
}

case "$1" in
"init")
    if [ $# -lt 3 ]; then
	usage;
    else
        init "$2" "$3"
    fi
    ;;

"init-gov")
    if [ $# -lt 4 ]; then
	usage;
    else
        init_gov "$2" "$3" "$4"
    fi
    ;;

"wipe")
    if [ ! "$3" = "" ]; then
        wipe $2 $3
    else
        wipe_all $2
    fi
    ;;

"clean")
    if [ ! "$2" = "" ]; then
        clean $2
    else
        clean_all
    fi
    ;;

"stop")
    echo -n "stopping..."
    if [ ! "$2" = "" ]; then
        NODE=$2
    else
        NODE=
    fi
    PIDS=`get_gmet_pids`
    if [ ! "$PIDS" = "" ]; then
	echo $PIDS | xargs -L1 sudo kill
    fi
    for i in {1..200}; do
	PIDS=`get_gmet_pids`
	[ "$PIDS" = "" ] && break
	echo -n "."
	sleep 1
    done
    PIDS=`get_gmet_pids`
    if [ ! "$PIDS" = "" ]; then
	echo $PIDS | xargs -L1 sudo kill -9
    fi
    echo "done."
    ;;

"start")
    if [ ! "$2" = "" ]; then
        start $2
	echo "started $2."
    else
        start_all
    fi
    ;;

"start-inner")
    if [ "$2" = "" ]; then
	usage;
    else
        start $2 inner
    fi
    ;;

"restart")
    if [ ! "$2" = "" ]; then
	$0 stop $2
        start $2
    else
	$0 stop
        start_all
    fi
    ;;

"start-nodes"|"restart-nodes"|"stop-nodes")
    if [ "${NODES}" = "" ]; then
	echo "NODES is not defined"
    fi
    do_nodes $1 ${NODES}
    ;;

"console")
    d=$(get_data_dir "$2")
    if [ ! -d $d ]; then
	usage; exit;
    fi
    RCJS=
    if [ -f "$d/rc.js" ]; then
        RCJS="--preload $d/rc.js"
    fi
    exec ${d}/bin/gmet ${RCJS} attach ipc:${d}/geth.ipc
    ;;

*)
    usage;
    ;;
esac

# EOF
