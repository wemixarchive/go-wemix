#!/bin/bash

bin/gwemix-local.sh init "" conf/config.json
bin/gwemix-local.sh start

wait_for_port() {
    local port="8588"
    while ! nc -z localhost "$port"; do
        sleep 1
    done
}

# 포트가 열릴 때까지 기다림
wait_for_port 8588

bin/gwemix.sh init-gov "" conf/config.json keystore/account1 test