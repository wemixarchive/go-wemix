#!/bin/bash

# Check required arguments
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 -a <nodekey_num>"
  exit 1
fi

# Option parsing
while getopts "a:" opt; do
  case ${opt} in
  a)
    NODEKEY_NUM=$OPTARG
    ;;
  \?)
    echo "Usage: $0 -a <nodekey_num>"
    exit 1
    ;;
  esac
done

# Check required arguments
if [ -z "$NODEKEY_NUM" ]; then
  exit 1
fi

mkdir local-docker-env/nodekey

# Create nodekeys
for ((i = 1; i <= NODEKEY_NUM; i++)); do
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