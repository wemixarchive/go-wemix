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
    echo "Usage: $0 -a <account_num>"
    exit 1
    ;;
  esac
done

# 필수 인수 확인
if [ -z "$ACCOUNT_NUM" ]; then
  exit 1
fi

mkdir geth
chmod 0700 nodekey
cp nodekey/nodekey"$ACCOUNT_NUM" geth/nodekey
