#!/bin/bash

# gwemix 서비스 시작
chmod a+x bin/gwemix

# Check if the boot node has started
check_boot_node_started() {
    local max_retries=30
    local retries=0
    local found=0

    echo "Checking if the boot node has started..."

    while [ $retries -lt $max_retries ]; do
        if nc -zv 172.16.237.11 8588; then
            echo "Connection successful"
            found=1
            break
        else
            echo "Waiting for connection to the boot node... Attempt $((retries+1))"
            sleep 1
        fi
        retries=$((retries+1))
    done

    if [ $found -eq 0 ]; then
        echo "Failed to connect to the boot node after $max_retries attempts."
        exit 1
    fi
}

check_boot_node_started

bin/gwemix wemix download-genesis --url http://172.16.237.11:8588 --out genesis.json
bin/gwemix --datadir /usr/local/wemix --metrics --http --http.addr 0.0.0.0 --mine --http.port 8588 --ws --ws.addr 0.0.0.0 --ws.port 8598 --syncmode full --gcmode archive --port 8589