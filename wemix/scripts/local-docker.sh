#!/bin/bash

# 필수 인수 확인
if [ "$#" -ne 6 ]; then
  echo "Usage: $0 -a <account_num> -b <branch> -r <repo>"
  exit 1
fi

# 옵션 파싱
while getopts "a:b:r:" opt; do
  case ${opt} in
  a)
    ACCOUNT_NUM=$OPTARG
    ;;
  b)
    BRANCH=$OPTARG
    ;;
  r)
    REPO=$OPTARG
    ;;
  \?)
    echo "Usage: $0 -a <account_num> -b <branch> -r <repo>"
    exit 1
    ;;
  esac
done

# 필수 인수 확인
if [ -z "$ACCOUNT_NUM" ] || [ -z "$BRANCH" ] || [ -z "$REPO" ]; then
  exit 1
fi

# key-gen.sh 실행
chmod +x key-gen.sh
./key-gen.sh -a "$ACCOUNT_NUM"

# config-gen.sh 실행
chmod +x config-gen.sh
./config-gen.sh -a "$ACCOUNT_NUM" -f config.json

# BRANCH와 REPO 정보를 입력으로 받아 docker-compose-gen.sh 실행
chmod +x docker-compose-gen.sh
./docker-compose-gen.sh -a "$ACCOUNT_NUM" -b "$BRANCH" -r "$REPO"

# docker-compose.yml 파일을 이용해 docker-compose build 및 up 실행
docker compose -f docker-compose.yml build --no-cache
docker compose -f docker-compose.yml up -d