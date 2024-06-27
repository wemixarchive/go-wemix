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

mkdir keystore
mkdir nodekey

# 계정 및 노드 키 생성
for ((i = 1; i <= ACCOUNT_NUM; i++)); do
  if [ -f "keystore/account$i" ]; then
    rm keystore/account$i
  fi
  yes "test" | head -n 2 | gwemix wemix new-account --out keystore/account$i
  
  if [ -f "nodekey/nodekey$i" ]; then
    rm nodekey/nodekey$i
  fi
  gwemix wemix new-nodekey --out nodekey/nodekey$i
done