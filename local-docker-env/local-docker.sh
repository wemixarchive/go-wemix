#!/bin/bash

# gwemix path를 환경변수에 추가
current_dir=$(pwd)
export PATH=$PATH:${current_dir}/go-wemix/build/bin

# 옵션 파싱
while getopts "a:v:" opt; do
  case ${opt} in
  a)
    ACCOUNT_NUM=$OPTARG
    ;;
  v)
    VERSION=$OPTARG
    ;;
  \?)
    echo "Usage: $0 -a <account_num>"
    exit 1
    ;;
  esac
done

# 필수 인수 확인
if [ -z "$ACCOUNT_NUM" ]; then
  echo "Account number is required. Usage: $0 -a <account_num>"
  exit 1
fi

# Set default version if not provided
if [ -z "$VERSION" ]; then
    VERSION="latest"
fi

# key-gen.sh 실행
chmod +x local-docker-env/key-gen.sh
./local-docker-env/key-gen.sh -a "$ACCOUNT_NUM" || { echo "Failed to execute key-gen.sh."; exit 1; }

# config-gen.sh 실행
chmod +x local-docker-env/config-gen.sh
./local-docker-env/config-gen.sh -a "$ACCOUNT_NUM" -f local-docker-env/config.json || { echo "Failed to execute config-gen.sh."; exit 1; }

# docker-compose-gen.sh 실행
chmod +x local-docker-env/docker-compose-gen.sh
./local-docker-env/docker-compose-gen.sh -a "$ACCOUNT_NUM" -v "$VERSION" || { echo "Failed to execute docker-compose-gen.sh."; exit 1; }

# Dockerfile.boot 및 Dockerfile.node 파일 복사
cp local-docker-env/Dockerfile.boot ./ || { echo "Failed to copy Dockerfile.boot."; exit 1; }
cp local-docker-env/Dockerfile.node ./ || { echo "Failed to copy Dockerfile.node."; exit 1; }

# docker-compose.yml 파일을 이용해 docker-compose build 및 up 실행
docker compose -f docker-compose.yml build --no-cache || { echo "Failed to build docker-compose."; exit 1; }
docker compose -f docker-compose.yml up -d || { echo "Failed to start docker-compose."; exit 1; }