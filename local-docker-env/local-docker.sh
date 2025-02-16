#!/bin/bash

# gwemix path를 환경변수에 추가
current_dir=$(pwd)
export PATH=$PATH:${current_dir}/build/bin

# 옵션 파싱
KEEP_NODEKEY=false
while getopts "a:v:-:" opt; do
  case ${opt} in
  a)
    NODE_NUM=$OPTARG
    ;;
  v)
    VERSION=$OPTARG
    ;;
  -)
    case "${OPTARG}" in
      keep-nodekey)
        KEEP_NODEKEY=true
        ;;
      *)
        echo "Usage: $0 -a <node_num> -v <version> [--keep-nodekey]"
        exit 1
        ;;
    esac
    ;;
  \?)
    echo "Usage: $0 -a <node_num> -v <version> [--keep-nodekey]"
    exit 1
    ;;
  esac
done
shift $((OPTIND -1))

# 필수 인수 확인
if [ -z "$NODE_NUM" ]; then
  echo "Node number is required. Usage: $0 -a <node_num>"
  exit 1
fi

# Set default version if not provided
if [ -z "$VERSION" ]; then
    VERSION="latest"
fi

# 기존에 실행중인 docker-compose 중지
docker compose -f docker-compose.yml down || { echo "Failed to stop docker-compose."; }

# gen-account.sh 실행
chmod +x local-docker-env/gen-account.sh
./local-docker-env/gen-account.sh -a "$NODE_NUM" || { echo "Failed to execute gen-account.sh."; exit 1; }

# keep node key 옵션이 없으면 gen-nodekey.sh 실행, 있다면 node num 만큼 nodekey1부터 nodeketnum 까지 있는지 확인
if [ "$KEEP_NODEKEY" = false ]; then
  chmod +x local-docker-env/gen-nodekey.sh
  ./local-docker-env/gen-nodekey.sh -a "$NODE_NUM" || { echo "Failed to execute gen-nodekey.sh."; exit 1; }
  echo "Nodekey generated successfully."
else
  for ((i = 1; i <= NODE_NUM; i++)); do
    if [ ! -f "local-docker-env/nodekey/nodekey$i" ]; then
      echo "Nodekey $i does not exist"
      exit 1
    fi
  done
  echo "Nodekey exists."
fi

# gen-config.sh 실행
chmod +x local-docker-env/gen-config.sh
./local-docker-env/gen-config.sh -a "$NODE_NUM" -f local-docker-env/config.json || { echo "Failed to execute gen-config.sh."; exit 1; }

# gen-docker-compose.sh 실행
chmod +x local-docker-env/gen-docker-compose.sh
./local-docker-env/gen-docker-compose.sh -a "$NODE_NUM" -f docker-compose.yml -v "$VERSION" || { echo "Failed to execute gen-docker-compose.sh."; exit 1; }

# Dockerfile.boot 및 Dockerfile.node 파일 복사
cp local-docker-env/Dockerfile.boot ./ || { echo "Failed to copy Dockerfile.boot."; exit 1; }
cp local-docker-env/Dockerfile.node ./ || { echo "Failed to copy Dockerfile.node."; exit 1; }

# docker-compose.yml 파일을 이용해 docker-compose build 및 up 실행
docker compose -f docker-compose.yml build --no-cache || { echo "Failed to build docker-compose."; exit 1; }
docker compose -f docker-compose.yml up -d || { echo "Failed to start docker-compose."; exit 1; }