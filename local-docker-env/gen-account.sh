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

mkdir local-docker-env/keystore

# Create accounts
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
done