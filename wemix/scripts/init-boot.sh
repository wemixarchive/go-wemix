#!/bin/bash

bin/gwemix.sh init "" conf/config.json
bin/gwemix.sh start &

wait_for_port() {
    local port="$1"
    local retry_interval=1
    local max_retries=30
    local retries=0

    while :; do
        # wget을 사용하여 localhost의 특정 포트에 HTTP 요청을 시도합니다.
        # -T 옵션은 타임아웃을 설정합니다. -O -는 출력을 무시합니다.
        if wget --spider -T 10 -O - http://localhost:"$port" >/dev/null 2>&1; then
            echo "Port $port is open."
            break
        else
            echo "Waiting for port $port to open. Attempt $((retries+1))/$max_retries."
            sleep "$retry_interval"
        fi

        retries=$((retries+1))
        if [ "$retries" -ge "$max_retries" ]; then
            echo "Failed to connect to port $port after $max_retries attempts."
            exit 1
        fi
    done
}

# 포트가 열릴 때까지 기다림
wait_for_port 8588

bin/gwemix.sh init-gov "" conf/config.json keystore/account1 test