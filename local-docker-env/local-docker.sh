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

# key-gen.sh 실행
chmod +x local-docker-env/key-gen.sh
./local-docker-env/key-gen.sh -a "$ACCOUNT_NUM"

# config-gen.sh 실행
chmod +x local-docker-env/config-gen.sh
./local-docker-env/config-gen.sh -a "$ACCOUNT_NUM" -f local-docker-env/config.json

# docker-compose-gen.sh 실행
chmod +x local-docker-env/docker-compose-gen.sh
./local-docker-env/docker-compose-gen.sh -a "$ACCOUNT_NUM"

# Dockerfile.boot 및 Dockerfile.node 파일 복사
cp local-docker-env/Dockerfile.boot ./
cp local-docker-env/Dockerfile.node ./

# docker-compose.yml 파일을 이용해 docker-compose build 및 up 실행
docker compose -f docker-compose.yml build --no-cache
docker compose -f docker-compose.yml up -d