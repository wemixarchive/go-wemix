#!/bin/bash

# 필수 인수 확인
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 -a <account_num>"
  exit 1
fi

# 옵션 파싱
while getopts "a:" opt; do
  case ${opt} in
  a)
    ACCOUNT_NUM=$OPTARG
    ;;
  \?)
    echo "Usage: $0 -a <account_num> -f <output_config_file>"
    exit 1
    ;;
  esac
done

# 필수 인수 확인
if [ -z "$ACCOUNT_NUM" ]; then
  exit 1
fi

mkdir local-docker-env/keystore
mkdir local-docker-env/nodekey

# 계정 및 노드 키 생성
for ((i = 1; i <= ACCOUNT_NUM; i++)); do
  if [ -f "local-docker-env/keystore/account$i" ]; then
    rm local-docker-env/keystore/account$i
  fi

  if yes "test" | head -n 2 | gwemix wemix new-account --out local-docker-env/keystore/account$i; then
    echo "Account $i created successfully"
  else
    echo "Failed to create account $i"
    exit 1
  fi
  
  if [ -f "local-docker-env/nodekey/nodekey$i" ]; then
    rm local-docker-env/nodekey/nodekey$i
  fi
  if gwemix wemix new-nodekey --out local-docker-env/nodekey/nodekey$i; then
    echo "Nodekey $i created successfully"
  else
    echo "Failed to create nodekey $i"
    exit 1
  fi
done