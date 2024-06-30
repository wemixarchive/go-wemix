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

# wemix 빌드
USE_ROCKSDB=NO make -f Makefile-local-env

# build/bin 디렉토리를 local-docker-env 디렉토리로 복사
cp -r build/bin/ local-docker-env/
echo "build/bin 디렉토리를 local-docker-env 디렉토리로 복사"

# key-gen.sh 실행
chmod +x local-docker-env/key-gen.sh
./local-docker-env/key-gen.sh -a "$ACCOUNT_NUM"

# config-gen.sh 실행
chmod +x local-docker-env/config-gen.sh
./local-docker-env/config-gen.sh -a "$ACCOUNT_NUM" -f local-docker-env/config.json

# docker-compose-gen.sh 실행
chmod +x local-docker-env/docker-compose-gen.sh
./local-docker-env/docker-compose-gen.sh -a "$ACCOUNT_NUM"

# docker-compose.yml 파일을 이용해 docker-compose build 및 up 실행
docker compose -f local-docker-env/docker-compose.yml build --no-cache
docker compose -f local-docker-env/docker-compose.yml up -d