#!/bin/bash

chmod a+x bin/gwemix.sh
bin/gwemix.sh init "" conf/config.json
bin/gwemix.sh start

# gwemix 서비스가 정상적으로 시작되었는지 확인
check_gwemix_started() {
    local log_file="logs/log" # gwemix 로그 파일 경로
    local success_message="HTTP server started"
    local max_retries=30
    local retries=0
    local found=0

    echo "Checking if gwemix has started..."

    while [ $retries -lt $max_retries ]; do
        if grep -q "$success_message" "$log_file"; then
            echo "gwemix has started successfully."
            found=1
            break
        else
            echo "Waiting for gwemix to start... Attempt $((retries+1))"
            sleep 1
        fi
        retries=$((retries+1))
    done

    if [ $found -eq 0 ]; then
        echo "Failed to confirm gwemix start after $max_retries attempts."
        exit 1
    fi
}

# gwemix 서비스가 정상적으로 시작되었는지 확인
check_gwemix_started

bin/gwemix.sh init-gov "" conf/config.json keystore/account1 test

echo "sleep 10 seconds for etcd init..."
sleep 10

# Open console and execute admin.etcdInit()
bin/gwemix.sh console <<EOF
admin.etcdInit()
EOF

echo "etcd init done."
