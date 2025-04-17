#!/bin/bash

# Check required arguments
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 -a <account_num>"
  exit 1
fi

# Option parsing
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

# Check required arguments
if [ -z "$ACCOUNT_NUM" ]; then
  exit 1
fi

mkdir geth
chmod 0700 nodekey
cp nodekey/nodekey"$ACCOUNT_NUM" geth/nodekey
