#!/bin/bash

# gwemix path를 환경변수에 추가
current_dir=$(pwd)
export PATH=$PATH:${current_dir}/go-wemix/build/bin

# 옵션 파싱
while getopts "a:b:r:v:" opt; do
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
  v)
    VERSION=$OPTARG
    ;;
  \?)
    echo "Usage: $0 -a <account_num> -b <branch> -r <repo>"
    exit 1
    ;;
  esac
done

# 필수 인수 확인
if [ -z "$ACCOUNT_NUM" ] || [ -z "$BRANCH" ] || [ -z "$REPO" ]; then
  echo "Missing required arguments. Usage: $0 -a <account_num> -b <branch> -r <repo>"
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

# BRANCH와 REPO 정보를 입력으로 받아 docker-compose-gen.sh 실행
chmod +x local-docker-env/docker-compose-gen-git.sh
./local-docker-env/docker-compose-gen-git.sh -a "$ACCOUNT_NUM" -b "$BRANCH" -r "$REPO" -v "$VERSION" || { echo "Failed to execute docker-compose-gen-git.sh."; exit 1; }

# Dockerfile.boot.git 및 Dockerfile.node.git 파일 복사
cp local-docker-env/Dockerfile.boot.git ./ || { echo "Failed to copy Dockerfile.boot.git."; exit 1; }
cp local-docker-env/Dockerfile.node.git ./ || { echo "Failed to copy Dockerfile.node.git."; exit 1; }

# docker-compose.yml 파일을 이용해 docker-compose build 및 up 실행
docker compose -f docker-compose.yml build --no-cache || { echo "Failed to build docker-compose."; exit 1; }
docker compose -f docker-compose.yml up -d || { echo "Failed to start docker-compose."; exit 1; }