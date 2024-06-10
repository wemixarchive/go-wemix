#!/bin/bash

cd $(dirname ${BASH_SOURCE[0]})/..

gwemix.sh start
sleep 5

RESULT=$(gwemix --exec "admin.wemixInfo.etcd" attach gwemix.ipc)

if [[ "$RESULT" == "Fatal:"* ]]; then
    echo "Failed to start gwemix!"
    exit 1
elif [[ "$RESULT" == "{}" ]]; then
    echo "Initailize governance"

    gwemix.sh init-gov wemix conf/config.json
    sleep 5

    gwemix --exec "admin.wemixInfo" attach gwemix.ipc
    gwemix --exec "admin.etcdInit()" attach gwemix.ipc # Initialize etcd
    sleep 5
    gwemix --exec "admin.wemixInfo.etcd" attach gwemix.ipc
fi

echo "Started gwemix!"