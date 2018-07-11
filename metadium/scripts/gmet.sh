#!/bin/bash

[ "$DIR" = "" ] && DIR=/opt

CHAIN_ID=101
CONSENSUS_METHOD=2
FIXED_GAS_LIMIT=3141592653
MAX_IDLE_BLOCK_INTERVAL=600
BLOCKS_PER_TURN=10
MINER_THREADS=1
[ "${METADIUM_ABI}" = "" ] && METADIUM_ABI=MetadiumAdmin.js

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

# void init(String node, String genesis_json, int port, String account_file)
function init_old ()
{
    NODE="$1"
    GENESIS="$2"
    PORT=$3
    ACCOUNT_FILE="$4"

    if [ ! -f "$GENESIS" ]; then
	echo "Cannot find genesis file: $2"
	return
    fi

    d=$(get_data_dir "$NODE")
    if [ -x "$d/bin/gmet" ]; then
        GMET="$d/bin/gmet"
    fi

    cd $d
    [ -d logs ] || mkdir -p logs

    echo "PORT=${PORT}" > $d/.rc

    [ "${GENESIS}" = "$d/.genesis.json" ] || /bin/cp ${GENESIS} $d/.genesis.json
    $GMET --datadir ${PWD} init $d/.genesis.json

    if [ -f "$ACCOUNT_FILE" ]; then
	cp "$ACCOUNT_FILE" $d/keystore/
    fi
}

# void init(String node, String config_json, int port)
function init ()
{
    NODE="$1"
    CONFIG="$2"
    PORT=$3

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

    if [ ! -f "${d}/conf/genesis-template.json" -o ! -f "${d}/conf/MetadiumAdmin-template.sol" ]; then
	echo "Cannot find template files."
	return 1
    fi

    [ -d "$d/logs" ] || mkdir -p "$d/logs"

    ${GMET} metadium genesis --data "$CONFIG" --genesis "$d/conf/genesis-template.json" --out "$d/genesis.json"
    [ $? = 0 ] || return $?
    ${GMET} metadium admin-contract --data "$CONFIG" --admin "$d/conf/MetadiumAdmin-template.sol" --out "$d/MetadiumAdmin.sol"
    [ $? = 0 ] || return $?

    echo "PORT=${PORT}" > $d/.rc
    ${GMET} --datadir ${PWD} init $d/genesis.json
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
	geth/transactions.rlp geth.ipc logs/* default.etcd
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
    if [ -x "$d/bin/logrot" ]; then
	LOGROT="$d/bin/logrot"
    else
	echo "Cannot find logrot"
	return
    fi

    if [ -f "$d/.rc" ]; then
	source "$d/.rc"
    fi

    MINE_OPT="--mine --minerthreads 1";
    if [ "${PORT}" != "" ]; then
	PORT_OPT="--port ${PORT}";
	RPC_PORT_OPT="--rpc --rpcaddr 0.0.0.0 --rpcport $(($PORT + 1))"
    else
	PORT_OPT=
	RPC_PORT_OPT=
    fi
    CHAIN_ID_OPT="--networkid ${CHAIN_ID}"

    if [ -f "$d/rc.js" ]; then
        RCJS="--preload $d/rc.js"
    fi

    TXPOOL_OPTS="--txpool.accountslots 100000 --txpool.globalslots 100000"
    [ "${TARGET_GAS_LIMIT}" != "" ] && \
	TARGET_GAS_LIMIT_OPT="--targetgaslimit ${TARGET_GAS_LIMIT}"

    METADIUM_OPTS="--consensusmethod ${CONSENSUS_METHOD} --fixedgaslimit ${FIXED_GAS_LIMIT} --maxidleblockinterval ${MAX_IDLE_BLOCK_INTERVAL} --blocksperturn ${BLOCKS_PER_TURN} --metadiumabi ${METADIUM_ABI}"

    cd $d
    $GMET --datadir ${PWD} ${CHAIN_ID_OPT} ${MINE_OPT} --nodiscover	\
	--metrics ${PORT_OPT} ${RPC_PORT_OPT} ${TXPOOL_OPTS}            \
	${TARGET_GAS_LIMIT_OPT}	${METADIUM_OPTS} ${RCJS} 2>&1		\
        | ${LOGROT} ${d}/logs/log 10M 5 &
}

function start_all ()
{
    for i in `/bin/ls -1 ${DIR}/`; do
        if [ ! -d "${DIR}/$i" -o ! -d "${DIR}/$i/geth" ]; then
            continue
        fi
        start $i
    done
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
    echo "Usage: `basename $0` [init <node> <config.json> <port> |
	clean [<node>] | wipe [<node>] | console [<node>] |
	[re]start [<node>] | stop [<node>] | [re]start-nodes | stop-nodes]

*-nodes uses NODES environment variable: [<host> <dir>]+
"
}

case "$1" in
"init")
    if [ $# -lt 4 ]; then
	usage;
    else
        init "$2" "$3" "$4" "$5"
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
    if [ ! "$2" = "" ]; then
        NODE=$2
    else
        NODE=
    fi
    for i in {1..10}; do
        ps axww | grep -v grep | grep -q "gmet.*datadir.*${NODE}"
        if [ ! $? = 0 ]; then
            break
        else
            ps axww | grep -v grep | grep "gmet.*datadir.*${NODE}" | awk '{print $1}' | xargs -L1 sudo kill
            sleep 3;
        fi
    done
    ps axww | grep -v grep | grep -q "gmet.*datadir.*${NODE}"
    if [ $? = 0 ]; then
        ps axww | grep -v grep | grep "gmet.*datadir.*${NODE}" | awk '{print $1}' | xargs -L1 sudo kill -9
    fi
    ;;

"start")
    if [ ! "$2" = "" ]; then
        start $2
    else
        start_all
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
